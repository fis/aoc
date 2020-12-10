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

// Package days contains the glue and tests for all AoC 2019 days.
package days

import (
	"github.com/fis/aoc-go/util"

	_ "github.com/fis/aoc-go/2019/day01"
	_ "github.com/fis/aoc-go/2019/day02"
	_ "github.com/fis/aoc-go/2019/day03"
	_ "github.com/fis/aoc-go/2019/day04"
	_ "github.com/fis/aoc-go/2019/day05"
	_ "github.com/fis/aoc-go/2019/day06"
	_ "github.com/fis/aoc-go/2019/day07"
	_ "github.com/fis/aoc-go/2019/day08"
	_ "github.com/fis/aoc-go/2019/day09"
	_ "github.com/fis/aoc-go/2019/day10"
	_ "github.com/fis/aoc-go/2019/day11"
	_ "github.com/fis/aoc-go/2019/day12"
	_ "github.com/fis/aoc-go/2019/day13"
	_ "github.com/fis/aoc-go/2019/day14"
	_ "github.com/fis/aoc-go/2019/day15"
	_ "github.com/fis/aoc-go/2019/day16"
	_ "github.com/fis/aoc-go/2019/day17"
	_ "github.com/fis/aoc-go/2019/day18"
	_ "github.com/fis/aoc-go/2019/day19"
	_ "github.com/fis/aoc-go/2019/day20"
	_ "github.com/fis/aoc-go/2019/day21"
	_ "github.com/fis/aoc-go/2019/day22"
	_ "github.com/fis/aoc-go/2019/day23"
	_ "github.com/fis/aoc-go/2019/day24"
	_ "github.com/fis/aoc-go/2019/day25"
)

func Solve(day int, path string) ([]string, error) {
	return util.CallSolver(day, path)
}
