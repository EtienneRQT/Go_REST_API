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
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

// createTables creates the events table in the database if it does not already exist.
// It executes a CREATE TABLE statement and handles any errors.
func createTables() {
	createEventsTable := `CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		 name TEXT NOT NULL, 
		 description TEXT NOT NULL,  
		 location TEXT NOT NULL, 
		 dateTime DATETIME NOT NULL, 
		 user_id INTEGER NOT NULL)`

	_, err := DB.Exec(createEventsTable)

	if err != nil {
		panic("Could not create events table.")
	}
}
