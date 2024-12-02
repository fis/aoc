// Copyright 2024 Google LLC
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

// Package day02 solves AoC 2024 day 2.
package day02

import (
	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/fn"
	"github.com/fis/aoc/util/ix"
)

func init() {
	glue.RegisterSolver(2024, 2, glue.LineSolver(glue.WithParser(parseReport, solve)))
}

func solve(reports [][]int) ([]string, error) {
	part1 := fn.CountIf(reports, isSafe)
	part2 := fn.CountIf(reports, isSafeDampened)
	return glue.Ints(part1, part2), nil
}

func parseReport(line string) ([]int, error) {
	return util.Ints(line), nil
}

func isSafe(report []int) bool {
	sign := 0
	for i := 1; i < len(report); i++ {
		d := report[i] - report[i-1]
		if ad := ix.Abs(d); ad < 1 || ad > 3 {
			return false
		}
		if sign == 0 {
			sign = ix.Sign(d)
		} else if sign != ix.Sign(d) {
			return false
		}
	}
	return true
}

func isSafeDampened(report []int) bool {
	dampened := make([]int, len(report)-1)
	for del := range report {
		copy(dampened[:del], report[:del])
		copy(dampened[del:], report[del+1:])
		if isSafe(dampened) {
			return true
		}
	}
	return false
}
