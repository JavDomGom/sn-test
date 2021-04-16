package models

import "time"

/* RecordMsg Structure for message. */
type RecordMsg struct {
	UserID             string    `bson:"userId" json:"userId,omitempty"`
	Message            string    `bson:"message" json:"message,omitempty"`
	Datetime           time.Time `bson:"datetime" json:"datetime,omitempty"`
	InReplyToMessageID string    `bson:"inReplyToMessageId" json:"inReplyToMessageId,omitempty"`
}
