package controller

import (
	"api/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Hello(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello world")
}

func GetUsers(c echo.Context) error {
	us := service.UserService{}
	return c.JSON(http.StatusOK, us.GetUsers())
}
