package notes

type Note struct {
	Id        int
	Content   string
	Type      NoteType
	Status    NoteStatus
	VirtualId int
}

type NoteType string
type NoteStatus string

const (
	TypePin  NoteType = "pin"
	TypeTodo NoteType = "todo"
	TypeMisc NoteType = "misc"

	StatusActive   NoteStatus = "active"
	StatusDone     NoteStatus = "done"
	StatusCanceled NoteStatus = "canceled"
)
