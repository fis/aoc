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

// Package day12 solves AoC 2021 day 12.
package day12

import (
	"fmt"
	"io"
	"strings"
	"unicode"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util/graph"
)

const inputRegexp = `^([^-]+)-([^-]+)$`

func init() {
	glue.RegisterSolver(2021, 12, glue.RegexpSolver{Solver: solve, Regexp: inputRegexp})
	glue.RegisterPlotter(2021, 12, "", glue.LinePlotter(plot), map[string]string{
		"ex1": edgesToText(ex1),
		"ex2": edgesToText(ex2),
		"ex3": edgesToText(ex3),
	})
}

func solve(edges [][]string) ([]string, error) {
	g := makeIntGraph(edges)
	p1 := g.countAllPaths(false)
	p2 := g.countAllPaths(true)
	return glue.Ints(p1, p2), nil
}

// quick and dirty integer-index graph with edge lists & a uint64 bitmap for visited small caves

type intGraph struct {
	verts map[string]int
	small []bool
	edges [][]int
}

func (g *intGraph) countAllPaths(allowTwice bool) int {
	start, end := g.verts["start"], g.verts["end"]
	smallCaves := uint64(1) << start
	return g.countPaths(start, end, allowTwice, smallCaves)
}

func (g *intGraph) countPaths(src, dst int, allowTwice bool, smallCaves uint64) (paths int) {
	if src == dst {
		return 1
	}
	for _, next := range g.edges[src] {
		small := g.small[next]
		usedAllowance := false
		if small {
			if smallCaves&(uint64(1)<<next) != 0 {
				if allowTwice {
					usedAllowance = true
					allowTwice = false
				} else {
					continue
				}
			}
			if !usedAllowance {
				smallCaves |= uint64(1) << next
			}
		}
		paths += g.countPaths(next, dst, allowTwice, smallCaves)
		if small {
			if usedAllowance {
				allowTwice = true
			} else {
				smallCaves &= ^(uint64(1) << next)
			}
		}
	}
	return paths
}

func makeIntGraph(lines [][]string) (g *intGraph) {
	verts := make(map[string]int)
	small := []bool(nil)
	for _, line := range lines {
		for _, vert := range line {
			if _, seen := verts[vert]; !seen {
				i := len(small)
				verts[vert] = i
				small = append(small, unicode.IsLower(rune(vert[0])))
			}
		}
	}
	start, end := verts["start"], verts["end"]
	edges := make([][]int, len(verts))
	for _, line := range lines {
		a, b := verts[line[0]], verts[line[1]]
		if a != end && b != start {
			edges[a] = append(edges[a], b)
		}
		if b != end && a != start {
			edges[b] = append(edges[b], a)
		}
	}
	return &intGraph{verts: verts, small: small, edges: edges}
}

// util/graph

func countAllPathsSparse(g *graph.Sparse, allowTwice bool) int {
	start, _ := g.V("start")
	end, _ := g.V("end")
	smallCaves := make([]bool, g.Len())
	smallCaves[start] = true
	return countPathsSparse(g, start, end, allowTwice, smallCaves)
}

func countPathsSparse(g *graph.Sparse, src, dst int, allowTwice bool, smallCaves []bool) (paths int) {
	if src == dst {
		return 1
	}
	for it := g.Succ(src); it.Valid(); it = g.Next(it) {
		next := it.Head()
		lower := unicode.IsLower(rune(g.Label(next)[0]))
		usedAllowance := false
		if lower {
			if smallCaves[next] {
				if allowTwice {
					usedAllowance = true
					allowTwice = false
				} else {
					continue
				}
			}
			if !usedAllowance {
				smallCaves[next] = true
			}
		}
		paths += countPathsSparse(g, next, dst, allowTwice, smallCaves)
		if lower {
			if usedAllowance {
				allowTwice = true
			} else {
				smallCaves[next] = false
			}
		}
	}
	return paths
}

func countAllPathsDense(g *graph.Dense, allowTwice bool) int {
	start, _ := g.V("start")
	end, _ := g.V("end")
	smallCaves := make([]bool, g.Len())
	smallCaves[start] = true
	return countPathsDense(g, start, end, allowTwice, smallCaves)
}

func countPathsDense(g *graph.Dense, src, dst int, allowTwice bool, smallCaves []bool) (paths int) {
	if src == dst {
		return 1
	}
	for it := g.Succ(src); it.Valid(); it = g.Next(it) {
		next := it.Head()
		lower := unicode.IsLower(rune(g.Label(next)[0]))
		usedAllowance := false
		if lower {
			if smallCaves[next] {
				if allowTwice {
					usedAllowance = true
					allowTwice = false
				} else {
					continue
				}
			}
			if !usedAllowance {
				smallCaves[next] = true
			}
		}
		paths += countPathsDense(g, next, dst, allowTwice, smallCaves)
		if lower {
			if usedAllowance {
				allowTwice = true
			} else {
				smallCaves[next] = false
			}
		}
	}
	return paths
}

func makeGraph[T any](edges [][]string, builder func(*graph.Builder) T) T {
	g := graph.NewBuilder()
	for _, edge := range edges {
		if edge[0] != "end" && edge[1] != "start" {
			g.AddEdgeL(edge[0], edge[1])
		}
		if edge[1] != "end" && edge[0] != "start" {
			g.AddEdgeL(edge[1], edge[0])
		}
	}
	return builder(g)
}

// plotting

var (
	ex1 = [][]string{
		{"start", "A"},
		{"start", "b"},
		{"A", "c"},
		{"A", "b"},
		{"b", "d"},
		{"A", "end"},
		{"b", "end"},
	}
	ex2 = [][]string{
		{"dc", "end"},
		{"HN", "start"},
		{"start", "kj"},
		{"dc", "start"},
		{"dc", "HN"},
		{"LN", "dc"},
		{"HN", "end"},
		{"kj", "sa"},
		{"kj", "HN"},
		{"kj", "dc"},
	}
	ex3 = [][]string{
		{"fs", "end"},
		{"he", "DX"},
		{"fs", "he"},
		{"start", "DX"},
		{"pj", "DX"},
		{"end", "zg"},
		{"zg", "sl"},
		{"zg", "pj"},
		{"pj", "he"},
		{"RW", "he"},
		{"fs", "DX"},
		{"pj", "RW"},
		{"zg", "RW"},
		{"start", "pj"},
		{"he", "WI"},
		{"zg", "he"},
		{"pj", "fs"},
		{"start", "RW"},
	}
)

func edgesToText(edges [][]string) string {
	b := strings.Builder{}
	for _, edge := range edges {
		fmt.Fprintf(&b, "%s-%s\n", edge[0], edge[1])
	}
	return b.String()
}

func plot(lines []string, w io.Writer) error {
	fmt.Fprintln(w, "graph G {")
	fmt.Fprintln(w, `  start [style="filled",fillcolor="#0f9d58",fontcolor="white"];`)
	fmt.Fprintln(w, `  end [style="filled",fillcolor="#db4437",fontcolor="white"];`)
	lower := map[string]bool{"start": true, "end": true}
	for _, line := range lines {
		parts := strings.Split(line, "-")
		for _, node := range parts {
			if unicode.IsLower(rune(node[0])) && !lower[node] {
				fmt.Fprintf(w, "  %s [style=\"filled\",fillcolor=\"#dddddd\"];\n", node)
				lower[node] = true
			}
		}
		fmt.Fprintf(w, "  %s -- %s;\n", parts[0], parts[1])
	}
	fmt.Fprintln(w, "}")
	return nil
}
