package main

import (
	"database/sql"
	"example/sticky/internal/flags"
	"example/sticky/internal/persistence"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type note struct {
	id      int
	content string
}

func add(content string, db *sql.DB) {
	stmt, err := db.Prepare(`
		INSERT INTO notes(id, content)
		VALUES(NULL, ?)
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(content)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Note successfully added.")
}

func get(id int, db *sql.DB) {
	stmt, err := db.Prepare(`
		SELECT id, content FROM notes
		WHERE id = ?
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	n := note{}
	err = stmt.QueryRow(id).Scan(&n.id, &n.content)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("id: %d - content: %s\n", n.id, n.content)
}

func list(db *sql.DB) {
	stmt, err := db.Prepare(`SELECT id, content FROM notes`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		n := note{}
		err = rows.Scan(&n.id, &n.content)
		if err != nil {
			log.Fatal()
		}
		fmt.Printf("id: %d - content: %s\n", n.id, n.content)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func del(id int, db *sql.DB) {
	stmt, err := db.Prepare(`DELETE FROM notes WHERE id = ?`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Note successfully deleted.")
}

func main() {
	f := flags.Parse()

	db := persistence.InitDb()
	defer db.Close()

	switch {
	case f.Add != "":
		add(f.Add, db)
	case f.Get != 0:
		get(f.Get, db)
	case f.List:
		list(db)
	case f.Del != 0:
		del(f.Del, db)
	default:
		list(db)
	}
}
