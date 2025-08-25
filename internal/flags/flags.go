package flags

import (
	"flag"
	"fmt"
	"os"

	"github.com/highseas-software/sticky/internal/formatter"
	"github.com/highseas-software/sticky/internal/notes"
)

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

	flag.BoolVar(&f.ListLists, "ls", false, "")
	flag.StringVar(&f.AddList, "la", "", "")
	flag.IntVar(&f.DelList, "ld", 0, "")

	flag.Usage = printUsageMenu

	flag.Parse()

	validateFlags(f)

	return f
}
