package types

type NoteStatus string

const (
	StatusPin     NoteStatus = "pin"
	StatusCross   NoteStatus = "cross"
	StatusDefault NoteStatus = "default"
)
