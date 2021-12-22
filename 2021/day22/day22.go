// Copyright 2021 Google LLC
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

// Package day22 solves AoC 2021 day 22.
package day22

import (
	"math"
	"strconv"

	"github.com/fis/aoc/glue"
)

const inputRegexp = `^(on|off) x=(-?\d+)\.\.(-?\d+),y=(-?\d+)\.\.(-?\d+),z=(-?\d+)\.\.(-?\d+)`

func init() {
	glue.RegisterSolver(2021, 22, glue.RegexpSolver{
		Solver: solve,
		Regexp: inputRegexp,
	})
}

func solve(lines [][]string) ([]string, error) {
	steps := parseInput(lines)
	tree := makeCubeTree()
	for _, step := range steps {
		tree.set(step.area, step.state)
	}
	p1 := tree.popCount(r3{p3{-50, -50, -50}, p3{51, 51, 51}})
	p2 := tree.popCount(allR3)
	return glue.Ints(p1, p2), nil
}

type bootStep struct {
	state cubeState
	area  r3
}

func parseInput(lines [][]string) (steps []bootStep) {
	steps = make([]bootStep, len(lines))
	for i, line := range lines {
		if line[0] == "on" {
			steps[i].state = cubesOn
		}
		minX, _ := strconv.Atoi(line[1])
		maxX, _ := strconv.Atoi(line[2])
		minY, _ := strconv.Atoi(line[3])
		maxY, _ := strconv.Atoi(line[4])
		minZ, _ := strconv.Atoi(line[5])
		maxZ, _ := strconv.Atoi(line[6])
		steps[i].area = r3{p3{int32(minX), int32(minY), int32(minZ)}, p3{int32(maxX) + 1, int32(maxY) + 1, int32(maxZ) + 1}}
	}
	return steps
}

type cubeTree struct {
	nodes []cubeTreeNode
}

type cubeTreeNode struct {
	state cubeState
	p     p3
	child uint32
}

const cubeTreeInitialSize = 256 * 1024

func makeCubeTree() cubeTree {
	return cubeTree{nodes: make([]cubeTreeNode, 1, cubeTreeInitialSize)}
}

func (ct *cubeTree) set(area r3, state cubeState) {
	ct.setNode(0, allR3, area, state)
}

func (ct *cubeTree) setNode(node uint32, bounds, area r3, state cubeState) {
	n := &ct.nodes[node]
	if area.min == bounds.min && area.max == bounds.max {
		n.state = state
		return
	}
	if n.state == cubesMixed {
		ct.descend(node, bounds, area, func(child uint32, subBounds, subArea r3) {
			ct.setNode(child, subBounds, subArea, state)
		})
		return
	}
	if state == n.state {
		return
	}
	if area.min != bounds.min {
		next := uint32(len(ct.nodes))
		ct.nodes = append(ct.nodes, make([]cubeTreeNode, 8)...)
		n = &ct.nodes[node]
		if n.state == cubesOn {
			ct.nodes[next+0].state, ct.nodes[next+1].state, ct.nodes[next+2].state, ct.nodes[next+3].state = cubesOn, cubesOn, cubesOn, cubesOn
			ct.nodes[next+4].state, ct.nodes[next+5].state, ct.nodes[next+6].state, ct.nodes[next+7].state = cubesOn, cubesOn, cubesOn, cubesOn
		}
		n.state, n.p, n.child = cubesMixed, area.min, next
		node, n = next+7, &ct.nodes[next+7]
	}
	if area.max != bounds.max {
		next := uint32(len(ct.nodes))
		ct.nodes = append(ct.nodes, make([]cubeTreeNode, 8)...)
		n = &ct.nodes[node]
		if n.state == cubesOn {
			ct.nodes[next+0].state, ct.nodes[next+1].state, ct.nodes[next+2].state, ct.nodes[next+3].state = cubesOn, cubesOn, cubesOn, cubesOn
			ct.nodes[next+4].state, ct.nodes[next+5].state, ct.nodes[next+6].state, ct.nodes[next+7].state = cubesOn, cubesOn, cubesOn, cubesOn
		}
		n.state, n.p, n.child = cubesMixed, area.max, next
		n = &ct.nodes[next]
	}
	n.state = state
}

func (ct *cubeTree) popCount(area r3) int {
	return ct.popCountNode(0, allR3, area)
}

