package flags

import (
	"flag"
)

func Parse() *Flags {
	f := &Flags{}

	flag.StringVar(&f.List, "l", "", "")
	flag.StringVar(&f.List, "list", "", "")

	flag.StringVar(&f.Add, "a", "", "")
	flag.StringVar(&f.Add, "add", "", "")

	flag.IntVar(&f.Del, "del", 0, "")
	flag.IntVar(&f.Del, "d", 0, "")

	flag.IntVar(&f.Mut, "m", 0, "")
	flag.IntVar(&f.Mut, "mut", 0, "")

	flag.BoolVar(&f.Pin, "p", false, "")
	flag.BoolVar(&f.Pin, "pin", false, "")
	flag.BoolVar(&f.Cross, "c", false, "")
	flag.BoolVar(&f.Cross, "cross", false, "")
	flag.BoolVar(&f.Red, "r", false, "")
	flag.BoolVar(&f.Red, "red", false, "")
	flag.BoolVar(&f.Green, "g", false, "")
	flag.BoolVar(&f.Green, "green", false, "")
	flag.BoolVar(&f.Blue, "b", false, "")
	flag.BoolVar(&f.Blue, "blue", false, "")
	flag.BoolVar(&f.Yellow, "y", false, "")
	flag.BoolVar(&f.Yellow, "yellow", false, "")

	flag.BoolVar(&f.GetAllLists, "ls", false, "")
	flag.BoolVar(&f.GetAllLists, "lists", false, "")
	flag.StringVar(&f.AddList, "la", "", "")
	flag.StringVar(&f.AddList, "list-add", "", "")
	flag.IntVar(&f.DelList, "ld", 0, "")
	flag.IntVar(&f.DelList, "list-del", 0, "")

	flag.Usage = printUsageMenu

	flag.Parse()

	validateFlags(f)

	return f
}
