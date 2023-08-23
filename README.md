# CRUD API using Go and postgresSQL 

This is a simple CRUD (Create, Reas, Update, Delete) API implementation in Go using the `gorilla/mux` router and PostgreSQL as the database.

## Setup
1. Clone the repository to your local machine:
   git clone https://github.com:fiificoder/mux-postgres.git

2. Install the required dependancies using `go get` :
   go get github.com/gorilla/mux
   go get github.com/lib/pq


3. Create a PostgreSQL database and update the connection string:
- Create a PostgreSQL database and a "trackList" table with appropriate columns (id, title, artist).
- Update the `connStr` variable in `db.go` with your PostgreSQL connection string.

## Running the API

1. Navigate to the project directory:
cd your-crud-api


2. Run the API using:
go run main.go


3. The API will be available at: `http://localhost:8080`

## Endpoints

- `GET /tracks`: Retrieve a list of all tracks.
- `GET /tracks/{id}`: Retrieve details of a specific track.
- `POST /track`: Create a new track.
- `PUT /tracks/{id}`: Update details of a track.
- `DELETE /tracks/{id}`: Delete a track.

## Usage

- To retrieve a list of all tracks:
curl http://localhost:8080/tracks


- To retrieve details of a specific track:
curl http://localhost:8080/books/1


- To create a new track:
curl -X POST -H "Content-Type: application/json" -d '{"title":"New Track","artist":"John Legend"}' http://localhost:8080/tracks


- To update details of a track:
curl -X PUT -H "Content-Type: application/json" -d '{"title":"Updated Title","artist":"Jane Smith"}' http://localhost:8080/tracks/1


- To delete a book:
curl -X DELETE http://localhost:8080/tracks/1



## Notes

- This is a basic example for educational purposes and may lack features like input validation.
- Make sure to adapt the code for production use, including proper error handling and security measures.




