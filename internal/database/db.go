package database

import (
	"database/sql"
	"log"
)

func InitDb() *sql.DB {
	db, err := sql.Open("sqlite3", "sticky.sqlite")
	if err != nil {
		log.Fatal(err)
	}

	stmt := `
	CREATE TABLE IF NOT EXISTS notes (
		id integer NOT NULL PRIMARY KEY,
		content TEXT,
		type TEXT DEFAULT 'misc',
		status TEXT DEFAULT NULL
	);
	`

	_, err = db.Exec(stmt)
	if err != nil {
		log.Printf("%q: %s\n", err, stmt)
	}

	return db
}
