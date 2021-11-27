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

// Package day13 solves AoC 2019 day 13.
package day13

import (
	"github.com/fis/aoc/2019/intcode"
	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2019, 13, intcode.Solver(solve))
}

func solve(prog []int64) ([]int64, error) {
	p1 := part1(prog)
	prog[0] = 2
	p2 := part2(prog)
	return []int64{p1, p2}, nil
}

func part1(prog []int64) int64 {
	out, _ := intcode.Run(prog, nil)
	var blocks int64
	for i := 2; i < len(out); i += 3 {
		if out[i] == 2 {
			blocks++
		}
	}
	return blocks
}

func part2(prog []int64) int64 {
	var (
		vm  intcode.VM
		tok intcode.WalkToken
	)
	var score, ball, paddle int64
	vm.Load(prog)
	for vm.Walk(&tok) {
		if tok.IsInput() {
			switch {
			case paddle < ball:
				tok.ProvideInput(1)
			case paddle > ball:
				tok.ProvideInput(-1)
			default:
				tok.ProvideInput(0)
			}
			continue
		}
		x := tok.ReadOutput()
		// ignore middle output (Y position)
		vm.Walk(&tok)
		vm.Walk(&tok)
		out := tok.ReadOutput()
		switch {
		case x == -1:
			score = out
		case x >= 0 && out == 3:
			paddle = x
		case x >= 0 && out == 4:
			ball = x
		}
	}
	return score
}
