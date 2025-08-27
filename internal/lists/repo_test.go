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

	os.Setenv("APP_ENV", "test")
	defer os.Unsetenv("APP_ENV")

	db := database.InitDb()
	return db
}

func TestGetAll(t *testing.T) {
	db := setupTestDB(t)
	repo := NewDBRepository(db)

	lists, err := repo.GetAll()
	if err != nil {
		t.Fatalf("GetAll returned error: %v", err)
	}

	if len(lists) == 0 {
		t.Errorf("expected default 'sticky' list, got none")
	}
}
