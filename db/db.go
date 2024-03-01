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

// createTables initializes the database by creating the users and events tables.
// It creates each table by executing a SQL CREATE TABLE statement.
// If there is an error creating a table, it panics.
func createTables() {
	createUsersTable := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		username UNIQUE TEXT NOT NULL,
		email UNIQUE TEXT NOT NULL,
		password TEXT NOT NULL)`

	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("Could not create users table.")
	}

	createEventsTable := `CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL, 
		description TEXT NOT NULL,  
		location TEXT NOT NULL, 
		dateTime DATETIME NOT NULL, 
		user_id INTEGER NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id))`

	_, err = DB.Exec(createEventsTable)

	if err != nil {
		panic("Could not create events table.")
	}
}
