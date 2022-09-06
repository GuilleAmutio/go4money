package main

import (
	"database/sql"
	"log"

	"github.com/guilleamutio/go4money/api"
	db "github.com/guilleamutio/go4money/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://postgres:mysecretpassword@localhost:5432/go4money?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	// Database connections
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Error while connecting to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Error while starting webserver:", err)
	}
}
