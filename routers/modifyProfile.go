package routers

import (
	"encoding/json"
	"net/http"

	"github.com/JavDomGom/sn-test/db"
	"github.com/JavDomGom/sn-test/models"
)

/* ModifyProfile modify a user's profile. */
func ModifyProfile(w http.ResponseWriter, r *http.Request) {

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Incorrect data. "+err.Error(), 400)
		return
	}

	var status bool

	status, err = db.ModifyProfile(t, IDUser)
	if err != nil {
		http.Error(w, "An error occurred while trying to modify the register. "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "The user's profile has not been modified. ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
