package daythree

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"www.advent.com/opener"
)

type window struct {
	start int
	end int
	data string
}

func (w *window) slide() bool {
	if w.end != len(w.data) - 1 {
		w.start += 1
		w.end += 1

		return true
	} else if w.start != w.end {
		w.start += 1

		return true
	}

	return false
}

func (w *window) show() string {
	return w.data[w.start:w.end]
}

func isMul(s string) bool {
	re := regexp.MustCompile(`^mul\(\d{1,3},\d{1,3}\)`)
	return re.MatchString(s)
}

func parseValues(str string) (int, int) {
	split := strings.Split(str, ",")
	first := strings.Split(split[0], "(")[1]
	second := strings.Split(split[1], ")")[0]

	f, err := strconv.Atoi(first)
	if err != nil {
		panic(err)
	}

	s, err := strconv.Atoi(second)
	if err != nil {
		panic(err)
	}

	return f, s
}

func PartOne() {
	lines := opener.MustReadFile("./day-3/input.txt")

	full := strings.Join(lines, "")

	w := window{
		data: full,
		start: 0,
		end: 12,
	}

	total := 0
	for w.slide() {
		s := w.show()
		fmt.Print(s, " ")
		if isMul(s) {
			x, y := parseValues(s)
			total += x * y
			fmt.Print(x, y)
		}
		fmt.Println()
	}

	fmt.Println(total)
}