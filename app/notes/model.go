package notes

import (
	"database/sql"
	"time"
)

type NoteSerializer struct {
	Id      int    `query:"id" form:"id" json:"id"`
	Content string `form:"content" json:"content"`
}

type Note struct {
	Id         int
	Content    string
	Created_at time.Time
	Updated_at time.Time
}

func SelectNotes(db *sql.DB, search string) ([]Note, error) {
	rows, err := db.Query(`
		SELECT id, content
		FROM notes 
		WHERE content LIKE '%' || $1 || '%'
		ORDER BY updated_at DESC
		`, search)
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

func InsertNote(db *sql.DB, note Note) (Note, error) {
	tx, err := db.Begin()
	if err != nil {
		return note, err
	}

	var id int
	err = tx.QueryRow(
		`INSERT INTO notes (content) VALUES (?) RETURNING id`,
		note.Content).Scan(&id)
	if err != nil {
		return note, err
	}

	if tx.Commit() != nil {
		return note, err
	}

	note.Id = id

	return note, nil
}

func UpdateNote(db *sql.DB, note Note) (Note, error) {
	tx, err := db.Begin()
	if err != nil {
		return note, err
	}

	_, err = tx.Exec(`
		UPDATE notes 
		SET content = $1
		WHERE id = $2`,
		note.Content,
		note.Id)
	if err != nil {
		return note, err
	}

	if tx.Commit() != nil {
		return note, err
	}

	return note, nil
}
