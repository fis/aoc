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

// Package day25 solves AoC 2023 day 25.
package day25

import (
	"fmt"
	"io"
	"slices"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2023, 25, glue.LineSolver(solve))
	glue.RegisterPlotter(2023, 25, "", glue.LinePlotter(plot), map[string]string{"ex": ex})
}

func solve(lines []string) ([]string, error) {
	g := parseGraph(lines)
	l, r, _ := findCut(g)
	return glue.Ints(l * r), nil
}

func findCut(g *sparseGraph) (leftSize, rightSize int, edges [3][2]int) {
	cf := &cutFinder{
		p:    make([]int, g.len()),
		q:    util.MakeQueue[int](512),
		seen: util.MakeFixedBitmap1D(g.len()),
	}
	for src := 0; src < g.len(); src++ {
		firstPath := cf.longestPath(g, src)
		for i := 0; i < len(firstPath)-1; i++ {
			g.delEdge(firstPath[i], firstPath[i+1])
			secondPath := cf.shortestPath(g, firstPath[i], firstPath[i+1])
			for j := 0; j < len(secondPath)-1; j++ {
				g.delEdge(secondPath[j], secondPath[j+1])
				thirdPath := cf.shortestPath(g, secondPath[j], secondPath[j+1])
				for k := 0; k < len(thirdPath)-1; k++ {
					g.delEdge(thirdPath[k], thirdPath[k+1])
					ls := cf.componentSize(g, thirdPath[k], thirdPath[k+1])
					if ls != -1 {
						rs := g.len() - ls
						edges = [3][2]int{
							{firstPath[i], firstPath[i+1]},
							{secondPath[j], secondPath[j+1]},
							{thirdPath[k], thirdPath[k+1]},
						}
						for _, e := range edges {
							g.addEdge(e[0], e[1])
						}
						return max(ls, rs), min(ls, rs), edges
					}
					g.addEdge(thirdPath[k], thirdPath[k+1])
				}
				g.addEdge(secondPath[j], secondPath[j+1])
			}
			g.addEdge(firstPath[i], firstPath[i+1])
		}
	}
	return -1, -1, edges
}

type cutFinder struct {
	p    []int
	q    util.Queue[int]
	seen util.FixedBitmap1D
}

func (cf *cutFinder) longestPath(g *sparseGraph, src int) []int {
	cf.p[src] = -1
	cf.q.Clear()
	cf.q.Push(src)
	cf.seen.Clear()
	cf.seen.Set(src)
	at := -1
	for !cf.q.Empty() {
		at = cf.q.Pop()
		for _, n := range g.edges[at] {
			if cf.seen.Get(n) {
				continue
			}
			cf.p[n] = at
			cf.seen.Set(n)
			cf.q.Push(n)
		}
	}
	path := []int{at}
	for cf.p[at] != -1 {
		path = append(path, cf.p[at])
		at = cf.p[at]
	}
	return path
}

func (cf *cutFinder) shortestPath(g *sparseGraph, src, dst int) []int {
	cf.p[src] = -1
	cf.q.Clear()
	cf.q.Push(src)
	cf.seen.Clear()
	cf.seen.Set(src)
	for !cf.q.Empty() {
		at := cf.q.Pop()
		for _, n := range g.edges[at] {
			if n == dst {
				path := []int{dst, at}
				for cf.p[at] != -1 {
					path = append(path, cf.p[at])
					at = cf.p[at]
				}
				return path
			}
			if cf.seen.Get(n) {
				continue
			}
			cf.p[n] = at
			cf.seen.Set(n)
			cf.q.Push(n)
		}
	}
	return nil
}

func (cf *cutFinder) componentSize(g *sparseGraph, src, dst int) (reachable int) {
	cf.q.Clear()
	cf.q.Push(src)
	cf.seen.Clear()
	cf.seen.Set(src)
	reachable = 1
	for !cf.q.Empty() {
		at := cf.q.Pop()
		for _, n := range g.edges[at] {
			if n == dst {
				return -1
			}
			if cf.seen.Get(n) {
				continue
			}
			cf.seen.Set(n)
			reachable++
			cf.q.Push(n)
		}
	}
	return reachable
}

type sparseGraph struct {
	edges  [][]int
	labels []string
}

func (g *sparseGraph) len() int { return len(g.edges) }

func (g *sparseGraph) addEdge(a, b int) {
	if n := max(a, b); n+1 > len(g.edges) {
		g.edges = append(g.edges, make([][]int, n+1-len(g.edges))...)
	}
	g.edges[a] = append(g.edges[a], b)
	g.edges[b] = append(g.edges[b], a)
}

func (g *sparseGraph) delEdge(a, b int) {
	ai, bi := slices.Index(g.edges[a], b), slices.Index(g.edges[b], a)
	an, bn := len(g.edges[a]), len(g.edges[b])
	if ai != an-1 {
		g.edges[a][ai] = g.edges[a][an-1]
	}
	if bi != bn-1 {
		g.edges[b][bi] = g.edges[b][bn-1]
	}
	g.edges[a] = g.edges[a][:an-1]
	g.edges[b] = g.edges[b][:bn-1]
}

func parseGraph(lines []string) (g *sparseGraph) {
	g = &sparseGraph{}
	v := make(util.LabelMap)
	for _, line := range lines {
		src, dsts, _ := strings.Cut(line, ": ")
		srcV := v.Get(src)
		for s := util.Splitter(dsts); !s.Empty(); {
			dst := s.Next(" ")
			dstV := v.Get(dst)
			g.addEdge(srcV, dstV)
		}
	}
	g.labels = make([]string, g.len())
	for k, v := range v {
		g.labels[v] = k
	}
	return g
}

// plotting

var ex = strings.TrimPrefix(`
jqt: rhn xhk nvd
rsh: frs pzl lsr
xhk: hfx
cmg: qnr nvd lhk bvb
rhn: xhk bvb hfx
bvb: xhk hfx
pzl: lsr hfx nvd
qnr: nvd
ntq: jqt hfx bvb xhk
nvd: lhk
lsr: lhk
rzs: qnr cmg lsr rsh
frs: qnr lhk lsr
`, "\n")

func plot(lines []string, w io.Writer) error {
	g := parseGraph(lines)
	_, _, edges := findCut(g)
	fmt.Fprintln(w, "graph G {")
	for v := 0; v < g.len(); v++ {
		fmt.Fprintf(w, "  v%d [label=\"%s\"];\n", v, g.labels[v])
	}
	for u := 0; u < g.len(); u++ {
		for _, v := range g.edges[u] {
			if v > u {
				fmt.Fprintf(w, "  v%d -- v%d", u, v)
				if slices.Contains(edges[:], [2]int{u, v}) || slices.Contains(edges[:], [2]int{v, u}) {
					fmt.Fprintf(w, " [color=\"red\"]")
				}
				fmt.Fprintln(w, ";")
			}
		}
	}
	fmt.Fprintln(w, "}")
	return nil
}
