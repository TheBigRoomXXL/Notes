package server

import (
	notes "notes/app/notes"
	"notes/app/users"
	"strings"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Run() {

	// Setup Echo
	e := echo.New()

	// Set Session for Authentification
	e.Use(session.Middleware(sessions.NewCookieStore(GenerateRandomBytes(32))))

	// Setup Database
	db := CreateDbConnection("notes.db")
	defer db.Close()
	e.Use(middlewareDb(db))

	// Set Logging to text instead of JSON
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${status}: ${method} \"${uri}\" \n",
	}))

	// Load template for use by handlers
	e.Renderer = T

	// Gzip everything comming from static
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: func(c echo.Context) bool {
			return !strings.Contains(c.Path(), "static")
		},
	}))

	// Routes
	e.Static("/static", "app/static")

	e.GET("/", notes.Index)
	e.GET("/notes", notes.GetNotes)
	e.POST("/notes", notes.PostNotes)
	e.PUT("/notes/:id", notes.PutNote)
	e.DELETE("/notes/:id", notes.DelNote)

	e.POST("/users/login", users.PostLogin)

	// Run
	e.Logger.Fatal(e.Start(":3000"))
}
