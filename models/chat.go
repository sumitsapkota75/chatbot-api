package models

import (
	"github.com/google/uuid"
	"gopkg.in/mgo.v2/bson"
)

type BINARY16 uuid.UUID
type UserMessage struct {
	Text string `json:"text"`
}

type Messages struct {
	Text   string `bson:"text" json:"text"`
	IsUser bool   `bson:"isUser" json:"isUser"`
}

type Conversation struct {
	ConversationId string     `bson:"conversation_id" json:"conversation_id"`
	Messages       []Messages `bson:"messages" json:"messages"`
	Email          string     `bson:"email" json:"email"`
}
type GetConversation struct {
	ConversationId bson.Binary `bson:"id" json:"conversationId"`
	Messages       []Messages  `bson:"messages" json:"messages"`
	Email          string      `bson:"email" json:"email"`
}
