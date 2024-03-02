package routes

import (
	"net/http"

	"example.com/api/models"
	"example.com/api/utils"
	"github.com/gin-gonic/gin"
)

// signup handles user registration. It binds the JSON body to a User model,
// saves the user to the database, and returns the saved user or any errors.
func signup(context *gin.Context) {
	var user models.User
	if err := context.BindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := user.Save(); err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusCreated, user)
}

// login handles user authentication. It binds the JSON body to a User model,
// validates the user credentials, and returns the authenticated user or any errors.
func login(context *gin.Context) {
	var user models.User
	if err := context.BindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := user.Login(); err != nil {
		context.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	token, err := utils.GenrateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": token, "message": "Logged in successfully"})
}
