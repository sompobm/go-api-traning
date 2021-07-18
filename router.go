package api

import (
	"api/controller"
	"fmt"

	"github.com/labstack/echo/v4"
)

func NewServer(port int) {

	e := echo.New()
	e.GET("/", controller.Hello)
	e.GET("/users", controller.GetUsers)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", port)))
}
