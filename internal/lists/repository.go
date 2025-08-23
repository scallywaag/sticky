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

	err := db.QueryRow(`
		SELECT l.id, l.name
		FROM state s
		JOIN lists l ON s.list_id = l.id
		WHERE s.key = 'active';
	`).Scan(&l.Id, &l.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no active list found: %w", err)
		}
		return nil, fmt.Errorf("query row failed: %w", err)
	}

	return l, nil
}
