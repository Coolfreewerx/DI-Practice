package repository

import (
	"api/posts/ent"
	"context"
	// "log"

	mPost "api/posts/model/post"
)

type postRepository struct {
	clientDB *ent.Client
}

func NewPostRepository(clientDB *ent.Client) postRepository{
	return postRepository{
		clientDB: clientDB,
	}
}

func (repo postRepository) GetAllPost(ctx context.Context) ([]mPost.Post, error) {
	posts := []mPost.Post{}
	postRepo, err := repo.clientDB.Post.Query().All(ctx)
	if err != nil {
		return posts, err
	}

	for _, post := range postRepo {
		posts = append(posts, mPost.Post{
				ID: post.ID,
				UserId: post.UserId,
				Title: post.Title,
				Body: post.Body,
		})
		
	}
	// log.Println(posts)
	return posts, nil
}