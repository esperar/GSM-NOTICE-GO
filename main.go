package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"goboard/handler"
	"log"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()

	e.POST("/api/signup", handler.SignUp)
	e.POST("/api/signin", handler.SignIn)
	e.Logger.Fatal(e.Start(":1323"))
}
