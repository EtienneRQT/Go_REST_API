package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes the provided password using bcrypt with a cost of 14.
// It returns the hashed password as a string, or panics if there is an error.
func HashPassword(password string) (string, error) {
	bcrypt, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		panic(err)
	}

	return string(bcrypt), nil
}

// CheckPasswordHash compares a plaintext password against a hashed password to check if they match.
// It returns true if the passwords match, false otherwise.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
