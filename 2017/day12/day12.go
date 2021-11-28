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

// Package day12 solves AoC 2017 day 12.
package day12

import (
	"fmt"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2017, 12, glue.LineSolver(solve))
}

func solve(lines []string) ([]int, error) {
	g, err := parseLines(lines)
	if err != nil {
		return nil, err
	}
	vertGroup, groupVerts := partition(g)
	p1 := len(groupVerts[vertGroup[g.V("0")]])
	p2 := len(groupVerts)
	return []int{p1, p2}, nil
}

func partition(g *util.Graph) (vertGroup map[int]int, groupVerts map[int][]int) {
	vertGroup = make(map[int]int)
	groupVerts = make(map[int][]int)
	g.RangeV(func(startV int) {
		if _, found := vertGroup[startV]; found {
			return
		}
		group := len(groupVerts)
		edge := []int{startV}
		for len(edge) > 0 {
			at := edge[len(edge)-1]
			edge = edge[:len(edge)-1]
			if _, found := vertGroup[at]; found {
				continue
			}
			vertGroup[at] = group
			groupVerts[group] = append(groupVerts[group], at)
			g.RangeSuccV(at, func(toV int) bool {
				if _, found := vertGroup[toV]; !found {
					edge = append(edge, toV)
				}
				return true
			})
		}
	})
	return vertGroup, groupVerts
}

func parseLines(lines []string) (*util.Graph, error) {
	g := &util.Graph{}
	for _, line := range lines {
		parts := strings.SplitN(line, " ", 3)
		if len(parts) != 3 || parts[1] != "<->" {
			return nil, fmt.Errorf("invalid line: %q: expected \"x <-> y, z, ...\"", line)
		}
		fromV := g.V(parts[0])
		for _, to := range strings.Split(parts[2], ", ") {
			toV := g.V(to)
			g.AddEdgeV(fromV, toV)
			g.AddEdgeV(toV, fromV)
		}
	}
	return g, nil
}
