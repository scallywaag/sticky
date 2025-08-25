package main

import (
	"log"

	"github.com/highseas-software/sticky/internal/config"
	"github.com/highseas-software/sticky/internal/database"
	"github.com/highseas-software/sticky/internal/flags"
	"github.com/highseas-software/sticky/internal/lists"
	"github.com/highseas-software/sticky/internal/notes"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	config.PrintAppEnv()
	f := flags.Parse()
	c := flags.ExtractColor(f)
	s := flags.GetNoteStatus(f)

	db := database.InitDb()
	defer db.Close()

	listsRepo := lists.NewDBRepository(db)
	listsService := lists.NewService(listsRepo)

	switch {
	case f.Add != "":
		err := notes.Add(f.Add, c, s, db)
		if err != nil {
			log.Fatal(err)
		}
	case f.List != "":
		err := notes.List(f.List, db)
		if err != nil {
			log.Fatal(err)
		}
	case f.Del != 0:
		err := notes.Del(f.Del, db)
		if err != nil {
			log.Fatal(err)
		}
	case f.Mut != 0:
		err := notes.Mut(f.Mut, c, s, db)
		if err != nil {
			log.Fatal(err)
		}
	case f.ListLists:
		err := listsService.GetAll()
		if err != nil {
			log.Fatal(err)
		}
	case f.AddList != "":
		err := lists.AddList(f.AddList, db)
		if err != nil {
			log.Fatal(err)
		}
	case f.DelList != 0:
		err := lists.DelList(f.DelList, db)
		if err != nil {
			log.Fatal(err)
		}
	default:
		err := notes.List(f.List, db)
		if err != nil {
			log.Fatal(err)
		}
	}
}
