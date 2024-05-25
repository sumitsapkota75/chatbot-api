package services

import (
	"context"
	"errors"

	"ai-backend/models"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type UserService struct {
	client *mongo.Client
}

// NewUserService creates a new instance of UserService
func NewUserService(client *mongo.Client) *UserService {
	return &UserService{client: client}
}

// CreateUser creates a new user
func (us *UserService) CreateUserService(user models.User) error {
	collection := us.client.Database("ai-db").Collection("users")

	// Check for existing user with the same email
	filter := bson.M{"email": user.Email}
	var existingUser models.User
	err := collection.FindOne(context.Background(), filter).Decode(&existingUser)

	// Handle different scenarios based on the find result
	if err == mongo.ErrNoDocuments {
		// No existing user found, proceed with creation
		_, err = collection.InsertOne(context.Background(), user)
		return err
	} else if err != nil {
		return err // Handle other errors during the find operation
	}

	// Existing user found, return an error
	return errors.New("user with email already exists")
}
