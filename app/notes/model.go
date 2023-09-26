package note

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

type Note struct {
	Id      int
	Content string
}

func SelectNotes(c echo.Context) ([]Note, error) {
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

func InsertNote(c echo.Context) ([]Note, error) {
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

func UpdateNote(c echo.Context) ([]Note, error) {
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
