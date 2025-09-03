package flags

import (
	"github.com/highseas-software/sticky/internal/formatter"
	"github.com/highseas-software/sticky/internal/types"
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

func GetNoteStatus(f *Flags) types.NoteStatus {
	switch {
	case f.Pin:
		return types.StatusPin
	case f.Cross:
		return types.StatusCross
	default:
		return ""
	}
}
