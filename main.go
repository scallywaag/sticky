package main

import (
	"flag"
	"fmt"
)

type flags struct {
	add  string
	get  int
	list bool
	del  int
}

func main() {
	f := new(flags)
	flag.StringVar(&f.add, "add", "", "add a note")
	flag.IntVar(&f.get, "get", 0, "get a note by id")
	flag.BoolVar(&f.list, "list", false, "list all notes")
	flag.IntVar(&f.del, "del", 0, "delete a note by id")
	flag.Parse()

	switch {
	case f.add != "":
		fmt.Println("flag add")
	case f.get != 0:
		fmt.Println("flag get")
	case f.list:
		fmt.Println("flag list")
	case f.del != 0:
		fmt.Println("flag del")
	default:
		fmt.Println("default - flag list")
	}
}
