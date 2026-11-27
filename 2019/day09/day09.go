// Package day09 solves AoC 2019 day 9.
package day09

import (
	"github.com/fis/aoc/2019/intcode"
	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2019, 9, intcode.Solver(solve))
}

func solve(prog []int64) ([]int64, error) {
	p1, _ := intcode.Run(prog, []int64{1})
	p2, _ := intcode.Run(prog, []int64{2})
	return []int64{p1[len(p1)-1], p2[len(p2)-1]}, nil
}
