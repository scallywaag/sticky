package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/scallywaag/sticky/internal/config"
	"github.com/scallywaag/sticky/internal/env"
	"github.com/scallywaag/sticky/internal/formatter"
)

func InitDb() *sql.DB {
	appEnv := config.GetAppEnv()

	var db *sql.DB
	var err error

	if appEnv == env.EnvTest {
		db, err = sql.Open("sqlite3", ":memory:")
		if err != nil {
			log.Fatalf("failed to open in-memory db: %v", err)
		}

		fmt.Println("Initialized in-memory SQLite database.")
	} else {
		dbPath := getDbPath(appEnv)
		dbExists := fileExists(dbPath)

		db, err = sql.Open("sqlite3", dbPath)
		if err != nil {
			log.Fatalf("failed to open db file: %v", err)
		}

		if !dbExists {
			message := fmt.Sprintf("Created 'sticky.db' database at: %s", dbPath)
			formatter.PrintColored(message, formatter.Blue)
		}
	}

	db.Exec("PRAGMA foreign_keys = ON")

	createTables(db)

	return db
}
