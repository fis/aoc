// Package day01 solves AoC 2018 day 1.
package day01

import (
	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2018, 1, glue.IntSolver(solve))
}

func solve(input []int) ([]string, error) {
	return glue.Ints(sum(input), findRep(input)), nil
}

func sum(changes []int) (final int) {
	for _, c := range changes {
		final += c
	}
	return final
}

func findRep(changes []int) int {
	cur, seen := 0, map[int]struct{}{0: {}}
	for {
		for _, c := range changes {
			cur += c
			if _, ok := seen[cur]; ok {
				return cur
			}
			seen[cur] = struct{}{}
		}
	}
}
