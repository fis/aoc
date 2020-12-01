// Package day09 solves AoC 2019 day 9.
package day09

import (
	"strconv"

	"github.com/fis/aoc-go/intcode"
)

func Solve(path string) ([]string, error) {
	prog, err := intcode.Load(path)
	if err != nil {
		return nil, err
	}

	p1, _ := intcode.Run(prog, []int64{1})
	p2, _ := intcode.Run(prog, []int64{2})

	return []string{
		strconv.FormatInt(p1[len(p1)-1], 10),
		strconv.FormatInt(p2[len(p2)-1], 10),
	}, nil
}
