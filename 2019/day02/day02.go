// Package day02 solves AoC 2019 day 2.
package day02

import (
	"github.com/fis/aoc/2019/intcode"
	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2019, 2, intcode.Solver(solve))
}

func solve(prog []int64) ([]int64, error) {
	return []int64{part1(prog), part2(prog)}, nil
}

func part1(prog []int64) int64 {
	return run(12, 2, prog)
}

func part2(prog []int64) int64 {
	for noun := range 100 {
		for verb := range 100 {
			if run(noun, verb, prog) == 19690720 {
				return int64(100*noun + verb)
			}
		}
	}
	panic("solution not found")
}

func run(noun, verb int, prog []int64) int64 {
	vm := intcode.VM{}
	vm.Load(prog)
	*vm.Mem(1) = int64(noun)
	*vm.Mem(2) = int64(verb)
	vm.Run(nil)
	return *vm.Mem(0)
}
