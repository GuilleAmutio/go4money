package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/guilleamutio/go4money/util"
	_ "github.com/lib/pq"
)

// Global variables
var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	// Load env variables
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// Database connections
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Error while connecting to db:", err)
	}

	// Queries and functions allowed
	testQueries = New(testDB)

	// Execute Go tests and then finish
	os.Exit(m.Run())
}
