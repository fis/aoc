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

// Package day24 solves AoC 2020 day 24.
package day24

import (
	"fmt"
	"strings"

	"github.com/fis/aoc-go/glue"
	"github.com/fis/aoc-go/util"
)

func init() {
	glue.RegisterSolver(2020, 24, glue.LineSolver(solve))
}

func solve(lines []string) ([]int, error) {
	paths, err := parsePaths(lines)
	if err != nil {
		return nil, err
	}

	m := newTileMap()
	m.flipAll(paths)
	part1 := m.countBlack()

	for i := 0; i < 100; i++ {
		m = m.evolve()
	}
	part2 := m.countBlack()

	return []int{part1, part2}, nil
}

/*
Hexagonal grid coordinate system:

(0,0)   (1,0)   (2,0)   (3,0)
    (0,1)   (1,1)   (2,1)
(0,2)   (1,2)   (2,2)   (3,2)
    (0,3)   (1,3)   (2,3)
*/

type direction int

const (
	dirE direction = iota
	dirSE
	dirSW
	dirW
	dirNW
	dirNE
)

func (d direction) move(p util.P) util.P {
	return p.Add(([2][6]util.P{
		{dirE: {1, 0}, dirSE: {0, 1}, dirSW: {-1, 1}, dirW: {-1, 0}, dirNW: {-1, -1}, dirNE: {0, -1}},
		{dirE: {1, 0}, dirSE: {1, 1}, dirSW: {0, 1}, dirW: {-1, 0}, dirNW: {0, -1}, dirNE: {1, -1}},
	})[p.Y&1][d])
}

func neighH(p util.P) [6]util.P {
	return [6]util.P{dirE.move(p), dirSE.move(p), dirSW.move(p), dirW.move(p), dirNW.move(p), dirNE.move(p)}
}

type tileMap struct {
	tiles    map[util.P]struct{}
	min, max util.P
}

func newTileMap() *tileMap {
	return &tileMap{tiles: make(map[util.P]struct{})}
}

func (m *tileMap) black(p util.P) bool {
	_, ok := m.tiles[p]
	return ok
}

func (m *tileMap) flip(p util.P) {
	if p.X-1 < m.min.X {
		m.min.X = p.X - 1
	}
	if p.X+1 > m.max.X {
		m.max.X = p.X + 1
	}
	if p.Y-1 < m.min.Y {
		m.min.Y = p.Y - 1
	}
	if p.Y+1 > m.max.Y {
		m.max.Y = p.Y + 1
	}
	if _, ok := m.tiles[p]; ok {
		delete(m.tiles, p)
	} else {
		m.tiles[p] = struct{}{}
	}
}

func (m *tileMap) flipPath(path []direction) {
	at := util.P{0, 0}
	for _, d := range path {
		at = d.move(at)
	}
	m.flip(at)
}

func (m *tileMap) flipAll(paths [][]direction) {
	for _, path := range paths {
		m.flipPath(path)
	}
}

func (m tileMap) countBlack() (n int) {
	return len(m.tiles)
}

func (m *tileMap) evolve() *tileMap {
	next := newTileMap()
	for y := m.min.Y; y <= m.max.Y; y++ {
		for x := m.min.X; x <= m.max.X; x++ {
			at := util.P{x, y}
			c := 0
			for _, n := range neighH(at) {
				if m.black(n) {
					c++
				}
			}
			if (m.black(at) && (c == 1 || c == 2)) || (!m.black(at) && c == 2) {
				next.flip(at)
			}
		}
	}
	return next
}

func parsePath(s string) (path []direction, err error) {
	names := []struct {
		name string
		d    direction
	}{{"e", dirE}, {"se", dirSE}, {"sw", dirSW}, {"w", dirW}, {"nw", dirNW}, {"ne", dirNE}}
next:
	for len(s) > 0 {
		for _, n := range names {
			if strings.HasPrefix(s, n.name) {
				path = append(path, n.d)
				s = s[len(n.name):]
				continue next
			}
		}
		return nil, fmt.Errorf("bad direction: %s", s)
	}
	return path, nil
}

func parsePaths(lines []string) (paths [][]direction, err error) {
	for _, line := range lines {
		path, err := parsePath(line)
		if err != nil {
			return nil, err
		}
		paths = append(paths, path)
	}
	return paths, nil
}
