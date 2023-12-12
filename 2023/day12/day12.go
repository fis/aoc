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
		var (
			row    []byte
			groups []int
		)
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
	// ways[i][j] ::= number of ways the last groups[i:] can fit in row[j:]
	ways := make([][]int, len(groups)+1)
	for i := 0; i <= len(groups); i++ {
		ways[i] = make([]int, len(row))
	}

	// base case: suffixes that can match no groups
	{
		w := 1
		for i := len(row) - 1; i >= 0; i-- {
			if row[i] == '#' {
				w = 0
			}
			ways[len(groups)][i] = w
		}
	}

	// recursive case
	for gi := len(groups) - 1; gi >= 0; gi-- {
		gs := groups[gi]
		for offset := len(row) - 1; offset >= 0; offset-- {
			w := 0
			if row[offset] != '#' && offset+1 < len(row) {
				w += ways[gi][offset+1]
			}
			if testFit(row[offset:], gs) {
				if offset+gs+1 < len(row) {
					w += ways[gi+1][offset+gs+1]
				} else if gi == len(groups)-1 {
					w++
				}
			}
			ways[gi][offset] = w
		}
	}

	return ways[0][0]
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
