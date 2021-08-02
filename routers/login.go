package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/tapiaw38/concejapi/db"
	"github.com/tapiaw38/concejapi/jwt"
	"github.com/tapiaw38/concejapi/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "The email or password are invalid "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "The user email is required", 400)
		return
	}
	document, exist := db.LoginAttempt(t.Email, t.Password)

	if exist == false {
		http.Error(w, "The email or password are invalid", 400)
		return
	}

	jwtKey, err := jwt.GenerateJWT(document)

	if err != nil {
		http.Error(w, "An error occurred while generating the token"+err.Error(), 400)
		return
	}

	response := models.LoginResponse{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

	//coquie
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
