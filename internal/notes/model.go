package notes

import (
	"github.com/highseas-software/sticky/internal/formatter"
	"github.com/highseas-software/sticky/internal/types"
)

type Note struct {
	Id      int
	Content string
	Color   formatter.Color
	Status  types.NoteStatus
}

func NewNote(content string, color formatter.Color, status types.NoteStatus) *Note {
	return &Note{
		Content: content,
		Color:   defaultColor(color),
		Status:  defaultStatus(status),
	}
}
