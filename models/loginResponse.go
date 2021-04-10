package models

/* LoginResponse Token that is returned with the login. */
type LoginResponse struct {
	Token string `json:"token,omitempty"`
}
