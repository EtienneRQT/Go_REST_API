package middlewares

import (
	"net/http"

	"github.com/EtienneRQT/Go_REST_API/utils"
	"github.com/gin-gonic/gin"
)

// Authenticate is a middleware function that authenticates requests.
// It checks for an Authorization header, verifies the token, and sets
// the user_id on the context if valid. Otherwise it aborts with 401 Unauthorized.
func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	userID, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	context.Set("userId", userID)
	context.Next()
}
