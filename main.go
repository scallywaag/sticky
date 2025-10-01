package main

import (
	"github.com/scallywaag/sticky/internal/config"
	"github.com/scallywaag/sticky/internal/database"
	"github.com/scallywaag/sticky/internal/flags"
	"github.com/scallywaag/sticky/internal/lists"
	"github.com/scallywaag/sticky/internal/notes"
	"github.com/scallywaag/sticky/internal/presentation"

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

	presentation.RunApp(f, listsService, notesService)
}
