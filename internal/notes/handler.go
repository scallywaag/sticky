package notes

import (
	"github.com/scallywaag/sticky/internal/formatter"
)

func noteColor(noteType NoteType, noteStatus NoteStatus) string {
	if noteType == TypePin {
		return formatter.Blue
	}

	if noteType == TypeTodo {
		switch noteStatus {
		case StatusActive:
			return formatter.Yellow
		case StatusDone:
			return formatter.Green
		case StatusCanceled:
			return formatter.Red
		}
	}

	return formatter.Default
}
