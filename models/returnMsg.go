package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* ReturnMsg Structure to return each message. */
type ReturnMsg struct {
	ID                 primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID             string             `bson:"userId" json:"userId,omitempty"`
	Message            string             `bson:"message" json:"message,omitempty"`
	Datetime           time.Time          `bson:"datetime" json:"datetime,omitempty"`
	InReplyToMessageID string             `bson:"inReplyToMessageId" json:"inReplyToMessageId,omitempty"`
	LikesCount         uint64             `bson:"likesCount" json:"likesCount"`
}
