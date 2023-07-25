package main

import (
	"log"
	"os"

	c "di-practice/controller"
	s "di-practice/service"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {

    db := &s.PostServiceDBImpl{}
	web := &s.PostServiceWebImpl{}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file", err)
	}

	e := echo.New()

	// Check dependency injection from json body request.
	e.GET("/check-posts", c.NewPostControllerWithOutService().HandleDI)

	// Change to database or web to check dependency injection.
	e.GET("/posts", c.NewPostController(web).GetPostsHandler)
	
	e.POST("/create-post", c.NewPostController(db).CreatePostHandler)

	e.Start(":" + os.Getenv("PORT"))
}
