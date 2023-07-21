package handler

import (
	"github.com/labstack/echo/v4"
	mPost "api/posts/model/post"
	mApi "api/posts/model/api"
)

type Handler interface {
	ShowAllPost(api mApi.API) echo.HandlerFunc
	GetAllPost() ([]mPost.Post, error)
	HandlerPosts(echo.Context) (error)
	GetAPI(string) (mApi.API, error)
}