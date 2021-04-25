package db

import (
	"context"
	"time"

	"github.com/JavDomGom/sn-test/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* RemoveFollow remove a specific follow. */
func RemoveFollow(follow models.Follow) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(DBname)
	col := db.Collection("follows")

	_, err := col.DeleteOne(ctx, follow)
	if err != nil {
		return false, err
	}

	// Get unfollowed and unfollower profiles to update their followers and followed counters.
	col = db.Collection("users")

	var (
		unfollowedProfile models.User
		unfollowerProfile models.User
	)

	followedObjID, _ := primitive.ObjectIDFromHex(follow.UserFollowedID)
	followerObjID, _ := primitive.ObjectIDFromHex(follow.UserID)

	// Get followed profile.
	filter := bson.M{"_id": followedObjID}
	col.FindOne(ctx, filter).Decode(&unfollowedProfile)

	// Get follower profile.
	filter = bson.M{"_id": followerObjID}
	col.FindOne(ctx, filter).Decode(&unfollowerProfile)

	// Decrement -1 followed count in unfollowerProfile.
	update := bson.D{{Key: "$set",
		Value: bson.D{
			{Key: "followingCount", Value: unfollowerProfile.FollowingCount - 1},
		},
	}}

	filter = bson.M{"_id": bson.M{"$eq": unfollowerProfile.ID}}

	_, err = col.UpdateOne(ctx, filter, update)
	if err != nil {
		return false, err
	}

	// Decrement -1 followers count in unfollowedProfile.
	update = bson.D{{Key: "$set",
		Value: bson.D{
			{Key: "followersCount", Value: unfollowedProfile.FollowersCount - 1},
		},
	}}

	filter = bson.M{"_id": bson.M{"$eq": unfollowedProfile.ID}}

	_, err = col.UpdateOne(ctx, filter, update)
	if err != nil {
		return false, err
	}

	return true, nil
}
