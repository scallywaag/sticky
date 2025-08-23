package formatter

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	Strike  = "\x1b[9m"
	Red     = "\x1b[31m"
	Green   = "\x1b[32m"
	Yellow  = "\x1b[33m"
	Blue    = "\x1b[34m"
	Default = "\x1b[39m"
	Reset   = "\x1b[0m"
)

func Print(content string, currId int, lastId int, color string, cross bool) {
	curr := strconv.Itoa(currId)
	last := strconv.Itoa(lastId)

	padDiff := getPadDiff(curr, last)
	maxPad := getPadDiff("", last)
	maxContent := 80 - len(last)
	sep := " - "
	sepNewline := strings.Repeat(" ", len(sep))

	lines := []string{}

	// Split content by existing newlines first
	paragraphs := strings.SplitSeq(content, "\n")

	for para := range paragraphs {
		if para == "" {
			lines = append(lines, "")
			continue
		}

		words := strings.Fields(para)
		line := ""

		for _, word := range words {
			if line == "" {
				line = word
			} else if len(line)+1+len(word) > maxContent {
				lines = append(lines, line)
				line = word
			} else {
				line += " " + word
			}
		}

		if line != "" {
			lines = append(lines, line)
		}
	}

	// Print first line with ID + padding
	if len(lines) > 0 {
		formatted := formatLine(lines[0], color, cross)
		newContent := fmt.Sprintf("%s%s%s\n", curr, sep, formatted)
		padded := leftPad(newContent, padDiff)
		fmt.Print(padded)
	}

	// Print remaining lines, aligned under first
	for i := 1; i < len(lines); i++ {
		formatted := formatLine(lines[i], color, cross)
		newContent := fmt.Sprintf("%s%s\n", sepNewline, formatted)
		padded := leftPad(newContent, maxPad)
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

func formatLine(content string, color string, cross bool) string {
	formatted := fmt.Sprintf("%s%s%s", color, content, Reset)
	if cross {
		formatted = fmt.Sprintf("%s%s", Strike, formatted)
	}
	return formatted
}
