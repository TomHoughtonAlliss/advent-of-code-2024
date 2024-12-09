package dayfive

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"www.advent.com/opener"
)

func (c *checker) validify(update string) string {
	split := strings.Split(update, ",")
	slices.SortFunc(split, func(a, b string) int {
		if a == b {
			return 0
		} else {
			r := c.getSharedRule(a, b)
			if r == fmt.Sprintf("%s|%s", a, b) {
				return -1
			} else {
				return 1
			}
		}
	})

	return strings.Join(split, ",")
}

func PartTwo() {
	lines := opener.MustReadFile("./day-5/input.txt")

	rules, updates := splitLines(lines)

	c := checker{
		rules: map[string][]string{},
	}

	c.init(rules)

	total := 0
	for _, update := range updates {
		if !c.updateIsValid(update) {
			fmt.Println(update)
			validified := c.validify(update)
			split := strings.Split(validified, ",")
			midpoint := len(split) / 2
			middle := split[midpoint]
			v, _ := strconv.Atoi(middle)
			total += v
		}
	}

	fmt.Println(total)
}