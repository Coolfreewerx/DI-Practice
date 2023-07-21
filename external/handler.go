package external

import (
	mHandler "api/posts/model/handler"
	mPosts "api/posts/model/post"
	mService "api/posts/model/service"
	mApi "api/posts/model/api"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)


type WebJson struct{}

type handler struct {
	transactionService mService.TransactionService
}

func NewHandler(transactionService mService.TransactionService) mHandler.Handler {
	return handler{
		transactionService: transactionService,
	}
}

// API Method Interface
func (h handler) GetAllPost() ([]mPosts.Post, error) {
	posts, err := h.transactionService.ShowAllPost(context.Background())
	if err != nil {
		return posts, err
	}
	return posts, nil
}

func (web WebJson) GetAllPost() ([]mPosts.Post, error) {
	req, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	posts := []mPosts.Post{}
	err = json.NewDecoder(req.Body).Decode(&posts)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

// Function for METHOD POST
func (h handler)HandlerPosts(c echo.Context) error {
	mode := struct{
		Type string `json:"type"`
	}{}
	if err := c.Bind(&mode); err != nil {
		log.Fatal(err)
		return err
	}
	
	if mode.Type == "DB" {
		ShowPosts(h, c) // Dependency Injection
	} else if mode.Type == "JSON" {
		web := WebJson{}
		ShowPosts(web, c) // Dependency Injection
	} else {
		c.String(http.StatusBadRequest, `type is "DB" or "JSON"`)
	}
	return nil
}

// Method Dependency injection
func ShowPosts(api mApi.API, c echo.Context) {
	posts, err := api.GetAllPost()
		if err != nil {
			 c.String(http.StatusInternalServerError, "Error: " + err.Error())
		}
		
		 c.JSON(http.StatusOK, posts)
}

// Function return struce
func (h handler) GetAPI(mode string) (mApi.API, error) {
	if mode == "DB" {
		return h, nil
	} else if mode == "JSON" {
		return &WebJson{}, nil
	}
	return nil, errors.New("Mode invalid")
}

// Method Dependency injection
func (h handler) ShowAllPost(api mApi.API) echo.HandlerFunc {
	return func(c echo.Context) error {
		posts, err := api.GetAllPost()
		if err != nil {
			return c.String(http.StatusInternalServerError, "Error: " + err.Error())
		}
		
		return c.JSON(http.StatusOK, posts)
	}
}