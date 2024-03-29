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
	"math"
	"slices"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2023, 25, glue.LineSolver(solve))
	glue.RegisterPlotter(2023, 25, "a", glue.LinePlotter(plotCut), map[string]string{"ex": ex})
	glue.RegisterPlotter(2023, 25, "b", glue.LinePlotter(plotEdges), map[string]string{"ex": ex})
}

func solve(lines []string) ([]string, error) {
	g := parseGraph(lines)
	l, r := findCutSizes(g)
	return glue.Ints(l * r), nil
}

func findCutSizes(g *sparseGraph) (leftSize, rightSize int) {
	cf := newCutFinder(g)
	src := 0
	for {
		dst, firstPath := cf.firstPath(src)
		if dst == -1 {
			break
		}
		for i := 0; i < len(firstPath)-1; i++ {
			g.delEdge(firstPath[i], firstPath[i+1])
		}
		secondPath := cf.shortestPath(src, dst)
		for i := 0; i < len(secondPath)-1; i++ {
			g.delEdge(secondPath[i], secondPath[i+1])
		}
		thirdPath := cf.shortestPath(src, dst)
		for i := 0; i < len(thirdPath)-1; i++ {
			g.delEdge(thirdPath[i], thirdPath[i+1])
			ls := cf.componentSize(thirdPath[i], thirdPath[i+1])
			if ls != -1 {
				rs := g.len() - ls
				return max(ls, rs), min(ls, rs)
			}
			g.addEdge(thirdPath[i], thirdPath[i+1])
		}
		paths := [3][]int{firstPath, secondPath}
		for _, p := range paths {
			for j := 0; j < len(p)-1; j++ {
				g.addEdge(p[j], p[j+1])
			}
		}
	}
	// fallback option
	edges := findCutEdges(g)
	for _, e := range edges {
		g.delEdge(e[0], e[1])
	}
	ls, rs := cf.componentSize(edges[0][0], edges[0][1]), cf.componentSize(edges[0][1], edges[0][0])
	return max(ls, rs), min(ls, rs)
}

func findCutEdges(g *sparseGraph) (edges [3][2]int) {
	cf := newCutFinder(g)
	src := 0
	for {
		dst, firstPath := cf.firstPath(src)
		if dst == -1 {
			panic("cut not found")
		}
		for i := 0; i < len(firstPath)-1; i++ {
			g.delEdge(firstPath[i], firstPath[i+1])
			secondPath := cf.shortestPath(firstPath[i], firstPath[i+1])
			for j := 0; j < len(secondPath)-1; j++ {
				g.delEdge(secondPath[j], secondPath[j+1])
				thirdPath := cf.shortestPath(secondPath[j], secondPath[j+1])
				for k := 0; k < len(thirdPath)-1; k++ {
					g.delEdge(thirdPath[k], thirdPath[k+1])
					if !cf.hasPath(thirdPath[k], thirdPath[k+1]) {
						edges = [3][2]int{
							{firstPath[i], firstPath[i+1]},
							{secondPath[j], secondPath[j+1]},
							{thirdPath[k], thirdPath[k+1]},
						}
						for _, e := range edges {
							g.addEdge(e[0], e[1])
						}
						return edges
					}
					g.addEdge(thirdPath[k], thirdPath[k+1])
				}
				g.addEdge(secondPath[j], secondPath[j+1])
			}
			g.addEdge(firstPath[i], firstPath[i+1])
		}
	}
}

type cutFinder struct {
	g       *sparseGraph
	p       []int
	q       util.Queue[int]
	seen    util.FixedBitmap1D
	nextDst int
}

func newCutFinder(g *sparseGraph) *cutFinder {
	return &cutFinder{
		g:       g,
		p:       make([]int, g.len()),
		q:       util.MakeQueue[int](512),
		seen:    util.MakeFixedBitmap1D(g.len()),
		nextDst: -1,
	}
}

