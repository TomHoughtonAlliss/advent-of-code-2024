package dayone

import (
	"fmt"

	"www.advent.com/opener"
)

func getFrequencyCount(nums []int) map[int]int {
	freq := make(map[int]int)
	for _, n := range nums {
		_, ok := freq[n]
		if ok {
			freq[n] += 1
		} else {
			freq[n] = 1
		}
	}
	return freq
}

func PartTwo() {
	lines := opener.MustReadFile("./day-1/input.txt")

	left, right := getIntArrays(lines)

	freq := getFrequencyCount(right)

	total := 0
	for _, n := range left {
		count := freq[n]
		total += n * count
	}

	fmt.Println(total)
}