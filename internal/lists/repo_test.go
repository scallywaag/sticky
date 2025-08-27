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

	os.Setenv("STICKY_ENV", "test")
	defer os.Unsetenv("STICKY_ENV")

	db := database.InitDb()
	return db
}

func getRepo(t *testing.T) *DBRepository {
	t.Helper()

	db := setupTestDB(t)
	repo := NewDBRepository(db)
	return repo
}

func loadFixture(t *testing.T, db *sql.DB) {
	t.Helper()
	_, err := db.Exec(listsSeed)
	if err != nil {
		t.Fatalf("failed to exec fixture: %v", err)
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

	_, err := repo.Add("test-list")
	if err != nil {
		t.Errorf("Add returned error: %v", err)
	}
}
