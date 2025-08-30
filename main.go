package main

import (
	"github.com/highseas-software/sticky/cmd/sticky"
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

	db := database.InitDb()
	defer db.Close()

	listsRepo := lists.NewDBRepository(db)
	listsService := lists.NewService(listsRepo)

	notesRepo := notes.NewDBRepository(db)
	notesService := notes.NewService(notesRepo, listsRepo)

	sticky.InitApp(f, listsService, notesService)
}
