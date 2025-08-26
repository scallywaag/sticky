package notes

import "github.com/highseas-software/sticky/internal/formatter"

type Note struct {
	Id      int
	Content string
	Color   formatter.Color
	Status  NoteStatus
}

type NoteStatus string

const (
	StatusPin     NoteStatus = "pin"
	StatusCross   NoteStatus = "cross"
	StatusDefault NoteStatus = "default"
)
