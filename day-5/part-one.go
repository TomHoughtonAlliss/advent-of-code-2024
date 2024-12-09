package dayfive

import (
	"fmt"
	"strconv"
	"strings"

	"www.advent.com/helpers"
	"www.advent.com/opener"
)

type checker struct {
	rules map[string][]string
}

func (c *checker) getSharedRule(a string, b string) string {
	aSet := c.rules[a]
	bSet := c.rules[b]

	intersection := helpers.GetIntersection(aSet, bSet)

	return intersection[0]
}

func (c *checker) ruleBroken(a string, i int, b string, j int) bool {
	if i == j {
		return false
	} else if i < j {
		r := c.getSharedRule(a, b)
		if r == fmt.Sprintf("%s|%s", b, a) {
			return true
		}
	} else {
		r := c.getSharedRule(a, b)
		if r == fmt.Sprintf("%s|%s", a, b) {
			return true
		}
	}

	return false
}

func (c *checker) updateIsValid(update string) bool {
	split := strings.Split(update, ",")
	for i, a := range split {
		for j, b := range split {
			if c.ruleBroken(a, i, b, j) {
				return false
			}
		}
	}

	return true
}

func (c *checker) init(rules []string) {
	for _, rule := range rules {
		split := strings.Split(rule, "|")
		a := split[0]
		b := split[1]

		c.rules[a] = append(c.rules[a], rule)
		c.rules[b] = append(c.rules[b], rule)
	}
}

func splitLines(lines []string) ([]string, []string) {
	for i, line := range lines {
		if line == "" {
			return lines[:i], lines[i + 1:]
		}
	}

	return []string{}, []string{}
}

func PartOne() {
	lines := opener.MustReadFile("./day-5/input.txt")

	rules, updates := splitLines(lines)

	c := checker{
		rules: map[string][]string{},
	}

	c.init(rules)

	total := 0
	for _, update := range updates {
		if c.updateIsValid(update) {
			split := strings.Split(update, ",")
			midpoint := len(split) / 2
			middle := split[midpoint]
			v, _ := strconv.Atoi(middle)
			total += v
		}
	}

	fmt.Println(total)
}