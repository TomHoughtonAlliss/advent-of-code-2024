package dayeight

import (
	"fmt"

	"www.advent.com/opener"
)

func (g *grid) makeLongAntinodes(c cell) {
	for _, row := range g.rows {
		for _, cell := range row {
			if areCandidates(c, cell) {
				g.setAntinode(c.x, c.y)
				g.setAntinode(cell.x, cell.y)

				dx := c.x - cell.x
				dy := c.y - cell.y

				accX := c.x + dx
				accY := c.y + dy
				for !g.outOfBounds(accX, accY) {
					g.setAntinode(accX, accY)

					accX += dx
					accY += dy
				}

				accX = cell.x - dx
				accY = cell.y - dy
				for !g.outOfBounds(accX, accY) {
					g.setAntinode(accX, accY)

					accX -= dx
					accY -= dy
				}
			}
		}
	}
}

func PartTwo() {
	lines := opener.MustReadFile("./day-8/input.txt")

	grid := initGrid(lines)

	for _, row := range grid.rows {
		for _, cell := range row {
			grid.makeLongAntinodes(cell)
		}
	}

	grid.print()
	fmt.Println(grid.countAntinodes())
}
