package lists

import (
	"errors"
)

// service
var (
	ErrNoLists = errors.New("You have no lists. Use -listadd or -la to get started.")
)

// repo
var (
	ErrNoActiveList   = errors.New("no active list found")
	ErrNoRowsAffected = errors.New("no rows affected")
)
