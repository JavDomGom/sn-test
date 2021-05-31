package db

import (
	"context"
	"time"

	"github.com/JavDomGom/sn-test/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* GetAllUsers Get all users who have a relation with me. */
func GetAllUsers(ID string, page int64, search string, userType string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(DBname)
	col := db.Collection("users")

	var results []*models.User

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"$or": []bson.M{
			{"name": bson.M{"$regex": `(?i)` + search}},
			{"lastName": bson.M{"$regex": `(?i)` + search}},
		},
	}

	cur, err := col.Find(ctx, query, findOptions)
	if err != nil {
		return results, false
	}

	var found, include bool

	for cur.Next(ctx) {
		var s models.User
		err := cur.Decode(&s)
		if err != nil {
			return results, false
		}

		var r models.Follow
		r.UserID = ID
		r.UserFollowedID = s.ID.Hex()

		include = false

		found, _ = CheckFollow(r)
		if userType == "noFollow" && !found {
			include = true
		}
		if userType == "follow" && found {
			include = true
		}

		if r.UserFollowedID == ID {
			include = false
		}

		if include {
			s.Password = ""
			s.Biography = ""
			s.WebSite = ""
			s.Location = ""
			s.Banner = ""
			s.Email = ""

			results = append(results, &s)
		}
	}

	err = cur.Err()
	if err != nil {
		return results, false
	}
	cur.Close(ctx)
	return results, true
}
