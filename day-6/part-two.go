package daysix

import (
	"fmt"

	"www.advent.com/opener"
)

func (r *room) beenHere(x int, y int) bool {
	c := r.cell(x, y)

	for _, d := range c.directions {
		if d == r.direction {
			return true
		}
	}

	return false
}

func (r *room) movePlus() (outside bool, looped bool) {
	dx, dy := toVector(r.direction)

	x := r.guardX + dx
	y := r.guardY + dy

	if r.outOfBounds(x, y) {
		return true, false
	} else if r.beenHere(x, y) {
		return false, true
	} else if r.cell(x, y).symbol == "#" {
		r.turn()

		if r.beenHere(r.guardX, r.guardY) {
			return false, true
		}

	} else {
		r.rows[r.guardY][r.guardX].directions = append(r.rows[r.guardY][r.guardX].directions, r.direction)

		r.guardX = x
		r.guardY = y
	}

	return outside, looped
}

func PartTwo() {
	lines := opener.MustReadFile("./day-6/input.txt")

	total := 0
	for y, line := range lines {
		for x := range line {
			r := initGrid(lines)
			if r.rows[y][x].symbol != "#" { // we're currently stood on a non-obstacle square
				fmt.Print("Trying: ", x, y, "\t")
				old := r.rows[y][x].symbol
				r.rows[y][x].symbol = "#"

				looped := false
				outside := false
				for !(looped || outside) {
					outside, looped = r.movePlus()
				}

				if looped {
					total += 1
					fmt.Println("Looped")
				} else {
					fmt.Println()
				}

				r.rows[y][x].symbol = old
			}
		}
	}

	fmt.Println(total)
}