package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/labstack/echo/v4/middleware"
	"goboard/handler"
	"goboard/test"
	"log"
	"os"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()

	e.POST("/api/signup", handler.SignUp)
	e.POST("/api/signin", handler.SignIn)
	e.GET("/api/getlist", test.MockData(), middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(os.Getenv("SECRET_KEY")),
		TokenLookup: "cookie:access-token"
	}))
	e.Logger.Fatal(e.Start(":1323"))
}
