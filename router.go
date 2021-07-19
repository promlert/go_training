package api

import (
	"api/controller"
	"fmt"

	"github.com/labstack/echo/v4"
)

func NewServer(port int) {
	r := echo.New()
	r.GET("/", controller.Hello)
	r.GET("/users", controller.GetUsers)
	r.Run(fmt.Sprintf(":%v", port))
}
