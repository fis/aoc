// Package day23 solves AoC 2023 day 23.
package day23

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/fn"
	"github.com/fis/aoc/util/graph"
)

func init() {
	glue.RegisterSolver(2023, 23, glue.FixedLevelSolver(solve))
	glue.RegisterPlotter(2023, 23, "a", plotter{tr: tracerA{}}, map[string]string{"ex": ex})
	glue.RegisterPlotter(2023, 23, "b", plotter{tr: tracerB{}, undir: true}, map[string]string{"ex": ex})
}

func solve(l *util.FixedLevel) ([]string, error) {
	g, startV, endV := deconstruct(l)
	p1 := longestPath(g)
	p2 := unsafeLongestPath(g, startV, endV)
	return glue.Ints(p1, p2), nil
}

func longestPath(g *graph.SparseW) (longest int) {
	d := make([]int, g.Len())
	for i := range d {
		d[i] = math.MaxInt
	}
	order := g.TopoSort(true)
	d[order[0]] = 0
	for _, u := range order {
		for v, w := range g.SuccW(u) {
			if d[v] > d[u]-w {
				d[v] = d[u] - w
			}
		}
	}
	return -fn.Min(d)
}

type vertex struct {
	degree uint32
	next   [4]struct{ v, w uint32 }
	seen   bool
}

var gateExits = []struct {
	d util.P
	c byte
}{{util.P{1, 0}, '>'}, {util.P{0, 1}, 'v'}, {util.P{-1, 0}, '<'}, {util.P{0, -1}, '^'}}

func deconstruct(l *util.FixedLevel) (g *graph.SparseW, startV, endV int) {
	gb := graph.NewBuilder()
	startV = gb.V("qS")
	endV = gb.V("qE")

	type path struct {
		at   util.P
		d    util.P
		srcV int
	}
	paths := util.QueueOf[path](16, path{at: util.P{bytes.IndexByte(l.Row(0), '.'), 1}, d: util.P{0, 1}, srcV: startV})
	gates := make(map[util.P]int)

nextPath:
	for !paths.Empty() {
		p := paths.Pop()

		at, d, steps := p.at, p.d, 0
	nextStep:
		for {
			steps++
			for _, newD := range []util.P{d, {-d.Y, d.X}, {d.Y, -d.X}} {
				n := at.Add(newD)
				if !l.InBounds(n.X, n.Y) {
					gb.AddEdgeW(p.srcV, endV, steps)
					continue nextPath
				}
				switch l.At(n.X, n.Y) {
				case '.':
					at, d = n, newD
					continue nextStep
				case '>', 'v', '<', '^':
					gate := n.Add(newD)
					dstV, ok := gates[gate]
					if !ok {
						dstV = gb.V(fmt.Sprintf("q%d", gb.Len()-1))
						gates[gate] = dstV
						for _, exit := range gateExits {
							if ep := gate.Add(exit.d); l.At(ep.X, ep.Y) == exit.c {
								paths.Push(path{at: ep, d: exit.d, srcV: dstV})
							}
						}
					}
					gb.AddEdgeW(p.srcV, dstV, steps+2)
					continue nextPath
				}
			}
			panic("dead end")
		}
	}

	return gb.SparseDigraphW(), startV, endV
}

// plotting

var ex = strings.TrimPrefix(`
#.#####################
#.......#########...###
#######.#########.#.###
###.....#.>.>.###.#.###
###v#####.#v#.###.#.###
###.>...#.#.#.....#...#
###v###.#.#.#########.#
###...#.#.#.......#...#
#####.#.#.#######.#.###
#.....#.#.#.......#...#
#.#####.#.#.#########v#
#.#...#...#...###...>.#
#.#.#v#######v###.###v#
#...#.>.#...>.>.#.###.#
#####v#.#.###v#.#.###.#
#.....#...#...#.#.#...#
#.#########.###.#.#.###
#...###...#...#...#.###
###.###.#.###v#####v###
#...#...#.#.>.>.#.>.###
#.###.###.#.###.#.#v###
#.....###...###...#...#
#####################.#
`, "\n")

type plotter struct {
	tr    tracer
	undir bool
}

