package handler

import (
	"github.com/labstack/echo/v4"
	"goboard/database"
	"goboard/helper"
	"goboard/models"
	_ "golang.org/x/crypto/ssh"
	"net/http"
)

func SignUp(c echo.Context) error {
	user := new(models.User)

	if err := c.Bind(user); err != nil {
		return sendJson(http.StatusBadRequest, "bad Request", c)
	}

	db := database.Connect()
	result := db.Find(&user, "email=?", user.Email)

	if result.RowsAffected != 0 {
		return sendJson(http.StatusBadRequest, "existing email", c)
	}

	hashPassword, err := helper.HashPassword(user.Password)
	if err != nil {
		return sendJson(http.StatusInternalServerError, err.Error(), c)
	}

	user.Password = hashPassword

	if err := db.Create(&user); err.Error != nil {
		return sendJson(http.StatusInternalServerError, "Failed SignUp", c)
	}

	return sendJson(http.StatusOK, "success", c)

}

func SignIn(c echo.Context) error {
	user := new(models.User)

	if err := c.Bind(user); err != nil {
		return sendJson(http.StatusBadRequest, "bad request", c)
	}

	password := user.Password

	db := database.Connect()
	result := db.Find(user, "email=?", user.Email)

	if result.RowsAffected == 0 {
		return echo.ErrBadRequest
	}

	response := helper.CheckPasswordHash(user.Password, password)

	var message string
	if !response {
		return echo.ErrUnauthorized
	} else {
		message = "Success"
	}

	return sendJson(http.StatusOK, message, c)
}

func sendJson(status int, message string, c echo.Context) error {
	return c.JSON(status, map[string]string{
		"message": message,
	})
}
