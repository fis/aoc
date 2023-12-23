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
)

func init() {
	glue.RegisterSolver(2023, 23, glue.FixedLevelSolver(solve))
	glue.RegisterPlotter(2023, 23, plotter{}, map[string]string{"ex": ex})
}

func solve(l *util.FixedLevel) ([]string, error) {
	g, startV, endV := deconstruct(l)
	p1 := longestPath(g)
	g.MakeUndirected()
	p2 := evenLongestPath(g, startV, endV)
	return glue.Ints(p1, p2), nil
}

func longestPath(g *util.Graph) (longest int) {
	d := make([]int, g.Len())
	for i := range d {
		d[i] = math.MaxInt
	}
	order := g.TopoSortV(true)
	d[order[0]] = 0
	for _, u := range order {
		g.RangeSuccV(u, func(v int) bool {
			w := g.W(u, v)
			if d[v] > d[u]-w {
				d[v] = d[u] - w
			}
			return true
		})
	}
	return -fn.Min(d)
}

func evenLongestPath(g *util.Graph, startV, endV int) (longest int) {
	sg := make([]vertex, g.Len())
	for u := range sg {
		g.RangeSuccV(u, func(v int) bool {
			if v != startV && v != endV {
				d := sg[u].degree
				sg[u].next[d].v, sg[u].next[d].w = uint32(v), uint32(g.W(u, v))
				sg[u].degree = d + 1
			}
			return true
		})
	}
	firstV, lastV := g.SuccV(startV, 0), g.PredV(endV, 0)
	wS, wE := g.W(startV, firstV), g.W(lastV, endV)
	return wS + int(bruteForce(sg, uint32(firstV), 0, uint32(lastV))) + wE
}

type vertex struct {
	degree uint32
	next   [4]struct{ v, w uint32 }
	seen   bool
}

func bruteForce(sg []vertex, atV, d, toV uint32) uint32 {
	if atV == toV {
		return d
	}
	sg[atV].seen = true
	maxD := uint32(0)
	for _, next := range sg[atV].next[:sg[atV].degree] {
		if sg[next.v].seen {
			continue
		}
		if nextD := bruteForce(sg, next.v, d+next.w, toV); nextD > maxD {
			maxD = nextD
		}
	}
	sg[atV].seen = false
	return maxD
}

var gateExits = []struct {
	d util.P
	c byte
}{{util.P{1, 0}, '>'}, {util.P{0, 1}, 'v'}, {util.P{-1, 0}, '<'}, {util.P{0, -1}, '^'}}

func deconstruct(l *util.FixedLevel) (g *util.Graph, startV, endV int) {
	g = &util.Graph{}
	startV = g.V("qS")
	endV = g.V("qE")

	type path struct {
		at   util.P
		d    util.P
		srcV int
	}
	paths := []path{{at: util.P{bytes.IndexByte(l.Row(0), '.'), 1}, d: util.P{0, 1}, srcV: startV}}
	gates := make(map[util.P]int)

nextPath:
	for len(paths) > 0 {
		p := paths[len(paths)-1]
		paths = paths[:len(paths)-1]

		at, d, steps := p.at, p.d, 0
	nextStep:
		for {
			steps++
			for _, newD := range []util.P{d, {-d.Y, d.X}, {d.Y, -d.X}} {
				n := at.Add(newD)
				if !l.InBounds(n.X, n.Y) {
					g.AddEdgeWV(p.srcV, endV, steps)
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
						dstV = g.V(fmt.Sprintf("q%d", g.Len()-1))
						gates[gate] = dstV
						for _, exit := range gateExits {
							if ep := gate.Add(exit.d); l.At(ep.X, ep.Y) == exit.c {
								paths = append(paths, path{at: ep, d: exit.d, srcV: dstV})
							}
						}
					}
					g.AddEdgeWV(p.srcV, dstV, steps+2)
					continue nextPath
				}
			}
			panic("dead end")
		}
	}

	return g, startV, endV
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

type plotter struct{}

func (plotter) Plot(r io.Reader, w io.Writer) error {
	input, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	l := util.ParseFixedLevel(input)
	g, _, _ := deconstruct(l)
	g.WriteDOT(w, "G", nil, nil)
	return nil
}
