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

// Package day04 solves AoC 2022 day 4.
package day04

import (
	"strconv"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util/fn"
)

func init() {
	glue.RegisterSolver(2022, 4, glue.ParsableLineSolver[[2]section]{
		Solver: solve,
		Parser: parsePair,
	})
}

func solve(pairs [][2]section) ([]string, error) {
	p1 := part1(pairs)
	p2 := part2(pairs)
	return glue.Ints(p1, p2), nil
}

func part1(pairs [][2]section) int {
	return fn.CountIf(pairs, func(s [2]section) bool { return s[0].contains(s[1]) || s[1].contains(s[0]) })
}

func part2(pairs [][2]section) int {
	return fn.CountIf(pairs, func(s [2]section) bool { return s[0].overlaps(s[1]) })
}

func parsePair(line string) ([2]section, error) {
	comma := strings.IndexByte(line, ',')
	return [2]section{parseSection(line[:comma]), parseSection(line[comma+1:])}, nil
}

type section struct {
	start, end int
}

func (s section) contains(t section) bool {
	return t.start >= s.start && t.end <= s.end
}

func (s section) overlaps(t section) bool {
	return t.start <= s.end && t.end >= s.start
}

func parseSection(spec string) section {
	dash := strings.IndexByte(spec, '-')
	start, _ := strconv.Atoi(spec[:dash])
	end, _ := strconv.Atoi(spec[dash+1:])
	return section{start, end}
}
