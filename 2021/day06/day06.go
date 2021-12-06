// Copyright 2021 Google LLC
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

// Package day06 solves AoC 2021 day 6.
package day06

import "github.com/fis/aoc/glue"

func init() {
	glue.RegisterSolver(2021, 6, glue.IntSolver(solve))
}

func solve(input []int) ([]string, error) {
	p1 := countFish(input, 80)
	p2 := countFish(input, 256)
	return glue.Ints(p1, p2), nil
}

func countFish(initial []int, days int) (total int) {
	counts := [9]int{}
	for _, i := range initial {
		counts[i]++
	}
	offset := 0
	for t := 0; t < days; t++ {
		counts[(offset+7)%9] += counts[offset]
		offset = (offset + 1) % 9
	}
	for _, c := range counts {
		total += c
	}
	return total
}
