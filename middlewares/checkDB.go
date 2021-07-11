package middlewares

import (
	"net/http"

	"github.com/tapiaw38/concejapi/db"
)

//CheckDB is a middlewares with chech db
func CheckDB(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == false {
			http.Error(w, "Conection db failed", 500)

			return
		}
		next.ServeHTTP(w, r)
	}
}
