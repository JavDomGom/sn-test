package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/JavDomGom/sn-test/models"
	"go.mongodb.org/mongo-driver/bson"
)

/* InsertMsg records a msg in database. */
func InsertMsg(t models.RecordMsg) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("test")
	col := db.Collection("messages")

	register := bson.M{
		"userId":   t.UserID,
		"message":  t.Message,
		"datetime": t.Datetime,
	}
	result, err := col.InsertOne(ctx, register)
	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
