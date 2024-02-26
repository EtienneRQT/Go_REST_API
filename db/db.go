package db

import (
	"database/sql"
	// Imports the sqlite3 driver from github.com/mattn/go-sqlite3.
	// This allows connecting to and querying SQLite databases.
	_ "github.com/mattn/go-sqlite3"
)

// DB is the application's SQL database.
var DB *sql.DB

// InitDB opens a connection to the sqlite3 database file api.db.
func InitDB() {
	DB, err := sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
}
