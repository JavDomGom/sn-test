package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* ReturnFollowersMsg Structure for returned message. */
type ReturnFollowersMsg struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID         string             `bson:"userId" json:"userId,omitempty"`
	UserFollowedID string             `bson:"userFollowedID" json:"userFollowedID,omitempty"`
	Msg            struct {
		Message  string    `bson:"message" json:"message,omitempty"`
		Datetime time.Time `bson:"datetime" json:"datetime,omitempty"`
		ID       string    `bson:"_id" json:"_id,omitempty"`
	}
}
