// Package day21 solves AoC 2019 day 21.
package day21

import (
	"strconv"

	"github.com/fis/aoc2019-go/intcode"
)

var input1 = []string{
	// jump if hole in any next three cells
	"NOT A T",
	"NOT T T",
	"AND B T",
	"AND C T",
	"NOT T J",
	// but only if ground available four cells in
	"AND D J",
	"WALK",
}

var input2 = []string{
	// jump if hole in any next three cells & if ground available
	"NOT A T",
	"NOT T T",
	"AND B T",
	"AND C T",
	"NOT T J",
	"AND D J",
	// don't jump if it's a trap
	"NOT E T",
	"NOT T T",
	"OR H T",
	"AND T J",
	"RUN",
}

func Solve(path string) ([]string, error) {
	prog, err := intcode.Load(path)
	if err != nil {
		return nil, err
	}

	p1 := run(prog, input1)
	p2 := run(prog, input2)

	return []string{strconv.FormatInt(p1, 10), strconv.FormatInt(p2, 10)}, nil
}

func run(prog []int64, input []string) int64 {
	out, _ := intcode.Run(prog, unlines(input))
	return out[len(out)-1]
}

// TODO: maybe make this available as a utility in intcode package.
func unlines(lines []string) []int64 {
	var out []int64
	for _, line := range lines {
		for _, r := range line {
			out = append(out, int64(r))
		}
		out = append(out, '\n')
	}
	return out
}
