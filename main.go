package main

import (
	"log"
	"os"

	c "di-practice/controller"
	s "di-practice/service"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/lib/pq"
)

// @title Echo Swagger For DI Practice API
// @version 1.0
// @description This is a server for dependency injection practice.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:1150
// @BasePath /
// @schemes http

//go:generate swag init -g docs\docs.go

func main() {

    db := &s.PostServiceDBImpl{}
	web := &s.PostServiceWebImpl{}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file", err)
	}

	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// *FOR TEST*

	// *TEST In New Branch*

	// Check dependency injection from json body request.
	e.GET("/check-posts", c.NewPostControllerWithOutService().HandleDI)

	// Change to database or web to check dependency injection.
	e.GET("api/posts", c.NewPostController(web).GetPostsHandler)
	
	// Create post to database from json body request.
	e.POST("api/posts/create", c.NewPostController(db).CreatePostHandler)

	e.Start(":" + os.Getenv("PORT"))
}
