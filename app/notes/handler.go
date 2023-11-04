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

// @Summary List your notes
// @Description  Perform a Full Text Search on your notes. Ordered by most recent update.
// @Tags notes
// @Accept       json, application/x-www-form-urlencoded
// @Produce      json, text/html
// @Param        search  query string false "FTS query"
// @Success      200  {array}   notes.Note
// @Failure      422  {object}  notes.Note
// @Router       /notes [get]
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

// @Summary Create a notes
// @Description  Save a note to the database and return the ID
// @Tags notes
// @Accept       json, application/x-www-form-urlencoded
// @Produce      json, text/html
// @Param        content  body string  true "The content of the note"
// @Success      200  {array}   notes.Note
// @Failure      422  {object}  notes.Note
// @Router       /notes [post]
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

// @Summary Update or delete a notes
// @Description  Default behavior is just to update the record in the DB and return 200 but if the content of the note is empty then it's treated as a DELETE request and return 204.
// @Tags notes
// @Accept       json, application/x-www-form-urlencoded
// @Produce      json, text/html
// @Param        id  path int  true "The note identifier"
// @Param        content  body string  true "The content of the note"
// @Success      200  {array}   notes.Note
// @Success      204  ""
// @Failure      422  {object}  notes.Note
// @Router       /notes/{id} [put]
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

// @Summary Delete a notes
// @Description  Remove a record from the database
// @Tags notes
// @Accept       json, application/x-www-form-urlencoded
// @Produce      json, text/html
// @Param        id  path int  true "The note identifier"
// @Success      204  ""
// @Failure      422  {object}  notes.Note
// @Router       /notes/{id} [delete]
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
