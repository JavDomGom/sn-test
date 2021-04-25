package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/JavDomGom/sn-test/models"
	"go.mongodb.org/mongo-driver/bson"
)

/* InsertMsg records a msg in database. */
func InsertMsg(msg models.RecordMsg) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(DBname)
	col := db.Collection("messages")

	register := bson.M{
		"userId":             msg.UserID,
		"message":            msg.Message,
		"datetime":           msg.Datetime,
		"inReplyToMessageId": msg.InReplyToMessageID,
	}
	result, err := col.InsertOne(ctx, register)
	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
