// Package days contains the glue and tests for all AoC 2019 days.
package days

import (
	"fmt"

	"github.com/fis/aoc2019-go/day01"
	"github.com/fis/aoc2019-go/day05"
	"github.com/fis/aoc2019-go/day18"
	"github.com/fis/aoc2019-go/day19"
)

var solvers = map[int]func(string) ([]string, error){
	1:  day01.Solve,
	5:  day05.Solve,
	18: day18.Solve,
	19: day19.Solve,
}

func Solve(day int, path string) ([]string, error) {
	solver, ok := solvers[day]
	if !ok {
		return nil, fmt.Errorf("unknown day: %d", day)
	}
	return solver(path)
}
