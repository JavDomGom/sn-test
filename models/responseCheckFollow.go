package models

/* ResponseCheckFollow returns true or false after consulting the follow between two users. */
type ResponseCheckFollow struct {
	Status bool `json:"status"`
}
