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

// Package day21 solves AoC 2019 day 21.
package day21

import (
	"github.com/fis/aoc-go/glue"
	"github.com/fis/aoc-go/intcode"
)

func init() {
	glue.RegisterSolver(2019, 21, intcode.Solver(solve))
}

var input1 = []string{
	// jump if hole in any next three cells
	"NOT A T",
	"NOT T T",
	"AND B T",
	"AND C T",
	"NOT T J",
	// but only if ground available four cells in
	"AND D J",
	"WALK",
}

var input2 = []string{
	// jump if hole in any next three cells & if ground available
	"NOT A T",
	"NOT T T",
	"AND B T",
	"AND C T",
	"NOT T J",
	"AND D J",
	// don't jump if it's a trap
	"NOT E T",
	"NOT T T",
	"OR H T",
	"AND T J",
	"RUN",
}

func solve(prog []int64) ([]int64, error) {
	p1 := run(prog, input1)
	p2 := run(prog, input2)
	return []int64{p1, p2}, nil
}

func run(prog []int64, input []string) int64 {
	out, _ := intcode.Run(prog, unlines(input))
	return out[len(out)-1]
}

// TODO: maybe make this available as a utility in intcode package.
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
