package services

import (
	"context"
	"time"

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

	currentTime := time.Now()

	// Check for existing document
	filter := bson.M{"conversation_id": conversationData.ConversationId}
	var existingDoc models.Conversation
	err := collection.FindOne(context.Background(), filter).Decode(&existingDoc)

	// Update object with conditional timestamp
	update := bson.M{
		"$set": bson.M{"email": conversationData.Email, "messages": conversationData.Messages},
	}

	// Include timestamp only on creation
	if err == mongo.ErrNoDocuments { // No document found, so it's a new conversation
		update["$set"] = bson.M{
			"email":      conversationData.Email,
			"time_stamp": currentTime,
			"messages":   conversationData.Messages,
		}
	} else if err != nil { // Handle other potential errors during document check
		return err
	}

	// Upsert: Update existing or insert a new document if none exists
	upsert := true
	_, err = collection.UpdateOne(
		context.Background(),
		filter, // Use same filter for checking and update
		update,
		&options.UpdateOptions{Upsert: &upsert},
	)

	// Handle different scenarios based on the result and error
	if err != nil {
		return err
	}

	return nil
}

func (us *SaveChatService) GetConversation(email string) ([]models.Conversation, error) {
	collection := us.client.Database("ai-db").Collection("conversations")

	// Define a filter to search for conversations by email
	filter := bson.M{"email": email}

	// Specify sort order by timestamp (descending)
	sort := bson.M{"time_stamp": -1} // -1 indicates descending order

	// Create a slice to hold the results
	var conversations []models.Conversation

	// Find all documents with sorting applied
	cursor, err := collection.Find(context.Background(), filter, &options.FindOptions{Sort: sort})
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

func (us *SaveChatService) GetConversationByID(conversationID string) (models.Conversation, error) {
	collection := us.client.Database("ai-db").Collection("conversations")

	// Define a filter to search for conversation by ID
	filter := bson.M{"conversation_id": conversationID}

	// Create a variable to hold the conversation
	var conversation models.Conversation

	// Find one document that matches the filter
	err := collection.FindOne(context.Background(), filter).Decode(&conversation)
	if err != nil {
		if err == mongo.ErrNoDocuments { // Handle case where no document is found
			return models.Conversation{}, err // Return an empty conversation struct and the error
		}
		return models.Conversation{}, err // Return an empty conversation struct and the error for other errors
	}

	return conversation, nil
}
