package db

import (
	"context"
	"time"

	"github.com/JavDomGom/sn-test/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* InsertFollow record relation in database. */
func InsertFollow(t models.Follow) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(DBname)
	col := db.Collection("follows")

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}

	// Get followed and follower profiles to update their followers and followed counters.
	usersCol := db.Collection("users")

	var (
		followedProfile models.User
		followerProfile models.User
	)

	followedObjID, _ := primitive.ObjectIDFromHex(t.UserFollowedID)
	followerObjID, _ := primitive.ObjectIDFromHex(t.UserID)

	// Get followed profile.
	condition := bson.M{"_id": followedObjID}
	usersCol.FindOne(ctx, condition).Decode(&followedProfile)

	// Get follower profile.
	condition = bson.M{"_id": followerObjID}
	usersCol.FindOne(ctx, condition).Decode(&followerProfile)

	// Increment +1 followed count in followerProfile.
	update := bson.D{{Key: "$set",
		Value: bson.D{
			{Key: "followingCount", Value: followerProfile.FollowingCount + 1},
		},
	}}

	filter := bson.M{"_id": bson.M{"$eq": followerProfile.ID}}

	_, err = usersCol.UpdateOne(ctx, filter, update)
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

	_, err = usersCol.UpdateOne(ctx, filter, update)
	if err != nil {
		return false, err
	}

	return true, nil
}
