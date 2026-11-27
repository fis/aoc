// Package day24 solves AoC 2021 day 24.
package day24

import (
	"strconv"
	"strings"

	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2021, 24, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	coeffs := parseCoeffs(lines)
	p1 := coeffs.findValid(true)
	p2 := coeffs.findValid(false)
	return []string{p1, p2}, nil
}

var (
	bigFirst   = [9]int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	smallFirst = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
)

func (coeffs *monadCoeffs) findValid(wantBig bool) (num string) {
	var digits [ndigits]byte
	pairs := coeffs.pairs()
	for _, pair := range pairs {
		d := coeffs[pair[0]].b + coeffs[pair[1]].a
		var d1, d2 int
		if wantBig {
			d1, d2 = min(9, 9-d), min(9, 9+d)
		} else {
			d1, d2 = max(1, 1-d), max(1, 1+d)
		}
		digits[pair[0]] = byte('0' + d1)
		digits[pair[1]] = byte('0' + d2)
	}
	return string(digits[:])
}

func (coeffs *monadCoeffs) pairs() (pairs [][2]int) {
	var stack []int
	for i, c := range coeffs {
		if !c.pop {
			stack = append(stack, i)
		} else if len(stack) > 0 {
			pairs = append(pairs, [2]int{stack[len(stack)-1], i})
			stack = stack[:len(stack)-1]
		} else {
			panic("coefficients do not pair up: stack empty")
		}
	}
	if len(stack) > 0 {
		panic("coefficients do not pair up: leftovers")
	}
	return pairs
}

const ndigits = 14

type monadCoeff struct {
	pop  bool
	a, b int
}

type monadCoeffs [ndigits]monadCoeff

func parseCoeffs(lines []string) (coeffs monadCoeffs) {
	for i := range ndigits {
		base := i * 18
		if words := strings.Split(lines[base+4], " "); words[2] == "26" {
			coeffs[i].pop = true
		}
		words := strings.Split(lines[base+5], " ")
		coeffs[i].a, _ = strconv.Atoi(words[2])
		words = strings.Split(lines[base+15], " ")
		coeffs[i].b, _ = strconv.Atoi(words[2])
	}
	return coeffs
}
