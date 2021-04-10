package routers

import (
	"encoding/json"
	"net/http"

	"github.com/JavDomGom/sn-test/db"
)

/* Profile extract values from a profile. */
func Profile(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "You must send the ID parameter", http.StatusBadRequest)
		return
	}

	profile, err := db.GetProfile(ID)
	if err != nil {
		http.Error(w, "An error occurred while trying to find the record. "+err.Error(), 400)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}
