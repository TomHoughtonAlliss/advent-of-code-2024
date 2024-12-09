package daytwo

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"www.advent.com/opener"
)

func toIntArray(chars []string) []int {
    ints := make([]int, len(chars))
    for i, c := range chars {
        n, _ := strconv.Atoi(c)
        ints[i] = n
    }

    return ints
}

func toIntArrays(stringArray []string) [][]int {
    intArrays := make([][]int, len(stringArray))
    for i, s := range stringArray {
        split := strings.Split(s, " ")
        intArrays[i] = toIntArray(split)
    }
    return intArrays
}

func getDifferences(nums []int) []int {
    difs := make([]int, len(nums) - 1)
    for i := range nums {
        if i != 0 {
            difs[i - 1] = nums[i] - nums[i - 1]
        }
    }
    return difs
}

func areSamePolarity(nums []int) bool {
    first := nums[0]
    for _, n := range nums {
        if first * n < 0 {
            return false
        }
    }

    return true
}

func areInAcceptableBounds(nums []int) bool {
    for _, n := range nums {
        if math.Abs(float64(n)) > 3 || math.Abs(float64(n)) < 1 {
            return false
        }
    }

    return true
} 

func PartOne() {
	lines := opener.MustReadFile("./day-2/input.txt")
    intArrays := toIntArrays(lines)

    difs := make([][]int, len(intArrays))
    for i, a := range intArrays {
        difs[i] = getDifferences(a)
    }

    count := 0
    for _, d := range difs {
        if areSamePolarity(d) {
            if areInAcceptableBounds(d) {
                count += 1
            }
        } 
    }

    fmt.Println(count)
}