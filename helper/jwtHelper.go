package helper

import (
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

func CreateJwt(Email string) (string, error) {
	mySigninKey := []byte(os.Getenv("SECRET_KEY"))

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
