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

// Package day03 solves AoC 2019 day 3.
package day03

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fis/aoc-go/glue"
	"github.com/fis/aoc-go/util"
)

func init() {
	glue.RegisterSolver(2019, 3, glue.LineSolver(solve))
}

func solve(lines []string) ([]int, error) {
	if len(lines) != 2 {
		return nil, fmt.Errorf("unexpected amount of lines: %v", lines)
	}
	p1, p2 := compute(lines[0], lines[1])
	return []int{p1, p2}, nil
}

func compute(w1, w2 string) (closest, best int) {
	m := make(map[util.P]int)
	walk(w1, func(x, y, s1 int) {
		p := util.P{x, y}
		if _, ok := m[p]; ok {
			return
		}
		m[p] = s1
	})
	closest, best = -1, -1
	walk(w2, func(x, y, s2 int) {
		s1, ok := m[util.P{x, y}]
		if !ok {
			return
		}
		td := abs(x) + abs(y)
		if closest == -1 || td < closest {
			closest = td
		}
		ts := s1 + s2
		if best == -1 || ts < best {
			best = ts
		}
	})
	return closest, best
}

var dir = map[byte]util.P{
	'L': {-1, 0},
	'R': {1, 0},
	'U': {0, -1},
	'D': {0, 1},
}

func walk(w string, cb func(x, y, s int)) {
	x, y, s := 0, 0, 0
	for _, op := range strings.Split(w, ",") {
		d := dir[op[0]]
		n, _ := strconv.Atoi(op[1:])
		for i := 0; i < n; i++ {
			x, y, s = x+d.X, y+d.Y, s+1
			cb(x, y, s)
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
