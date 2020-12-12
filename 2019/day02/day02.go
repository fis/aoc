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

// Package day02 solves AoC 2019 day 2.
package day02

import (
	"github.com/fis/aoc-go/glue"
	"github.com/fis/aoc-go/intcode"
)

func init() {
	glue.RegisterSolver(2019, 2, intcode.Solver(solve))
}

func solve(prog []int64) ([]int64, error) {
	return []int64{part1(prog), part2(prog)}, nil
}

func part1(prog []int64) int64 {
	return run(12, 2, prog)
}

func part2(prog []int64) int64 {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			if run(noun, verb, prog) == 19690720 {
				return int64(100*noun + verb)
			}
		}
	}
	panic("solution not found")
}

func run(noun, verb int, prog []int64) int64 {
	vm := intcode.VM{}
	vm.Load(prog)
	*vm.Mem(1) = int64(noun)
	*vm.Mem(2) = int64(verb)
	vm.Run(nil)
	return *vm.Mem(0)
}
