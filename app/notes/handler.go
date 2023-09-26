package note

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	notes, err := SelectNotes(c)
	if err != nil {
		return err
	}
	err = c.Render(http.StatusOK, "index.html", notes)
	return err
}

func GetNotes(c echo.Context) error {
	notes, err := SelectNotes(c)
	if err != nil {
		return err
	}
	err = c.Render(http.StatusOK, "notes.html", notes)
	return err
}

func PostNote(c echo.Context) error {
	notes, err := InsertNote(c)
	if err != nil {
		return err
	}
	err = c.Render(http.StatusOK, "newNote.html", notes)
	return err
}

func PutNote(c echo.Context) error {
	notes, err := UpdateNote(c)
	if err != nil {
		return err
	}
	err = c.Render(http.StatusOK, "updatedNote.html", notes)
	return err
}
