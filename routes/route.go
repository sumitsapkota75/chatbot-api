package routers

import (
	"ai-backend/controller"
	"log"

	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()
	router.Use(gin.Logger())

	router.POST("/chat", controller.ChatWithGemini)
	// router.POST("/fetchHistory", handlers.HandleFetchHistory)

	log.Println("\033[93mChatGPT started. Press CTRL+C to quit.\033[0m")
	router.Run()
}
