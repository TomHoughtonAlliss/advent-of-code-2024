package dayfour

import (
	"fmt"

	"www.advent.com/opener"
)

func getLines() []string {
	return opener.MustReadFile("./day-4/input.txt")
}

var (
	directions = [][][]int{
		{{1, 0}, {2, 0}, {3, 0}},       // right
		{{-1, 0}, {-2, 0}, {-3, 0}},    // left
		{{0, 1}, {0, 2}, {0, 3}},       // up
		{{0, -1}, {0, -2}, {0, -3}},    // down
		{{1, 1}, {2, 2}, {3, 3}},       // right up
		{{-1, 1}, {-2, 2}, {-3, 3}},    // left-up
		{{1, -1}, {2, -2}, {3, -3}},    // right down
		{{-1, -1}, {-2, -2}, {-3, -3}}, // left-down
	}
)

func validCoords(lines []string, x int, y int) bool {
	if x < 0 || x >= len(lines[0]) {
		return false
	}

	if y < 0 || y >= len(lines) {
		return false
	}

	return true
}

func checkCoordinate(lines []string, x int, y int, coordSet [][]int) bool {
	match := "MAS"
	for i, coords := range coordSet {
		dx, dy := coords[0], coords[1]
		newX := x + dx
		newY := y + dy
		if !(validCoords(lines, newX, newY) && lines[newY][newX] == match[i]) {
			return false
		}
	}

	return true
}

func checkXPosition(lines []string, x int, y int) int {
	count := 0
	
	for _, direction := range directions {
		if checkCoordinate(lines, x, y, direction) {
			count += 1
		}
	}

	return count
}

func PartOne() {
	lines := getLines()

	total := 0
	for y, row := range lines {
		for x, char := range row {
			if string(char) == "X" {
				total += checkXPosition(lines, x, y)
			}
		}
	}

	fmt.Println(total)
}
