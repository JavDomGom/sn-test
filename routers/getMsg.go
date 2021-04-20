package routers

import (
	"encoding/json"
	"net/http"

	"github.com/JavDomGom/sn-test/db"
)

/* GetMsg Get a message by id. */
func GetMsg(w http.ResponseWriter, r *http.Request) {
	IDMessage := r.URL.Query().Get("id")
	if len(IDMessage) < 1 {
		http.Error(w, "You must send the id parameter.", http.StatusBadRequest)
		return
	}

	response, correct := db.GetMsg(IDMessage)
	if !correct {
		http.Error(w, "Error getting message.", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
