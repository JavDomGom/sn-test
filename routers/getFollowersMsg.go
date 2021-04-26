package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/JavDomGom/sn-test/db"
)

/* GetFollowersMsg returns all messages from my followers. */
func GetFollowersMsg(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "You must send the page parameter.", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "You must send the page parameter as integer greater than zero.", http.StatusBadRequest)
		return
	}

	response, correct := db.GetFollowersMsg(IDUser, page)
	if !correct {
		http.Error(w, "Error reading messages.", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
