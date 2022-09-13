package main

import (
	"context"
	"github.com/guilleamutio/go4money/ent"
	"github.com/guilleamutio/go4money/util"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	// Load env variables
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("failed loading config", err)
	}

	// Open connection with database
	client, err := ent.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Start server
}
