package jwt

import (
	"time"

	"github.com/JavDomGom/sn-test/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/* GenerateJWT Encrypt with JWT. */
func GenerateJWT(t models.User) (string, error) {

	privateKey := []byte("aaaaaaaaaaaaaaaaa")

	payload := jwt.MapClaims{
		"email":       t.Email,
		"name":        t.Name,
		"lastName":    t.LastName,
		"dateOfBirth": t.DateOfBirth,
		"biography":   t.Biography,
		"location":    t.Location,
		"webSite":     t.WebSite,
		"_id":         t.ID.Hex(),
		"exp":         time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return tokenString, err
	}
	return tokenString, nil
}
