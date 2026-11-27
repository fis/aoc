// Package day03 solves AoC 2025 day 3.
package day03

import (
	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util/fn"
)

func init() {
	glue.RegisterSolver(2025, 3, glue.LineSolver(solve))
}

func solve(banks []string) ([]string, error) {
	p1 := fn.SumF(banks, maxJoltage2)
	p2 := fn.SumF(banks, func(bank string) int { return maxJoltageN(bank, 12) })
	return glue.Ints(p1, p2), nil
}

func maxJoltage2(bank string) int {
	n := len(bank)
	if n < 2 {
		return -1
	}
	i1, d1 := maxDigit(bank[:n-1])
	_, d2 := maxDigit(bank[i1+1:])
	return 10*d1 + d2
}

func maxJoltageN(bank string, c int) int {
	n := len(bank)
	if n < c {
		return -1
	}
	at, joltage := 0, 0
	for ; c > 0; c-- {
		i, d := maxDigit(bank[at : n-(c-1)])
		joltage = 10*joltage + d
		at += i + 1
	}
	return joltage
}

func maxDigit(s string) (maxI, maxD int) {
	for i, b := range []byte(s) {
		if d := int(b - '0'); d > maxD {
			maxI, maxD = i, d
		}
	}
	return
}
