package middlewares

import (
	"net/http"

	"github.com/tapiaw38/concejapi/routers"
)

//ValidateTooken middleware
func ValidateToken(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.TokenProcess(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error in the token"+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
