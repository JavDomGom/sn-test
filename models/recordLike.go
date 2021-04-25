package models

/* RecordLike model. */
type RecordLike struct {
	UserID    string   `bson:"userId" json:"userId"`
	LikesList []string `bson:"likesList" json:"likesList"`
}
