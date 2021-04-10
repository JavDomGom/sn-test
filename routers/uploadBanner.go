package routers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/JavDomGom/sn-test/db"
	"github.com/JavDomGom/sn-test/models"
)

/* UploadBanner upload banner to server. */
func UploadBanner(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("banner")
	var extension = filepath.Ext(handler.Filename)
	var filename string = "uploads/banners/" + IDUser + extension

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Failed to upload image! "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Failed to save banner to database! "+err.Error(), http.StatusBadRequest)
		return
	}

	var username models.User
	var status bool

	username.Banner = IDUser + extension
	status, err = db.ModifyProfile(username, IDUser)
	if err != nil || status == false {
		http.Error(w, "Failed to save banner to database! "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
