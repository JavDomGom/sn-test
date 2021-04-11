package routers

import (
	"net/http"

	"github.com/JavDomGom/sn-test/db"
	"github.com/JavDomGom/sn-test/models"
)

/* RemoveFollow remove a specific relation. */
func RemoveFollow(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var t models.Follow
	t.UserID = IDUser
	t.UserFollowedID = ID

	status, err := db.RemoveFollow(t)
	if err != nil {
		http.Error(w, "Failed to remove relation from database! "+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "Relation couldn't removed: "+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
