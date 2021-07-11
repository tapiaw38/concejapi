package routers

import (
	"encoding/json"
	"net/http"

	"github.com/tapiaw38/concejapi/db"
	"github.com/tapiaw38/concejapi/models"
)

// Register is a function that registers
func Register(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error in received data."+err.Error(), 400)

		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email is required.", 400)

		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "The password must have more than 6 digits.", 400)

		return
	}

	_, found, _ := db.CheackingUser(t.Email)
	if found == true {
		http.Error(w, "There is already a registered user with this email.", 400)

		return
	}

	_, status, err := db.InsertRegister(t)
	if err != nil {
		http.Error(w, "An error occurred while registering the user."+err.Error(), 400)

		return
	}

	if status == false {
		http.Error(w, "An error occurred while registering the user."+err.Error(), 400)

		return
	}

	w.WriteHeader(http.StatusCreated)

}
