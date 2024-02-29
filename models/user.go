package models

import (
	"example.com/api/db"
	"example.com/api/utils"
)

// User represents a user account with ID, username, email and password.
// User is used for user account management and authentication.
type User struct {
	ID       int64  `binding:"required"`
	Username string `binding:"required"`
	Email    string `binding:"required"`
	Password string `binding:"required"`
}


// Save inserts a new user record into the database.
// It hashes the password before inserting for security.
// Returns any errors that occur when preparing/executing the INSERT statement.
func (user *User) Save() error {
	query := "INSERT INTO users (username, email, password) VALUES (?,?,?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	hashedPassword := utils.HashPassword(user.Password)
	result, err := stmt.Exec(user.Username, user.Email, hashedPassword)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	user.ID = id

	return err
}
