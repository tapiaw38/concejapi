package routers

import (
	"encoding/json"
	"net/http"

	"github.com/tapiaw38/concejapi/db"
	"github.com/tapiaw38/concejapi/models"
)

//func ModifiedProfile
func ModifiedProfile(w http.ResponseWriter, r *http.Request) {

	var user models.User
	var status bool

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "data is failed"+err.Error(), 400)
		return
	}

	status, err = db.ModifiedRegister(user, UserID)

	if err != nil {
		http.Error(w, "Error register no modified"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "data is undefined"+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
