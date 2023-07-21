package main

import (
	"log"
	"net/http"
	"os"

	_handler "api/posts/external"
	_repo "api/posts/repository"
	_service "api/posts/service"
	"api/posts/progresql"

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

	// init struct to get dataset
	web, err := handler.GetAPI("JSON") // "DB" api for Database, "JSON" api for jsonplaceholder
	if err != nil {
		log.Fatal(err)
	}
	
	e.GET("/api/posts", handler.ShowAllPost(web)) 			// Dependency Injection
	// e.GET("/api/posts/db", handler.ShowAllPost(handler)) // Dependency Injection

	e.POST("/api/posts", handler.HandlerPosts) // METHOD POST body {"type" : "DB"} or {"type" : "JSON"}

	e.Start(":" + os.Getenv("PORT"))
}
