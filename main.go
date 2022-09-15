package main

import (
	"github.com/guilleamutio/go4money/cmd"
	"github.com/guilleamutio/go4money/database"
	"github.com/guilleamutio/go4money/util"
	"log"
)

func main() {
	// Load env variables
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("failed loading config", err)
	}

	// Open DB connection
	db, err := database.OpenDatabase(config)
	if err != nil {
		log.Fatal("failed while connecting to the database", err)
	}

	// Create webserver
	server := cmd.NewServer(db)

	// Run webserver
	err = server.StartServer(config.ServerAddress)
	if err != nil {
		log.Fatal("failed while starting the server", err)
	}
}
