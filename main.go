package main

import (
	// "context"
	"log"
	"net/http"
	"os"

	// "api/posts/ent"
	_handler "api/posts/external"
	"api/posts/progresql"
	_repo "api/posts/repository"
	_service "api/posts/service"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	// init Database
	clientDB, err := progresql.InitDatabase()
	if err != nil {
		log.Fatal(err)
		return
	}

	defer progresql.CloseDatabase(clientDB)
	

	// init Repository
	userRepo := _repo.NewUserRepository(clientDB)
	postRepo := _repo.NewPostRepository(clientDB)

	// init Service 
	service := _service.NewPostServie(postRepo, userRepo)

	// init External
	handler := _handler.NewHandler(service)

	e := echo.New()

	e.GET("/home", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello Home")
	})

	web := _handler.WebJson{}
	

	e.GET("/api/posts/test", handler.ShowAllPostHandler())
	e.GET("/api/posts/json", _handler.ShowAllPost(web))
	e.GET("/api/posts/db", _handler.ShowAllPost(handler))
	e.POST("/api/posts", handler.HandlerPosts)


	e.Start(":" + os.Getenv("PORT"))
}
