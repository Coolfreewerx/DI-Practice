package controller

import (
	"di-practice/service"
	m "di-practice/model"
	s "di-practice/service"

	"net/http"

	"github.com/labstack/echo/v4"
)

type PostController struct {
	postService s.PostService
	newPost     s.PostServiceDBImpl
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
func (c *PostController) GetPostsHandler(context echo.Context) error {
	posts, err := c.postService.GetPosts()
	if err != nil {
		return context.String(http.StatusInternalServerError, "internal server error")
	}
	return context.JSON(http.StatusOK, posts)
}

// Create a post to database.
func (c *PostController) CreatePostHandler(context echo.Context) error {
    post := new(m.Post)
    if err := context.Bind(post); err != nil {
        return context.String(http.StatusBadRequest, "invalid request body")
    }

    createdPost, err := c.newPost.SavePost(post)
    if err != nil {
        return context.String(http.StatusInternalServerError, "internal server error")
    }

    return context.JSON(http.StatusCreated, createdPost)
}