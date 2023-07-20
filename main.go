package main

import (
	"log"
	"net/http"
	"os"

	"di-practice/controller"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)





func HandleDI() echo.HandlerFunc {

	db := &controller.DB{}
	web := &controller.WebAPI{}

	handler :=  func(c echo.Context) error {
		input := controller.DatasourceInput{}

		if err := c.Bind(&input); err != nil {
			c.JSON(http.StatusBadRequest, "invalid JSON format")
		}

		if input.Workspace == "db" {
			controller.GetPostsHandler(db)
			
		} else if input.Workspace == "web" {
			controller.GetPostsHandler(web)
			
		} else {
			c.JSON(http.StatusBadRequest, "invalid workspace")
		}

	return c.String(http.StatusOK, "Hello, DI!")
	}

	return handler


}

func main() {

	db := &controller.DB{}
	// web := &controller.WebAPI{}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file", err)
	}

	e := echo.New()

	e.GET("/reply", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/posts", controller.GetPostsHandler(db))

	e.POST("/check-di", HandleDI())

	e.GET("/reply-di", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, DI!")
	})

	e.Start(":" + os.Getenv("PORT"))


}


