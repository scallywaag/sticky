package presentation

import (
	"errors"
	"log"

	"github.com/highseas-software/sticky/internal/formatter"
	"github.com/highseas-software/sticky/internal/lists"
	"github.com/highseas-software/sticky/internal/notes"
)

func fetchAllNotes(listName string, notesService *notes.Service) (*notes.NotesResult, error) {
	res, err := notesService.GetAll(listName)
	if err != nil {
		switch {
		case errors.Is(err, lists.UserErrNoLists),
			errors.Is(err, lists.UserErrInexistentList),
			errors.Is(err, notes.UserErrNoNotes):
			formatter.PrintColored(err.Error(), formatter.Yellow)
			return nil, nil
		default:
			return nil, err
		}
	}
	return res, nil
}

func handleGetAllNotes(listName string, notesService *notes.Service) {
	res, err := fetchAllNotes(listName, notesService)
	if err != nil {
		log.Fatal(err)
	}
	if res == nil {
		return
	}

	formatter.ClearScreen()
	formatter.PrintListHeader(res.ActiveListName, res.NotesCount)
	printNotes(res.NotesList, res.NotesCount)
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

	res, err := fetchAllNotes(listName, notesService)
	if err != nil {
		log.Fatal(err)
	}
	if res == nil {
		return
	}

	formatter.ClearScreen()
	formatter.PrintListHeader(res.ActiveListName, res.NotesCount)
	printNotes(res.NotesList, res.NotesCount)
	formatter.PrintColored("\nNote successfully added.", formatter.Yellow)
}

func handleDeleteNotes(id int, notesService *notes.Service) {
	listName, err := notesService.Delete(id)
	if err != nil {
		if errors.Is(err, notes.UserErrNoNotes) || errors.Is(err, notes.UserErrInvalidDel) {
			formatter.PrintColored(err.Error(), formatter.Yellow)
			return
		}
		log.Fatal(err)
	}

	res, err := notesService.GetAll(listName)
	if err != nil {
		switch {
		case errors.Is(err, lists.UserErrNoLists),
			errors.Is(err, lists.UserErrInexistentList):
			formatter.PrintColored(err.Error(), formatter.Yellow)
		case errors.Is(err, notes.UserErrNoNotes):
			formatter.ClearScreen()
			formatter.PrintColored("Note successfully deleted.", formatter.Yellow)
			return
		}
	}

	formatter.ClearScreen()
	formatter.PrintListHeader(res.ActiveListName, res.NotesCount)
	printNotes(res.NotesList, res.NotesCount)
	formatter.PrintColored("\nNote successfully deleted.", formatter.Yellow)
}

func handleMutateNotes(id int, color formatter.Color, status notes.NoteStatus, notesService *notes.Service) {
	listName, err := notesService.Update(id, color, status)
	if err != nil {
		if errors.Is(err, notes.UserErrNoNotes) || errors.Is(err, notes.UserErrInvalidMut) {
			formatter.PrintColored(err.Error(), formatter.Yellow)
			return
		}
		log.Fatal(err)
	}

	res, err := fetchAllNotes(listName, notesService)
	if err != nil {
		log.Fatal(err)
	}
	if res == nil {
		return
	}

	formatter.ClearScreen()
	formatter.PrintListHeader(res.ActiveListName, res.NotesCount)
	printNotes(res.NotesList, res.NotesCount)
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

	activeListName, err := listsService.GetActiveOrSetFirst()
	if err != nil {
		log.Fatal(err)
	}

	formatter.ClearScreen()
	formatter.PrintListHeader("lists", count)
	for _, l := range allLists {
		style := formatter.StatusDefault
		if l.Name == activeListName {
			style = formatter.StatusBold
		}
		formatter.PrintContent(l.Name, l.Id, count, formatter.Default, style)
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

func printNotes(notesList []notes.Note, notesCount int) {
	for _, note := range notesList {
		style := formatter.StatusDefault
		if note.Status == notes.StatusCross {
			style = formatter.StatusCross
		}
		formatter.PrintContent(note.Content, note.Id, notesCount, note.Color, style)
	}
}
