package notes

import (
	"errors"
)

// service
var (
	UserErrNoNotes    = errors.New("You have no notes. Use -add or -a to add one.")
	UserErrInvalidDel = errors.New("The note picked for deletion does not exist. To view existing notes use -list or -l.")
	UserErrInvalidMut = errors.New("The note picked for mutation does not exist. To view existing notes use -list or -l.")
)

// repo
var (
	ErrNoRowsAffected = errors.New("no rows affected")
	ErrNoRowsResult   = errors.New("no rows in result set")
)
