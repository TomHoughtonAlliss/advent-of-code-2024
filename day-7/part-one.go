package dayseven

import (
	"fmt"
	"strconv"
	"strings"

	"www.advent.com/opener"
)

func parse(line string) (int, []int) {
	split := strings.Split(line, ": ")
	total, _ := strconv.Atoi(split[0])

	strOperands := strings.Split(split[1], " ")
	operands := make([]int, len(strOperands))
	for i, val := range strOperands {
		op, _ := strconv.Atoi(val)
		operands[i] = op
	}

	return total, operands
}

func permuteOperators(length int, symbols []string) []string {
	var generate func(int, string)
	var results []string

	generate = func(n int, current string) {
		if n == 0 {
			results = append(results, current)
			return
		}
		for _, symbol := range symbols {
			generate(n - 1, current + symbol)
		}
	}

	results = []string{}
	generate(length, "")
	return results
}

func calc(f int, s int, op string) int {
	switch op {
	case "+":
		return f + s
	case "*":
		return f * s
	case "|": // alias for ||
		out := strconv.Itoa(f) + strconv.Itoa(s)
		i, _ := strconv.Atoi(out)
		return i
	default:
		panic("unknown operator: " + op)
	}
}

func opSetWorks(target int, operands []int, operators string) bool {
	acc := operands[0]
	for i, b := range operators {
		o := string(b)
		v := operands[i + 1]

		acc = calc(acc, v, o)
	}

	return acc == target
}

func works(target int, operands []int, operationOrders []string) bool {
	for _, operationSet := range operationOrders {
		if opSetWorks(target, operands, operationSet) {
			return true
		}
	}

	return false
}

func PartOne() {
	lines := opener.MustReadFile("./day-7/input.txt")

	total := 0
	for _, line := range lines {
		target, operands := parse(line)

		operationOrders := permuteOperators(len(operands) - 1, []string{"+", "*", "|"})
		
		if works(target, operands, operationOrders) {
			total += target
		}
	}

	fmt.Println(total)
}