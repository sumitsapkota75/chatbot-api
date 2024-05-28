package routers

import (
	"ai-backend/controller"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(cors.Default())

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Codex server up and running..."})
	})

	// get all conversation of user
	router.GET("/get-conversation", controller.GetConversation)
	// get a single conversation
	router.GET("/single-chat/:id", controller.GetConversationByID)
	// create chat with gemini sdk
	router.POST("/chat", controller.ChatWithGemini)
	// save new chats to the database
	router.POST("/save-conversation", controller.SaveConversation)

	log.Println("\033[93mChatGPT started. Press CTRL+C to quit.\033[0m")
	router.Run()
}
