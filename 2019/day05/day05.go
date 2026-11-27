// Package day05 solves AoC 2019 day 5.
package day05

import (
	"github.com/fis/aoc/2019/intcode"
	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2019, 5, intcode.Solver(solve))
}

func solve(prog []int64) ([]int64, error) {
	return []int64{part1(prog), part2(prog)}, nil
}

func part1(prog []int64) int64 {
	out, _ := intcode.Run(prog, []int64{1})
	for _, i := range out {
		util.Diagf("out: %d\n", i)
	}
	return out[len(out)-1]
}

func part2(prog []int64) int64 {
	out, _ := intcode.Run(prog, []int64{5})
	return out[0]
}
