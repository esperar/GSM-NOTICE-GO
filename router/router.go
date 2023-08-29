package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"goboard/handler"
	"goboard/test"
	"net/http"
	"os"
)

func Router() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/healthy", func(c echo.Context) error {
		return c.String(http.StatusOK, "Health Check OK!!")
	})

	e.POST("/api/signup", handler.SignUp)
	e.POST("/api/signin", handler.SignIn)
	e.GET("/api/getlist", test.MockData(), middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte(os.Getenv("SECRET_KEY")),
		TokenLookup: "cookie:access-token",
	}))

	return e
}
