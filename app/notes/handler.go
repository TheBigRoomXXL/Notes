package notes

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func Index(c echo.Context) error {
	db := c.Get("db").(*sql.DB)
	notes, err := SelectNotes(db, "")
	if err != nil {
		log.Error(err)
		return err
	}

	err = c.Render(http.StatusOK, "index", notes)
	if err != nil {
		log.Error(err)
	}
	return err
}

func GetNotes(c echo.Context) error {
	db := c.Get("db").(*sql.DB)
	notes, err := SelectNotes(db, "")
	if err != nil {
		return err
	}

	err = c.Render(http.StatusOK, "notes", notes)
	if err != nil {
		log.Error(err)
	}
	return err
}

func PostNotes(c echo.Context) error {
	var validInput NoteSerializer
	err := c.Bind(&validInput)
	if err != nil {
		return err
	}

	db := c.Get("db").(*sql.DB)
	note, err := InsertNote(db, Note{
		Content: validInput.Content,
	})
	if err != nil {
		return err
	}

	err = c.Render(http.StatusOK, "note", note)
	if err != nil {
		log.Error(err)
	}
	return err
}

func PutNote(c echo.Context) error {
	db := c.Get("db").(*sql.DB)
	note := Note{Id: 0, Content: c.FormValue("content")}
	note, err := UpdateNote(db, note)
	if err != nil {
		return err
	}

	err = c.Render(http.StatusOK, "note", note)
	if err != nil {
		log.Error(err)
	}
	return err
}
