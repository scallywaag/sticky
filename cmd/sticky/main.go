package main

import (
	"example/sticky/internal/database"
	"example/sticky/internal/flags"
	"example/sticky/internal/notes"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	f := flags.Parse()
	flags.ValidateFlags(f)

	db := database.InitDb()
	defer db.Close()

	switch {
	case f.Add != "":
		err := notes.Add(f.Add, notes.TypeMisc, "", db)
		if err != nil {
			log.Fatal(err)
		}
	case f.Get != 0:
		err := notes.Get(f.Get, db)
		if err != nil {
			log.Fatal(err)
		}
	case f.List:
		err := notes.List(db)
		if err != nil {
			log.Fatal(err)
		}
	case f.Del != 0:
		err := notes.Del(f.Del, db)
		if err != nil {
			log.Fatal(err)
		}
	default:
		notes.List(db)
	}
}
