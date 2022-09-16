package main

import (
	"github.com/guilleamutio/go4money/cmd"
	"github.com/guilleamutio/go4money/database"
	"github.com/guilleamutio/go4money/util"
	"log"
)

// @title           Swagger go4money API
// @version         1.0
// @description     This is a UI for go4money API.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
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
