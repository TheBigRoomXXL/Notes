package notes

import (
	"database/sql"
	"strings"
	"time"
)

// This struct is only used to generate the openapi documentation
type JustContent struct {
	Content string `form:"content" json:"content"`
}

type noteSerializer struct {
	Id      int    `param:"id" form:"id" json:"id"`
	Content string `form:"content" json:"content"`
}

type noteSearch struct {
	Search string `query:"search" form:"search" json:"search"`
}

type Note struct {
	Id         int
	Content    string
	Created_at time.Time // Managed by db with default value
	Updated_at time.Time // Managed by db with trigger
}

func SelectNotes(db *sql.DB, query noteSearch) ([]Note, error) {
	keywords := getKeywords(query)

	rows, err := buildFtsQuery(db, keywords)
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

	err = tx.Commit()
	if err != nil {
		return note, err
	}

	return note, nil
}

func DeleteNote(db *sql.DB, note Note) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		DELETE FROM notes 
		WHERE id = $1
	`, note.Id)

	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func getKeywords(query noteSearch) []interface{} {
	raw := strings.Fields(query.Search)
	keywords := []interface{}{}
	for _, keyword := range raw {
		keywords = append(keywords, "%"+keyword+"%")
	}
	return keywords
}

func buildFtsQuery(db *sql.DB, keywords []interface{}) (*sql.Rows, error) {
	var stmt strings.Builder
	stmt.WriteString(`
		SELECT id, content
		FROM notes 
	`)

	if len(keywords) > 0 {
		stmt.WriteString(` WHERE content LIKE ?`)

	}

	if len(keywords) > 1 {
		for i := 0; i < len(keywords)-1; i++ {
			stmt.WriteString(` AND content LIKE ?`)
		}
	}

	stmt.WriteString(` ORDER BY updated_at DESC`)
	rows, err := db.Query(stmt.String(), keywords...)

	return rows, err
}
