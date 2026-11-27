// Package day11 solves AoC 2025 day 11.
package day11

import (
	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/graph"
)

func init() {
	glue.RegisterSolver(2025, 11, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	g := buildGraph(lines)
	p1 := newPathCounter(g, "you").countTo("out")
	p2 := newValidPathCounter(g, "svr").countTo("out")
	return glue.Ints(p1, p2), nil
}

func buildGraph(lines []string) *graph.Sparse {
	gb := graph.NewBuilder()
	for _, line := range lines {
		words := util.Words(line)
		if len(words) < 2 {
			continue
		}
		tail := words[0]
		if len(tail) < 2 || tail[len(tail)-1] != ':' {
			continue
		}
		tail = tail[:len(tail)-1]
		for _, head := range words[1:] {
			gb.AddEdgeL(head, tail) // reversed direction because we want to iterate over predecessors, not successors
		}
	}
	return gb.SparseDigraph()
}

type pathCounter struct {
	g *graph.Sparse
	p []int
}

func newPathCounter(g *graph.Sparse, src string) *pathCounter {
	p := make([]int, g.Len())
	srcV, _ := g.V(src)
	p[srcV] = 1
	return &pathCounter{g: g, p: p}
}

func (pc *pathCounter) countTo(dst string) (paths int) {
	dstV, ok := pc.g.V(dst)
	if !ok {
		return -1
	}
	return pc.countToV(dstV)
}

func (pc *pathCounter) countToV(v int) (paths int) {
	if p := pc.p[v]; p != 0 {
		if p < 0 {
			p = 0
		}
		return p
	}
	for u := range pc.g.Succ(v) { // actually predecessors
		paths += pc.countToV(u)
	}
	if paths == 0 {
		pc.p[v] = -1
	} else {
		pc.p[v] = paths
	}
	return paths
}

type validPathCounter struct {
	g          *graph.Sparse
	fftV, dacV int
	p          []pathInfo
}

type pathInfo struct {
	set      bool
	total    int
	withFFT  int
	withDAC  int
	withBoth int
}

func newValidPathCounter(g *graph.Sparse, src string) *validPathCounter {
	p := make([]pathInfo, g.Len())
	srcV, _ := g.V(src)
	fftV, _ := g.V("fft")
	dacV, _ := g.V("dac")
	p[srcV] = pathInfo{set: true, total: 1, withFFT: 0, withDAC: 0, withBoth: 0}
	return &validPathCounter{g: g, fftV: fftV, dacV: dacV, p: p}
}

func (pc *validPathCounter) countTo(dst string) (paths int) {
	dstV, ok := pc.g.V(dst)
	if !ok {
		return -1
	}
	_, _, _, paths = pc.countToV(dstV)
	return paths
}

func (pc *validPathCounter) countToV(v int) (total, withFFT, withDAC, withBoth int) {
	if pi := pc.p[v]; pi.set {
		return pi.total, pi.withFFT, pi.withDAC, pi.withBoth
	}
	for u := range pc.g.Succ(v) { // actually predecessors
		uTotal, uFFT, uDAC, uBoth := pc.countToV(u)
		total, withFFT, withDAC, withBoth = total+uTotal, withFFT+uFFT, withDAC+uDAC, withBoth+uBoth
	}
	switch v {
	case pc.fftV:
		withFFT, withBoth = total, withDAC
	case pc.dacV:
		withDAC, withBoth = total, withFFT
	}
	pc.p[v] = pathInfo{set: true, total: total, withFFT: withFFT, withDAC: withDAC, withBoth: withBoth}
	return total, withFFT, withDAC, withBoth
}
