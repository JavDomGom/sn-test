package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/JavDomGom/sn-test/db"
	"github.com/JavDomGom/sn-test/jwt"
	"github.com/JavDomGom/sn-test/models"
)

/* Login Do login. */
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Invalid username or password: "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "Invalid email address.", 400)
		return
	}
	document, exists := db.LoginAttemp(t.Email, t.Password)
	if !exists {
		http.Error(w, "Invalid username or password.", 400)
		return
	}

	jwtKey, err := jwt.GenerateJWT(document)
	if err != nil {
		http.Error(w, "An error occurred while trying to generate the token: "+err.Error(), 400)
		return
	}

	resp := models.LoginResponse{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
