package presentation

import (
	"errors"
	"log"

	f "github.com/highseas-software/sticky/internal/formatter"
	l "github.com/highseas-software/sticky/internal/lists"
	n "github.com/highseas-software/sticky/internal/notes"
)

func handleGetAllNotes(listName string, notesService *n.Service) {
	err := notesService.GetAll(listName)
	if err != nil {
		if errors.Is(err, l.UserErrNoLists) || errors.Is(err, l.UserErrInexistentList) {
			f.PrintColored(err.Error(), f.Yellow)
		} else {
			log.Fatal(err)
		}
	}
}

func handleAddNotes(content string, color f.Color, status n.NoteStatus, notesService *n.Service) {
	err := notesService.Add(content, color, status)
	if err != nil {
		if errors.Is(err, l.UserErrNoLists) {
			f.PrintColored(err.Error(), f.Yellow)
		} else {
			log.Fatal(err)
		}
	}
}

func handleDeleteNotes(id int, notesService *n.Service) {
	err := notesService.Delete(id)
	if err != nil {
		log.Fatal(err)
	}
}

func handleMutateNotes(id int, color f.Color, status n.NoteStatus, notesService *n.Service) {
	err := notesService.Update(id, color, status)
	if err != nil {
		log.Fatal(err)
	}
}

func handleGetAllLists(listsService *l.Service) {
	lists, count, err := listsService.GetAll()
	if err != nil {
		if errors.Is(err, l.UserErrNoLists) {
			f.PrintColored(err.Error(), f.Yellow)
		} else {
			log.Fatal(err)
		}
	}

	f.ClearScreen()
	f.PrintListHeader("lists", count)
	for _, l := range lists {
		f.PrintContent(l.Name, l.Id, count, f.Default, false)
	}
}

func handleAddList(listName string, listsService *l.Service) {
	if err := listsService.Add(listName); err != nil {
		log.Fatal(err)
	}

	handleGetAllLists(listsService)

	f.PrintColored("\nList successfully added.", f.Yellow)
}

func handleDeleteList(listId int, listsService *l.Service) {
	err := listsService.Delete(listId)
	if err != nil {
		if errors.Is(err, l.UserErrNoLists) {
			f.PrintColored(err.Error(), f.Yellow)
			return
		} else if errors.Is(err, l.ErrNoListsExist) {
			f.PrintColored("\nList successfully deleted. No lists remain.", f.Yellow)
			return
		} else {
			log.Fatal(err)
		}
	}

	handleGetAllLists(listsService)

	f.PrintColored("\nList successfully deleted.", f.Yellow)
}
