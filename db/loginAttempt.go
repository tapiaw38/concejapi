package db

import (
	"github.com/tapiaw38/concejapi/models"
	"golang.org/x/crypto/bcrypt"
)

//Login attempt
func LoginAttempt(email string, password string) (models.User, bool) {
	user, find, _ := CheackingUser(email)

	if find == false {
		return user, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)

	if err != nil {
		return user, false
	}

	return user, true

}
