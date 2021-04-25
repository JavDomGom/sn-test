package routers

import (
	"net/http"

	"github.com/JavDomGom/sn-test/db"
	"github.com/JavDomGom/sn-test/models"
)

/* RemoveFollow remove a specific follow. */
func RemoveFollow(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "You must send the id parameter.", http.StatusBadRequest)
		return
	}

	var follow models.Follow

	follow.UserID = IDUser
	follow.UserFollowedID = ID

	status, err := db.RemoveFollow(follow)
	if err != nil {
		http.Error(w, "Failed to remove follow from database! "+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "Follow couldn't removed: "+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
