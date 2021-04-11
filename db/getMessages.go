package db

import (
	"context"
	"time"

	"github.com/JavDomGom/sn-test/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* GetMessages returns all messages from a user's profile. */
func GetMessages(ID string, page int64) ([]*models.ReturnMsg, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database(DBname)
	col := db.Collection("messages")

	var messages []*models.ReturnMsg

	condition := bson.M{
		"userId": ID,
	}

	opts := options.Find()
	opts.SetLimit(20)
	opts.SetSort(bson.D{{Key: "datetime", Value: -1}})
	opts.SetSkip((page - 1) * 20)

	cursor, err := col.Find(ctx, condition, opts)
	if err != nil {
		return messages, false
	}

	for cursor.Next(context.TODO()) {

		var msg models.ReturnMsg
		err := cursor.Decode(&msg)
		if err != nil {
			return messages, false
		}
		messages = append(messages, &msg)
	}
	return messages, true
}
