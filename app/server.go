package app

import (
	note "notes/app/notes"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Run() {

	// Setup Echo
	e := echo.New()

	// Load template for use by handlers
	e.Renderer = T

	//Text Loging
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${status}: ${method} \"${uri}\" \n",
	}))

	// Setup Database
	db := CreateDbConnection("notes.db")
	defer db.Close()
	e.Use(middlewareDb(db))

	// Gzip everything comming from static
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: func(c echo.Context) bool {
			return !strings.Contains(c.Path(), "static")
		},
	}))

	// Routes
	e.Static("/static", "app/static")

	e.GET("/", note.Index)
	e.GET("/notes", note.GetNotes)
	e.POST("/notes", note.PostNotes)
	e.PUT("/notes/:id", note.PutNote)
	e.DELETE("/notes/:id", note.DelNote)

	// Run
	e.Logger.Fatal(e.Start(":3000"))
}
