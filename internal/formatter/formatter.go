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
	curr := strconv.Itoa(currId)
	last := strconv.Itoa(lastId)

	padDiff := getPadDiff(curr, last)
	maxContent := 80 - len(last)

	// Split content into lines with word wrapping
	lines := []string{}
	line := ""
	word := ""

	for _, c := range content {
		if c != ' ' {
			word += string(c)
		} else {
			if len(line)+len(word)+1 > maxContent {
				lines = append(lines, line)
				line = word + " "
			} else {
				line += word + " "
			}
			word = ""
		}
	}

	// Handle last word
	if word != "" {
		if len(line)+len(word) > maxContent {
			lines = append(lines, line)
			line = word
		} else {
			line += word
		}
	}

	// Push the final line if it has content
	if line != "" {
		lines = append(lines, line)
	}

	// Print first line with proper padding
	if len(lines) > 0 {
		formatted := fmt.Sprintf("%s%s - %s%s\n", color, curr, lines[0], Reset)
		padded := leftPad(formatted, padDiff)
		fmt.Print(padded)
	}

	// Print remaining lines, aligned
	for i := 1; i < len(lines); i++ {
		formatted := fmt.Sprintf("%s   %s%s\n", color, lines[i], Reset)
		padded := leftPad(formatted, getPadDiff("", last))
		fmt.Print(padded)
	}
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
