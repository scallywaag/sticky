package database

import (
	"database/sql"
	"log"
)

func execStmt(db *sql.DB, stmt string) {
	if _, err := db.Exec(stmt); err != nil {
		log.Fatalf("Failed to execute statement: %v\nSQL: %s", err, stmt)
	}
}

func createTables(db *sql.DB) {
	execStmt(db, ListsSQL)
	execStmt(db, NotesSQL)
	execStmt(db, StateSQL)
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
