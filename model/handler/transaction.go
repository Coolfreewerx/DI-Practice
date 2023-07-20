package handler

import (
	"github.com/labstack/echo/v4"
	mPost "api/posts/model/post"
)

type Handler interface {
	ShowAllPostHandler() echo.HandlerFunc
	GetAllPost() ([]mPost.Post, error)
	HandlerPosts(echo.Context) (error)
}