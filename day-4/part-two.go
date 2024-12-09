package dayfour

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	newDirections = [][][]int{
		{
			{1, 1}, // top-right
			{0, 0}, // centre
			{-1, -1}, // bottom-left
		},
		{
			{-1, 1}, // top-left
			{0, 0}, // centre
			{1, -1}, // bottom-right
		},
	}
)

func isValidMAS(s string) bool {
	matched, err := regexp.MatchString(`^(MAS|SAM)$`, s)
	if err != nil {
		panic(err)
	}
	return matched
}

func checkAPosition(lines []string, x int, y int) bool {
	toCheck := make([]string, 2)
	for j, direction := range newDirections {
		extract := make([]string, 3)
		for i, coords := range direction {
			newX, newY := x + coords[0], y + coords[1]
			if !validCoords(lines, newX, newY) {
				return false
			}

			extract[i] = string(lines[newY][newX])
		}
		toCheck[j] = strings.Join(extract, "")
	}

	for _, s := range toCheck {
		if !isValidMAS(s) {
			return false
		}
	}

	return true
}

func PartTwo() {
	lines := getLines()

	total := 0
	for y, row := range lines {
		for x, char := range row {
			if string(char) == "A" {
				if checkAPosition(lines, x, y) {
					total += 1
				}
			}
		}
	}

	fmt.Println(total)
}