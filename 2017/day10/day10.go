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

// Package day10 solves AoC 2017 day 10.
package day10

import (
	"fmt"
	"strconv"

	"github.com/fis/aoc/2017/knot"
	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2017, 10, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	if len(lines) != 1 {
		return nil, fmt.Errorf("expected 1 line of input, got %d", len(lines))
	}
	ints := util.Ints(lines[0])
	lengths := make([]byte, len(ints))
	for i, n := range ints {
		lengths[i] = byte(n)
	}
	p1 := part1(knot.N, lengths)
	p2 := knot.Hash(knot.N, knot.Rounds, lines[0])
	return []string{strconv.Itoa(p1), fmt.Sprintf("%x", p2)}, nil
}

func part1(N int, lengths []byte) int {
	list := knot.List(N)
	knot.Round(0, 0, list, lengths)
	return int(list[0]) * int(list[1])
}
