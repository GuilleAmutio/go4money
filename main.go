package main

import (
	"database/sql"
	"log"

	"github.com/guilleamutio/go4money/api"
	db "github.com/guilleamutio/go4money/db/sqlc"
	"github.com/guilleamutio/go4money/util"
	_ "github.com/lib/pq"
)

func main() {
	// Load env variables
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	// Database connections
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Error while connecting to db:", err)
	}

	// Create server
	store := db.NewStore(conn)
	server := api.NewServer(store)

	// Start Server
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Error while starting webserver:", err)
	}
}
