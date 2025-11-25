package flags

import (
	"fmt"
	"os"

	"github.com/scallywaag/sticky/internal/formatter"
)

func printUsageMenu() {
	formatter.ClearScreen()
	fmt.Fprintf(os.Stderr, "Usage of sticky:\n")
	fmt.Fprintf(os.Stderr, "\n* operations\n")
	fmt.Fprintf(os.Stderr, "  -l, -list <listname> string\n\tlist all notes from <listname>\n")
	fmt.Fprintf(os.Stderr, "  -a, -add <content> string\n\tadd note <content>\n")
	fmt.Fprintf(os.Stderr, "  -d, -del <id> int\n\tdelete note by <id>\n")
	fmt.Fprintf(os.Stderr, "  -m, -mut <id> int\n\tmutate existing note <id>\n")

	fmt.Fprintf(os.Stderr, "\n* mutations (toggle)\n\tuseable with -a or -m flag\n")
	fmt.Fprintf(os.Stderr, "  -p, -pin bool\n\tpin note - send to top of list\n")
	fmt.Fprintf(os.Stderr, "  -c, -cross bool\n\tcross note - send to bottom of list\n")
	fmt.Fprintf(os.Stderr, "  -r, -red bool\n\tcolor note red\n")
	fmt.Fprintf(os.Stderr, "  -g, -green bool\n\tcolor note green\n")
	fmt.Fprintf(os.Stderr, "  -b, -blue bool\n\tcolor note blue\n")
	fmt.Fprintf(os.Stderr, "  -y, -yellow bool\n\tcolor note yellow\n")

	fmt.Fprintf(os.Stderr, "\n* lists\n")
	fmt.Fprintf(os.Stderr, "  -ls, -lists bool\n\tshow all existing lists\n")
	fmt.Fprintf(os.Stderr, "  -la, -list-add <listname> string\n\tadd new list <listname>\n")
	fmt.Fprintf(os.Stderr, "  -ld, -list-del <id> int\n\tdelete list by <id>\n")

	fmt.Fprintf(os.Stderr, "\n*\n")
	fmt.Fprintf(os.Stderr, "  -h\n\topen this help menu\n")
}
