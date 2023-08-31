package helper

import (
	"errors"
	"fmt"
	"goboard/database"
	"goboard/models"
	"gorm.io/gorm"
	"net/http"
)

func GetCurrentUserId(w http.ResponseWriter, r *http.Request) (int, error) {
	cookie, err := r.Cookie("access-token")

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return 0, errors.New("authentication failed")
	}

	token := cookie.Value

	email, err := VerifyToken(token)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return 0, errors.New("authentication failed")
	}

	db := database.Connect()
	user := models.User{}
	result := db.Where("email = ?", email).First(user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return 0, fmt.Errorf("user not found")
		}
		return 0, result.Error
	}

	return user.Id, nil
}
