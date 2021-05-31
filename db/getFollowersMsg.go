package db

import (
	"context"
	"time"

	"github.com/JavDomGom/sn-test/models"
	"go.mongodb.org/mongo-driver/bson"
)

/* GetFollowersMsg Gets messages from my followers. */
func GetFollowersMsg(ID string, page int) ([]models.ReturnFollowersMsg, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(DBname)
	col := db.Collection("follows")

	skip := (page - 1) * 20

	filter := make([]bson.M, 0)
	filter = append(filter, bson.M{"$match": bson.M{"userId": ID}})
	filter = append(filter, bson.M{
		"$lookup": bson.M{
			"from":         "messages",
			"localField":   "userFollowedID",
			"foreignField": "userId",
			"as":           "msg",
		}})
	filter = append(filter, bson.M{"$unwind": "$msg"})
	filter = append(filter, bson.M{"$sort": bson.M{"msg.datetime": -1}})
	filter = append(filter, bson.M{"$skip": skip})
	filter = append(filter, bson.M{"$limit": 20})

	cursor, _ := col.Aggregate(ctx, filter)
	var result []models.ReturnFollowersMsg
	err := cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}
	return result, true
}
