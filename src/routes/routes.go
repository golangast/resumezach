package routes

import (
	home "github.com/golangast/resumezach/src/handler/get/home"
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	e.GET("/", home.Home)

}
