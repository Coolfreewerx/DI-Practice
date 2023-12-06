package service

import (
	"context"
	"di-practice/ent"
	"encoding/json"

	"log"
	"net/http"
	"os"

	m "di-practice/model"
)

type PostService interface {
	GetPosts() ([]m.Post, error)
}

type PostServiceDBImpl struct {}

// GetPosts from database.
// GetPosts from database.
func (s *PostServiceDBImpl) GetPosts() ([]m.Post, error) {

	client, err := ent.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	ctx := context.Background()

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	
	if err := createDataSet(client, ctx) ; err != nil {
		log.Fatal(err)
	}

	db_posts, err := client.User.Query().All(ctx)
	if err != nil {
		log.Fatalf("failed query posts: %v", err)
	}


	posts := []m.Post{}
	for _, post := range db_posts {
	 	posts = append(posts, m.Post{
	 		UserId: post.UserId,
	 		ID: post.ID,
	 		Title: post.Title,
	 		Body: post.Body,
	 	})
	}

	return posts, nil
}

type PostServiceWebImpl struct {}

// GetPosts from web.
// GetPosts from web.
func (s *PostServiceWebImpl) GetPosts() ([]m.Post, error) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var posts []m.Post
	err = json.NewDecoder(resp.Body).Decode(&posts)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

// Generate Dataset.
func createDataSet(client *ent.Client, ctx context.Context) error {
	if (client.User.Query().CountX(ctx) != 0) {
		return nil
	}

	numOfUser, numOfPost := 2, 3

	for u := 1 ; u <= numOfUser ; u++ {
		for p := 1 ; p <= numOfPost ; p++ {
			err := client.User.Create().
				SetUserId(u).
				SetBody("est rerum tempore vitae\nsequi sint nihil reprehenderit dolor beatae ea dolores neque\nfugiat blanditiis voluptate porro vel nihil molestiae ut reiciendis\nqui aperiam non debitis possimus qui neque nisi nulla").
				SetTitle("sunt aut facere repellat provident occaecati excepturi optio reprehenderit").
				Exec(ctx)
			if err != nil {
				log.Fatal(err)
				return err
			}
		}
	}

	return nil
}

// SavePost to database.
func (s *PostServiceDBImpl) SavePost(post *m.Post) (*m.Post, error) {
	client, err := ent.Open("postgres", os.Getenv("POSTGRES_URL"))
 	if err != nil {
 		log.Fatalf("failed opening connection to postgres: %v", err)
 	}
 	defer client.Close()
 	ctx := context.Background()

 	db_post, err := client.User.Create().
 		SetUserId(post.UserId).
 		SetTitle(post.Title).
 		SetBody(post.Body).
 		Save(ctx)
 	if err != nil {
 		log.Fatalf("failed creating post: %v", err)
 	}

 	return &m.Post{
 		UserId: db_post.UserId,
		ID: db_post.ID,
 		Title: db_post.Title,
 		Body: db_post.Body,
 	}, nil
}
