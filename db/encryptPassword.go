package db

import "golang.org/x/crypto/bcrypt"

/* EncryptPassword Function to encrypt the user's password. */
func EncryptPassword(pass string) (string, error) {
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	return string(bytes), err
}
