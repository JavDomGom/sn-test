package routers

import (
	"encoding/json"
	"net/http"

	"github.com/JavDomGom/sn-test/db"
)

/* GetLikes returns all likes by user id.*/
func GetLikes(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("userId")
	if len(ID) < 1 {
		http.Error(w, "You must send the id parameter.", http.StatusBadRequest)
		return
	}

	response, correct := db.GetLikes(ID)
	if !correct {
		http.Error(w, "Error getting likes.", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
