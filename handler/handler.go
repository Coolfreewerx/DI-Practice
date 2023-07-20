package handler

import (
	"context"
	"di-practice/ent"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type API interface {
	GetPosts() ([]Post, error)
}

type Post struct {
	UserId 	int    		`json:"userId"`
	ID    	int    		`json:"id"`
	Title 	string 		`json:"title"`
	Body  	string 		`json:"body"`
}

type DB struct{}

func (db *DB) GetPosts() ([]Post, error) {

	client, err := ent.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	ctx := context.Background()

	db_posts, err := client.User.Query().All(ctx)
	if err != nil {
		log.Fatalf("failed query posts: %v", err)
	}

	
	posts := []Post{}
	for _, post := range db_posts {
	 	posts = append(posts, Post{
	 		UserId: post.UserId,
	 		ID: post.ID,
	 		Title: post.Title,
	 		Body: post.Body,
	 	})
	}

	return posts, nil
}


type WebAPI struct{}

func (api *WebAPI) GetPosts() ([]Post, error) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var posts []Post
	err = json.NewDecoder(resp.Body).Decode(&posts)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

type Application struct {
	db *DB
	web_api *WebAPI
}
 
func NewApplication(db *DB, web_api *WebAPI) *Application {
	return &Application{
		db:  db,
		web_api: web_api,
	}
}

func (app *Application) TestGetPostsHandler() string {
	dbData, err := app.db.GetPosts()
	posts, err := app.web_api.GetPosts()
	// posts, err := app.db.GetPosts()
	if err != nil {
		fmt.Println(err)
	}

	result := fmt.Sprintf("ข้อมูลฐานข้อมูล: %s\n\nข้อมูลจาก API:\n", dbData)
	for _, post := range posts {
		result += fmt.Sprintf("ID: %d\nUserID: %d\nTitle: %s\nBody: %s\n\n", post.ID, post.UserId, post.Title, post.Body)
	}

	return result
}

func GetPostsHandler(api API) echo.HandlerFunc {
	return func(c echo.Context) error {
		posts, err := api.GetPosts()
		if err != nil {
			log.Println(err)
			return c.String(http.StatusInternalServerError, "เกิดข้อผิดพลาด")
		}

		return c.JSON(http.StatusOK, posts)
	}
}