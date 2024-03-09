package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// CheckPasswordHash compares a plaintext password against a bcrypt hashed password.
// It returns true if the passwords match, false otherwise.
// This is used to verify a user's password during login.
func TestHashPassword(t *testing.T) {
	password := "password123"

	hashed, err := HashPassword(password)

	if err != nil {
		t.Errorf("Error hashing password: %v", err)
	}

	assert.NotEmpty(t, hashed)
	assert.NotEqual(t, password, hashed)

	isValid := CheckPasswordHash(password, hashed)
	assert.True(t, isValid)
}

// TestCheckPasswordHashInvalid tests checking a password hash
// against an invalid hash. It verifies that CheckPasswordHash()
// returns false when the hash does not match.
func TestCheckPasswordHashInvalid(t *testing.T) {
	password := "password123"
	invalidHash := "invalidhash"

	isValid := CheckPasswordHash(password, invalidHash)

	assert.False(t, isValid)
}

// TestCheckPasswordHashWrongPassword tests checking a hashed password
// against an incorrect plaintext password. It verifies that
// CheckPasswordHash() returns false when the provided password does
// not match the hashed password.
func TestCheckPasswordHashWrongPassword(t *testing.T) {
	password := "password123"
	wrongPassword := "wrongpassword"
	hashed, err := HashPassword(password)

	if err != nil {
		t.Errorf("Error hashing password: %v", err)
	}

	isValid := CheckPasswordHash(wrongPassword, hashed)

	assert.False(t, isValid)
}
