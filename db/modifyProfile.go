package db

import (
	"context"
	"time"

	"github.com/JavDomGom/sn-test/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* ModifyProfile modify the user's profile. */
func ModifyProfile(u models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(DBname)
	col := db.Collection("users")

	profile := make(map[string]interface{})
	if len(u.Name) > 0 {
		profile["name"] = u.Name
	}
	if len(u.LastName) > 0 {
		profile["lastName"] = u.LastName
	}
	if !u.DateOfBirth.IsZero() {
		profile["dateOfBirth"] = u.DateOfBirth
	}
	if len(u.Avatar) > 0 {
		profile["avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		profile["banner"] = u.Banner
	}
	if len(u.Biography) > 0 {
		profile["biography"] = u.Biography
	}
	if len(u.Location) > 0 {
		profile["location"] = u.Location
	}
	if len(u.WebSite) > 0 {
		profile["webSite"] = u.WebSite
	}

	update := bson.M{
		"$set": profile,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filter, update)
	if err != nil {
		return false, err
	}

	return true, nil
}
