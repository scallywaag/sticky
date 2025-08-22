package flags

import (
	"flag"
)

type Flags struct {
	// operations
	List string
	Add  string
	Del  int

	// mutations
	Pin   int
	Cross int

	// formatting
	Red    bool
	Green  bool
	Yellow bool
	Blue   bool
}

func Parse() *Flags {
	f := &Flags{}

	flag.StringVar(&f.List, "l", "", "list all notes")
	flag.StringVar(&f.Add, "a", "", "add a note")
	flag.IntVar(&f.Del, "d", 0, "delete a note by id")

	flag.IntVar(&f.Pin, "p", 0, "pin note - send to top of list (toggle)")
	flag.IntVar(&f.Cross, "c", 0, "cross note - send to bottom of list (toggle)")

	flag.BoolVar(&f.Red, "r", false, "color the note red")
	flag.BoolVar(&f.Green, "g", false, "color the note green")
	flag.BoolVar(&f.Blue, "b", false, "color the note blue")
	flag.BoolVar(&f.Yellow, "y", false, "color the note yellow")

	flag.Parse()

	return f
}
