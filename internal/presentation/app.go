package presentation

import (
	"errors"
	"log"

	f "github.com/highseas-software/sticky/internal/flags"
	"github.com/highseas-software/sticky/internal/formatter"
	"github.com/highseas-software/sticky/internal/lists"
	"github.com/highseas-software/sticky/internal/notes"

	_ "github.com/mattn/go-sqlite3"
)

func RunApp(flags *f.Flags, listsService *lists.Service, notesService *notes.Service) {
	switch {
	case flags.List != "":
		handleGetAllNotes(flags.List, notesService)
	case flags.Add != "":
		handleAddNotes(flags.Add, f.ExtractColor(flags), f.GetNoteStatus(flags), notesService)
	case flags.Del != 0:
		handleDeleteNotes(flags.Del, notesService)
	case flags.Mut != 0:
		handleMutateNotes(flags.Mut, f.ExtractColor(flags), f.GetNoteStatus(flags), notesService)
	case flags.GetAllLists:
		handleGetAllLists(listsService)

	case flags.AddList != "":
		handleAddList(flags.AddList, listsService)

	case flags.DelList != 0:
		handleDeleteList(flags.DelList, listsService)

	default:
		err := notesService.GetAll(flags.List)
		if err != nil {
			if errors.Is(err, lists.UserErrNoLists) {
				formatter.PrintColored(err.Error(), formatter.Yellow)
			} else {
				log.Fatal(err)
			}
		}
	}
}
