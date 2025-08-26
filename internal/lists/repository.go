package lists

import (
	"database/sql"
	"fmt"

	"github.com/highseas-software/sticky/internal/formatter"
)

func ListLists(db *sql.DB) error {
	var count int
	err := db.QueryRow(CountSQL).Scan(&count)
	if err != nil {
		return fmt.Errorf("query row failed: %w", err)
	}

	stmt, err := db.Prepare(GetAllSQL)
	if err != nil {
		return fmt.Errorf("prepare failed: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	formatter.ClearScreen()
	formatter.PrintListHeader("lists", count)
	for rows.Next() {
		l := List{}

		err = rows.Scan(&l.Id, &l.Name)
		if err != nil {
			return fmt.Errorf("scan failed: %w", err)
		}

		formatter.Print(l.Name, l.Id, count, formatter.Default, false)
	}

	if err = rows.Err(); err != nil {
		return fmt.Errorf("rows error: %w", err)
	}
	return nil
}

func AddList(name string, db *sql.DB) error {
	stmt, err := db.Prepare(AddSQL)
	if err != nil {
		return fmt.Errorf("prepare failed: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(name)
	if err != nil {
		return fmt.Errorf("exec failed: %w", err)
	}

	ListLists(db)
	formatter.PrintColored("\nList successfully added.", formatter.Yellow)
	return nil
}

func DelList(id int, db *sql.DB) error {
	stmt, err := db.Prepare(DeleteSQL)
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

	ListLists(db)
	formatter.PrintColored("\nList successfully deleted.", formatter.Yellow)
	return nil
}

func GetActiveList(db *sql.DB) (*List, error) {
	l := &List{}
	err := db.QueryRow(GetActiveSQL).Scan(&l.Id, &l.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no active list found: %w", err)
		}
		return nil, fmt.Errorf("query row failed: %w", err)
	}

	return l, nil
}

func SetActiveList(name string, db *sql.DB) (*List, error) {
	var listId int
	err := db.QueryRow(GetIdByNameSQL, name).Scan(&listId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("list %q does not exist", name)
		}
		return nil, fmt.Errorf("failed to query list: %w", err)
	}

	_, err = db.Exec(SetActiveSQL, listId)
	if err != nil {
		return nil, fmt.Errorf("failed to set active list: %w", err)
	}

	l := &List{
		Id:   listId,
		Name: name,
	}
	return l, nil
}
