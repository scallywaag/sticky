package flags

import "flag"

type Flags struct {
	Add  string
	Get  int
	List bool
	Del  int
}

func Parse() *Flags {
	f := &Flags{}
	flag.StringVar(&f.Add, "add", "", "add a note")
	flag.IntVar(&f.Get, "get", 0, "get a note by id")
	flag.BoolVar(&f.List, "list", false, "list all notes")
	flag.IntVar(&f.Del, "del", 0, "delete a note by id")
	flag.Parse()
	return f
}
