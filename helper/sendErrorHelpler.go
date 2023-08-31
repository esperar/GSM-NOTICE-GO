package helper

import "github.com/labstack/echo/v4"

func SendToJson(status int, message string, c echo.Context) error {
	return c.JSON(status, map[string]string{
		"message": message,
	})
}
