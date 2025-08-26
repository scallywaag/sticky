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

	notesRepo := notes.NewDBRepository(db)
	notesService := notes.NewService(notesRepo, listsRepo)

	switch {
	case f.Add != "":
		err := notesService.Add(f.Add, c, s)
		if err != nil {
			log.Fatal(err)
		}
	case f.List != "":
		err := notesService.GetAll(f.List)
		if err != nil {
			log.Fatal(err)
		}
	case f.Del != 0:
		err := notesService.Delete(f.Del)
		if err != nil {
			log.Fatal(err)
		}
	case f.Mut != 0:
		err := notesService.Update(f.Mut, c, s)
		if err != nil {
			log.Fatal(err)
		}
	case f.GetAllLists:
		err := listsService.GetAll()
		if err != nil {
			log.Fatal(err)
		}
	case f.AddList != "":
		err := listsService.Add(f.AddList)
		if err != nil {
			log.Fatal(err)
		}
	case f.DelList != 0:
		err := listsService.Delete(f.DelList)
		if err != nil {
			log.Fatal(err)
		}
	default:
		err := notesService.GetAll(f.List)
		if err != nil {
			log.Fatal(err)
		}
	}
}
