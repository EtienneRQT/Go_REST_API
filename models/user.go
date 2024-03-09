package models

import (
	"database/sql"
	"errors"

	"github.com/EtienneRQT/Go_REST_API/db"
	"github.com/EtienneRQT/Go_REST_API/utils"
)

// User represents a user account with ID, email and password.
// User is used for user account management and authentication.
type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

// Save inserts a new user record into the database.
// It hashes the password before inserting for security.
// Returns any errors that occur when preparing/executing the INSERT statement.
func (user *User) Save() error {
	query := "INSERT INTO users (email, password) VALUES (?,?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	if hashedPassword == "" {
		return errors.New("Password is required")
	}

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
	query := "SELECT password, id FROM users WHERE email = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var requiredPassword string
	err = stmt.QueryRow(user.Email).Scan(&requiredPassword, &user.ID)
	if err == sql.ErrNoRows {
		return errors.New("Invalid credentials")
	} else if err != nil {
		return err
	}

	if !utils.CheckPasswordHash(user.Password, requiredPassword) {
		return errors.New("Invalid credentials")
	}

	return nil
}
