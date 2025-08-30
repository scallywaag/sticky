package presentation

import (
	"errors"
	"log"

	"github.com/highseas-software/sticky/internal/formatter"
	l "github.com/highseas-software/sticky/internal/lists"
)

func handleAddList(listName string, listsService *l.Service) {
	if err := listsService.Add(listName); err != nil {
		log.Fatal(err)
	}

	getAllLists(listsService)

	formatter.PrintColored("\nList successfully added.", formatter.Yellow)
}

func getAllLists(listsService *l.Service) {
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
