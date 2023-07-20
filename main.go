package main

import (
	"log"
	"os"

	"di-practice/controller"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {

/* 	db := &controller.DB{}
	web := &controller.WebAPI{} */

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file", err)
	}

	e := echo.New()

	e.POST("/check-di", controller.HandleDI())

	e.Start(":" + os.Getenv("PORT"))

/* 	e.GET("/reply", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/posts", controller.GetPostsHandler(db)) */
}


