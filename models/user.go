package models

import (
	"database/sql"
	"errors"

	"example.com/api/db"
	"example.com/api/utils"
)

// User represents a user account with ID, email and password.
// User is used for user account management and authentication.
type User struct {
	ID       int64  `binding:"required"`
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

// Save inserts a new user record into the database.
// It hashes the password before inserting for security.
// Returns any errors that occur when preparing/executing the INSERT statement.
func (user *User) Save() error {
	query := "INSERT INTO users (email, password) VALUES (?,?,?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	hashedPassword := utils.HashPassword(user.Password)
	result, err := stmt.Exec(user.Email, hashedPassword)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	user.ID = id

	return err
}

// Login authenticates a user by email and password. It compares the
// hashed password to verify the correct password was provided.
// Returns ErrNoRows if no user with the email exists, an invalid
// password error if the password does not match, or any other
// errors that occur.
func (user *User) Login() error {
	query := "SELECT password FROM users WHERE email = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	hashedPassword := utils.HashPassword(user.Password)
	var requiredPassword string
	err = stmt.QueryRow(user.Email).Scan(&requiredPassword)
	if err == sql.ErrNoRows {
		return errors.New("No user with this email exists")
	} else if err != nil {
		return err
	}

	if !utils.CheckPasswordHash(requiredPassword, hashedPassword) {
		return errors.New("Invalid password")
	}

	return nil
}
