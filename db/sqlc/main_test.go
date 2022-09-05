package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:mysecretpassword@localhost:5432/go4money?sslmode=disable"
)

// Global variables
var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	// Database connections
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Error while connecting to db:", err)
	}

	// Queries and functions allowed
	testQueries = New(testDB)

	// Execute Go tests and then finish
	os.Exit(m.Run())
}
