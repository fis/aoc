// Package day04 solves AoC 2019 day 4.
package day04

import (
	"fmt"
	"strconv"

	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2019, 4, glue.IntSolver(solve))
}

func solve(inputRange []int) ([]string, error) {
	if len(inputRange) != 2 {
		return nil, fmt.Errorf("expected [low, high], got %v", inputRange)
	}

	a1, a2 := 0, 0
	for n := inputRange[0]; n <= inputRange[1]; n++ {
		digits := strconv.Itoa(n)
		if validate1([]byte(digits)) {
			a1++
		}
		if validate2([]byte(digits)) {
			a2++
		}
	}

	return glue.Ints(a1, a2), nil
}

func validate1(digits []byte) bool {
	double := false
	for i := range len(digits) - 1 {
		a, b := digits[i], digits[i+1]
		if b < a {
			return false
		}
		if b == a {
			double = true
		}
	}
	return double
}

func validate2(digits []byte) bool {
	double := false
	for i := range len(digits) - 1 {
		a, b := digits[i], digits[i+1]
		if b < a {
			return false
		}
		if b == a && (i == 0 || digits[i-1] != a) && (i+2 == len(digits) || digits[i+2] != b) {
			double = true
		}
	}
	return double
}
