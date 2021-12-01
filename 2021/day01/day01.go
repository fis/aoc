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

// Package day01 solves AoC 2021 day 1.
package day01

import "github.com/fis/aoc/glue"

func init() {
	glue.RegisterSolver(2021, 1, glue.IntSolver(solve))
}

func solve(depths []int) ([]string, error) {
	p1, p2 := increases(depths)
	return glue.Ints(p1, p2), nil
}

func increases(depths []int) (lag1, lag3 int) {
	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			lag1++
		}
		if i >= 3 && depths[i] > depths[i-3] {
			lag3++
		}
	}
	return lag1, lag3
}
