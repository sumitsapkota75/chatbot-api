package controller

import (
	db "ai-backend/helpers"
	"ai-backend/models"
	"ai-backend/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SaveConversation(c *gin.Context) {
	var conversation models.Conversation
	if err := c.ShouldBindJSON(&conversation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	conversationService := services.NewSaveChatService(db.Client)
	if err := conversationService.SaveChatService(conversation); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Conversation saved successfully", "conversation": conversation})
}

func GetConversation(c *gin.Context) {
	email := c.Query("email")
	// Check if email parameter is missing
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing email parameter"})
		return
	}
	conversationService := services.NewSaveChatService(db.Client)
	conversation, err := conversationService.GetConversation(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(conversation)
	c.JSON(http.StatusOK, gin.H{"data": conversation})
}
