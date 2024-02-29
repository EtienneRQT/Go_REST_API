package routes

import (
	"net/http"

	"example.com/api/models"
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
