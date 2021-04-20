package db

import (
	"context"
	"time"

	"github.com/JavDomGom/sn-test/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* GetMsg Get a message by id. */
func GetMsg(ID string) (models.ReturnMsg, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database(DBname)
	col := db.Collection("messages")

	var msg models.ReturnMsg
	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condition).Decode(&msg)
	if err != nil {
		return msg, false
	}
	return msg, true
}
