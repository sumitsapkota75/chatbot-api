package helpers

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

// ConnectToMongoDB establishes a connection to the MongoDB database
func ConnectToMongoDB() (*mongo.Client, error) {
	fmt.Println("Connecting to MongoDB...")
	ctx := context.Background()
	db_username := os.Getenv("DB_USERNAME")
	db_password := os.Getenv("DB_PASSWORD")
	connectionString := fmt.Sprintf("mongodb+srv://%s:%s@ai-db.nfdl7hv.mongodb.net/?retryWrites=true&w=majority&appName=ai-db", db_username, db_password)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Ping the database to verify successful connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	log.Println("Connected to MongoDB successfully!")
	Client = client
	return client, nil
}
