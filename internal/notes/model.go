package notes

type Note struct {
	Id      int
	Content string
	Color   string
	Status  NoteStatus
}

type NoteStatus string

const (
	StatusPin     NoteStatus = "pin"
	StatusCross   NoteStatus = "cross"
	StatusDefault NoteStatus = "default"
)
