package helper

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(secretKey, email string) (string, error) {
	// generate new token using method HS256
	token := jwt.New(jwt.SigningMethodHS256)

	// get claim token
	claims := token.Claims.(jwt.MapClaims)

	// add claim email to token
	claims["email"] = email

	// add expired time
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	// sigining token with secret key
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	// get token signed
	return tokenString, nil
}