package testutils

import (
	"database/sql"
	"os"
	"testing"

	"github.com/highseas-software/sticky/internal/database"
	"github.com/highseas-software/sticky/internal/env"
)

func SetupTestDB(t *testing.T) *sql.DB {
	t.Helper()

	if err := os.Setenv(env.StickyEnvVar, string(env.EnvTest)); err != nil {
		t.Fatalf("failed to set %s: %v", env.StickyEnvVar, err)
	}
	defer func() {
		if err := os.Unsetenv(env.StickyEnvVar); err != nil {
			t.Fatalf("failed to unset %s: %v", env.StickyEnvVar, err)
		}
	}()

	db := database.InitDb()
	// TODO: better error handling in InitDb and handle here as well
	return db
}

func LoadFixture(t *testing.T, db *sql.DB, sqlContent string) {
	t.Helper()

	if db == nil {
		t.Fatal("cannot load fixture into nil db")
	}

	result, err := db.Exec(sqlContent)
	if err != nil {
		t.Fatalf("failed to exec fixture SQL: %v", err)
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		t.Logf("warning: fixture executed but affected 0 rows")
	}
}

func GetRepo[T any](t *testing.T, constructor func(*sql.DB) T) T {
	t.Helper()
	db := SetupTestDB(t)
	return constructor(db)
}
