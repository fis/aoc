// Package day05 solves AoC 2019 day 5.
package day05

import (
	"strconv"

	"github.com/fis/aoc2019-go/intcode"
	"github.com/fis/aoc2019-go/util"
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
