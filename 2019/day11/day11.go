// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package day11 solves AoC 2019 day 11.
package day11

import (
	"strconv"

	"github.com/fis/aoc/2019/intcode"
	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2019, 11, intcode.SolverS(solve))
}

func solve(prog []int64) ([]string, error) {
	return append(
		[]string{strconv.Itoa(part1(prog))},
		part2(prog)...), nil
}

func part1(prog []int64) int {
	level := util.SparseLevel(util.P{0, 0}, ' ')
	run(prog, level)
	painted := 0
	level.Range(func(_, _ int, _ byte) { painted++ })
	return painted
}

func part2(prog []int64) []string {
	level := util.SparseLevel(util.P{0, 0}, ' ')
	level.Set(0, 0, '#')
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
	for vm.Walk(&t) {
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
