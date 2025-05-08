package notes

import (
	"database/sql"
	"fmt"
	"log"
)

func Add(content string, db *sql.DB) error {
	stmt, err := db.Prepare(`
		INSERT INTO notes(id, content)
		VALUES(NULL, ?)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(content)
	if err != nil {
		return err
	}

	fmt.Println("Note successfully added.")
	return nil
}

func Get(id int, db *sql.DB) error {
	stmt, err := db.Prepare(`
		SELECT id, content FROM notes
		WHERE id = ?
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	n := Note{}
	err = stmt.QueryRow(id).Scan(&n.Id, &n.Content)
	if err != nil {
		return err
	}

	fmt.Printf("id: %d - content: %s\n", n.Id, n.Content)
	return nil
}

func List(db *sql.DB) error {
	stmt, err := db.Prepare(`SELECT id, content FROM notes`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		n := Note{}
		err = rows.Scan(&n.Id, &n.Content)
		if err != nil {
			log.Fatal()
		}
		fmt.Printf("id: %d - content: %s\n", n.Id, n.Content)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return nil
}

func Del(id int, db *sql.DB) error {
	stmt, err := db.Prepare(`DELETE FROM notes WHERE id = ?`)
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
