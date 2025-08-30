package presentation

import (
	"errors"
	"log"

	"github.com/highseas-software/sticky/internal/formatter"
	l "github.com/highseas-software/sticky/internal/lists"
)

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
