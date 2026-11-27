// Package day01 solves AoC 2019 day 1.
package day01

import (
	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2019, 1, glue.IntSolver(solve))
}

func solve(modules []int) ([]string, error) {
	return glue.Ints(part1(modules), part2(modules)), nil
}

func part1(modules []int) int {
	sum := 0
	for _, m := range modules {
		sum += moduleFuel(m)
	}
	return sum
}

func moduleFuel(w int) int {
	return w/3 - 2
}

func part2(modules []int) int {
	sum := 0
	for _, m := range modules {
		sum += moduleFuelComplete(m)
	}
	return sum
}

func moduleFuelComplete(w int) int {
	total := 0
	for {
		f := moduleFuel(w)
		if f <= 0 {
			return total
		}
		total += f
		w = f
	}
}
