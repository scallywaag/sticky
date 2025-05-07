package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type flags struct {
	add  string
	get  int
	list bool
	del  int
}

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

func main() {
	f := new(flags)
	flag.StringVar(&f.add, "add", "", "add a note")
	flag.IntVar(&f.get, "get", 0, "get a note by id")
	flag.BoolVar(&f.list, "list", false, "list all notes")
	flag.IntVar(&f.del, "del", 0, "delete a note by id")
	flag.Parse()

	db, err := sql.Open("sqlite3", "sticky.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt := `
	CREATE TABLE IF NOT EXISTS notes (
		id integer NOT NULL PRIMARY KEY,
		content TEXT
	);
	`

	_, err = db.Exec(stmt)
	if err != nil {
		log.Printf("%q: %s\n", err, stmt)
	}

	switch {
	case f.add != "":
		add(f.add, db)
	case f.get != 0:
		get(f.get, db)
	case f.list:
		fmt.Println("flag list")
	case f.del != 0:
		fmt.Println("flag del")
	default:
		fmt.Println("default - flag list")
	}
}
