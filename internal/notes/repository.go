package notes

import (
	"database/sql"
	"fmt"

	"github.com/highseas-software/sticky/internal/formatter"
)

func List(db *sql.DB) error {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM notes").Scan(&count)
	if err != nil {
		return fmt.Errorf("query row failed: %w", err)
	}

	stmt, err := db.Prepare(`
		SELECT id, content, color
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

		err = rows.Scan(&n.Id, &n.Content, &n.Color)
		if err != nil {
			return fmt.Errorf("scan failed: %w", err)
		}

		formatter.Print(n.Content, n.Id, count, n.Color)
	}

	if err = rows.Err(); err != nil {
		return fmt.Errorf("rows error: %w", err)
	}
	return nil
}

func Add(content string, color string, db *sql.DB) error {
	stmt, err := db.Prepare(`
		INSERT INTO notes(id, content, color)
		VALUES(NULL, ?, ?)
	`)
	if err != nil {
		return fmt.Errorf("prepare failed: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(content, color)
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
