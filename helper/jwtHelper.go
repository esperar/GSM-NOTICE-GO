package helper

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

var mySigninKey = []byte(os.Getenv("SECRET_KEY"))

func CreateJwt(Email string) (string, error) {

	aToken := jwt.New(jwt.SigningMethodHS256)
	claims := aToken.Claims.(jwt.MapClaims)
	claims["Email"] = Email
	claims["exp"] = time.Now().Add(time.Minute * 20).Unix()

	token, err := aToken.SignedString(mySigninKey)

	if err != nil {
		return "", err
	}

	return token, err
}

func VerifyToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return mySigninKey, nil
	})

	if err != nil {
		return 0, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, errors.New("invalid token")
	}

	email := claims["email"].(string)

	return email, nil
}
