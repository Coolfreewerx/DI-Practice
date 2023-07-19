package main

import (
	"log"
	"net/http"
	"os"


	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {

	db := &DB{}
	// api := &WebAPI{}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file", err)
	}

	e := echo.New()


	e.GET("/reply", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/posts", GetPostsHandler(db))


	e.Start(":" + os.Getenv("PORT"))
}

