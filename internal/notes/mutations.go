package notes

import (
	"github.com/scallywaag/sticky/internal/formatter"
)

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

func defaultColor(newColor formatter.Color) formatter.Color {
	if newColor == "" {
		return formatter.Default
	}
	return newColor
}

func defaultStatus(newStatus NoteStatus) NoteStatus {
	if newStatus == "" {
		return StatusDefault
	}
	return newStatus
}
