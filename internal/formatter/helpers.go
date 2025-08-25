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

func formatLine(content string, color string, cross bool) string {
	formatted := fmt.Sprintf("%s%s%s", color, content, Reset)
	if cross {
		formatted = fmt.Sprintf("%s%s", Strike, formatted)
	}
	return formatted
}
