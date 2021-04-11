package db

import (
	"context"
	"time"

	"github.com/JavDomGom/sn-test/models"
)

/* RemoveFollow remove a specific relation. */
func RemoveFollow(t models.Follow) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(DBname)
	col := db.Collection("follows")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
