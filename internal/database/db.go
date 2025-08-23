package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
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

func getDbPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	dbPath := filepath.Join(homeDir, ".local", "share", "sticky", "sticky.db")
	if err := os.MkdirAll(filepath.Dir(dbPath), 0700); err != nil {
		log.Fatalf("Error creating directory: %v", err)
	}

	return dbPath
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}

func execStmt(db *sql.DB, stmt string) {
	if _, err := db.Exec(stmt); err != nil {
		log.Fatalf("Failed to execute statement: %v\nSQL: %s", err, stmt)
	}
}

func createTables(db *sql.DB) {
	// Lists
	execStmt(db, `
		CREATE TABLE IF NOT EXISTS lists (
			id INTEGER PRIMARY KEY,
			name TEXT UNIQUE NOT NULL
		);
	`)

	// Notes
	execStmt(db, `
		CREATE TABLE IF NOT EXISTS notes (
			id INTEGER PRIMARY KEY,
			content TEXT NOT NULL,
			color TEXT,
			status TEXT NOT NULL DEFAULT 'default',
			list_id INTEGER NOT NULL DEFAULT 1,
			FOREIGN KEY(list_id) REFERENCES lists(id) ON DELETE CASCADE
		);
	`)

	// State
	execStmt(db, `
		CREATE TABLE IF NOT EXISTS state (
			key TEXT PRIMARY KEY,
			list_id INTEGER NOT NULL,
			FOREIGN KEY(list_id) REFERENCES lists(id) ON DELETE CASCADE
		);
	`)
}

func ensureDefaultList(db *sql.DB) error {
	var count int
	err := db.QueryRow(`SELECT COUNT(*) FROM lists`).Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		// No lists exist -> create "sticky"
		res, err := db.Exec(`INSERT INTO lists (name) VALUES ('sticky')`)
		if err != nil {
			return err
		}

		id, err := res.LastInsertId()
		if err != nil {
			return err
		}

		// Set sticky as active in state table
		_, err = db.Exec(`
			INSERT INTO state (key, list_id)
			VALUES ('active', ?)
		`, id)
		if err != nil {
			return err
		}
	}

	// If count > 0, do nothing (leave existing lists and active state as-is)
	return nil
}
