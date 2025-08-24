package notes

import "github.com/highseas-software/sticky/internal/formatter"

func mutateColor(current formatter.Color, newColor formatter.Color) formatter.Color {
	if newColor == "" {
		return current
	}
	if current == newColor {
		return formatter.Default
	}
	return newColor
}

func toggleStatus(current NoteStatus, newStatus NoteStatus) NoteStatus {
	if newStatus == "" {
		return current
	}
	if current == newStatus {
		return StatusDefault
	}
	return newStatus
}
