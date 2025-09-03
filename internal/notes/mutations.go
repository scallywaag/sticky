package notes

import (
	"github.com/highseas-software/sticky/internal/formatter"
	"github.com/highseas-software/sticky/internal/types"
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

func toggleStatus(current types.NoteStatus, newStatus types.NoteStatus) types.NoteStatus {
	if newStatus == "" {
		return current
	}
	if current == newStatus {
		return types.StatusDefault
	}
	return newStatus
}

func defaultColor(newColor formatter.Color) formatter.Color {
	if newColor == "" {
		return formatter.Default
	}
	return newColor
}

func defaultStatus(newStatus types.NoteStatus) types.NoteStatus {
	if newStatus == "" {
		return types.StatusDefault
	}
	return newStatus
}
