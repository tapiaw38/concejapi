package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/tapiaw38/concejapi/models"
)

//GenerateJWT generated a token
func GenerateJWT(t models.User) (string, error) {

	myKey := []byte("bXlTZWNyZXRQYXNzd29yZG15U2VjcmV0UGFzc3dvcmQK")

	payload := jwt.MapClaims{
		"_id":       t.Id.Hex(),
		"FirstName": t.FirstName,
		"LastName":  t.LastName,
		"email":     t.Email,
		"DateBirth": t.DateBirth,
		"Banner":    t.Banner,
		"Avatar":    t.Avatar,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
