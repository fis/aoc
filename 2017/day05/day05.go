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

// Package day05 solves AoC 2017 day 5.
package day05

import "github.com/fis/aoc/glue"

func init() {
	glue.RegisterSolver(2017, 5, glue.IntSolver(solve))
}

func solve(input []int) ([]string, error) {
	p1 := part1(append(input[:0:0], input...))
	p2 := part2(append(input[:0:0], input...))
	return glue.Ints(p1, p2), nil
}

func part1(offsets []int) (steps int) {
	ip := 0
	for ip >= 0 && ip < len(offsets) {
		off := offsets[ip]
		offsets[ip]++
		ip += off
		steps++
	}
	return steps
}

func part2(offsets []int) (steps int) {
	ip := 0
	for ip >= 0 && ip < len(offsets) {
		off := offsets[ip]
		if off >= 3 {
			offsets[ip]--
		} else {
			offsets[ip]++
		}
		ip += off
		steps++
	}
	return steps
}
