package models

/* Follow model. */
type Follow struct {
	UserID         string `bson:"userId" json:"userId"`
	UserFollowedID string `bson:"userFollowedID" json:"userFollowedID"`
}
