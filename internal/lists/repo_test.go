package lists

import (
	"database/sql"
	"os"
	"testing"

	"github.com/highseas-software/sticky/internal/database"

	_ "github.com/mattn/go-sqlite3"

	_ "embed"
)

//go:embed testdata/lists_seed.sql
var listsSeed string

func setupTestDB(t *testing.T) *sql.DB {
	t.Helper()

	if err := os.Setenv("STICKY_ENV", "test"); err != nil {
		t.Fatalf("failed to set STICKY_ENV: %v", err)
	}
	defer func() {
		if err := os.Unsetenv("STICKY_ENV"); err != nil {
			t.Fatalf("failed to unset STICKY_ENV: %v", err)
		}
	}()

	db := database.InitDb()
	return db
}

func getRepo(t *testing.T) *DBRepository {
	t.Helper()

	db := setupTestDB(t)
	return NewDBRepository(db)
}

func loadFixture(t *testing.T, db *sql.DB) {
	t.Helper()

	if db == nil {
		t.Fatal("cannot load fixture into nil db")
	}

	result, err := db.Exec(listsSeed)
	if err != nil {
		t.Fatalf("failed to exec fixture SQL: %v", err)
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		t.Logf("warning: fixture executed but affected 0 rows")
	}
}

func TestGetAll_WithDefaultList(t *testing.T) {
	repo := getRepo(t)

	lists, err := repo.GetAll()
	if err != nil {
		t.Fatalf("GetAll returned error: %v", err)
	}

	if len(lists) == 0 {
		t.Errorf("expected default 'sticky' list, got none")
	}
}

func TestGetAll_WithFixture(t *testing.T) {
	repo := getRepo(t)
	loadFixture(t, repo.db)

	lists, err := repo.GetAll()
	if err != nil {
		t.Fatalf("GetAll returned error: %v", err)
	}

	if len(lists) != 3 {
		t.Errorf("expected 3 lists from fixture, got %d", len(lists))
	}
}

func TestAdd(t *testing.T) {
	repo := getRepo(t)

	id, err := repo.Add("test-list")
	if err != nil {
		t.Fatalf("Add returned error: %v", err)
	}

	if id <= 0 {
		t.Errorf("expected valid id > 0, got %d", id)
	}
}
