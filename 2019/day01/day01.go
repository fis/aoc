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

// Package day01 solves AoC 2019 day 1.
package day01

import (
	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2019, 1, glue.IntSolver(solve))
}

func solve(modules []int) ([]int, error) {
	return []int{part1(modules), part2(modules)}, nil
}

func part1(modules []int) int {
	sum := 0
	for _, m := range modules {
		sum += moduleFuel(m)
	}
	return sum
}

func moduleFuel(w int) int {
	return w/3 - 2
}

func part2(modules []int) int {
	sum := 0
	for _, m := range modules {
		sum += moduleFuelComplete(m)
	}
	return sum
}

func moduleFuelComplete(w int) int {
	total := 0
	for {
		f := moduleFuel(w)
		if f <= 0 {
			return total
		}
		total += f
		w = f
	}
}
