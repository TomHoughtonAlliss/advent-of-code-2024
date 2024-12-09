package daynine

import (
	"fmt"
	"strconv"

	"www.advent.com/opener"
)

type block struct {
	id     int
	length int
}

func (b *block) toArray() []int {
	arr := make([]int, b.length)
	for i := range b.length {
		arr[i] = b.id
	}
	return arr
}

type disc struct {
	blocks []block
}

func (d *disc) getHightestID() int {
	var highest int
	for _, b := range d.blocks {
		if b.id > highest {
			highest = b.id
		}
	}
	return highest
}

func (d *disc) idToIndex(id int) int {
	for i, b := range d.blocks {
		if b.id == id {
			return i
		}
	}

	return -1
}

func (d *disc) toArray() []int {
	arr := make([]int, 0)
	for _, b := range d.blocks {
		arr = append(arr, b.toArray()...)
	}
	return arr
}

func (d *disc) findFirstGap(length int) int {
	for i, b := range d.blocks {
		if b.length >= length && b.id == -1 {
			return i
		}
	}

	return -1
}

func (d *disc) insert(b block, i int) {
	newBlocks := make([]block, 0)

	newBlocks = append(newBlocks, d.blocks[:i]...)
	newBlocks = append(newBlocks, b)
	newBlocks = append(newBlocks, d.blocks[i:]...)

	d.blocks = newBlocks
}

func (d *disc) swap(blockIndex int) bool {
	b := d.blocks[blockIndex]

	gapIndex := d.findFirstGap(b.length)
	if gapIndex == -1 || gapIndex > blockIndex || b.id == -1 {
		return false
	}

	d.blocks[blockIndex].id, d.blocks[gapIndex].id = d.blocks[gapIndex].id, d.blocks[blockIndex].id

	dif := d.blocks[gapIndex].length - d.blocks[blockIndex].length

	var newBlock block
	var add bool
	if dif > 0 {
		newBlock = block{
			id:     -1,
			length: dif,
		}
		add = true
	}

	d.blocks[gapIndex].length = d.blocks[blockIndex].length

	if add {
		d.insert(newBlock, gapIndex+1)
	}

	return true
}

func initDisc(line []int) disc {
	blocks := make([]block, len(line))

	for i, v := range line {
		var b block

		if i%2 == 0 {
			id := i / 2
			b = block{
				id:     id,
				length: v,
			}
		} else {
			b = block{
				id:     -1,
				length: v,
			}
		}

		blocks[i] = b
	}

	return disc{
		blocks: blocks,
	}
}

func toIntArray(line string) []int {
	ints := make([]int, len(line))
	for i, c := range line {
		s := string(c)
		v, _ := strconv.Atoi(s)
		ints[i] = v
	}

	return ints
}

func PartTwo() {
	lines := opener.MustReadFile("./day-9/input.txt")

	intLine := toIntArray(lines[0])

	d := initDisc(intLine)

	max := d.getHightestID()
	for j := range max + 1 {
		id := max - j
		i := d.idToIndex(id)
		if !(d.blocks[i].id == -1) {
			d.swap(i)
		}
	}

	fmt.Println(calculateChecksum(d.toArray()))
}
