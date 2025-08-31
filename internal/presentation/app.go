package presentation

import (
	f "github.com/highseas-software/sticky/internal/flags"
	l "github.com/highseas-software/sticky/internal/lists"
	n "github.com/highseas-software/sticky/internal/notes"

	_ "github.com/mattn/go-sqlite3"
)

func RunApp(flags *f.Flags, listsService *l.Service, notesService *n.Service) {
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
		handleGetAllNotes(flags.List, notesService)
	}
}
