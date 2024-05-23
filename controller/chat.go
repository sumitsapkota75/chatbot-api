package controller

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"

	helper "ai-backend/helpers"
	models "ai-backend/models"
)

func ChatWithGemini(c *gin.Context) {
	var message models.UserMessage
	err := c.ShouldBindJSON(&message)
	if err != nil {
		log.Printf("Error processing request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	model := client.GenerativeModel("gemini-pro")
	response, err := model.GenerateContent(ctx, genai.Text(message.Text))
	if err != nil {
		log.Printf("Error getting response from Gemini: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	formattedResponse := helper.FormatResponse(response)
	c.JSON(http.StatusOK, gin.H{
		"response": formattedResponse,
	})
}
