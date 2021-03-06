package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/JavDomGom/sn-test/db"
)

/* ReadMsg read messages. */
func ReadMsg(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "You must send the id parameter.", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "You must send the page parameter.", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "You must send the page parameter with a value greater than zero.", http.StatusBadRequest)
		return
	}

	pag := int64(page)
	response, correct := db.GetMessages(ID, pag)
	if !correct {
		http.Error(w, "Error reading messages.", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
