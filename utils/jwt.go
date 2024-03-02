package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "super-secret-key"

// GenrateToken generates a JWT token for the given email and userID,
// with a 24 hour expiration.
func GenrateToken(email string, userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"email":  email,
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}
