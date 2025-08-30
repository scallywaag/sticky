package presentation

import (
	"errors"
	"log"

	"github.com/highseas-software/sticky/internal/formatter"
	"github.com/highseas-software/sticky/internal/lists"
	l "github.com/highseas-software/sticky/internal/lists"
)

func handleAddList(listName string, listsService *lists.Service) {
	if err := listsService.Add(listName); err != nil {
		log.Fatal(err)
	}

	l, count, err := listsService.GetAll()
	GetAllLists(l, count, err)

	formatter.PrintColored("\nList successfully added.", formatter.Yellow)
}

func GetAllLists(lists []l.List, count int, err error) {
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
