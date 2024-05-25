package controller

import (
	db "ai-backend/helpers"
	"ai-backend/models"
	"ai-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUserController(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userService := services.NewUserService(db.Client)
	if err := userService.CreateUserService(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
