package main

import (
	"example/sticky/internal/database"
	"example/sticky/internal/flags"
	"example/sticky/internal/notes"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	f := flags.Parse()

	db := database.InitDb()
	defer db.Close()

	switch {
	case f.Add != "":
		notes.Add(f.Add, db)
	case f.Get != 0:
		notes.Get(f.Get, db)
	case f.List:
		notes.List(db)
	case f.Del != 0:
		notes.Del(f.Del, db)
	default:
		notes.List(db)
	}
}
