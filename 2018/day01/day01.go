// Copyright 2020 Google LLC
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

// Package day01 solves AoC 2018 day 1.
package day01

import (
	"github.com/fis/aoc-go/glue"
)

func init() {
	glue.RegisterSolver(2018, 1, glue.IntSolver(solve))
}

func solve(input []int) ([]int, error) {
	return []int{sum(input), findRep(input)}, nil
}

func sum(changes []int) (final int) {
	for _, c := range changes {
		final += c
	}
	return final
}

func findRep(changes []int) int {
	cur, seen := 0, map[int]struct{}{0: {}}
	for {
		for _, c := range changes {
			cur += c
			if _, ok := seen[cur]; ok {
				return cur
			}
			seen[cur] = struct{}{}
		}
	}
}
