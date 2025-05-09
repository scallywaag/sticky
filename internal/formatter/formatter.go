package formatter

import (
	"example/sticky/internal/colors"
	"fmt"
	"strconv"
)

func Print(content string, curr int, last int) {
	d := getPadDiff(strconv.Itoa(curr), strconv.Itoa(last))
	s := fmt.Sprintf("%s%d - %s\n", colors.Blue, curr, content)
	ps := leftPad(s, d)
	fmt.Printf("%s", ps)
}

func getPadDiff(a string, b string) int {
	la := len(a)
	lb := len(b)
	diff := 0

	if la < lb {
		diff = lb - la
	}

	return diff
}

func leftPad(s string, l int) string {
	ps := ""

	for range l {
		ps += " "
	}

	return ps + s
}
