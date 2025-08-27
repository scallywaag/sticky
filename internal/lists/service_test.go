package lists

import (
	"fmt"
	"testing"

	"github.com/highseas-software/sticky/internal/formatter"
)

type mockRepo struct {
	count int
	lists []List
	err   error
}

func (m *mockRepo) Count() (int, error) {
	return m.count, m.err
}

func (m *mockRepo) GetAll() ([]List, error) {
	return m.lists, m.err
}

func (m *mockRepo) Add(name string) (int, error)                 { return 0, nil }
func (m *mockRepo) Delete(id int) error                          { return nil }
func (m *mockRepo) GetActive() (*List, error)                    { return nil, nil }
func (m *mockRepo) SetActive(id int, name string) (*List, error) { return nil, nil }
func (m *mockRepo) GetId(name string) (int, error)               { return 0, nil }

func TestGetAll(t *testing.T) {
	cleared := false
	printed := []string{}

	oldClear, oldHeader, oldPrint := clearScreen, printListHeader, printContent
	defer func() {
		clearScreen, printListHeader, printContent = oldClear, oldHeader, oldPrint
	}()

	clearScreen = func() { cleared = true }
	printListHeader = func(title string, count int) { printed = append(printed, fmt.Sprintf("%s:%d", title, count)) }
	printContent = func(name string, id, count int, color formatter.Color, active bool) {
		printed = append(printed, fmt.Sprintf("%s:%d", name, id))
	}

	mock := &mockRepo{
		count: 1,
		lists: []List{{Id: 1, Name: "default"}},
	}
	service := NewService(mock)

	if err := service.GetAll(); err != nil {
		t.Fatalf("GetAll failed: %v", err)
	}

	if !cleared {
		t.Errorf("expected ClearScreen to be called")
	}
	if len(printed) != 2 {
		t.Errorf("expected 2 prints, got %d", len(printed))
	}
}
