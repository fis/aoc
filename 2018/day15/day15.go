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

// Package day15 solves AoC 2018 day 15.
package day15

import (
	"sort"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2018, 15, glue.LevelSolver{Solver: solve, Empty: '#'})
}

func solve(level *util.Level) ([]int, error) {
	part1 := combat(level.Copy(), 3, false)
	part2 := findEdge(level)
	return []int{part1, part2}, nil
}

func findEdge(level *util.Level) int {
	right, rightOut := 4, -1
	for {
		rightOut = combat(level.Copy(), right, true)
		if rightOut != -1 {
			break
		}
		right *= 2
	}
	left := right / 2

	for right-left > 1 {
		mid := left + (right-left)/2
		midOut := combat(level.Copy(), mid, true)
		if midOut != -1 {
			right, rightOut = mid, midOut
		} else {
			left = mid
		}
	}

	return rightOut
}

type unit struct {
	kind  byte // 'E' or 'G'
	pos   util.P
	hp    int
	power int
}

func combat(level *util.Level, elfPower int, holyElves bool) (outcome int) {
	var units []unit
	level.Range(func(x, y int, b byte) {
		if b == 'E' || b == 'G' {
			pow := 3
			if b == 'E' {
				pow = elfPower
			}
			units = append(units, unit{kind: b, pos: util.P{x, y}, hp: 200, power: pow})
		}
	})

	type distP struct {
		p    util.P
		dist int
	}
	nearest, fringe := []distP(nil), []distP(nil)
	inRange := []int(nil)

	for round := 0; ; round++ {
		sort.Slice(units, func(i, j int) bool { return lessP(units[i].pos, units[j].pos) })
		for ui, u := range units {
			if u.hp <= 0 {
				continue
			}

			targets, inRange := 0, inRange[:0]
			open := make(map[util.P]struct{})
			for ti, t := range units {
				if t.kind != u.kind && t.hp > 0 {
					targets++
					if util.DistM(u.pos, t.pos) == 1 {
						inRange = append(inRange, ti)
					}
					for _, n := range t.pos.Neigh() {
						if level.At(n.X, n.Y) == '.' {
							open[util.P{n.X, n.Y}] = struct{}{}
						}
					}
				}
			}
			if targets == 0 {
				totalHP := 0
				for _, t := range units {
					if t.hp > 0 {
						totalHP += t.hp
					}
				}
				return round * totalHP
			}
			if len(open) == 0 && len(inRange) == 0 {
				continue
			}

			if len(inRange) == 0 {
				var seen map[util.P]struct{}

				nearest, fringe = nearest[:0], fringe[:0]
				fringe = append(fringe, distP{p: u.pos, dist: 0})
				seen = map[util.P]struct{}{}
				for len(fringe) > 0 && (len(nearest) == 0 || fringe[0].dist <= nearest[0].dist) {
					c := fringe[0]
					fringe = fringe[1:]
					if _, ok := open[c.p]; ok {
						nearest = append(nearest, c)
						continue
					}
					for _, n := range c.p.Neigh() {
						if level.At(n.X, n.Y) != '.' {
							continue
						} else if _, ok := seen[n]; ok {
							continue
						}
						seen[n] = struct{}{}
						fringe = append(fringe, distP{p: n, dist: c.dist + 1})
					}
				}
				if len(nearest) == 0 {
					continue
				}
				sort.Slice(nearest, func(i, j int) bool { return lessP(nearest[i].p, nearest[j].p) })
				chosen := nearest[0].p

				nearest, fringe = nearest[:0], fringe[:0]
				fringe = append(fringe, distP{p: chosen, dist: 0})
				seen = map[util.P]struct{}{chosen: {}}
				for len(fringe) > 0 && (len(nearest) == 0 || fringe[0].dist <= nearest[0].dist) {
					c := fringe[0]
					fringe = fringe[1:]
					if util.DistM(c.p, u.pos) == 1 {
						nearest = append(nearest, c)
						continue
					}
					for _, n := range c.p.Neigh() {
						if level.At(n.X, n.Y) != '.' {
							continue
						} else if _, ok := seen[n]; ok {
							continue
						}
						seen[n] = struct{}{}
						fringe = append(fringe, distP{p: n, dist: c.dist + 1})
					}
				}
				sort.Slice(nearest, func(i, j int) bool { return lessP(nearest[i].p, nearest[j].p) })
				step := nearest[0].p

				level.Set(u.pos.X, u.pos.Y, '.')
				u.pos = step
				level.Set(u.pos.X, u.pos.Y, u.kind)
				units[ui] = u

				for ti, t := range units {
					if t.kind != u.kind && t.hp > 0 && util.DistM(u.pos, t.pos) == 1 {
						inRange = append(inRange, ti)
					}
				}
			}

			if len(inRange) == 0 {
				continue
			}
			sort.Slice(inRange, func(i, j int) bool {
				ui, uj := units[inRange[i]], units[inRange[j]]
				return ui.hp < uj.hp || (ui.hp == uj.hp && lessP(ui.pos, uj.pos))
			})
			ti := inRange[0]
			units[ti].hp -= u.power
			if units[ti].hp <= 0 {
				if holyElves && units[ti].kind == 'E' {
					return -1
				}
				tp := units[ti].pos
				level.Set(tp.X, tp.Y, '.')
			}
		}

		j := 0
		for i, u := range units {
			if u.hp > 0 {
				if j < i {
					units[j] = units[i]
				}
				j++
			}
		}
		if j < len(units) {
			units = units[:j]
		}
	}
}

func lessP(a, b util.P) bool {
	return a.Y < b.Y || (a.Y == b.Y && a.X < b.X)
}
