// Package day09 solves AoC 2019 day 9.
package day09

import (
	"strconv"

	"github.com/fis/aoc2019-go/intcode"
)

func Solve(path string) ([]string, error) {
	prog, err := intcode.Load(path)
	if err != nil {
		return nil, err
	}

	vm := intcode.VM{}
	vm.Load(prog)
	p1 := vm.Run([]int64{1})
	vm.Load(prog)
	p2 := vm.Run([]int64{2})

	return []string{
		strconv.FormatInt(p1[len(p1)-1], 10),
		strconv.FormatInt(p2[len(p2)-1], 10),
	}, nil
}
