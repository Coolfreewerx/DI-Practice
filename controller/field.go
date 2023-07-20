package controller

type DatasourceInput struct {
	Workspace string `json:"workspace"`
}

type DatasourceOuput struct {
	Error 	string		`json:"error"`
	Data 	[]Post 		`json:"data"`
}

type Post struct {
	UserId 	int    		`json:"userId"`
	ID    	int    		`json:"id"`
	Title 	string 		`json:"title"`
	Body  	string 		`json:"body"`
}

type API interface {
	GetPosts() ([]Post, error)
}

type DB struct{}
type WebAPI struct{}


