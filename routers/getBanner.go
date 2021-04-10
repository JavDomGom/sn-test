package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/JavDomGom/sn-test/db"
)

/* GetBanner sends banner to HTTP. */
func GetBanner(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "You must send the id parameter.", http.StatusBadRequest)
		return
	}

	profile, err := db.GetProfile(ID)
	if err != nil {
		http.Error(w, "User not found.", http.StatusBadRequest)
		return
	}

	OpenFile, err := os.Open("uploads/banners/" + profile.Banner)
	if err != nil {
		http.Error(w, "Image not found.", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "Error copying image.", http.StatusBadRequest)
	}
}
