package main

import (
	"log"
	"net/http"
	"os"

	"di-practice/handler"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {

	// db := &handler.DB{}
	web := &handler.WebAPI{}
	// app := handler.NewApplication(db, web)

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file", err)
	}

	e := echo.New()

	e.GET("/reply", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/posts", handler.GetPostsHandler(web))


	e.Start(":" + os.Getenv("PORT"))
}

