package notes

import (
	"database/sql"
	"fmt"

	"github.com/highseas-software/sticky/internal/formatter"
)

func List(db *sql.DB) error {
	var count int
	err := db.QueryRow(CountNotesSQL).Scan(&count)
	if err != nil {
		return fmt.Errorf("query row failed: %w", err)
	}

	stmt, err := db.Prepare(ListNotesSQL)
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

		err = rows.Scan(&n.VirtualId, &n.Content, &n.Color, &n.Status)
		if err != nil {
			return fmt.Errorf("scan failed: %w", err)
		}

		cross := n.Status == StatusCross
		formatter.Print(n.Content, n.VirtualId, count, n.Color, cross)
	}

	if err = rows.Err(); err != nil {
		return fmt.Errorf("rows error: %w", err)
	}
	return nil
}

func Add(content string, color formatter.Color, status NoteStatus, db *sql.DB) error {
	stmt, err := db.Prepare(AddNoteSQL)
	if err != nil {
		return fmt.Errorf("prepare failed: %w", err)
	}
	defer stmt.Close()

	c := color
	if c == "" {
		c = formatter.Default
	}

	s := status
	if s == "" {
		s = StatusDefault
	}

	_, err = stmt.Exec(content, c, s)
	if err != nil {
		return fmt.Errorf("exec failed: %w", err)
	}

	List(db)
	formatter.PrintColored("\nNote successfully added.", formatter.Yellow)
	return nil
}

func Del(id int, db *sql.DB) error {
	stmt, err := db.Prepare(DeleteNoteSQL)
	if err != nil {
		return fmt.Errorf("prepare failed: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("exec failed: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("delete failed: no rows affected")
	}

	List(db)
	formatter.PrintColored("\nNote successfully deleted.", formatter.Yellow)
	return nil
}

func Mut(id int, color formatter.Color, status NoteStatus, db *sql.DB) error {
	var currentColor formatter.Color
	var currentStatus NoteStatus
	err := db.QueryRow(
		GetMutationsSQL,
		id,
	).Scan(&currentColor, &currentStatus)
	if err != nil {
		return fmt.Errorf("query row failed: %w", err)
	}

	stmt, err := db.Prepare(MutateNoteSQL)
	if err != nil {
		return fmt.Errorf("prepare failed: %w", err)
	}
	defer stmt.Close()

	c := mutateColor(currentColor, color)
	x := toggleStatus(currentStatus, status)
	result, err := stmt.Exec(c, x, id)
	if err != nil {
		return fmt.Errorf("exec failed: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("delete failed: no rows affected")
	}

	List(db)
	formatter.PrintColored("\nNote successfully mutated.", formatter.Yellow)
	return nil
}

func mutateColor(current formatter.Color, newColor formatter.Color) formatter.Color {
	if newColor == "" {
		return current
	}
	if current == newColor {
		return formatter.Default
	}
	return newColor
}

func toggleStatus(current NoteStatus, newStatus NoteStatus) NoteStatus {
	if newStatus == "" {
		return current
	}
	if current == newStatus {
		return StatusDefault
	}
	return newStatus
}
