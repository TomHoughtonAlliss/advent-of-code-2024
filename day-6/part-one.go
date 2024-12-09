package daysix

import (
	"fmt"
	"strings"

	"www.advent.com/opener"
)

type cell struct {
	symbol     string
	directions []string
	loopable   bool
}

type room struct {
	rows      [][]cell
	guardX    int
	guardY    int
	direction string // one of U,R,D,L
}

func (r *room) fill(x int, y int, sym string) {
	r.rows[y][x].symbol = sym
	r.rows[y][x].directions = append(r.rows[y][x].directions, r.direction)
}

func (r *room) countVisited() int {
	count := 0
	for _, row := range r.rows {
		for _, cell := range row {
			if cell.symbol == "X" || cell.symbol == "^" {
				count += 1
			}
		}
	}
	return count
}

func (r *room) outOfBounds(x int, y int) bool {
	if y < 0 || y >= len(r.rows) {
		return true
	}

	if x < 0 || x >= len(r.rows[0]) {
		return true
	}

	return false
}

func (r *room) print() {
	for _, row := range r.rows {
		for _, cell := range row {
			var toPrint string
			if cell.loopable {
				toPrint = "O"
			} else if len(cell.directions) != 0 {
				toPrint = cell.directions[0]
			} else {
				toPrint = cell.symbol
			}
			fmt.Print(toPrint)
		}
		fmt.Println()
	}
}

func (r *room) cell(x int, y int) cell {
	if r.outOfBounds(x, y) {
		return cell{}
	}
	return r.rows[y][x]
}

func (r *room) turn() {
	switch r.direction {
	case "U":
		r.direction = "R"
	case "R":
		r.direction = "D"
	case "D":
		r.direction = "L"
	case "L":
		r.direction = "U"
	default:
		r.direction = ""
	}
}

func (r *room) move() bool {
	dx, dy := toVector(r.direction)

	x := r.guardX + dx
	y := r.guardY + dy

	if r.outOfBounds(x, y) {
		r.print()
		return false
	} else if r.cell(x, y).symbol == "#" {
		r.turn()
	} else {
		r.guardX = x
		r.guardY = y
		r.fill(x, y, "X")
	}

	return true
}

func toVector(dir string) (int, int) {
	switch dir {
	case "U":
		return 0, -1
	case "R":
		return 1, 0
	case "D":
		return 0, 1
	case "L":
		return -1, 0
	default:
		return 0, 0
	}
}

func initGrid(lines []string) room {
	r := room{
		rows: make([][]cell, len(lines)),
	}

	for y, line := range lines {
		split := strings.Split(line, "")
		row := make([]cell, len(split))

		for i, elt := range split {
			row[i].symbol = elt
		}

		r.rows[y] = row

		for x, cell := range row {
			if cell.symbol == "^" {
				r.guardX = x
				r.guardY = y
				r.rows[y][x].directions = append(r.rows[y][x].directions, "U")
			}
		}
	}

	r.direction = "U"

	return r
}

func PartOne() {
	lines := opener.MustReadFile("./day-6/input.txt")

	r := initGrid(lines)

	for r.move() {
	}

	fmt.Println(r.countVisited())
}
