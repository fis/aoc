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

// Package day11 solves AoC 2017 day 11.
package day11

import (
	"fmt"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util/ix"
)

func init() {
	glue.RegisterSolver(2017, 11, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	if len(lines) != 1 {
		return nil, fmt.Errorf("expected 1 line of input, got %d", len(lines))
	}
	steps := strings.Split(lines[0], ",")
	p1, p2, err := distances(steps)
	if err != nil {
		return nil, err
	}
	return glue.Ints(p1, p2), nil
}

func distances(steps []string) (lastD, maxD int, err error) {
	q, r := 0, 0
	for _, step := range steps {
		d, ok := directions[step]
		if !ok {
			return 0, 0, fmt.Errorf("invalid direction: %q", step)
		}
		q += d.q
		r += d.r
		if dist := distance(q, r); dist > maxD {
			maxD = dist
		}
	}
	return distance(q, r), maxD, nil
}

var directions = map[string]struct{ q, r int }{
	"nw": {q: -1, r: 0},
	"n":  {q: 0, r: -1},
	"ne": {q: 1, r: -1},
	"se": {q: 1, r: 0},
	"s":  {q: 0, r: 1},
	"sw": {q: -1, r: 1},
}

func distance(q, r int) int {
	return (ix.Abs(q) + ix.Abs(q+r) + ix.Abs(r)) / 2
}
