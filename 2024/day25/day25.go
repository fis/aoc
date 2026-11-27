// Package day25 solves AoC 2024 day 25.
package day25

import (
	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2024, 25, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	locks, keys := parseInput(lines)
	part1 := countPossible(locks, keys)
	return glue.Ints(part1), nil
}

func countPossible(locks, keys []schematic) (count int) {
	for _, lock := range locks {
		for _, key := range keys {
			if lock&key == 0 {
				count++
			}
		}
	}
	return count
}

type schematic uint64

const (
	schW = 5
	schH = 7
)

func parseInput(lines []string) (locks, keys []schematic) {
	for i := 0; i < len(lines); i += schH + 1 {
		var sch schematic
		for j := range schH {
			for _, b := range []byte(lines[i+j]) {
				sch <<= 1
				if b == '#' {
					sch |= 1
				}
			}
		}
		if sch&((1<<schW)-1) == 0 {
			locks = append(locks, sch)
		} else {
			keys = append(keys, sch)
		}
	}
	return locks, keys
}
