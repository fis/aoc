// Package day05 solves AoC 2017 day 5.
package day05

import "github.com/fis/aoc/glue"

func init() {
	glue.RegisterSolver(2017, 5, glue.IntSolver(solve))
}

func solve(input []int) ([]string, error) {
	p1 := part1(append(input[:0:0], input...))
	p2 := part2(append(input[:0:0], input...))
	return glue.Ints(p1, p2), nil
}

func part1(offsets []int) (steps int) {
	ip := 0
	for ip >= 0 && ip < len(offsets) {
		off := offsets[ip]
		offsets[ip]++
		ip += off
		steps++
	}
	return steps
}

func part2(offsets []int) (steps int) {
	ip := 0
	for ip >= 0 && ip < len(offsets) {
		off := offsets[ip]
		if off >= 3 {
			offsets[ip]--
		} else {
			offsets[ip]++
		}
		ip += off
		steps++
	}
	return steps
}
