package lists

import (
	"database/sql"
	"os"
	"testing"

	"github.com/highseas-software/sticky/internal/database"

	_ "github.com/mattn/go-sqlite3"
)

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

func loadFixture(t *testing.T, db *sql.DB, path string) {
	t.Helper()

	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read fixture file: %v", err)
	}

	_, err = db.Exec(string(content))
	if err != nil {
		t.Fatalf("failed to execute fixture SQL: %v", err)
	}
}

func TestGetAll(t *testing.T) {
	repo := getRepo(t)

	lists, err := repo.GetAll()
	if err != nil {
		t.Fatalf("GetAll returned error: %v", err)
	}

	if len(lists) == 0 {
		t.Errorf("expected default 'sticky' list, got none")
	}
}

func TestAdd(t *testing.T) {
	repo := getRepo(t)

	_, err := repo.Add("test-list")
	if err != nil {
		t.Errorf("Add returned error: %v", err)
	}
}
