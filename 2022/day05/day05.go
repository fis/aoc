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

// Package day05 solves AoC 2022 day 5.
package day05

import (
	"fmt"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/fn"
	"golang.org/x/exp/slices"
)

func init() {
	glue.RegisterSolver(2022, 5, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	stacks, moves := parseInput(lines)
	stacks1 := applyMoves(stacks, moves, move.apply9000)
	stacks2 := applyMoves(stacks, moves, move.apply9001)
	p1 := tops(stacks1)
	p2 := tops(stacks2)
	return []string{p1, p2}, nil
}

func applyMoves(stacks [][]byte, moves []move, f func(move, [][]byte)) [][]byte {
	stacks = fn.Map(stacks, slices.Clone[[]byte])
	for _, m := range moves {
		f(m, stacks)
	}
	return stacks
}

func tops(stacks [][]byte) string {
	return string(fn.Map(stacks, func(s []byte) byte { return s[len(s)-1] }))
}

type move struct {
	count, from, to int
}

func (m move) apply9000(stacks [][]byte) {
	from, to := stacks[m.from], stacks[m.to]
	for i := 0; i < m.count; i++ {
		t := from[len(from)-1]
		from = from[:len(from)-1]
		to = append(to, t)
	}
	stacks[m.from], stacks[m.to] = from, to
}

func (m move) apply9001(stacks [][]byte) {
	from, to := stacks[m.from], stacks[m.to]
	fromTop, toTop := len(from), len(to)
	to = append(to, make([]byte, m.count)...)
	copy(to[toTop:], from[fromTop-m.count:])
	from = from[:fromTop-m.count]
	stacks[m.from], stacks[m.to] = from, to
}

func parseInput(lines []string) (stacks [][]byte, moves []move) {
	split := slices.Index(lines, "")

	stacks = make([][]byte, len(util.Words(lines[split-1])))
	for i := split - 2; i >= 0; i-- {
		line := lines[i]
		for j, k := 1, 0; j < len(line); j, k = j+4, k+1 {
			b := line[j]
			if b != ' ' {
				stacks[k] = append(stacks[k], line[j])
			}
		}
	}

	for _, line := range lines[split+1:] {
		var m move
		fmt.Sscanf(line, "move %d from %d to %d", &m.count, &m.from, &m.to)
		m.from, m.to = m.from-1, m.to-1
		moves = append(moves, m)
	}

	return stacks, moves
}
