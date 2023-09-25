package main

import (
	"fmt"
	"net/http"
	view "notes/src/views"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Renderer = view.T
	e.Static("/static", "src/views/static")
	e.GET("/", Hello)

	e.Logger.Info(e.Start(":3000"))
}

func Hello(c echo.Context) error {
	err := c.Render(http.StatusOK, "index.html", "World")
	fmt.Println(err)
	return err
}