func (ct *cubeTree) popCountNode(node uint32, bounds, area r3) (cubes int) {
	n := &ct.nodes[node]
	switch n.state {
	case cubesOff:
		return 0
	case cubesOn:
		sx, sy, sz := area.max.x-area.min.x, area.max.y-area.min.y, area.max.z-area.min.z
		return int(sx) * int(sy) * int(sz)
	default:
		ct.descend(node, bounds, area, func(child uint32, subBounds, subArea r3) {
			cubes += ct.popCountNode(child, subBounds, subArea)
		})
		return cubes
	}
}

func (ct *cubeTree) descend(node uint32, bounds, area r3, f func(child uint32, subBounds, subArea r3)) {
	p, child := ct.nodes[node].p, ct.nodes[node].child
	if area.min.x < p.x && area.min.y < p.y && area.min.z < p.z {
		subBounds := r3{bounds.min, p}
		subArea := r3{area.min, p3{min(area.max.x, p.x), min(area.max.y, p.y), min(area.max.z, p.z)}}
		f(child+0, subBounds, subArea)
	}
	if area.max.x >= p.x && area.min.y < p.y && area.min.z < p.z {
		subBounds := r3{p3{p.x, bounds.min.y, bounds.min.z}, p3{bounds.max.x, p.y, p.z}}
		subArea := r3{p3{max(area.min.x, p.x), area.min.y, area.min.z}, p3{area.max.x, min(area.max.y, p.y), min(area.max.z, p.z)}}
		f(child+1, subBounds, subArea)
	}
	if area.min.x < p.x && area.max.y >= p.y && area.min.z < p.z {
		subBounds := r3{p3{bounds.min.x, p.y, bounds.min.z}, p3{p.x, bounds.max.y, p.z}}
		subArea := r3{p3{area.min.x, max(area.min.y, p.y), area.min.z}, p3{min(area.max.x, p.x), area.max.y, min(area.max.z, p.z)}}
		f(child+2, subBounds, subArea)
	}
	if area.max.x >= p.x && area.max.y >= p.y && area.min.z < p.z {
		subBounds := r3{p3{p.x, p.y, bounds.min.z}, p3{bounds.max.x, bounds.max.y, p.z}}
		subArea := r3{p3{max(area.min.x, p.x), max(area.min.y, p.y), area.min.z}, p3{area.max.x, area.max.y, min(area.max.z, p.z)}}
		f(child+3, subBounds, subArea)
	}
	if area.min.x < p.x && area.min.y < p.y && area.max.z >= p.z {
		subBounds := r3{p3{bounds.min.x, bounds.min.y, p.z}, p3{p.x, p.y, bounds.max.z}}
		subArea := r3{p3{area.min.x, area.min.y, max(area.min.z, p.z)}, p3{min(area.max.x, p.x), min(area.max.y, p.y), area.max.z}}
		f(child+4, subBounds, subArea)
	}
	if area.max.x >= p.x && area.min.y < p.y && area.max.z >= p.z {
		subBounds := r3{p3{p.x, bounds.min.y, p.z}, p3{bounds.max.x, p.y, bounds.max.z}}
		subArea := r3{p3{max(area.min.x, p.x), area.min.y, max(area.min.z, p.z)}, p3{area.max.x, min(area.max.y, p.y), area.max.z}}
		f(child+5, subBounds, subArea)
	}
	if area.min.x < p.x && area.max.y >= p.y && area.max.z >= p.z {
		subBounds := r3{p3{bounds.min.x, p.y, p.z}, p3{p.x, bounds.max.y, bounds.max.z}}
		subArea := r3{p3{area.min.x, max(area.min.y, p.y), max(area.min.z, p.z)}, p3{min(area.max.x, p.x), area.max.y, area.max.z}}
		f(child+6, subBounds, subArea)
	}
	if area.max.x >= p.x && area.max.y >= p.y && area.max.z >= p.z {
		subBounds := r3{p, bounds.max}
		subArea := r3{p3{max(area.min.x, p.x), max(area.min.y, p.y), max(area.min.z, p.z)}, area.max}
		f(child+7, subBounds, subArea)
	}
}

type cubeState uint32

const (
	cubesOff   cubeState = 0
	cubesOn    cubeState = 1
	cubesMixed cubeState = 2
)

type p3 struct {
	x, y, z int32
}

type r3 struct {
	min, max p3
}

var (
	minP3 = p3{math.MinInt32, math.MinInt32, math.MinInt32}
	maxP3 = p3{math.MaxInt32, math.MaxInt32, math.MaxInt32}
	allR3 = r3{minP3, maxP3}
)

func min(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

func max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}
