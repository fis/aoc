// Package days contains the glue and tests for all AoC 2019 days.
package days

import (
	"fmt"

	"github.com/fis/aoc2019-go/day01"
	"github.com/fis/aoc2019-go/day02"
	"github.com/fis/aoc2019-go/day03"
	"github.com/fis/aoc2019-go/day04"
	"github.com/fis/aoc2019-go/day05"
	"github.com/fis/aoc2019-go/day06"
	"github.com/fis/aoc2019-go/day07"
	"github.com/fis/aoc2019-go/day08"
	"github.com/fis/aoc2019-go/day09"
	"github.com/fis/aoc2019-go/day10"
	"github.com/fis/aoc2019-go/day11"
	"github.com/fis/aoc2019-go/day12"
	"github.com/fis/aoc2019-go/day13"
	"github.com/fis/aoc2019-go/day14"
	"github.com/fis/aoc2019-go/day15"
	"github.com/fis/aoc2019-go/day16"
	"github.com/fis/aoc2019-go/day17"
	"github.com/fis/aoc2019-go/day18"
	"github.com/fis/aoc2019-go/day19"
	"github.com/fis/aoc2019-go/day20"
	"github.com/fis/aoc2019-go/day21"
	"github.com/fis/aoc2019-go/day22"
)

var solvers = map[int]func(string) ([]string, error){
	1:  day01.Solve,
	2:  day02.Solve,
	3:  day03.Solve,
	4:  day04.Solve,
	5:  day05.Solve,
	6:  day06.Solve,
	7:  day07.Solve,
	8:  day08.Solve,
	9:  day09.Solve,
	10: day10.Solve,
	11: day11.Solve,
	12: day12.Solve,
	13: day13.Solve,
	14: day14.Solve,
	15: day15.Solve,
	16: day16.Solve,
	17: day17.Solve,
	18: day18.Solve,
	19: day19.Solve,
	20: day20.Solve,
	21: day21.Solve,
	22: day22.Solve,
}

func Solve(day int, path string) ([]string, error) {
	solver, ok := solvers[day]
	if !ok {
		return nil, fmt.Errorf("unknown day: %d", day)
	}
	return solver(path)
}
