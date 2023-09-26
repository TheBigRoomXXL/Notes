package main

import (
	"database/sql"
	"net/http"
	"notes/views"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Setup Database
	db := CreateDbConnection("notes.db")
	defer db.Close()

	// Setup Echo
	e := echo.New()
	e.Renderer = views.T
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${status}: ${method} \"${uri}\" \n",
	}))
	e.Use(middlewareDb(db))

	//Routes
	e.Static("/static", "views/static")
	e.GET("/", Hello)

	e.Logger.Fatal(e.Start(":3000"))
}

type Note struct {
	Id      int
	Content string
}

func Hello(c echo.Context) error {
	notes, err := GetData(c)
	if err != nil {
		return err
	}
	err = c.Render(http.StatusOK, "index.html", notes)
	return err
}

func GetData(c echo.Context) ([]Note, error) {
	db := c.Get("db").(*sql.DB)
	rows, err := db.Query("SELECT * FROM notes ")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	notes := []Note{}
	for rows.Next() {
		var note Note
		if err := rows.Scan(&note.Id, &note.Content); err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return notes, nil
}
