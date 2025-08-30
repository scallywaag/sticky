package presentation

import (
	"errors"
	"log"

	"github.com/highseas-software/sticky/internal/formatter"
	l "github.com/highseas-software/sticky/internal/lists"
)

func handleGetAllLists(listsService *l.Service) {
	lists, count, err := listsService.GetAll()
	if err != nil {
		if errors.Is(err, l.UserErrNoLists) {
			formatter.PrintColored(err.Error(), formatter.Yellow)
		} else {
			log.Fatal(err)
		}
	}

	formatter.ClearScreen()
	formatter.PrintListHeader("lists", count)
	for _, l := range lists {
		formatter.PrintContent(l.Name, l.Id, count, formatter.Default, false)
	}
}

func handleAddList(listName string, listsService *l.Service) {
	if err := listsService.Add(listName); err != nil {
		log.Fatal(err)
	}

	handleGetAllLists(listsService)

	formatter.PrintColored("\nList successfully added.", formatter.Yellow)
}

func handleDeleteList(listId int, listsService *l.Service) {
	err := listsService.Delete(listId)
	if err != nil {
		if errors.Is(err, l.UserErrNoLists) {
			formatter.PrintColored(err.Error(), formatter.Yellow)
			return
		} else if errors.Is(err, l.ErrNoListsExist) {
			formatter.PrintColored("\nList successfully deleted. No lists remain.", formatter.Yellow)
			return
		} else {
			log.Fatal(err)
		}
	}

	handleGetAllLists(listsService)

	formatter.PrintColored("\nList successfully deleted.", formatter.Yellow)
}
