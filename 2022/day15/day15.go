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

// Package day15 solves AoC 2022 day 15.
package day15

import (
	"fmt"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/fn"
	"github.com/fis/aoc/util/ix"
)

func init() {
	glue.RegisterSolver(2022, 15, glue.LineSolver(glue.WithParser(parseReading, solve)))
}

func solve(data []reading) ([]string, error) {
	p1 := part1(data, 2000000)
	p2b := part2(data)
	return glue.Ints(p1, p2b), nil
}

func part1(data []reading, y int) int {
	var covered intervalSet
	for _, r := range data {
		bd := util.DistM(r.sensor, r.beacon)
		yd := ix.Abs(y - r.sensor.Y)
		if w := bd - yd; w >= 0 {
			if r.beacon.Y == y {
				if r.sensor.X-w < r.beacon.X {
					covered = covered.add(r.sensor.X-w, r.beacon.X-1)
				}
				if r.sensor.X+w > r.beacon.X {
					covered = covered.add(r.beacon.X+1, r.sensor.X+w)
				}
			} else {
				covered = covered.add(r.sensor.X-w, r.sensor.X+w)
			}
		}
	}
	return covered.size()
}

func part2(data []reading) int {
	var qt quadTree = quadTreeLeaf(false)
	for _, r := range data {
		d := util.DistM(r.sensor, r.beacon)
		min := util.P{r.sensor.X - r.sensor.Y - d, r.sensor.X + r.sensor.Y - d}
		max := util.P{r.sensor.X - r.sensor.Y + d + 1, r.sensor.X + r.sensor.Y + d + 1}
		qt = qt.set(min, max, util.MinP, util.MaxP, true)
	}
	gap, _ := qt.findGap(util.MinP, util.MaxP)
	x, y := (gap.X+gap.Y)/2, (-gap.X+gap.Y)/2
	return 4000000*x + y
}

// part 1 & part 2: common parsing

type reading struct {
	sensor util.P
	beacon util.P
}

func parseReading(line string) (reading, error) {
	var sx, sy, bx, by int
	if _, err := fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by); err != nil {
		return reading{}, fmt.Errorf("bad reading: %q: %w", line, err)
	}
	return reading{sensor: util.P{sx, sy}, beacon: util.P{bx, by}}, nil
}

// 1D interval set for part 1

type interval struct {
	start, end int
}

func (iv interval) size() int {
	return iv.end - iv.start + 1
}

func (iv interval) contains(x int) bool {
	return x >= iv.start && x <= iv.end
}

type intervalSet []interval

func (is intervalSet) size() int {
	return fn.SumF(is, interval.size)
}

func (is intervalSet) add(start, end int) intervalSet {
	if len(is) == 0 {
		return intervalSet{interval{start, end}}
	}
	pos := 0
	for pos < len(is) && start > is[pos].start {
		pos++
	}
	is = append(is, interval{})
	copy(is[pos+1:], is[pos:])
	is[pos] = interval{start, end}
	for pos := 0; pos+1 < len(is); {
		if is[pos+1].start <= is[pos].end+1 {
			is[pos].end = ix.Max(is[pos].end, is[pos+1].end)
			copy(is[pos+1:], is[pos+2:])
			is = is[:len(is)-1]
		} else {
			pos++
		}
	}
	return is
}

// 2D quadtree for part 2

type quadTree interface {
	set(min, max, treeMin, treeMax util.P, v bool) quadTree
	findGap(min, max util.P) (p util.P, ok bool)
}

type quadTreeNode struct {
	p util.P
	q [2][2]quadTree
}

type quadTreeLeaf bool

func (n *quadTreeNode) set(min, max, treeMin, treeMax util.P, v bool) quadTree {
	if min.X < n.p.X && min.Y < n.p.Y {
		qmax := util.P{ix.Min(max.X, n.p.X), ix.Min(max.Y, n.p.Y)}
		n.q[0][0] = n.q[0][0].set(min, qmax, treeMin, n.p, v)
	}
	if max.X > n.p.X && min.Y < n.p.Y {
		qmin, qmax := util.P{ix.Max(min.X, n.p.X), min.Y}, util.P{max.X, ix.Min(max.Y, n.p.Y)}
		n.q[0][1] = n.q[0][1].set(qmin, qmax, util.P{n.p.X, treeMin.Y}, util.P{treeMax.X, n.p.Y}, v)
	}
	if min.X < n.p.X && max.Y > n.p.Y {
		qmin, qmax := util.P{min.X, ix.Max(min.Y, n.p.Y)}, util.P{ix.Min(max.X, n.p.X), max.Y}
		n.q[1][0] = n.q[1][0].set(qmin, qmax, util.P{treeMin.X, n.p.Y}, util.P{n.p.X, treeMax.Y}, v)
	}
	if max.X > n.p.X && max.Y > n.p.Y {
		qmin := util.P{ix.Max(min.X, n.p.X), ix.Max(min.Y, n.p.Y)}
		n.q[1][1] = n.q[1][1].set(qmin, max, n.p, treeMax, v)
	}
	return n
}

func (lv quadTreeLeaf) set(min, max, treeMin, treeMax util.P, v bool) quadTree {
	if bool(lv) == v {
		return lv
	}
	var qt quadTree = quadTreeLeaf(v)
	if min != treeMin {
		qt = &quadTreeNode{p: min, q: [2][2]quadTree{{lv, lv}, {lv, qt}}}
	}
	if max != treeMax {
		qt = &quadTreeNode{p: max, q: [2][2]quadTree{{qt, lv}, {lv, lv}}}
	}
	return qt
}

func (n *quadTreeNode) findGap(min, max util.P) (p util.P, ok bool) {
	if gap, ok := n.q[0][0].findGap(min, n.p); ok {
		return gap, true
	}
	if gap, ok := n.q[0][1].findGap(util.P{n.p.X, min.Y}, util.P{max.X, n.p.Y}); ok {
		return gap, true
	}
	if gap, ok := n.q[1][0].findGap(util.P{min.X, n.p.Y}, util.P{n.p.X, max.Y}); ok {
		return gap, true
	}
	if gap, ok := n.q[1][1].findGap(n.p, max); ok {
		return gap, true
	}
	return util.P{}, false
}

func (lv quadTreeLeaf) findGap(min, max util.P) (p util.P, ok bool) {
	if !lv && max.X == min.X+1 && max.Y == min.Y+1 {
		return min, true
	}
	return util.P{}, false
}
