package routes

import (
	"fmt"
	"net/http"
	"rest-api/models"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Inavild body. Error message: %s", err)})
		return
	}

	createdUser, err := user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Couldn't save user. Error message: %s", err)})
		return

	}

	context.JSON(http.StatusOK, createdUser)
}
