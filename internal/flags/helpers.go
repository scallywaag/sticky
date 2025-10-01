package flags

import (
	"github.com/scallywaag/sticky/internal/formatter"
	"github.com/scallywaag/sticky/internal/notes"
)

func ExtractColor(f *Flags) formatter.Color {
	switch {
	case f.Red:
		return formatter.Red
	case f.Green:
		return formatter.Green
	case f.Blue:
		return formatter.Blue
	case f.Yellow:
		return formatter.Yellow
	default:
		return ""
	}
}

func GetNoteStatus(f *Flags) notes.NoteStatus {
	switch {
	case f.Pin:
		return notes.StatusPin
	case f.Cross:
		return notes.StatusCross
	default:
		return ""
	}
}
