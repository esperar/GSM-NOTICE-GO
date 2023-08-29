package test

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func MockData() echo.HandlerFunc {
	return func(c echo.Context) error {
		list := map[string]string{
			"1": "희망이",
			"2": "주홍이",
			"3": "운린이",
		}
		return c.JSON(http.StatusOK, list)
	}
}
