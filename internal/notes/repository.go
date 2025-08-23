package notes

import (
	"database/sql"
	"fmt"

	"github.com/highseas-software/sticky/internal/formatter"
)

func List(db *sql.DB) error {
	stmt, err := db.Prepare(`
		SELECT id, content
		FROM notes 
	`)
	if err != nil {
		return fmt.Errorf("prepare failed: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		n := Note{}

		err = rows.Scan(&n.Id, &n.Content)
		if err != nil {
			return fmt.Errorf("scan failed: %w", err)
		}

		formatter.Print(n.Content, n.Id, 10, formatter.Blue)
	}

	if err = rows.Err(); err != nil {
		return fmt.Errorf("rows error: %w", err)
	}
	return nil
}

func Add(content string, db *sql.DB) error {
	stmt, err := db.Prepare(`
		INSERT INTO notes(id, content)
		VALUES(NULL, ?)
	`)
	if err != nil {
		return fmt.Errorf("prepare failed: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(content)
	if err != nil {
		return fmt.Errorf("exec failed: %w", err)
	}

	fmt.Println("Note successfully added.")
	return nil
}

func Del(id int, db *sql.DB) error {
	stmt, err := db.Prepare(`
		DELETE FROM notes WHERE id = ?
	`)
	if err != nil {
		return fmt.Errorf("prepare failed: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("prepare failed: %w", err)
	}

	fmt.Println("Note successfully deleted.")
	return nil
}
