package routes

import (
	"github.com/golangast/resumezach/src/handler/get/home"
	"github.com/labstack/echo"
)

func Routes(e *echo.Echo) {
	e.GET("/", home.Home)

}
