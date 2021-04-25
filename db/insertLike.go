package db

import (
	"context"
	"time"

	"github.com/JavDomGom/sn-test/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* InsertLike record like in database. */
func InsertLike(IDUser, IDMsg string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(DBname)
	col := db.Collection("likes")

	var (
		like   models.Like
		filter bson.M
		update bson.D
	)

	// Get like object if exist.
	filter = bson.M{"userId": IDUser}
	err := col.FindOne(ctx, filter).Decode(&like)

	// If like doesn't exist create it.
	if err != nil {
		like.UserID = IDUser
		like.LikesList = append(like.LikesList, IDMsg)

		_, err = col.InsertOne(ctx, like)
		if err != nil {
			return false, err
		}
	} else {
		// Append IDMsg to likesList.
		update := bson.D{{Key: "$set",
			Value: bson.D{
				{Key: "likesList", Value: append(like.LikesList, IDMsg)},
			},
		}}

		filter = bson.M{"userId": bson.M{"$eq": like.UserID}}
		_, err = col.UpdateOne(ctx, filter, update)
		if err != nil {
			return false, err
		}
	}

	// Get liked message to update likes counter.
	col = db.Collection("messages")

	var likedMsg models.ReturnMsg

	ObjIdMsg, _ := primitive.ObjectIDFromHex(IDMsg)

	// Get liked message.
	filter = bson.M{"_id": ObjIdMsg}
	col.FindOne(ctx, filter).Decode(&likedMsg)

	// Increment +1 likes count in message.
	update = bson.D{{Key: "$set",
		Value: bson.D{
			{Key: "likesCount", Value: likedMsg.LikesCount + 1},
		},
	}}

	filter = bson.M{"_id": bson.M{"$eq": likedMsg.ID}}
	_, err = col.UpdateOne(ctx, filter, update)
	if err != nil {
		return false, err
	}

	return true, nil
}
