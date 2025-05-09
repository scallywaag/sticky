package notes

import (
	"database/sql"
	f "example/sticky/internal/formatter"
	"fmt"
	"log"
)

func Add(content string, noteType NoteType, noteStatus NoteStatus, db *sql.DB) error {
	stmt, err := db.Prepare(`
		INSERT INTO notes(id, content, type, status)
		VALUES(NULL, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var ns sql.NullString
	if noteType == TypeTodo {
		ns = sql.NullString{String: string(noteStatus), Valid: true}
	} else {
		ns = sql.NullString{Valid: false}
	}

	_, err = stmt.Exec(content, noteType, ns)
	if err != nil {
		return err
	}

	fmt.Println("Note successfully added.")
	return nil
}

func Get(id int, db *sql.DB) error {
	stmt, err := db.Prepare(`
		WITH ordered_notes AS (
			SELECT
				ROW_NUMBER() OVER (ORDER by id) AS virtual_id,
				content,
				type,
				status
			FROM notes
		)
		SELECT virtual_id, content, type, status
		FROM ordered_notes
		WHERE virtual_id = ?
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	n := Note{}
	err = stmt.QueryRow(id).Scan(&n.VirtualId, &n.Content, &n.Type, &n.Status)
	if err != nil {
		return err
	}

	f.Print(n.Content, n.VirtualId, 11)
	return nil
}

func List(db *sql.DB) error {
	stmt, err := db.Prepare(`
		WITH ordered_notes AS (
			SELECT
				ROW_NUMBER() OVER (ORDER by id) AS virtual_id,
				content,
				type,
				status
			FROM notes
		)
		SELECT virtual_id, content, type, status
		FROM ordered_notes
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		n := Note{}

		var statusStr sql.NullString
		err = rows.Scan(&n.VirtualId, &n.Content, &n.Type, &statusStr)

		if statusStr.Valid {
			n.Status = NoteStatus(statusStr.String)
		} else {
			n.Status = ""
		}

		if err != nil {
			log.Fatal()
		}
		f.Print(n.Content, n.VirtualId, 11)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return nil
}

func Del(id int, db *sql.DB) error {
	stmt, err := db.Prepare(`
		WITH ordered_notes AS (
			SELECT
				ROW_NUMBER() OVER (ORDER BY id) as virtual_id,
				id
			FROM notes
		)
		DELETE FROM notes
		WHERE id = (SELECT id FROM ordered_notes WHERE virtual_id = ?)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	fmt.Println("Note successfully deleted.")
	return nil
}

func UpdateTodo(id int, status NoteStatus, db *sql.DB) error {
	stmt, err := db.Prepare(`
		WITH ordered_notes AS (
			SELECT
				ROW_NUMBER() OVER (ORDER BY id) as virtual_id,
				id,
				type
			FROM notes
		)
		UPDATE notes
		SET status = ?
		WHERE id = (SELECT id FROM ordered_notes WHERE virtual_id = ? AND type = 'todo')
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(status, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no update performed: note not found or type not 'todo'")
	}

	fmt.Println("Note successfully updated.")
	return nil
}
