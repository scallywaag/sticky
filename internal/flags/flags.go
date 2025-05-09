package flags

import "flag"

type Flags struct {
	Add   string
	Get   int
	List  bool
	Del   int
	Pin   bool
	Todo  bool
	Done  int
	Cross int
}

func Parse() *Flags {
	f := &Flags{}

	flag.StringVar(&f.Add, "add", "", "add a note")
	flag.IntVar(&f.Get, "get", 0, "get a note by id")
	flag.BoolVar(&f.List, "list", false, "list all notes")
	flag.IntVar(&f.Del, "del", 0, "delete a note by id")

	flag.BoolVar(&f.Pin, "pin", false, "pin note to top of list")
	flag.BoolVar(&f.Todo, "todo", false, "mark as todo note")
	flag.IntVar(&f.Done, "done", 0, "set todo note as done")
	flag.IntVar(&f.Cross, "cross", 0, "set todo note as canceled")

	flag.Parse()

	return f
}
