package flags

import (
	"flag"
	"fmt"
	"os"

	"github.com/highseas-software/sticky/internal/formatter"
	"github.com/highseas-software/sticky/internal/notes"
)

type Flags struct {
	// operations
	List string
	Add  string
	Del  int
	Mut  int

	// mutations
	Pin    bool
	Cross  bool
	Red    bool
	Green  bool
	Yellow bool
	Blue   bool
}

func Parse() *Flags {
	f := &Flags{}

	flag.StringVar(&f.List, "l", "", "")
	flag.StringVar(&f.Add, "a", "", "")
	flag.IntVar(&f.Del, "d", 0, "")
	flag.IntVar(&f.Mut, "m", 0, "")

	flag.BoolVar(&f.Pin, "p", false, "")
	flag.BoolVar(&f.Cross, "c", false, "")
	flag.BoolVar(&f.Red, "r", false, "")
	flag.BoolVar(&f.Green, "g", false, "")
	flag.BoolVar(&f.Blue, "b", false, "")
	flag.BoolVar(&f.Yellow, "y", false, "")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of sticky:\n")
		fmt.Fprintf(os.Stderr, "\n* operations\n")
		fmt.Fprintf(os.Stderr, "  -l <listname> string\n\tlist all notes from <listname>\n")
		fmt.Fprintf(os.Stderr, "  -a <content> string\n\tadd note <content>\n")
		fmt.Fprintf(os.Stderr, "  -d <id> int\n\tdelete note by <id>\n")
		fmt.Fprintf(os.Stderr, "  -m <id> int\n\tmutate existing note <id>\n")

		fmt.Fprintf(os.Stderr, "\n* mutations (toggle)\n\tuseable with -a or -m flag\n")
		fmt.Fprintf(os.Stderr, "  -p bool\n\tpin note - send to top of list\n")
		fmt.Fprintf(os.Stderr, "  -c bool\n\tcross note - send to bottom of list\n")
		fmt.Fprintf(os.Stderr, "  -r bool\n\tcolor note red\n")
		fmt.Fprintf(os.Stderr, "  -g bool\n\tcolor note green\n")
		fmt.Fprintf(os.Stderr, "  -b bool\n\tcolor note blue\n")
		fmt.Fprintf(os.Stderr, "  -y bool\n\tcolor note yellow\n")
	}

	flag.Parse()

	validateFlags(f)

	return f
}

func validateFlags(f *Flags) {
	opCount := 0
	if f.Add != "" {
		opCount++
	}
	if f.List != "" {
		opCount++
	}
	if f.Del > 0 {
		opCount++
	}
	if f.Mut > 0 {
		opCount++
	}

	formatCount := 0
	if f.Red {
		formatCount++
	}
	if f.Green {
		formatCount++
	}
	if f.Blue {
		formatCount++
	}
	if f.Yellow {
		formatCount++
	}

	mutCount := 0
	if f.Cross {
		mutCount++
	}
	if f.Pin {
		mutCount++
	}

	if opCount > 1 {
		fmt.Println("Error: only one of -l, -a, -d, -m can be used at a time.")
		os.Exit(1)
	}

	if formatCount > 1 {
		fmt.Println("Error: only one of -r, -g, -b, -y can be used at a time.")
		os.Exit(1)
	}

	if mutCount > 1 {
		fmt.Println("Error: only one of -p, -c can be used at a time.")
		os.Exit(1)
	}

	if (f.Red || f.Green || f.Blue || f.Yellow || f.Cross) && (f.Add == "" && f.Mut == 0) {
		fmt.Println("Error: formatting requires an operation like -a or -m.")
		os.Exit(1)
	}
}

func ExtractColor(f *Flags) formatter.Color {
	switch {
	case f.Red:
		return formatter.Red
	case f.Green:
		return formatter.Green
	case f.Blue:
		return formatter.Blue
	case f.Yellow:
		return formatter.Yellow
	default:
		return ""
	}
}

func GetNoteStatus(f *Flags) notes.NoteStatus {
	switch {
	case f.Pin:
		return notes.StatusPin
	case f.Cross:
		return notes.StatusCross
	default:
		return ""
	}
}
