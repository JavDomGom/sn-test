package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/JavDomGom/sn-test/db"
	"github.com/JavDomGom/sn-test/models"
)

/* RecordMsg record a message in database. */
func RecordMsg(w http.ResponseWriter, r *http.Request) {
	var message models.Msg

	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, "Error in the received data. "+err.Error(), 400)
		return
	}

	msg := models.RecordMsg{
		UserID:             IDUser,
		Message:            message.Message,
		Datetime:           time.Now(),
		InReplyToMessageID: message.InReplyToMessageID,
	}

	_, status, err := db.InsertMsg(msg)
	if err != nil {
		http.Error(w, "An error occurred while trying to insert the register. "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "The message could not be inserted.", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
