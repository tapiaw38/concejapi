package routers

import (
	"encoding/json"
	"net/http"

	"github.com/tapiaw38/concejapi/db"
)

//ViewProfile
func ViewProfile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "you must send the id parameter", http.StatusBadRequest)
		return
	}

	profile, err := db.SearchProfile(ID)

	if err != nil {
		http.Error(w, "An error occurred while trying to find the record"+err.Error(), 400)

	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}
