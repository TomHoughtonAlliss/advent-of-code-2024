package daythree

import (
	"fmt"
	"regexp"

	"strings"

	"www.advent.com/opener"
)

func isDo(s string) bool {
	matched, _ := regexp.MatchString(`^do\(\)`, s)
	return matched
}

func isDont(s string) bool {
	matched, _ := regexp.MatchString(`^don't\(\)`, s)
	return matched
}

func PartTwo() {
	lines := opener.MustReadFile("./day-3/input.txt")

	full := strings.Join(lines, "")

	w := window{
		data: full,
		start: 0,
		end: 12,
	}

	do := true

	total := 0
	for w.slide() {
		s := w.show()
		if isDo(s) {
			do = true
		} else if isDont(s) {
			do = false
		} else if isMul(s) && do {
			x, y := parseValues(s)
			total += x * y
		}
	}

	fmt.Println(total)
}