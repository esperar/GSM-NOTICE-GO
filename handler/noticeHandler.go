package handler

import (
	"github.com/labstack/echo/v4"
	"goboard/database"
	"goboard/helper"
	models "goboard/models"
	"net/http"
)

func CreateNotice(w http.ResponseWriter, r *http.Request, c echo.Context) error {

	notice := new(models.Notice)

	if err := c.Bind(notice); err != nil {
		return helper.SendToJson(http.StatusBadRequest, "bad Request", c)
	}

	currentUserId, err := helper.GetCurrentUserId(w, r)

	if err != nil {
		return helper.SendToJson(http.StatusInternalServerError, "user not found", c)
	}

	db := database.Connect()

	notice.CreatedBy = currentUserId

	if err := db.Create(&notice); err.Error != nil {
		return helper.SendToJson(http.StatusInternalServerError, "Failed Create Notice", c)
	}

	return helper.SendToJson(http.StatusCreated, "success", c)
}

func GetAllNotices(e echo.Context, w http.ResponseWriter) error {

	db := database.Connect()

	var notices []models.Notice
	db.Find(&notices)

	return e.JSON(http.StatusOK, notices)
}
