package service

import (
	"context"

	mPost "api/posts/model/post"
	
)

type TransactionService interface {
	ShowAllPost(ctx context.Context) ([]mPost.Post, error)
}