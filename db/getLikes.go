package db

import (
	"context"
	"time"

	"github.com/JavDomGom/sn-test/models"
	"go.mongodb.org/mongo-driver/bson"
)

/* GetLikes returns all likes by user id. */
func GetLikes(ID string) (models.Like, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database(DBname)
	col := db.Collection("likes")

	var like models.Like

	filter := bson.M{
		"userId": ID,
	}

	err := col.FindOne(ctx, filter).Decode(&like)
	if err != nil {
		return like, false
	}
	return like, true
}
