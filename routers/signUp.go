package routers

import (
	"encoding/json"
	"net/http"

	"github.com/JavDomGom/sn-test/db"
	"github.com/JavDomGom/sn-test/models"
)

/* SignUp Function to register a user in database.*/
func SignUp(w http.ResponseWriter, r *http.Request) {

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error in the received data. "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email is a required data.", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "You must specify a password of at least 6 characters. ", 400)
		return
	}

	_, found, _ := db.CheckIfUserAlreadyExists(t.Email)
	if found {
		http.Error(w, "There is already a registered user with this email.", 400)
		return
	}

	_, status, err := db.InsertUser(t)
	if err != nil {
		http.Error(w, "An error occurred while trying to register the user: "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "User could not be saved in database: ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
