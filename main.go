package main

import (
	"ai-backend/helpers"
	routers "ai-backend/routes"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	// load env variables
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}
	// connect to mongoDB database
	helpers.ConnectToMongoDB()
	// initialize routes
	routers.Router()
}
