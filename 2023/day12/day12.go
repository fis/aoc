// Copyright 2023 Google LLC
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

// Package day12 solves AoC 2023 day 12.
package day12

import (
	"fmt"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/fn"
)

func init() {
	glue.RegisterSolver(2023, 12, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	data, err := fn.MapE(lines, parseLine)
	if err != nil {
		return nil, err
	}
	p1 := countAllWays(data)
	data = unfold(data, 5)
	p2 := countAllWays(data)
	return glue.Ints(p1, p2), nil
}

func countAllWays(data []record) int {
	return fn.SumF(data, func(r record) int { return countWays(r.row, r.groups) })
}

func unfold(data []record, factor int) []record {
	unfolded := make([]record, len(data))
	for i, r := range data {
		row := make([]byte, 0, factor*len(r.row)+factor-1)
		groups := make([]int, 0, factor*len(r.groups))
		for j := 0; j < factor; j++ {
			if j > 0 {
				row = append(row, '?')
			}
			row = append(row, r.row...)
			groups = append(groups, r.groups...)
		}
		unfolded[i] = record{row: row, groups: groups}
	}
	return unfolded
}

func countWays(row []byte, groups []int) int {
	return memoize(row, groups).solve(0, 0)
}

type memoized struct {
	row    []byte
	groups []int
	gmin   []int // gmin[i] = minimum size to place groups[i:]
	ways   []int // flattened (len(groups)+1) x (len(row)+1) array, 0 for unsolved, 1+w for solved
}

func memoize(row []byte, groups []int) *memoized {
	gmin := make([]int, len(groups))
	for gi, gm := len(groups)-1, -1; gi >= 0; gi-- {
		gm += groups[gi] + 1
		gmin[gi] = gm
	}
	ways := make([]int, (len(groups)+1)*(len(row)+1))
	return &memoized{row: row, groups: groups, gmin: gmin, ways: ways}
}

func (m *memoized) solve(gi, offset int) (w int) {
	if w := m.ways[gi*(len(m.row)+1)+offset]; w > 0 {
		return w - 1
	}
	w = 0
	if gi == len(m.groups) {
		// no groups left to place: see if suffix can be empty
		if offset == len(m.row) {
			w = 1
		} else if m.row[offset] == '#' {
			w = 0
		} else {
			w = m.solve(gi, offset+1)
		}
	} else if len(m.row[offset:]) < m.gmin[gi] {
		// some groups left to place, but no chance to place them
		w = 0
	} else {
		// some groups left to place and might still make it: try to place next group here or later
		if gs := m.groups[gi]; testFit(m.row[offset:], gs) {
			if offset+gs+1 < len(m.row) {
				w += m.solve(gi+1, offset+gs+1)
			} else if gi == len(m.groups)-1 {
				w++
			}
		}
		if m.row[offset] != '#' && offset+1 < len(m.row) {
			w += m.solve(gi, offset+1)
		}
	}
	m.ways[gi*(len(m.row)+1)+offset] = w + 1
	return w
}

func testFit(row []byte, groupSize int) bool {
	if groupSize > len(row) {
		return false
	}
	for i := 0; i < groupSize; i++ {
		if row[i] == '.' {
			return false
		}
	}
	return len(row) == groupSize || row[groupSize] != '#'
}

type record struct {
	row    []byte
	groups []int
}

func parseLine(line string) (record, error) {
	row, groups, ok := strings.Cut(line, " ")
	if !ok {
		return record{}, fmt.Errorf("bad line: %q", line)
	}
	return record{row: []byte(row), groups: util.Ints(groups)}, nil
}
