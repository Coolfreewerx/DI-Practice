package controller

import (
	"di-practice/service"
	m "di-practice/model"
	s "di-practice/service"

	"net/http"

	"github.com/labstack/echo/v4"
)

type PostController struct {
	postService service.PostService
}

func NewPostController(postService service.PostService) *PostController {
    return &PostController{
        postService: postService,
    }
}

func NewPostControllerWithOutService() *PostController {
    return &PostController{}
}

// HandleDI check dependency injection from json body request.
func (c *PostController) HandleDI(context echo.Context) error {

	input 	:= 	m.DatasourceInput{}
	db  	:=  &s.PostServiceDBImpl{}
	web 	:= 	&s.PostServiceWebImpl{}

	if err := context.Bind(&input); err != nil {
		context.JSON(http.StatusBadRequest, "invalid JSON format")
	}

	if input.Workspace == "DB" {
		c.postService = db
		c.GetPostsHandler(context)
			
	} else if input.Workspace == "Web" {
		c.postService = web
		c.GetPostsHandler(context)
			
	} else {
		return context.JSON(http.StatusBadRequest, "invalid workspace")
	}

	return nil
}

// GetPostsHandler get posts from database or web.
// @Summary Show a posts.
// @Description get a posts from database or web.
// @Tags Posts
// @Accept */*
// @Produce json
// @Success 200 {interface} interface{} "Success response"
// @Router /posts [get]
func (c *PostController) GetPostsHandler(context echo.Context) error {
	posts, err := c.postService.GetPosts()
	if err != nil {
		return context.String(http.StatusInternalServerError, "internal server error")
	}
	return context.JSON(http.StatusOK, posts)
}
