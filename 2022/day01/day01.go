// Copyright 2022 Google LLC
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

// Package day01 solves AoC 2022 day 1.
package day01

import (
	"slices"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/fn"
)

func init() {
	glue.RegisterSolver(2022, 1, glue.ChunkSolver(solve))
}

func solve(chunks []string) ([]string, error) {
	data := fn.Map(chunks, util.Ints)
	p1, p2 := maxCalories(data)
	return glue.Ints(p1, p2), nil
}

func maxCalories(data [][]int) (top, top3 int) {
	sums := fn.Map(data, fn.Sum[[]int])
	slices.Sort(sums)
	n := len(sums)
	return sums[n-1], sums[n-1] + sums[n-2] + sums[n-3]
}
