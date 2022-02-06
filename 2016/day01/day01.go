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

// Package day01 solves AoC 2016 day 1.
package day01

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2016, 1, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	if len(lines) != 1 {
		return nil, fmt.Errorf("expected 1 line, got %d", len(lines))
	}
	steps := strings.Split(lines[0], ", ")
	p1, p2 := find(steps)
	return glue.Ints(p1, p2), nil
}

func find(steps []string) (end, hq int) {
	at, d := util.P{0, 0}, util.P{0, -1}
	seen := map[util.P]struct{}{{0, 0}: {}}
	hq = -1
	for _, step := range steps {
		switch step[0] {
		case 'L':
			d.X, d.Y = d.Y, -d.X
		case 'R':
			d.X, d.Y = -d.Y, d.X
		}
		n, _ := strconv.Atoi(step[1:])
		for i := 0; i < n; i++ {
			at.X += d.X
			at.Y += d.Y
			if hq >= 0 {
				continue
			}
			if _, ok := seen[at]; ok {
				hq = util.DistM(util.P{0, 0}, at)
				seen = nil
			} else {
				seen[at] = struct{}{}
			}
		}
	}
	end = util.DistM(util.P{0, 0}, at)
	return end, hq
}
