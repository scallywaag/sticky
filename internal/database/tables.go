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
	execStmt(db, DefaultStateSQL)
}
