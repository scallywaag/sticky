package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/highseas-software/sticky/internal/config"
	"github.com/highseas-software/sticky/internal/env"
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
			fmt.Println("Created 'sticky.db' database at: " + dbPath)
		}
	}

	db.Exec("PRAGMA foreign_keys = ON")

	createTables(db)

	return db
}
