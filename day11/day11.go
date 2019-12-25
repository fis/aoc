// Package day11 solves AoC 2019 day 11.
package day11

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
	return append(
		[]string{strconv.Itoa(part1(prog))},
		part2(prog)...), nil
}

func part1(prog []int64) int {
	level := util.ParseLevelString(" ", ' ')
	run(prog, level)
	painted := 0
	level.Range(func(_, _ int, _ byte) { painted++ })
	return painted
}

func part2(prog []int64) []string {
	level := util.ParseLevelString("#", ' ')
	run(prog, level)
	var rows []string
	min, max := level.Bounds()
	w := max.X - min.X + 1
	for y := min.Y; y <= max.Y; y++ {
		row := make([]byte, w)
		for x, i := min.X, 0; x <= max.X; x, i = x+1, i+1 {
			row[i] = level.At(x, y)
		}
		rows = append(rows, string(row))
	}
	return rows
}

func run(prog []int64, level *util.Level) {
	vm := intcode.VM{}
	vm.Load(prog)

	x, y, dx, dy := 0, 0, 0, -1
	var t intcode.WalkToken
	for {
		vm.Walk(&t)
		if t.IsEmpty() {
			return
		}
		if level.At(x, y) == '#' {
			t.ProvideInput(1)
		} else {
			t.ProvideInput(0)
		}
		vm.Walk(&t)
		if t.ReadOutput() == 1 {
			level.Set(x, y, '#')
		} else {
			level.Set(x, y, '.')
		}
		vm.Walk(&t)
		if t.ReadOutput() == 1 {
			dx, dy = -dy, dx
		} else {
			dx, dy = dy, -dx
		}
		x, y = x+dx, y+dy
	}
}
