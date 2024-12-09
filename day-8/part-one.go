package dayeight

import (
	"fmt"
	"strings"

	"www.advent.com/opener"
)

type cell struct {
	isAntinode bool
	signal     string
	x          int
	y          int
}

type grid struct {
	rows [][]cell
}

func (g *grid) print() {
	for _, row := range g.rows {
		for _, cell := range row {
			if cell.signal != "." {
				fmt.Print(cell.signal)
			} else if cell.isAntinode {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g *grid) outOfBounds(x int, y int) bool {
	if y < 0 || y >= len(g.rows) {
		return true
	}

	if x < 0 || x >= len(g.rows[0]) {
		return true
	}

	return false
}

func (g *grid) setAntinode(x int, y int) {
	if !g.outOfBounds(x, y) {
		g.rows[y][x].isAntinode = true
	}
}

func (g *grid) makeAntinodes(c cell) {
	for _, row := range g.rows {
		for _, cell := range row {
			if areCandidates(c, cell) {
				dx := c.x - cell.x
				dy := c.y - cell.y

				g.setAntinode(c.x + dx, c.y + dy)
				g.setAntinode(cell.x - dx, cell.y - dy)
			}
		}
	}
}

func (g *grid) countAntinodes() int {
	count := 0
	
	for _, row := range g.rows {
		for _, cell := range row {
			if cell.isAntinode {
				count += 1
			}
		}
	}

	return count
}

func sameCell(c1 cell, c2 cell) bool {
	return c1.x == c2.x && c1.y == c2.y
}

func sameSignal(c1 cell, c2 cell) bool {
	return c1.signal != "." && (c1.signal == c2.signal)
}

func areCandidates(c1 cell, c2 cell) bool {
	return !sameCell(c1, c2) && sameSignal(c1, c2)
}

func initGrid(lines []string) grid {
	rows := make([][]cell, len(lines))

	for i, line := range lines {
		cells := make([]cell, len(line))
		freqs := strings.Split(line, "")

		for j, f := range freqs {
			cells[j] = cell{
				signal: f,
				x:      j,
				y:      i,
			}
		}
		rows[i] = cells
	}

	return grid{
		rows: rows,
	}
}

func PartOne() {
	lines := opener.MustReadFile("./day-8/input.txt")

	grid := initGrid(lines)

	for _, row := range grid.rows {
		for _, cell := range row {
			grid.makeAntinodes(cell)
		}
	}

	grid.print()
	fmt.Println(grid.countAntinodes())
}
