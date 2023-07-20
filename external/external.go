package external

import (
	"github.com/labstack/echo/v4"
)
const (
	ROUTE_GET  = "GET"
	ROUTE_POST = "POST"
)

var methodRoutes map[string]map[string] echo.HandlerFunc = InitMethodRoutes()

func InitMethodRoutes() map[string]map[string] echo.HandlerFunc {
	output := map[string]map[string]echo.HandlerFunc{}
	output[ROUTE_GET] = make(map[string]echo.HandlerFunc)
	output[ROUTE_POST] = make(map[string]echo.HandlerFunc)
	return output
}