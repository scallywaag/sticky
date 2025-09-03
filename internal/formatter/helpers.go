package formatter

import (
	"fmt"
	"strings"
)

func getPadDiff(a, b string) int {
	if len(a) < len(b) {
		return len(b) - len(a)
	}

	return 0
}

func leftPad(str string, spaces int) string {
	return strings.Repeat(" ", spaces) + str
}

func formatLine(content string, color Color, status ContentStatus) string {
	formatted := fmt.Sprintf("%s%s%s", color, content, Reset)

	switch status {
	case StatusCross:
		formatted = fmt.Sprintf("%s%s", Strike, formatted)
	case StatusBold:
		formatted = fmt.Sprintf("%s%s", Bold, formatted)
	}

	return formatted
}
