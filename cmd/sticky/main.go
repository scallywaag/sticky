package main

import (
	"fmt"

	"github.com/highseas-software/sticky/internal/database"
	"github.com/highseas-software/sticky/internal/notes"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("\x1b[34mHello World! ...in blue\x1b[0m")

	db := database.InitDb()
	defer db.Close()

	notes.Add("test string", db)
	notes.Add("some more testing", db)
	notes.List(db)
}
