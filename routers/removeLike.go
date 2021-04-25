package routers

import (
	"net/http"

	"github.com/JavDomGom/sn-test/db"
)

/* RemoveLike remove a specific like. */
func RemoveLike(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "You must send the id parameter.", http.StatusBadRequest)
		return
	}

	status, err := db.RemoveLike(IDUser, ID)
	if err != nil {
		http.Error(w, "Failed to remove like from database! "+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "Like couldn't removed: "+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
