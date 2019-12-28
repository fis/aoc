// Package day02 solves AoC 2019 day 2.
package day02

import (
	"strconv"

	"github.com/fis/aoc2019-go/intcode"
)

func Solve(path string) ([]string, error) {
	prog, err := intcode.Load(path)
	if err != nil {
		return nil, err
	}

	return []string{
		strconv.FormatInt(part1(prog), 10),
		strconv.FormatInt(part2(prog), 10),
	}, nil
}

func part1(prog []int64) int64 {
	return run(12, 2, prog)
}

func part2(prog []int64) int64 {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
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
