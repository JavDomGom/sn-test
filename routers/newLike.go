package routers

import (
	"net/http"

	"github.com/JavDomGom/sn-test/db"
)

/* NewLike register a new like. */
func NewLike(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "You must send the id parameter.", http.StatusBadRequest)
		return
	}

	status, err := db.InsertLike(IDUser, ID)
	if err != nil {
		http.Error(w, "Failed to save like to database! "+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "Like couldn't be saved: "+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
