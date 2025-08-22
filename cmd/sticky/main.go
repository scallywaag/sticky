package main

import (
	"log"

	"github.com/highseas-software/sticky/internal/database"
	"github.com/highseas-software/sticky/internal/flags"
	"github.com/highseas-software/sticky/internal/notes"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	f := flags.Parse()

	db := database.InitDb()
	defer db.Close()

	switch {
	case f.Add != "":
		err := notes.Add(f.Add, db)
		if err != nil {
			log.Fatal(err)
		}
	case f.List != "":
		err := notes.List(db)
		if err != nil {
			log.Fatal(err)
		}
	default:
		notes.List(db)
	}
}
