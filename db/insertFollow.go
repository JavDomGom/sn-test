package db

import (
	"context"
	"time"

	"github.com/JavDomGom/sn-test/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* InsertFollow record follow in database. */
func InsertFollow(follow models.Follow) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(DBname)
	col := db.Collection("follows")

	_, err := col.InsertOne(ctx, follow)
	if err != nil {
		return false, err
	}

	// Get followed and follower profiles to update their followers and followed counters.
	col = db.Collection("users")

	var (
		followedProfile models.User
		followerProfile models.User
	)

	followedObjID, _ := primitive.ObjectIDFromHex(follow.UserFollowedID)
	followerObjID, _ := primitive.ObjectIDFromHex(follow.UserID)

	// Get followed profile.
	filter := bson.M{"_id": followedObjID}
	col.FindOne(ctx, filter).Decode(&followedProfile)

	// Get follower profile.
	filter = bson.M{"_id": followerObjID}
	col.FindOne(ctx, filter).Decode(&followerProfile)

	// Increment +1 followed count in followerProfile.
	update := bson.D{{Key: "$set",
		Value: bson.D{
			{Key: "followingCount", Value: followerProfile.FollowingCount + 1},
		},
	}}

	filter = bson.M{"_id": bson.M{"$eq": followerProfile.ID}}
	_, err = col.UpdateOne(ctx, filter, update)
	if err != nil {
		return false, err
	}

	// Increment +1 followers count in followedProfile.
	update = bson.D{{Key: "$set",
		Value: bson.D{
			{Key: "followersCount", Value: followedProfile.FollowersCount + 1},
		},
	}}

	filter = bson.M{"_id": bson.M{"$eq": followedProfile.ID}}
	_, err = col.UpdateOne(ctx, filter, update)
	if err != nil {
		return false, err
	}

	return true, nil
}
