package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func getTracks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	query := `SELECT id, title, artist FROM trackList`
	rows, err := Db.Query(query)
	if err != nil {
		fmt.Println("Error retreiving tracks:", err)
	}
	defer rows.Close()

	var Tracks []Track

	for rows.Next() {
		var track Track

		err := rows.Scan(&track.ID, &track.Title, &track.Artist)
		if err != nil {
			fmt.Println("Could not scan tracks:", err)
			return
		}
		Tracks = append(Tracks, track)
	}
	json.NewEncoder(w).Encode(Tracks)
}

func getTrack(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]
	row := Db.QueryRow("SELECT id, title, artist FROM trackList WHERE id=$1", id)

	var track Track
	err := row.Scan(&track.ID, &track.Title, &track.Artist)
	if err != nil {
		fmt.Println("Could not scan tracks:", err)
		return
	}
	json.NewEncoder(w).Encode(track)
}

func deleteTrack(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]
	_, err := Db.Exec("DELETE FROM trackList WHERE id=$1", id)
	if err != nil {
		fmt.Println("ERROR; Cannot delete track from TrackList:", err)
	}
}

func updateTrack(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	var track Track
	json.NewDecoder(r.Body).Decode(&track)

	_, err := Db.Exec("UPDATE trackList SET title=$1 artist=$2 WHERE id=$3", track.Title, track.Artist, id)
	if err != nil {
		fmt.Println("ERROR; Cannot update tack in trackList:", err)
	}
	w.WriteHeader(http.StatusOK)
}

func createTrack(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var track Track
	json.NewDecoder(r.Body).Decode(&track)

	_, err := Db.Exec("INSERT INTO trackList (title, artist) VALUES($1, $2)", track.Title, track.Artist)
	if err != nil {
		fmt.Println("ERROR; Cannot vreate new track:", err)
	}
}
