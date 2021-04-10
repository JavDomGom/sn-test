package middleware

import (
	"net/http"

	"github.com/JavDomGom/sn-test/db"
)

/* CheckDB middleware to check database status.*/
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !db.CheckConnection() {
			http.Error(w, "Lost connection to the database.", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
