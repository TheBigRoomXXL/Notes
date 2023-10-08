package app

import (
	note "notes/app/notes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Run() {
	// Setup Database
	db := CreateDbConnection("notes.db")
	defer db.Close()

	// Setup Echo
	e := echo.New()
	e.Renderer = T
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${status}: ${method} \"${uri}\" \n",
	}))
	e.Use(middlewareDb(db))

	//Routes
	e.Static("/static", "app/static")
	e.GET("/", note.Index)
	e.GET("/notes", note.GetNotes)
	e.POST("/notes", note.PostNotes)
	e.PUT("/notes/:id", note.PutNote)

	e.Logger.Fatal(e.Start(":3000"))
}