func (cf *cutFinder) firstPath(src int) (dst int, path []int) {
	switch cf.nextDst {
	case -1:
		cf.nextDst = 0
		path = cf.longestPath(src)
		return path[0], path
	case src:
		cf.nextDst++
		fallthrough
	default:
		if cf.nextDst == cf.g.len() {
			return -1, nil
		}
		dst = cf.nextDst
		cf.nextDst++
		return dst, cf.shortestPath(src, dst)
	}
}

func (cf *cutFinder) longestPath(src int) []int {
	g := cf.g
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

func (cf *cutFinder) shortestPath(src, dst int) []int {
	g := cf.g
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

func (cf *cutFinder) hasPath(src, dst int) bool {
	g := cf.g
	cf.q.Clear()
	cf.q.Push(src)
	cf.seen.Clear()
	cf.seen.Set(src)
	for !cf.q.Empty() {
		at := cf.q.Pop()
		for _, n := range g.edges[at] {
			if n == dst {
				return true
			}
			if cf.seen.Get(n) {
				continue
			}
			cf.seen.Set(n)
			cf.q.Push(n)
		}
	}
	return false
}

func (cf *cutFinder) componentSize(src, dst int) (reachable int) {
	g := cf.g
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

func (cf *cutFinder) edgeUsage(src int, counts map[[2]int]int) {
	g := cf.g
	cf.p[src] = -1
	cf.q.Clear()
	cf.q.Push(src)
	cf.seen.Clear()
	cf.seen.Set(src)
	for !cf.q.Empty() {
		at := cf.q.Pop()
		for _, n := range g.edges[at] {
			if cf.seen.Get(n) {
				continue
			}
			cf.p[n] = at
			cf.seen.Set(n)
			cf.q.Push(n)
		}
	}
	for end := 0; end < g.len(); end++ {
		v := end
		for cf.p[v] != -1 {
			u := cf.p[v]
			e := [2]int{min(u, v), max(u, v)}
			counts[e] = counts[e] + 1
			v = u
		}
	}
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

func plotCut(lines []string, w io.Writer) error {
	g := parseGraph(lines)
	edges := findCutEdges(g)
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

func plotEdges(lines []string, w io.Writer) error {
	g := parseGraph(lines)
	counts := make(map[[2]int]int)
	cf := newCutFinder(g)
	for v := 0; v < g.len(); v++ {
		cf.edgeUsage(v, counts)
	}
	vCounts, minC, maxC := make(map[int]int), math.MaxInt, math.MinInt
	for e, c := range counts {
		minC = min(minC, c)
		maxC = max(maxC, c)
		vCounts[e[0]] = vCounts[e[0]] + c
		vCounts[e[1]] = vCounts[e[1]] + c
	}
	fmt.Println("graph G {")
	fmt.Printf("  bgcolor=\"black\";")
	for v := 0; v < g.len(); v++ {
		c := vCounts[v] / len(g.edges[v])
		r, g, b := colorOf(c, minC, maxC)
		fmt.Printf("  v%d [label=\"\",shape=point,color=\"#%02x%02x%02x\"];\n", v, r, g, b)
	}
	for u := 0; u < g.len(); u++ {
		for _, v := range g.edges[u] {
			if v > u {
				c := counts[[2]int{u, v}]
				r, g, b := colorOf(c, minC, maxC)
				fmt.Printf("  v%d -- v%d [color=\"#%02x%02x%02x\"];\n", u, v, r, g, b)
			}
		}
	}
	fmt.Println("}")
	return nil
}

func colorOf(c, minC, maxC int) (r, g, b int) {
	lc, lmin, lmax := math.Log(float64(c)), math.Log(float64(minC)), math.Log(float64(maxC))
	v := 3 * float64(lc-lmin) / float64(lmax-lmin)
	switch {
	case v >= 2:
		r, g, b = 255, 255, colorComp(v-2)
	case v >= 1:
		r, g, b = 255, colorComp(v-1), 0
	default:
		r, g, b = colorComp(v), 0, 0
	}
	return r, g, b
}

func colorComp(v float64) int {
	return int(math.Round(255 * v))
}
