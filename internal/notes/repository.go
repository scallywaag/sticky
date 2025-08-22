package notes

import (
	"database/sql"
	"fmt"
)

func Add(content string, db *sql.DB) error {
	stmt, err := db.Prepare(`
		INSERT INTO notes(id, content)
		VALUES(NULL, ?)
	`)
	if err != nil {
		fmt.Println("Error in insert prepare statement")
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec("test note")
	if err != nil {
		fmt.Println("Error inserting note")
		return err
	}

	fmt.Println("Note successfully added.")
	return nil
}

func Get(id int, db *sql.DB) error {
	stmt, err := db.Prepare(`
		SELECT id, content
		FROM notes 
		WHERE id = ?
	`)
	if err != nil {
		fmt.Println("Error in select prepare statement")
		return err
	}
	defer stmt.Close()

	n := Note{}

	err = stmt.QueryRow(id).Scan(&n.Id, &n.Content)
	if err != nil {
		fmt.Println("Error selecting note")
		return err
	}

	fmt.Printf("The note:\nid: %d\ncontent: %s\n", n.Id, n.Content)
	return nil
}
