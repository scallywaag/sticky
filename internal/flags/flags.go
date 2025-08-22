package flags

import (
	"flag"
)

type Flags struct {
	// operations
	List string
	Add  string
	Del  string

	// mutations
	Pin   string
	Cross string

	// formatting
	Red    string
	Green  string
	Yellow string
	Blue   string
}

func Parse() *Flags {
	f := &Flags{}

	flag.StringVar(&f.List, "list", "", "list all notes")
	flag.StringVar(&f.Add, "add", "", "add a note")
	flag.StringVar(&f.Del, "del", "", "delete a note by id")

	flag.StringVar(&f.Pin, "pin", "", "pin note - send to top of list (toggle)")
	flag.StringVar(&f.Cross, "cross", "", "cross note - send to bottom of list (toggle)")

	flag.StringVar(&f.Red, "red", "", "color the note red")
	flag.StringVar(&f.Green, "green", "", "color the note green")
	flag.StringVar(&f.Yellow, "yellow", "", "color the note yellow")
	flag.StringVar(&f.Blue, "blue", "", "color the note blue")

	flag.Parse()

	return f
}
