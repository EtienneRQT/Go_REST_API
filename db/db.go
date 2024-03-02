package db

import (
	"database/sql"
	"fmt"

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

const (
	createUsersTable = `CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT, 
        email TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL);`

	createEventsTable = `CREATE TABLE IF NOT EXISTS events (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL, 
        description TEXT NOT NULL,  
        location TEXT NOT NULL, 
        dateTime DATETIME NOT NULL, 
        user_id INTEGER NOT NULL,
        FOREIGN KEY (user_id) REFERENCES users(id));`
)

// createTables creates the users and events tables if they do not already exist.
// It executes a combined SQL statement containing CREATE TABLE IF NOT EXISTS for both tables.
// If there is an error executing the SQL, it will panic with the error message.
func createTables() {
	createTablesQuery := createUsersTable + "\n" + createEventsTable
	_, err := DB.Exec(createTablesQuery)

	if err != nil {
		panic(fmt.Sprintf("Could not create tables: %v", err))
	}
}
