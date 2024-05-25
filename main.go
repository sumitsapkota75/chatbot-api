package main

import (
	"ai-backend/helpers"
	routers "ai-backend/routes"
)

func main() {
	helpers.ConnectToMongoDB()
	routers.Router()
}
