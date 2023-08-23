package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/tracks", getTracks).Methods("GET")
	router.HandleFunc("/tracks/{id}", getTrack).Methods("GET")
	router.HandleFunc("/track", createTrack).Methods("POST")
	router.HandleFunc("/tracks/{id}", updateTrack).Methods("PUT")
	router.HandleFunc("/tracks/{id}", deleteTrack).Methods("DELETE")

	fmt.Println("Starting sever at port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
