package handler

import "net/http"

func Home(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", nil)
}
