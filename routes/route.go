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
	router.GET("/get-conversation", controller.GetConversation)
	router.POST("/chat", controller.ChatWithGemini)

	router.POST("/save-conversation", controller.SaveConversation)

	log.Println("\033[93mChatGPT started. Press CTRL+C to quit.\033[0m")
	router.Run()
}
