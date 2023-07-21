package api

import (
	mPost "api/posts/model/post"
)

type API interface {
	GetAllPost() ([]mPost.Post, error)
}