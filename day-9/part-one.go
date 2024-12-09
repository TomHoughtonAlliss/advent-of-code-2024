package daynine

import (
	"fmt"
	"strconv"

	"www.advent.com/opener"
)

func getIDs(length int, id int) []int {
	b := make([]int, length)
	for i := range length {
		b[i] = id
	}

	return b
}

func getBlanks(length int) []int {
	b := make([]int, length)
	for i := range length {
		b[i] = -1
	}

	return b
}

func findFirstBlank(arr []int) int {
	for i, elt := range arr {
		if elt == -1 {
			return i
		}
	}

	return -1
}

func findLastDigit(arr []int) int {
	for j := range arr {
		i := len(arr) - j - 1

		if arr[i] != -1 {
			return i
		}
	}

	return -1
}

func swap(arr []int) ([]int, bool) {
	blankIndex := findFirstBlank(arr)
	lastIndex := findLastDigit(arr)

	if blankIndex > lastIndex {
		return arr, false
	}

	arr[blankIndex], arr[lastIndex] = arr[lastIndex], arr[blankIndex]

	return arr, true
}

func calculateChecksum(arr []int) int {
	total := 0
	for i, elt := range arr {
		if elt != -1 {
			total += i * elt
		}
	}

	return total
}

func PartOne() {
	lines := opener.MustReadFile("./day-9/input.txt")
	line := lines[0]

	unpacked := []int{}
	for i, char := range line {
		c := string(char)

		v, _ := strconv.Atoi(c)

		if i % 2 == 0 {
			id := i / 2
			unpacked = append(unpacked, getIDs(v, id)...)
		} else {
			unpacked = append(unpacked, getBlanks(v)...)
		}
	}

	cont := true
	for cont {
		unpacked, cont = swap(unpacked)
	}

	checksum := calculateChecksum(unpacked)

	fmt.Println(checksum)
}
