package lists

import (
	"database/sql"
	"fmt"
)

func ListLists(db *sql.DB) error {
	fmt.Println("not implemented")

	return nil
}

func AddList(db *sql.DB) error {
	fmt.Println("not implemented")

	return nil
}

func DelList(db *sql.DB) error {
	fmt.Println("not implemented")

	return nil
}

func GetActiveList(db *sql.DB) (*List, error) {
	l := &List{}

	err := db.QueryRow(GetActiveListSQL).Scan(&l.Id, &l.Name)
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
	err := db.QueryRow(GetListIdByNameSQL, name).Scan(&listId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("list %q does not exist", name)
		}
		return nil, fmt.Errorf("failed to query list: %w", err)
	}

	_, err = db.Exec(SetActiveListSQL, listId)
	if err != nil {
		return nil, fmt.Errorf("failed to set active list: %w", err)
	}

	l := &List{
		Id:   listId,
		Name: name,
	}
	return l, nil
}
