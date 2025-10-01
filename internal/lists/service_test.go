package lists

import (
	"fmt"
	"testing"

	"github.com/scallywaag/sticky/internal/formatter"
)

func TestListsService(t *testing.T) {
	t.Run("GetAll", func(t *testing.T) {
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

		mock := &MockRepo{
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
	})

	t.Run("Add", func(t *testing.T) {
		cleared := false
		printed := []string{}
		colored := []string{}

		oldClear, oldHeader, oldPrint, oldColored := clearScreen, printListHeader, printContent, printColored
		defer func() {
			clearScreen, printListHeader, printContent, printColored =
				oldClear, oldHeader, oldPrint, oldColored
		}()

		clearScreen = func() { cleared = true }
		printListHeader = func(title string, count int) {
			printed = append(printed, fmt.Sprintf("header:%s:%d", title, count))
		}
		printContent = func(name string, id, count int, color formatter.Color, active bool) {
			printed = append(printed, fmt.Sprintf("list:%s:%d", name, id))
		}
		printColored = func(msg string, color formatter.Color) {
			colored = append(colored, msg)
		}

		mock := &MockRepo{
			count: 2,
			lists: []List{
				{Id: 1, Name: "default"},
				{Id: 2, Name: "todo"},
			},
		}
		service := NewService(mock)

		if err := service.Add("todo"); err != nil {
			t.Fatalf("Add failed: %v", err)
		}

		if mock.addName != "todo" {
			t.Errorf("expected Add to be called with 'todo', got %q", mock.addName)
		}

		if !mock.addCalled {
			t.Error("expected Add to call repo.Add")
		}
		if !mock.getAllCalled {
			t.Error("expected Add to call repo.GetAll")
		}

		if !cleared {
			t.Errorf("expected ClearScreen to be called")
		}
		if len(printed) != 3 {
			t.Errorf("expected 3 prints (1 header, 2 lists), got %d", len(printed))
		}
		if len(colored) != 1 || colored[0] != "\nList successfully added." {
			t.Errorf("expected success message, got %+v", colored)
		}
	})
}
