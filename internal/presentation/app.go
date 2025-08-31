package presentation

import (
	"github.com/highseas-software/sticky/internal/flags"
	"github.com/highseas-software/sticky/internal/lists"
	"github.com/highseas-software/sticky/internal/notes"

	_ "github.com/mattn/go-sqlite3"
)

func RunApp(f *flags.Flags, listsService *lists.Service, notesService *notes.Service) {
	switch {
	case f.List != "":
		handleGetAllNotes(f.List, notesService)
	case f.Add != "":
		handleAddNotes(f.Add, flags.ExtractColor(f), flags.GetNoteStatus(f), notesService)
	case f.Del != 0:
		handleDeleteNotes(f.Del, notesService)
	case f.Mut != 0:
		handleMutateNotes(f.Mut, flags.ExtractColor(f), flags.GetNoteStatus(f), notesService)
	case f.GetAllLists:
		handleGetAllLists(listsService)

	case f.AddList != "":
		handleAddList(f.AddList, listsService)

	case f.DelList != 0:
		handleDeleteList(f.DelList, listsService)

	default:
		handleGetAllNotes(f.List, notesService)
	}
}
