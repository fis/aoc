// Package day12 solves AoC 2017 day 12.
package day12

import (
	"fmt"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/graph"
)

func init() {
	glue.RegisterSolver(2017, 12, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	g, err := buildGraph(lines)
	if err != nil {
		return nil, err
	}
	vertGroup, groupVerts := partition(g)
	zero, ok := g.V("0")
	if !ok {
		return nil, fmt.Errorf("no vertex 0 in graph")
	}
	p1 := len(groupVerts[vertGroup[zero]])
	p2 := len(groupVerts)
	return glue.Ints(p1, p2), nil
}

func partition(g *graph.Sparse) (vertGroup map[int]int, groupVerts map[int][]int) {
	vertGroup = make(map[int]int)
	groupVerts = make(map[int][]int)
	for startV := range g.Len() {
		if _, found := vertGroup[startV]; found {
			continue
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
			for toV := range g.Succ(at) {
				if _, found := vertGroup[toV]; !found {
					edge = append(edge, toV)
				}
			}
		}
	}
	return vertGroup, groupVerts
}

func buildGraph(lines []string) (*graph.Sparse, error) {
	g := graph.NewBuilder()
	for _, line := range lines {
		from, tos, ok := strings.Cut(line, " <-> ")
		if !ok {
			return nil, fmt.Errorf("bad line: missing <->: %s", line)
		}
		u := g.V(from)
		for s := util.Splitter(tos); !s.Empty(); {
			to := s.Next(", ")
			v := g.V(to)
			g.AddEdge(u, v)
		}
	}
	return g.SparseDigraph(), nil
}
