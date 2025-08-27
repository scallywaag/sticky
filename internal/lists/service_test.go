package lists

import (
	"fmt"
	"testing"

	"github.com/highseas-software/sticky/internal/formatter"
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

		mock := &MockListRepo{
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
}
