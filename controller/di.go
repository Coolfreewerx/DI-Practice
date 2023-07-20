package controller

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

func GetPostsHandler(api API) echo.HandlerFunc {


	fmt.Println("enter to GetPostsHandler")
	fmt.Println("api: ", api)

	handler :=  func(c echo.Context) error {
		posts, err := api.GetPosts()
		fmt.Println(posts)
		if err != nil {
			log.Println(err)
			return c.String(http.StatusInternalServerError, "เกิดข้อผิดพลาด")
		}

		return c.JSON(http.StatusOK, posts)
	}

	return handler
}