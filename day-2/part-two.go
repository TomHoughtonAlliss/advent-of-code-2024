package daytwo

import (
	"fmt"

	"www.advent.com/opener"
)

func getSubsets[T any](arr []T) [][]T {
	subsets := make([][]T, len(arr))
	for i := range arr {
		subset := make([]T, 0)
		subset = append(subset, arr[:i]...)
		subset = append(subset, arr[i+1:]...)

	subsets[i] = subset
	}
	
	return subsets
}

func checkArray(difs []int) bool {
	if !areSamePolarity(difs) {
		return false
	}
	if !areInAcceptableBounds(difs) {
		return false
	} 

	return true
}

func checkSetOfDifs(difSet [][]int) bool {
	for _, dif := range difSet {
		if checkArray(dif) {
			return true
		}
	}

	return false
}

func getSetOfDifs(intArray []int) [][]int {
	subsets := getSubsets(intArray)

	difs := make([][]int, len(subsets))
	for i, s := range subsets {
		difs[i] = getDifferences(s)
	}

	return difs
}

func PartTwo() {
	lines := opener.MustReadFile("./day-2/input.txt")
	intArrays := toIntArrays(lines)
	
	count := 0
	for _, a := range intArrays {
		difs := getSetOfDifs(a)
		if checkSetOfDifs(difs) {
			count += 1
		}
	}

    fmt.Println(count)
}