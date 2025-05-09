package flags

import (
	"example/sticky/internal/notes"
	"flag"
	"fmt"
	"os"
)

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

func ValidateFlags(f *Flags) {
	opCount := 0
	if f.Add != "" {
		opCount++
	}
	if f.Get > 0 {
		opCount++
	}
	if f.List {
		opCount++
	}
	if f.Del > 0 {
		opCount++
	}
	if f.Done > 0 {
		opCount++
	}
	if f.Cross > 0 {
		opCount++
	}

	if opCount > 1 {
		fmt.Println("Error: only one of --add, --get, --list, --del, --done, --cross can be used at a time.")
		os.Exit(1)
	}

	if (f.Pin || f.Todo) && f.Add == "" {
		fmt.Println("Error: --pin and --todo can only be used with --add.")
		os.Exit(1)
	}

	if f.Done > 0 && f.Cross > 0 {
		fmt.Println("Error: can only use one of --done or --cross.")
		os.Exit(1)
	}
}

func GetNoteType(f *Flags) notes.NoteType {
	if f.Pin {
		return notes.TypePin
	}

	if f.Todo {
		return notes.TypeTodo
	}

	return notes.TypeMisc
}

func GetNoteStatus(f *Flags) notes.NoteStatus {
	if f.Done > 0 {
		return notes.StatusDone
	}

	if f.Cross > 0 {
		return notes.StatusCanceled
	}

	return notes.StatusActive
}
