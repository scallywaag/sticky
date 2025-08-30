package lists

import (
	"errors"
)

// service
var (
	UserErrNoLists        = errors.New("You have no lists. Use -list-add or -la to get started.")
	UserErrInexistentList = errors.New("This list does not exist. Use -lists or -ls to see existing lists.")
)

// repo
var (
	ErrNoActiveList   = errors.New("no active list found")
	ErrNoRowsAffected = errors.New("no rows affected")
	ErrNoListsExist   = errors.New("no lists exist")
	ErrListInexistent = errors.New("inexistent list")
)
