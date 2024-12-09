package dayone

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"www.advent.com/opener"
)

func toInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func getIntArrays(lines []string) ([]int, []int) {
	left := make([]int, len(lines))
	right := make([]int, len(lines))
	for i, line := range lines {
		split := strings.Split(line, " ")
		
		left[i] = toInt(split[0])
		right[i] = toInt(split[len(split) - 1])
	}

	return left, right
}

func PartOne() {
	lines := opener.MustReadFile("./day-1/input.txt")

	left, right := getIntArrays(lines)

	slices.Sort(left)
	slices.Sort(right)

	var total float64 = 0
	for i  := range left {
		total += math.Abs(float64(left[i] - right[i]))
	}

	fmt.Println(int(total))
}