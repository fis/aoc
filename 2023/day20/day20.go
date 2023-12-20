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

// Package day20 solves AoC 2023 day 20.
package day20

import (
	"fmt"
	"io"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/ix"
)

func init() {
	glue.RegisterSolver(2023, 20, glue.LineSolver(solve))
	glue.RegisterPlotter(2023, 20, glue.LinePlotter(plot), nil)
}

func solve(lines []string) ([]string, error) {
	g, _, err := parseGraph(lines)
	if err != nil {
		return nil, err
	}
	p1 := stepN(g, 1000)
	p2, err := analyze(g)
	if err != nil {
		return nil, err
	}
	return glue.Ints(p1, p2), nil
}

func stepN(g *graph, steps int) int {
	state := make([]uint16, len(g.modules))
	q := util.MakeQueue[pulse](64)
	low, high := 0, 0
	for i := 0; i < steps; i++ {
		q.Clear()
		l, h := step(g, state, &q)
		low, high = low+l, high+h
	}
	return low * high
}

func step(g *graph, state []uint16, q *util.Queue[pulse]) (lowCount, highCount int) {
	lowCount = 1
	for _, e := range g.broadcast {
		q.Push(pulse{high: false, e: e})
	}
	for !q.Empty() {
		p := q.Pop()
		if p.high {
			highCount++
		} else {
			lowCount++
		}
		dst := p.e.dst
		dstM := &g.modules[dst]
		switch dstM.typ {
		case modFlipFlop:
			if !p.high {
				state[dst] ^= 1
				high := state[dst] == 1
				for _, e := range dstM.outputs {
					q.Push(pulse{high: high, e: e})
				}
			}
		case modConjunction:
			if p.high {
				state[dst] |= 1 << p.e.idx
			} else {
				state[dst] &^= 1 << p.e.idx
			}
			high := state[dst] != 1<<dstM.numInputs-1
			for _, e := range dstM.outputs {
				q.Push(pulse{high: high, e: e})
			}
		}
	}
	return lowCount, highCount
}

func analyze(g *graph) (lcm int, err error) {
	counts := make([]int, len(g.broadcast))
	chain := make([]byte, len(g.modules))
	for i, e := range g.broadcast {
		chain = chain[:0]
		at, j := e.dst, 0
		for g.modules[at].typ == modFlipFlop {
			m := &g.modules[at]
			switch len(m.outputs) {
			case 1:
				at = m.outputs[0].dst
				if g.modules[at].typ == '&' {
					counts[i] |= 1 << j
				}
			case 2:
				o1, o2 := m.outputs[0].dst, m.outputs[1].dst
				m1, m2 := &g.modules[o1], &g.modules[o2]
				if m1.typ != modFlipFlop {
					o1, m1, m2 = o2, m2, m1
				}
				if m1.typ != modFlipFlop || m2.typ != modConjunction {
					return 0, fmt.Errorf("unexpected types of outputs for a chain flip-flop: %q %q", m1.typ, m2.typ)
				}
				at = o1
				counts[i] |= 1 << j
			default:
				return 0, fmt.Errorf("unexpected number of outputs for a chain flip-flop: %d", len(m.outputs))
			}
			j++
		}
	}
	lcm = 1
	for _, c := range counts {
		lcm = ix.LCM(lcm, c)
	}
	return lcm, nil
}

type graph struct {
	modules   []module
	broadcast []edge
}

type module struct {
	typ       byte
	numInputs byte
	outputs   []edge
}

type edge struct {
	dst int
	idx int
}

type pulse struct {
	high bool
	e    edge
}

const (
	modUntyped     = 0
	modFlipFlop    = '%'
	modConjunction = '&'
)

func parseGraph(lines []string) (*graph, util.LabelMap, error) {
	modules := make([]module, len(lines))
	broadcast := []edge(nil)

	labels := make(util.LabelMap)
	for _, line := range lines {
		srcLabel, dsts, ok := strings.Cut(line, " -> ")
		if !ok {
			return nil, nil, fmt.Errorf("missing '->' separator: %q", line)
		}

		s := util.Splitter(dsts)
		edges := make([]edge, s.Count(", "))
		for i := range edges {
			dstLabel := s.Next(", ")
			dst := labels.Get(dstLabel)
			if modules[dst].numInputs == 16 {
				return nil, nil, fmt.Errorf("too many inputs (%d) for module %s", modules[dst].numInputs, dstLabel)
			}
			edges[i] = edge{dst: dst, idx: int(modules[dst].numInputs)}
			modules[dst].numInputs++
		}

		if srcLabel == "broadcaster" {
			broadcast = edges
			continue
		}

		if len(srcLabel) < 2 || (srcLabel[0] != modFlipFlop && srcLabel[0] != modConjunction) {
			return nil, nil, fmt.Errorf("bad source: %q", srcLabel)
		}
		src := labels.Get(srcLabel[1:])
		modules[src].typ = srcLabel[0]
		modules[src].outputs = edges
	}

	return &graph{modules: modules, broadcast: broadcast}, labels, nil
}

// plotting

func plot(lines []string, w io.Writer) error {
	g, labelMap, err := parseGraph(lines)
	if err != nil {
		return err
	}
	labels := make([]string, len(labelMap))
	for label, idx := range labelMap {
		labels[idx] = label
	}

	fmt.Fprintln(w, "digraph G {")
	fmt.Fprintln(w, "  bc [label=\"BC\",shape=doublecircle];")
	for _, e := range g.broadcast {
		fmt.Fprintf(w, "  bc -> n%d;\n", e.dst)
	}
	for i, m := range g.modules {
		attr := ""
		switch m.typ {
		case modUntyped:
			fmt.Fprintf(w, "  n%d [label=\"%s\",shape=doublecircle];", i, labels[i])
		case modFlipFlop:
			attr = "color=\"#4285f4\""
			fmt.Fprintf(w, "  n%d [label=\"%c%s\",shape=circle,%s];\n", i, m.typ, labels[i], attr)
		case modConjunction:
			attr = "color=\"#db4437\""
			fmt.Fprintf(w, "  n%d [label=\"%c%s\",shape=circle,%s];\n", i, m.typ, labels[i], attr)
		}
		for _, e := range m.outputs {
			fmt.Fprintf(w, "  n%d -> n%d [%s];\n", i, e.dst, attr)
		}
	}
	fmt.Fprintln(w, "}")

	return nil
}
