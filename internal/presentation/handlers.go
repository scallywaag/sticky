package presentation

import (
	"errors"
	"log"

	"github.com/highseas-software/sticky/internal/formatter"
	"github.com/highseas-software/sticky/internal/lists"
	"github.com/highseas-software/sticky/internal/notes"
)

func handleGetAllNotes(listName string, notesService *notes.Service) {
	notesList, count, listName, err := notesService.GetAll(listName)

	if err != nil {
		if errors.Is(err, lists.UserErrNoLists) || errors.Is(err, lists.UserErrInexistentList) {
			formatter.PrintColored(err.Error(), formatter.Yellow)
			return
		} else {
			log.Fatal(err)
		}
	}

	formatter.ClearScreen()
	formatter.PrintListHeader(listName, count)

	for _, note := range notesList {
		cross := note.Status == notes.StatusCross
		formatter.PrintContent(note.Content, note.Id, count, note.Color, cross)
	}
}

func handleAddNotes(content string, color formatter.Color, status notes.NoteStatus, notesService *notes.Service) {
	listName, err := notesService.Add(content, color, status)
	if err != nil {
		if errors.Is(err, lists.UserErrNoLists) {
			formatter.PrintColored(err.Error(), formatter.Yellow)
			return
		} else {
			log.Fatal(err)
		}
	}

	handleGetAllNotes(listName, notesService)
	formatter.PrintColored("\nNote successfully added.", formatter.Yellow)
}

func handleDeleteNotes(id int, notesService *notes.Service) {
	listName, err := notesService.Delete(id)
	if err != nil {
		log.Fatal(err)
	}

	handleGetAllNotes(listName, notesService)
	formatter.PrintColored("\nNote successfully deleted.", formatter.Yellow)
}

func handleMutateNotes(id int, color formatter.Color, status notes.NoteStatus, notesService *notes.Service) {
	listName, err := notesService.Update(id, color, status)
	if err != nil {
		log.Fatal(err)
	}

	handleGetAllNotes(listName, notesService)
	formatter.PrintColored("\nNote successfully mutated.", formatter.Yellow)
}

func handleGetAllLists(listsService *lists.Service) {
	allLists, count, err := listsService.GetAll()
	if err != nil {
		if errors.Is(err, lists.UserErrNoLists) {
			formatter.PrintColored(err.Error(), formatter.Yellow)
			return
		} else {
			log.Fatal(err)
		}
	}

	formatter.ClearScreen()
	formatter.PrintListHeader("lists", count)
	for _, l := range allLists {
		formatter.PrintContent(l.Name, l.Id, count, formatter.Default, false)
	}
}

func handleAddList(listName string, listsService *lists.Service) {
	if err := listsService.Add(listName); err != nil {
		log.Fatal(err)
	}

	handleGetAllLists(listsService)
	formatter.PrintColored("\nList successfully added.", formatter.Yellow)
}

func handleDeleteList(listId int, listsService *lists.Service) {
	err := listsService.Delete(listId)
	if err != nil {
		if errors.Is(err, lists.UserErrNoLists) {
			formatter.PrintColored(err.Error(), formatter.Yellow)
			return
		} else if errors.Is(err, lists.ErrNoListsExist) {
			formatter.PrintColored("\nList successfully deleted. No lists remain.", formatter.Yellow)
			return
		} else {
			log.Fatal(err)
		}
	}

	handleGetAllLists(listsService)
	formatter.PrintColored("\nList successfully deleted.", formatter.Yellow)
}
