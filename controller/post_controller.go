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

func (c *PostController) HandleDI (context echo.Context) error {

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

func (c *PostController) GetPostsHandler(context echo.Context) error {
	posts, err := c.postService.GetPosts()
	if err != nil {
		return context.String(http.StatusInternalServerError, "internal server error")
	}
	return context.JSON(http.StatusOK, posts)
}
