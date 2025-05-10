package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/scallywaag/sticky/internal/formatter"
)

func InitDb() *sql.DB {
	dbPath := getDbPath()
	dbExists := fileExists(dbPath)

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	createNotesTable(db)

	if !dbExists {
		fmt.Println(formatter.Blue + "Created 'sticky.db' database at: " + dbPath + formatter.Reset)
	}

	return db
}

func getDbPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	dbPath := filepath.Join(homeDir, ".local", "share", "sticky", "sticky.db")
	if err := os.MkdirAll(filepath.Dir(dbPath), os.ModePerm); err != nil {
		log.Fatalf("Error creating directory: %v", err)
	}

	return dbPath
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}

func createNotesTable(db *sql.DB) {
	stmt := `
	CREATE TABLE IF NOT EXISTS notes (
		id integer NOT NULL PRIMARY KEY,
		content TEXT,
		type TEXT DEFAULT 'misc',
		status TEXT DEFAULT NULL
	);
	`

	_, err := db.Exec(stmt)
	if err != nil {
		log.Printf("%q: %s\n", err, stmt)
	}
}
