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

// Package day09 solves AoC 2018 day 9.
package day09

import (
	"fmt"

	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2018, 9, glue.LineSolver(solve))
}

func solve(lines []string) ([]int, error) {
	var players, marbles int
	if _, err := fmt.Sscanf(lines[0], "%d players; last marble is worth %d points", &players, &marbles); err != nil {
		return nil, err
	}

	p1 := game(players, marbles)
	p2 := game(players, 100*marbles)

	return []int{p1, p2}, nil
}

func game(players, marbles int) (maxScore int) {
	m := &marble{num: 0}
	m.ccw, m.cw = m, m

	scores := make([]int, players)
	for turn := 1; turn <= marbles; turn++ {
		if turn%23 != 0 {
			m = m.cw.insert(turn)
		} else {
			pl := turn % players
			scores[pl] += turn
			m = m.ccw.ccw.ccw.ccw.ccw.ccw.ccw
			scores[pl] += m.num
			m = m.remove()
		}
	}

	maxScore = scores[0]
	for _, s := range scores[1:] {
		if s > maxScore {
			maxScore = s
		}
	}
	return maxScore
}

type marble struct {
	num     int
	ccw, cw *marble
}

// insert injects a new marble clockwise from the receiver, and returns it.
func (m *marble) insert(num int) *marble {
	nm := &marble{num: num, ccw: m, cw: m.cw}
	nm.ccw.cw, nm.cw.ccw = nm, nm
	return nm
}

// remove removes the receiver from the ring, and returns its (former) clockwise neighbour.
func (m *marble) remove() *marble {
	nm := m.cw
	m.ccw.cw, m.cw.ccw = m.cw, m.ccw
	return nm
}
