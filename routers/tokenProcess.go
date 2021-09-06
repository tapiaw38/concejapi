package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/tapiaw38/concejapi/db"
	"github.com/tapiaw38/concejapi/models"
)

/*VARIABLES RETURNED BY THE MODEL
THAT WILL BE USED IN ALL END POINTS*/

var Email string
var UserID string

//TokenProcess
func TokenProcess(tk string) (*models.Claim, bool, string, error) {
	myKey := []byte("bXlTZWNyZXRQYXNzd29yZG15U2VjcmV0UGFzc3dvcmQK")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("token format invalid")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})

	if err == nil {
		_, user, _ := db.CheackingUser(claims.Email)

		if user == true {
			Email = claims.Email
			UserID = claims.Id.Hex()
		}

		return claims, user, UserID, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("Token invalid")
	}

	return claims, false, string(""), err
}
