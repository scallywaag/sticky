package formatter

import (
	"fmt"
	"strconv"
)

const (
	Red     = "\x1b[31m"
	Green   = "\x1b[32m"
	Yellow  = "\x1b[33m"
	Blue    = "\x1b[34m"
	Default = "\x1b[39m"
	Reset   = "\x1b[0m"
)

func Print(content string, curr int, last int, color string) {
	d := getPadDiff(strconv.Itoa(curr), strconv.Itoa(last))
	s := fmt.Sprintf("%s%d - %s%s\n", color, curr, content, Reset)
	ps := leftPad(s, d)
	fmt.Printf("%s", ps)
}

func getPadDiff(a, b string) int {
	if len(a) < len(b) {
		return len(b) - len(a)
	}

	return 0
}

func leftPad(s string, l int) string {
	ps := ""

	for range l {
		ps += " "
	}

	return ps + s
}
