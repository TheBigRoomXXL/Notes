package notes

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	db := c.Get("db").(*sql.DB)
	notes, err := SelectNotes(db, noteSearch{})
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "index", notes)
}

func GetNotes(c echo.Context) error {
	var validInput noteSearch
	err := c.Bind(&validInput)
	if err != nil {
		return err
	}

	db := c.Get("db").(*sql.DB)
	notes, err := SelectNotes(db, validInput)
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "notes", notes)

}

func PostNotes(c echo.Context) error {
	var validInput noteSerializer
	err := c.Bind(&validInput)
	if err != nil {
		return err
	}

	note, err := InsertNote(
		c.Get("db").(*sql.DB),
		Note{Content: validInput.Content},
	)
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "note", note)
}

// Handler PUT request.
// Default behavior is just to update the record in the DB and return 200
// but if the content of the note is empty then it's treated as a DELETE
// request and return 204.
func PutNote(c echo.Context) error {

	var validInput noteSerializer
	err := c.Bind(&validInput)
	if err != nil {
		return err
	}

	if validInput.Content == "" {
		return handleDelete(c, validInput)
	}

	return handleUpdate(c, validInput)
}

func DelNote(c echo.Context) error {
	var validInput noteSerializer
	err := c.Bind(&validInput)
	if err != nil {
		return err
	}

	return handleDelete(c, validInput)
}

func handleUpdate(c echo.Context, validatedInput noteSerializer) error {
	note, err := UpdateNote(
		c.Get("db").(*sql.DB),
		Note{Id: validatedInput.Id, Content: validatedInput.Content},
	)
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "note", note)
}

func handleDelete(c echo.Context, validatedInput noteSerializer) error {
	err := DeleteNote(
		c.Get("db").(*sql.DB),
		Note{Id: validatedInput.Id, Content: validatedInput.Content},
	)
	if err != nil {
		return err
	}

	return c.NoContent(204)
}
