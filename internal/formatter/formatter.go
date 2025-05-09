package formatter

import (
	"fmt"
	"strconv"
)

func Print(content string, firstId int, lastId int) {
	s := ""
	ps := getPaddingSize(firstId, lastId)

	for i := 0; i < ps; i++ {
		s += " "
	}

	fmt.Printf("%s%d - %s\n", s, firstId, content)
}

func getPaddingSize(firstId int, lastId int) int {
	f := len(strconv.Itoa(firstId))
	l := len(strconv.Itoa(lastId))

	diff := 0
	if f < l {
		diff = l - f
	}

	return diff
}
