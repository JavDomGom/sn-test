package models

/* Msg Capture the received message from the body. */
type Msg struct {
	Message            string `bson:"message" json:"message"`
	InReplyToMessageID string `bson:"inReplyToMessageId" json:"inReplyToMessageId"`
}
