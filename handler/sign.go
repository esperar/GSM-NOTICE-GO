package handler

import (
	"github.com/labstack/echo/v4"
	"goboard/database"
	helper "goboard/helper"
	"goboard/models"
	_ "golang.org/x/crypto/ssh"
	"net/http"
	"time"
)

func SignUp(c echo.Context) error {
	user := new(models.User)

	if err := c.Bind(user); err != nil {
		return helper.SendToJson(http.StatusBadRequest, "bad Request", c)
	}

	db := database.Connect()
	result := db.Find(&user, "email=?", user.Email)

	if result.RowsAffected != 0 {
		return helper.SendToJson(http.StatusBadRequest, "existing email", c)
	}

	hashPassword, err := helper.HashPassword(user.Password)
	if err != nil {
		return helper.SendToJson(http.StatusInternalServerError, err.Error(), c)
	}

	user.Password = hashPassword

	if err := db.Create(&user); err.Error != nil {
		return helper.SendToJson(http.StatusInternalServerError, "Failed SignUp", c)
	}

	return helper.SendToJson(http.StatusOK, "success", c)

}

func SignIn(c echo.Context) error {
	user := new(models.User)

	if err := c.Bind(user); err != nil {
		return helper.SendToJson(http.StatusBadRequest, "bad request", c)
	}

	password := user.Password

	db := database.Connect()
	result := db.Find(user, "email=?", user.Email)

	if result.RowsAffected == 0 {
		return echo.ErrBadRequest
	}

	response := helper.CheckPasswordHash(user.Password, password)

	if !response {
		return echo.ErrUnauthorized
	}

	accessToken, err := helper.CreateJwt(user.Email)
	if err != nil {
		return echo.ErrInternalServerError
	}

	cookie := new(http.Cookie)
	cookie.Name = "access-token"
	cookie.Value = accessToken
	cookie.HttpOnly = true
	cookie.Expires = time.Now().Add(time.Hour * 24)

	c.SetCookie(cookie)

	return helper.SendToJson(http.StatusOK, "Login Success! accessToken = "+accessToken, c)
}
