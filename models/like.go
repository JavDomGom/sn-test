package models

/* Like model. */
type Like struct {
	UserID    string   `bson:"userId" json:"userId"`
	LikesList []string `bson:"likesList" json:"likesList"`
}
