package services

import (
	"context"
	"errors"

	"ai-backend/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type SaveChatService struct {
	client *mongo.Client
}

// NewSaveChatService creates a new instance of SaveChatService
func NewSaveChatService(client *mongo.Client) *SaveChatService {
	return &SaveChatService{client: client}
}

func (scs *SaveChatService) SaveChatService(conversationData models.Conversation) error {
	collection := scs.client.Database("ai-db").Collection("conversations")

	// Define an update object to push new messages
	update := bson.M{
		"$push": bson.M{"messages": bson.M{"$each": conversationData.Messages}},
		"$set":  bson.M{"email": conversationData.Email}, // Update email field
	}

	// Upsert: Update existing or insert a new document if none exists
	upsert := true
	result, err := collection.UpdateOne(
		context.Background(),
		bson.M{"conversation_id": conversationData.ConversationId},
		update,
		&options.UpdateOptions{Upsert: &upsert},
	)

	// Handle different scenarios based on the result and error
	if err != nil {
		return err
	}

	// Check if document was modified (updated or inserted)
	if result.ModifiedCount == 0 && result.MatchedCount == 0 {
		return errors.New("unexpected: document not modified or matched")
	}

	return nil
}

func (us *SaveChatService) GetConversation(email string) ([]models.Conversation, error) {
	collection := us.client.Database("ai-db").Collection("conversations")

	// Define a filter to search for conversations by email
	filter := bson.M{"email": email}

	// Create a slice to hold the results
	var conversations []models.Conversation

	// Find all documents that match the filter
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	// Decode all documents into the conversations slice
	err = cursor.All(context.Background(), &conversations)
	if err != nil {
		return nil, err
	}

	return conversations, nil
}
