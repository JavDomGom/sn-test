package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/JavDomGom/sn-test/db"
)

/* GetUsers returns a list with all users. */
func GetUsers(w http.ResponseWriter, r *http.Request) {
	userType := r.URL.Query().Get("userType")
	if len(userType) < 1 {
		http.Error(w, "You must send the userType parameter.", http.StatusBadRequest)
		return
	}

	page := r.URL.Query().Get("page")
	if len(page) < 1 {
		http.Error(w, "You must send the page parameter.", http.StatusBadRequest)
		return
	}

	// This parameter is optional.
	search := r.URL.Query().Get("search")

	pagTmp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "You must send the type parameter as integer greater than zero.", http.StatusBadRequest)
		return
	}

	pag := int64(pagTmp)

	result, status := db.GetAllUsers(IDUser, pag, search, userType)
	if status == false {
		http.Error(w, "Error reading users.", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
