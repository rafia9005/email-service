package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/rafia9005/email-service/internal/handler"
)

func main() {
	e := echo.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Emails service is running!")
	})

	e.POST("/emails", handler.SendEmail)

	e.Logger.Fatal(e.Start(":8080"))
}
