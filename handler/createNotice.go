package handler

import (
	"github.com/labstack/echo/v4"
	"goboard/database"
	"goboard/helper"
	"goboard/models"
	"net/http"
)

func CreateNotice(c echo.Context) error {

	notice := new(models.Notice)

	if err := c.Bind(notice); err != nil {
		return helper.SendToJson(http.StatusBadRequest, "bad Request", c)
	}

	db := database.Connect()

	if err := db.Create(&notice); err.Error != nil {
		return helper.SendToJson(http.StatusInternalServerError, "Failed Create Notice", c)
	}

	return helper.SendToJson(http.StatusCreated, "success", c)
}
