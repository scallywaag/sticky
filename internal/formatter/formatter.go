package formatter

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	Red     = "\x1b[31m"
	Green   = "\x1b[32m"
	Yellow  = "\x1b[33m"
	Blue    = "\x1b[34m"
	Default = "\x1b[39m"
	Reset   = "\x1b[0m"
)

func Print(content string, currId int, lastId int, color string) {
	padDiff := getPadDiff(strconv.Itoa(currId), strconv.Itoa(lastId))
	formattedContent := fmt.Sprintf("%s%d - %s%s\n", color, currId, content, Reset)
	paddedContent := leftPad(formattedContent, padDiff)
	fmt.Printf("%s", paddedContent)
}

func getPadDiff(a, b string) int {
	if len(a) < len(b) {
		return len(b) - len(a)
	}

	return 0
}

func leftPad(str string, spaces int) string {
	return strings.Repeat(" ", spaces) + str
}
