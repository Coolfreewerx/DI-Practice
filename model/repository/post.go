package repository

import (
	"context"

	mPost "api/posts/model/post"
)

type PostRepository interface {
	GetAllPost(ctx context.Context) ([]mPost.Post, error)
}
