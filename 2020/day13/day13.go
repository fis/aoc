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

// Package day13 solves AoC 2020 day 13.
package day13

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/fis/aoc-go/glue"
)

func init() {
	glue.RegisterSolver(2020, 13, glue.LineSolver(solve))
}

func solve(lines []string) ([]int, error) {
	earliest, buses, err := parseNotes(lines)
	if err != nil {
		return nil, err
	}
	id, wait := nextBus(earliest, buses)
	best := bestTime(buses)
	return []int{id * wait, best}, nil
}

func parseNotes(lines []string) (earliest int, buses []int, err error) {
	if len(lines) != 2 {
		return 0, nil, fmt.Errorf("expected 2 lines, got %d", len(lines))
	}
	if earliest, err = strconv.Atoi(lines[0]); err != nil {
		return 0, nil, fmt.Errorf("earliest timestamp not a number: %q: %w", lines[0], err)
	}
	for _, name := range strings.Split(lines[1], ",") {
		if name == "x" {
			buses = append(buses, -1)
		} else if id, err := strconv.Atoi(name); err == nil {
			buses = append(buses, id)
		} else {
			return 0, nil, fmt.Errorf("bad bus ID: %q", name)
		}
	}
	return earliest, buses, nil
}

func nextBus(earliest int, buses []int) (id, wait int) {
	id, wait = -1, math.MaxInt32
	for _, b := range buses {
		if b <= 0 {
			continue
		}
		n := (earliest + b - 1) / b * b
		w := n - earliest
		if w < wait {
			id, wait = b, w
		}
	}
	return id, wait
}

type constraint struct{ off, mod int }

func bestTime(buses []int) int {
	var cs []constraint
	for i, b := range buses {
		if b <= 0 {
			continue
		}
		cs = append(cs, constraint{off: ((-i)%b + b) % b, mod: b})
	}
	for len(cs) > 1 {
		i := len(cs) - 2
		cs[i] = merge(cs[i], cs[i+1])
		cs = cs[:i+1]
	}
	return cs[0].off
}

func merge(ca, cb constraint) constraint {
	if ca.mod < cb.mod {
		ca, cb = cb, ca
	}
	for a := ca.off; ; a += ca.mod {
		if a%cb.mod == cb.off {
			return constraint{off: a, mod: ca.mod * cb.mod}
		}
	}
}
