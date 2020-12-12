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

// Package day17 solves AoC 2019 day 17.
package day17

import (
	"github.com/fis/aoc-go/glue"
	"github.com/fis/aoc-go/intcode"
	"github.com/fis/aoc-go/util"
)

func init() {
	glue.RegisterSolver(2019, 17, intcode.Solver(solve))
}

var input = []string{
	"A,A,B,C,B,C,B,C,C,A",
	"R,8,L,4,R,4,R,10,R,8",
	"L,12,L,12,R,8,R,8",
	"R,10,R,4,R,4",
	"n",
}

func solve(prog []int64) ([]int64, error) {
	level := capture(prog)
	p1 := crosses(level)

	prog[0] = 2
	vm := intcode.VM{}
	vm.Load(prog)
	out := vm.Run(unlines(input))

	return []int64{int64(p1), out[len(out)-1]}, nil
}

func capture(prog []int64) *util.Level {
	out, _ := intcode.Run(prog, nil)
	level := util.ParseLevelString("", '.')
	x, y := 0, 0
	for _, v := range out {
		switch v {
		case 10:
			x, y = 0, y+1
		case '#':
			level.Set(x, y, '#')
			fallthrough
		default:
			x++
		}
	}
	return level
}

func crosses(level *util.Level) int {
	sum := 0
	level.Range(func(x, y int, _ byte) {
		for _, n := range (util.P{x, y}).Neigh() {
			if level.At(n.X, n.Y) != '#' {
				return
			}
		}
		sum += x * y
	})
	return sum
}

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
