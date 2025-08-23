package notes

import (
	"database/sql"
	"fmt"

	"github.com/highseas-software/sticky/internal/formatter"
	"github.com/highseas-software/sticky/internal/lists"
)

func List(name string, db *sql.DB) error {
	var activeList *lists.List
	var err error

	if name != "" {
		activeList, err = lists.SetActiveList(name, db)
		if err != nil {
			return fmt.Errorf("failed to set active list: %w", err)
		}

	} else {
		activeList, err = lists.GetActiveList(db)
		if err != nil {
			return fmt.Errorf("failed to get active list: %w", err)
		}
	}

	var count int
	err = db.QueryRow(CountNotesSQL, activeList.Id).Scan(&count)
	if err != nil {
		return fmt.Errorf("query row failed: %w", err)
	}

	stmt, err := db.Prepare(ListNotesSQL)
	if err != nil {
		return fmt.Errorf("prepare failed: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(activeList.Id)
	if err != nil {
		return fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	formatter.PrintListHeader(activeList.Name, count)
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
	activeList, err := lists.GetActiveList(db)
	if err != nil {
		return fmt.Errorf("failed to get active list: %w", err)
	}

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

	_, err = stmt.Exec(content, c, s, activeList.Id)
	if err != nil {
		return fmt.Errorf("exec failed: %w", err)
	}

	List(activeList.Name, db)
	formatter.PrintColored("\nNote successfully added.", formatter.Yellow)
	return nil
}

func Del(id int, db *sql.DB) error {
	activeList, err := lists.GetActiveList(db)
	if err != nil {
		return fmt.Errorf("failed to get active list: %w", err)
	}

	stmt, err := db.Prepare(DeleteNoteSQL)
	if err != nil {
		return fmt.Errorf("prepare failed: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(activeList.Id, id)
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

	List(activeList.Name, db)
	formatter.PrintColored("\nNote successfully deleted.", formatter.Yellow)
	return nil
}

func Mut(id int, color formatter.Color, status NoteStatus, db *sql.DB) error {
	activeList, err := lists.GetActiveList(db)
	if err != nil {
		return fmt.Errorf("failed to get active list: %w", err)
	}

	var currentColor formatter.Color
	var currentStatus NoteStatus
	err = db.QueryRow(
		GetMutationsSQL,
		activeList.Id,
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
	result, err := stmt.Exec(activeList.Id, c, x, id)
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

	List(activeList.Name, db)
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
