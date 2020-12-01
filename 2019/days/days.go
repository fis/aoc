// Package days contains the glue and tests for all AoC 2019 days.
package days

import (
	"fmt"

	"github.com/fis/aoc-go/2019/day01"
	"github.com/fis/aoc-go/2019/day02"
	"github.com/fis/aoc-go/2019/day03"
	"github.com/fis/aoc-go/2019/day04"
	"github.com/fis/aoc-go/2019/day05"
	"github.com/fis/aoc-go/2019/day06"
	"github.com/fis/aoc-go/2019/day07"
	"github.com/fis/aoc-go/2019/day08"
	"github.com/fis/aoc-go/2019/day09"
	"github.com/fis/aoc-go/2019/day10"
	"github.com/fis/aoc-go/2019/day11"
	"github.com/fis/aoc-go/2019/day12"
	"github.com/fis/aoc-go/2019/day13"
	"github.com/fis/aoc-go/2019/day14"
	"github.com/fis/aoc-go/2019/day15"
	"github.com/fis/aoc-go/2019/day16"
	"github.com/fis/aoc-go/2019/day17"
	"github.com/fis/aoc-go/2019/day18"
	"github.com/fis/aoc-go/2019/day19"
	"github.com/fis/aoc-go/2019/day20"
	"github.com/fis/aoc-go/2019/day21"
	"github.com/fis/aoc-go/2019/day22"
	"github.com/fis/aoc-go/2019/day23"
	"github.com/fis/aoc-go/2019/day24"
	"github.com/fis/aoc-go/2019/day25"
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
	23: day23.Solve,
	24: day24.Solve,
	25: day25.Solve,
}

func Solve(day int, path string) ([]string, error) {
	solver, ok := solvers[day]
	if !ok {
		return nil, fmt.Errorf("unknown day: %d", day)
	}
	return solver(path)
}
