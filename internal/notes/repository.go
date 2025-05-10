package notes

import (
	"database/sql"
	"fmt"
	f "github.com/scallywaag/sticky/internal/formatter"
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
				ROW_NUMBER() OVER (
					ORDER BY
						CASE
							WHEN type = 'pin' THEN 1
							WHEN type = 'todo' AND status = 'active' THEN 2
							WHEN type = 'misc' THEN 3
							WHEN type = 'todo' AND status = 'done' THEN 4
							WHEN type = 'todo' AND status = 'canceled' THEN 5
							ELSE 6
						END,
						id
				) AS virtual_id,
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

	var statusStr sql.NullString
	n := Note{}

	err = stmt.QueryRow(id).Scan(&n.VirtualId, &n.Content, &n.Type, &statusStr)

	if statusStr.Valid {
		n.Status = NoteStatus(statusStr.String)
	} else {
		n.Status = ""
	}

	if err != nil {
		return err
	}

	color := noteColor(n.Type, n.Status)
	f.Print(n.Content, n.VirtualId, 11, color)
	return nil
}

func List(db *sql.DB) error {
	stmt, err := db.Prepare(`
		WITH ordered_notes AS (
			SELECT
				ROW_NUMBER() OVER (
					ORDER BY
						CASE
							WHEN type = 'pin' THEN 1
							WHEN type = 'todo' AND status = 'active' THEN 2
							WHEN type = 'misc' THEN 3
							WHEN type = 'todo' AND status = 'done' THEN 4
							WHEN type = 'todo' AND status = 'canceled' THEN 5
							ELSE 6
						END,
						id
				) AS virtual_id,
				content,
				type,
				status
			FROM notes
		)
		SELECT virtual_id, content, type, status
		FROM ordered_notes
	`)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var statusStr sql.NullString
		n := Note{}

		err = rows.Scan(&n.VirtualId, &n.Content, &n.Type, &statusStr)

		if statusStr.Valid {
			n.Status = NoteStatus(statusStr.String)
		} else {
			n.Status = ""
		}

		if err != nil {
			log.Fatal()
		}

		color := noteColor(n.Type, n.Status)
		f.Print(n.Content, n.VirtualId, 11, color)
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
				ROW_NUMBER() OVER (
					ORDER BY
						CASE
							WHEN type = 'pin' THEN 1
							WHEN type = 'todo' AND status = 'active' THEN 2
							WHEN type = 'misc' THEN 3
							WHEN type = 'todo' AND status = 'done' THEN 4
							WHEN type = 'todo' AND status = 'canceled' THEN 5
							ELSE 6
						END,
						id
				) AS virtual_id,
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
				ROW_NUMBER() OVER (
					ORDER BY
						CASE
							WHEN type = 'pin' THEN 1
							WHEN type = 'todo' AND status = 'active' THEN 2
							WHEN type = 'misc' THEN 3
							WHEN type = 'todo' AND status = 'done' THEN 4
							WHEN type = 'todo' AND status = 'canceled' THEN 5
							ELSE 6
						END,
						id
				) AS virtual_id,
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
