package notes

import (
	"errors"
)

// service
var (
	UserErrNoNotes = errors.New("You have no notes. Use -add or -a to add one.")
)

// repo
var (
	ErrNoRowsAffected = errors.New("no rows affected")
	ErrNoRowsResult   = errors.New("no rows in result set")
)
