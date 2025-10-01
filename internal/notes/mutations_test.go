package notes

import (
	"testing"

	"github.com/scallywaag/sticky/internal/formatter"
)

func TestMutateColor(t *testing.T) {
	tests := []struct {
		name     string
		current  formatter.Color
		newColor formatter.Color
		expected formatter.Color
	}{
		{
			name:     "returns current when newColor is empty",
			current:  formatter.Green,
			newColor: "",
			expected: formatter.Green,
		},
		{
			name:     "returns Default when newColor equals current",
			current:  formatter.Blue,
			newColor: formatter.Blue,
			expected: formatter.Default,
		},
		{
			name:     "returns newColor when newColor is different",
			current:  formatter.Red,
			newColor: formatter.Yellow,
			expected: formatter.Yellow,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mutateColor(tt.current, tt.newColor)
			if got != tt.expected {
				t.Errorf("mutateColor(%q, %q) = %q; want %q",
					tt.current, tt.newColor, got, tt.expected)
			}
		})
	}
}

func TestToggleStatus(t *testing.T) {
	tests := []struct {
		name      string
		current   NoteStatus
		newStatus NoteStatus
		expected  NoteStatus
	}{
		{
			name:      "returns current when newStatus is empty",
			current:   StatusCross,
			newStatus: "",
			expected:  StatusCross,
		},
		{
			name:      "returns StatusDefault when newStatus equals current",
			current:   StatusPin,
			newStatus: StatusPin,
			expected:  StatusDefault,
		},
		{
			name:      "returns newStatus when newStatus is different",
			current:   StatusPin,
			newStatus: StatusCross,
			expected:  StatusCross,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toggleStatus(tt.current, tt.newStatus)
			if got != tt.expected {
				t.Errorf("toggleStatus(%q, %q) = %q; want %q",
					tt.current, tt.newStatus, got, tt.expected)
			}
		})
	}
}
