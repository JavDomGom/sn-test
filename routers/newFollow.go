package routers

import (
	"net/http"

	"github.com/JavDomGom/sn-test/db"
	"github.com/JavDomGom/sn-test/models"
)

/* NewFollow register a new follow between two users. */
func NewFollow(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "You must send the id parameter.", http.StatusBadRequest)
		return
	}

	if IDUser == ID {
		http.Error(w, "You cannot follow yourself!", http.StatusBadRequest)
		return
	}

	var follow models.Follow
	follow.UserID = IDUser
	follow.UserFollowedID = ID

	status, err := db.InsertFollow(follow)
	if err != nil {
		http.Error(w, "Failed to save follow to database! "+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "Follow couldn't be saved: "+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
