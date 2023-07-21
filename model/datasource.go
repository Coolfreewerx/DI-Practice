package model

type DatasourceInput struct {
	Workspace string `json:"workspace"`
}

type DatasourceOuput struct {
	Error 	string		`json:"error"`
	Data 	[]Post 		`json:"data"`
}