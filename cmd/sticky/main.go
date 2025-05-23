package main

import (
	"github.com/scallywaag/sticky/internal/database"
	"github.com/scallywaag/sticky/internal/flags"
	"github.com/scallywaag/sticky/internal/notes"
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
		t := flags.GetNoteType(f)
		s := flags.GetNoteStatus(f)
		err := notes.Add(f.Add, t, s, db)
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
	case f.Done != 0:
		err := notes.UpdateTodo(f.Done, notes.StatusDone, db)
		if err != nil {
			log.Fatal(err)
		}
	case f.Cross != 0:
		err := notes.UpdateTodo(f.Cross, notes.StatusCanceled, db)
		if err != nil {
			log.Fatal(err)
		}
	default:
		notes.List(db)
	}
}
