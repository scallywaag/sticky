package lists

import (
	"testing"

	"github.com/scallywaag/sticky/internal/testutils"

	_ "github.com/mattn/go-sqlite3"

	_ "embed"
)

//go:embed testdata/lists_seed.sql
var listsSeed string

func getRepo(t *testing.T) *DBRepository {
	return testutils.GetRepo(t, NewDBRepository)
}

func TestListsRepo(t *testing.T) {
	// t.Run("GetAll with default list", func(t *testing.T) {
	// 	repo := getRepo(t)
	//
	// 	lists, err := repo.GetAll()
	// 	if err != nil {
	// 		t.Fatalf("GetAll returned error: %v", err)
	// 	}
	//
	// 	if len(lists) == 0 {
	// 		t.Errorf("want 1 list, got none")
	// 	}
	//
	// 	wantId := 1
	// 	wantName := "sticky"
	// 	got := lists[0]
	// 	if got.Id != wantId || got.Name != wantName {
	// 		t.Errorf(
	// 			"want '%d - %s', got '%d - %s'",
	// 			wantId, wantName, got.Id, got.Name,
	// 		)
	// 	}
	// })

	t.Run("GetAll with fixture", func(t *testing.T) {
		repo := getRepo(t)
		testutils.LoadFixture(t, repo.db, listsSeed)

		lists, err := repo.GetAll()
		if err != nil {
			t.Fatalf("GetAll returned error: %v", err)
		}

		if len(lists) != 3 {
			t.Errorf("want 3 lists from fixture, got %d", len(lists))
		}
	})

	t.Run("Add", func(t *testing.T) {
		repo := getRepo(t)

		id, err := repo.Add("test-list")
		if err != nil {
			t.Fatalf("Add returned error: %v", err)
		}

		if id <= 0 {
			t.Errorf("id > 0, got %d", id)
		}
	})

	t.Run("Delete", func(t *testing.T) {
		repo := getRepo(t)
		testutils.LoadFixture(t, repo.db, listsSeed)

		err := repo.Delete(1)
		if err != nil {
			t.Fatalf("Delete returned error: %v", err)
		}
	})

	// t.Run("GetActive", func(t *testing.T) {
	// 	repo := getRepo(t)
	//
	// 	l, err := repo.GetActive()
	// 	if err != nil {
	// 		t.Fatalf("GetActive returned error: %v", err)
	// 	}
	//
	// 	wantId := 1
	// 	wantName := "sticky"
	// 	if l.Id != wantId || l.Name != wantName {
	// 		t.Errorf(
	// 			"want '%d - %s', got '%d - %s'",
	// 			wantId, wantName, l.Id, l.Name,
	// 		)
	// 	}
	// })

	t.Run("GetActive with fixture", func(t *testing.T) {
		repo := getRepo(t)
		testutils.LoadFixture(t, repo.db, listsSeed)

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
	})

	t.Run("SetActive with fixture", func(t *testing.T) {
		repo := getRepo(t)
		testutils.LoadFixture(t, repo.db, listsSeed)

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
	})

	// t.Run("Count", func(t *testing.T) {
	// 	repo := getRepo(t)
	//
	// 	want := 1
	// 	got, err := repo.Count()
	// 	if err != nil {
	// 		t.Fatalf("Count returned error: %v", err)
	// 	}
	//
	// 	if want != got {
	// 		t.Errorf("want %d, got %d", want, got)
	// 	}
	// })

	t.Run("Count with fixture", func(t *testing.T) {
		repo := getRepo(t)
		testutils.LoadFixture(t, repo.db, listsSeed)

		want := 3
		got, err := repo.Count()
		if err != nil {
			t.Fatalf("Count returned error: %v", err)
		}

		if want != got {
			t.Errorf("want %d, got %d", want, got)
		}
	})
}
