package database

import (
	"database/sql"
	"fmt"
	"log"
)

func InitDb() *sql.DB {
	dbPath := getDbPath()
	dbExists := fileExists(dbPath)

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	db.Exec("PRAGMA foreign_keys = ON")

	createTables(db)
	ensureDefaultList(db)

	if !dbExists {
		fmt.Println("Created 'sticky.db' database at: " + dbPath)
	}

	return db
}
