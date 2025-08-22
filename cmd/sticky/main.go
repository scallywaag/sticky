package main

import (
	"fmt"
	"github.com/highseas-software/sticky/internal/database"

	_ "github.com/mattn/go-sqlite3"
)

type Note struct {
	Id      int
	Content string
}

func main() {
	fmt.Println("\x1b[34mHello World! ...in blue\x1b[0m")

	db := database.InitDb()
	defer db.Close()

	stmt, err := db.Prepare(`
		INSERT INTO notes(id, content)
		VALUES(NULL, ?)
	`)
	if err != nil {
		fmt.Println("Error in insert prepare statement")
	}
	defer stmt.Close()

	_, err = stmt.Exec("test note")
	if err != nil {
		fmt.Println("Error inserting note")
	}

	fmt.Println("Note successfully added.")

	stmt, err = db.Prepare(`
		SELECT id, content
		FROM notes 
		WHERE id = ?
	`)
	if err != nil {
		fmt.Println("Error in select prepare statement")
	}
	defer stmt.Close()

	n := Note{}

	err = stmt.QueryRow(1).Scan(&n.Id, &n.Content)
	if err != nil {
		fmt.Println("Error selecting note")
	}
	fmt.Printf("The note:\nid: %d\ncontent: %s\n", n.Id, n.Content)
}
