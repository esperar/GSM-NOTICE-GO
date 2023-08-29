package main

import (
	"github.com/joho/godotenv"
	_ "github.com/labstack/echo/v4/middleware"
	"goboard/router"
	"log"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := router.Router()

	e.Logger.Fatal(e.Start(":1323"))
}
