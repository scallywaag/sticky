package lists

import (
	"errors"
)

// service
var (
	UserErrNoLists = errors.New("You have no lists. Use -list-add or -la to get started.")
)

// repo
var (
	ErrNoActiveList   = errors.New("no active list found")
	ErrNoRowsAffected = errors.New("no rows affected")
)
