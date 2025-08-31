package notes

import "github.com/highseas-software/sticky/internal/formatter"

type Note struct {
	Id      int
	Content string
	Color   formatter.Color
	Status  NoteStatus
}

func NewNote(content string, color formatter.Color, status NoteStatus) *Note {
	return &Note{
		Content: content,
		Color:   defaultColor(color),
		Status:  defaultStatus(status),
	}
}

type NoteStatus string

const (
	StatusPin     NoteStatus = "pin"
	StatusCross   NoteStatus = "cross"
	StatusDefault NoteStatus = "default"
)
