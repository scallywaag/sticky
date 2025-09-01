package notes

import (
	"errors"
)

// service
var (
	UserErrNoNotes = errors.New("You have no notes. Use -add or -a to add one.")
)
