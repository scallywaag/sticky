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
		t.Errorf("want 1 list, got none")
	}

	wantId := 1
	wantName := "sticky"
	got := lists[0]
	if got.Id != wantId || got.Name != wantName {
		t.Errorf(
			"want '%d - %s', got '%d - %s'",
			wantId, wantName, got.Id, got.Name,
		)
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
		t.Errorf("want 3 lists from fixture, got %d", len(lists))
	}
}

func TestAdd(t *testing.T) {
	repo := getRepo(t)

	id, err := repo.Add("test-list")
	if err != nil {
		t.Fatalf("Add returned error: %v", err)
	}

	if id <= 0 {
		t.Errorf("id > 0, got %d", id)
	}
}

func TestDelete(t *testing.T) {
	repo := getRepo(t)
	loadFixture(t, repo.db)

	err := repo.Delete(1)
	if err != nil {
		t.Fatalf("Delete returned error: %v", err)
	}
}

func TestGetActive(t *testing.T) {
	repo := getRepo(t)

	l, err := repo.GetActive()
	if err != nil {
		t.Fatalf("GetActive returned error: %v", err)
	}

	wantId := 1
	wantName := "sticky"
	if l.Id != wantId || l.Name != wantName {
		t.Errorf(
			"want '%d - %s', got '%d - %s'",
			wantId, wantName, l.Id, l.Name,
		)
	}
}

func TestGetActive_WithFixture(t *testing.T) {
	repo := getRepo(t)
	loadFixture(t, repo.db)

	l, err := repo.GetActive()
	if err != nil {
		t.Fatalf("GetActive returned error: %v", err)
	}

	wantId := 2
	wantName := "work"
	if l.Id != wantId || l.Name != wantName {
		t.Errorf(
			"want '%d - %s', got '%d - %s'",
			wantId, wantName, l.Id, l.Name,
		)
	}
}

func TestSetActive_WithFixture(t *testing.T) {
	repo := getRepo(t)
	loadFixture(t, repo.db)

	wantId := 2
	wantName := "work"

	l, err := repo.SetActive(wantId, wantName) // seems redundant to send both
	if err != nil {
		t.Fatalf("SetActive returned error: %v", err)
	}

	if l.Id != wantId || l.Name != wantName {
		t.Errorf(
			"want '%d - %s', got '%d - %s'",
			wantId, wantName, l.Id, l.Name,
		)
	}
}

func TestCount(t *testing.T) {
	repo := getRepo(t)

	want := 1
	got, err := repo.Count()
	if err != nil {
		t.Fatalf("Count returned error: %v", err)
	}

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
