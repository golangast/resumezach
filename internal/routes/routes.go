package routes

func Routes(e *echo.Echo) {
	e.GET("/", home.Home)

}
