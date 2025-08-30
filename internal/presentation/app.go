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

var (
	clearScreen     = formatter.ClearScreen
	printListHeader = formatter.PrintListHeader
	printContent    = formatter.PrintContent
	printColored    = formatter.PrintColored
)

func RunApp(flags *f.Flags, listsService *lists.Service, notesService *notes.Service) {
	color := f.ExtractColor(flags)
	status := f.GetNoteStatus(flags)

	switch {
	case flags.Add != "":
		err := notesService.Add(flags.Add, color, status)
		if err != nil {
			if errors.Is(err, lists.UserErrNoLists) {
				formatter.PrintColored(err.Error(), formatter.Yellow)
			} else {
				log.Fatal(err)
			}
		}
	case flags.List != "":
		err := notesService.GetAll(flags.List)
		if err != nil {
			if errors.Is(err, lists.UserErrNoLists) || errors.Is(err, lists.UserErrInexistentList) {
				formatter.PrintColored(err.Error(), formatter.Yellow)
			} else {
				log.Fatal(err)
			}
		}
	case flags.Del != 0:
		err := notesService.Delete(flags.Del)
		if err != nil {
			log.Fatal(err)
		}
	case flags.Mut != 0:
		err := notesService.Update(flags.Mut, color, status)
		if err != nil {
			log.Fatal(err)
		}
	case flags.GetAllLists:
		l, count, err := listsService.GetAll()

		clearScreen()
		printListHeader("lists", count)

		for _, l := range l {
			printContent(l.Name, l.Id, count, formatter.Default, false)
		}

		if err != nil {
			if errors.Is(err, lists.UserErrNoLists) {
				formatter.PrintColored(err.Error(), formatter.Yellow)
			} else {
				log.Fatal(err)
			}
		}
	case flags.AddList != "":
		err := listsService.Add(flags.AddList)
		if err != nil {
			log.Fatal(err)
		}
	case flags.DelList != 0:
		err := listsService.Delete(flags.DelList)
		if err != nil {
			if errors.Is(err, lists.UserErrNoLists) {
				formatter.PrintColored(err.Error(), formatter.Yellow)
			} else {
				log.Fatal(err)
			}
		}
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
