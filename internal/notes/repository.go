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

		err = rows.Scan(&n.Id, &n.Content, &n.Color, &n.Cross)
		if err != nil {
			return fmt.Errorf("scan failed: %w", err)
		}

		formatter.Print(n.Content, n.Id, count, n.Color, n.Cross)
	}

	if err = rows.Err(); err != nil {
		return fmt.Errorf("rows error: %w", err)
	}
	return nil
}

func Add(content string, color string, cross bool, db *sql.DB) error {
	stmt, err := db.Prepare(AddNoteSQL)
	if err != nil {
		return fmt.Errorf("prepare failed: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(content, color, cross)
	if err != nil {
		return fmt.Errorf("exec failed: %w", err)
	}

	formatter.PrintColored("Note successfully added.", formatter.Yellow)
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

func Mut(id int, color string, cross bool, db *sql.DB) error {
	var currentColor string
	var currentCross bool
	err := db.QueryRow(
		GetMutationsSQL,
		id,
	).Scan(&currentColor, &currentCross)
	if err != nil {
		return fmt.Errorf("query row failed: %w", err)
	}

	stmt, err := db.Prepare(MutateNoteSQL)
	if err != nil {
		return fmt.Errorf("prepare failed: %w", err)
	}
	defer stmt.Close()

	c := mutateColor(currentColor, color)
	x := toggleCross(currentCross, cross)
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

func mutateColor(current string, newColor string) string {
	if current == newColor {
		return formatter.Default
	}
	return newColor
}

func toggleCross(current bool, toggle bool) bool {
	if toggle {
		return !current
	}
	return current
}