func (p plotter) Plot(r io.Reader, w io.Writer) error {
	const (
		highlightColor = `"#db4437"`
		subdueColor    = `"gray75"`
	)

	input, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	l := util.ParseFixedLevel(input)
	g, startV, endV := deconstruct(l)
	longestV, longestE := p.tr.trace(g, startV, endV)
	nodeAttr := func(v int) map[string]string {
		m := map[string]string{}
		if _, ok := longestV[v]; ok {
			m["color"] = highlightColor
			m["fontcolor"] = highlightColor
		} else if p.undir {
			m["color"] = subdueColor
			m["fontcolor"] = subdueColor
		}
		return m
	}
	edgeAttr := func(u, v int) map[string]string {
		m := map[string]string{}
		_, ok1 := longestE[[2]int{u, v}]
		_, ok2 := longestE[[2]int{v, u}]
		if ok1 || ok2 {
			m["color"] = highlightColor
			m["fontcolor"] = highlightColor
		} else if p.undir {
			m["color"] = subdueColor
			m["fontcolor"] = subdueColor
		}
		if p.undir {
			m["dir"] = "none"
		}
		return m
	}
	graph.WriteDOT(g, w, "G", true, nodeAttr, edgeAttr)
	return nil
}

type tracer interface {
	trace(g *graph.SparseW, startV, endV int) (longestV map[int]struct{}, longestE map[[2]int]struct{})
}

type tracerA struct{}

func (tracerA) trace(g *graph.SparseW, startV, endV int) (longestV map[int]struct{}, longestE map[[2]int]struct{}) {
	d := make([]int, g.Len())
	p := make([]int, g.Len())
	for i := range d {
		d[i] = math.MaxInt
		p[i] = -1
	}
	order := g.TopoSort(true)
	d[order[0]] = 0
	for _, u := range order {
		for v, w := range g.SuccW(u) {
			if d[v] > d[u]-w {
				d[v] = d[u] - w
				p[v] = u
			}
		}
	}
	longestV = make(map[int]struct{})
	longestE = make(map[[2]int]struct{})
	for v := order[len(order)-1]; v != -1; v = p[v] {
		longestV[v] = struct{}{}
		if u := p[v]; u != -1 {
			longestE[[2]int{u, v}] = struct{}{}
		}
	}
	return longestV, longestE
}

type tracerB struct {
	bestD uint32
	chain *backRef
}

type backRef struct {
	v uint32
	p *backRef
}

func (tr tracerB) trace(g *graph.SparseW, startV, endV int) (longestV map[int]struct{}, longestE map[[2]int]struct{}) {
	sg := make([]vertex, g.Len())
	for u := range sg {
		for v, w := range g.SuccW(u) {
			if u != startV && u != endV && v != startV && v != endV {
				d := sg[u].degree
				sg[u].next[d].v, sg[u].next[d].w = uint32(v), uint32(w)
				sg[u].degree = d + 1
				d = sg[v].degree
				sg[v].next[d].v, sg[v].next[d].w = uint32(u), uint32(w)
				sg[v].degree = d + 1
			}
		}
	}
	firstV, lastV := g.SuccI(startV, 0), g.PredI(endV, 0)
	tr.bruteForce(sg, uint32(firstV), 0, uint32(lastV), nil)
	longestV = map[int]struct{}{startV: {}, endV: {}}
	longestE = map[[2]int]struct{}{{startV, firstV}: {}, {lastV, endV}: {}}
	for at := tr.chain; at != nil; at = at.p {
		longestV[int(at.v)] = struct{}{}
		if at.p != nil {
			longestE[[2]int{int(at.p.v), int(at.v)}] = struct{}{}
		}
	}
	return longestV, longestE
}

func (tr *tracerB) bruteForce(sg []vertex, atV, d, toV uint32, br *backRef) {
	br = &backRef{v: atV, p: br}
	if atV == toV {
		if d > tr.bestD {
			tr.bestD = d
			tr.chain = br
		}
		return
	}
	sg[atV].seen = true
	for _, next := range sg[atV].next[:sg[atV].degree] {
		if sg[next.v].seen {
			continue
		}
		tr.bruteForce(sg, next.v, d+next.w, toV, br)
	}
	sg[atV].seen = false
}
