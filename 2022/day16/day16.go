// Copyright 2022 Google LLC
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

// Package day16 solves AoC 2022 day 16.
package day16

import (
	"fmt"
	"io"
	"math/bits"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/fn"
)

func init() {
	glue.RegisterSolver(2022, 16, glue.LineSolver(glue.WithParser(ParseValveScan, solve)))
	glue.RegisterPlotter(2022, 16, "", glue.LinePlotter(plot), map[string]string{"ex": strings.Join(ExampleScan, "\n")})
}

func solve(scan []ValveScan) ([]string, error) {
	sum := Preprocess(scan)
	p1 := findOne(sum, 30)
	p2 := findTwo(sum, 26)
	return glue.Ints(p1, p2), nil
}

type ValveScan struct {
	name     string
	flowRate int
	tunnels  []string
}

type ValveSummary struct {
	FlowRates []int
	InitDist  []int
	Dist      [][]int
}

func Preprocess(scan []ValveScan) (sum ValveSummary) {
	allValves, nonzeroValves := make(map[string]*ValveScan), make(map[string]int)
	for i, v := range scan {
		allValves[v.name] = &scan[i]
		if v.flowRate != 0 {
			nonzeroValves[v.name] = len(sum.FlowRates)
			sum.FlowRates = append(sum.FlowRates, v.flowRate)
		}
	}

	n := len(sum.FlowRates)
	sum.InitDist = make([]int, n)
	sum.Dist = fn.MapRange(0, n, func(int) []int { return make([]int, n) })

	for _, from := range append(fn.Keys(nonzeroValves), "AA") {
		fromN, fromNonzero := nonzeroValves[from]
		type path struct {
			at string
			d  int
		}
		seen := map[string]struct{}{from: {}}
		q := []path{{from, 0}}
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			if pn, ok := nonzeroValves[p.at]; ok {
				if fromNonzero {
					sum.Dist[fromN][pn] = p.d
				} else {
					sum.InitDist[pn] = p.d
				}
			}
			for _, n := range allValves[p.at].tunnels {
				if _, ok := seen[n]; !ok {
					seen[n] = struct{}{}
					q = append(q, path{n, p.d + 1})
				}
			}
		}
	}

	return sum
}

const (
	nValves    = 15 // using some fixed-size arrays for performance
	nValveSets = 1 << nValves
)

func findOne(sum ValveSummary, maxT int) (bestP int) {
	q := util.NewBucketQ[state](32)
	for i := uint16(0); i < nValves; i++ {
		t := sum.InitDist[i] + 1
		q.Push(t, state{at: i, open: 1 << i, pressure: uint32((maxT - t) * sum.FlowRates[i])})
	}

	seen := [nValves][nValveSets]stateTracker{}
	for q.Len() > 0 {
		pt, p := q.Pop()
		if pr := int(p.pressure); pr > bestP {
			bestP = pr
		}
		for i := uint16(0); i < nValves; i++ {
			open := p.open | (1 << i)
			if open == p.open {
				continue
			}
			t := pt + sum.Dist[p.at][i] + 1
			if t >= maxT {
				continue
			}
			next := state{at: i, open: open, pressure: p.pressure + uint32((maxT-t)*sum.FlowRates[i])}
			if seen[i][open].keep(t, next) {
				q.Push(t, next)
			}
		}
	}

	return bestP
}

func findTwo(sum ValveSummary, maxT int) (bestP int) {
	q := util.NewBucketQ[state](32)
	for i := uint16(0); i < nValves; i++ {
		t := sum.InitDist[i] + 1
		q.Push(t, state{at: i, open: 1 << i, pressure: uint32((maxT - t) * sum.FlowRates[i])})
	}

	seen := [nValves][nValveSets]stateTracker{}
	maxP := [nValveSets]uint32{}
	for q.Len() > 0 {
		pt, p := q.Pop()
		if pr := p.pressure; pr > maxP[p.open] {
			maxP[p.open] = pr
		}
		for i := uint16(0); i < nValves; i++ {
			open := p.open | (1 << i)
			if open == p.open {
				continue
			}
			t := pt + sum.Dist[p.at][i] + 1
			if t >= maxT {
				continue
			}
			next := state{at: i, open: open, pressure: p.pressure + uint32((maxT-t)*sum.FlowRates[i])}
			if seen[i][open].keep(t, next) {
				q.Push(t, next)
			}
		}
	}

	for setSize := 2; setSize <= nValves; setSize++ {
		// There's probably a faster way to iterate integers with a given bit count, but this is fine for 15 bits.
		firstSet := uint16((1 << setSize) - 1)
		lastSet := firstSet << (nValves - setSize)
		for open := firstSet; open <= lastSet; open++ {
			if maxP[open] != 0 || bits.OnesCount16(open) != setSize {
				continue
			}
			for i := 0; i < nValves; i++ {
				lessOpen := open &^ (1 << i)
				if lessOpen != open && maxP[lessOpen] > maxP[open] {
					maxP[open] = maxP[lessOpen]
				}
			}
		}
	}

	for open1 := uint16(0); open1 < nValveSets; open1++ {
		open2 := (nValveSets - 1) ^ open1
		if p := int(maxP[open1] + maxP[open2]); p > bestP {
			bestP = p
		}
	}

	return bestP
}

type state struct {
	at       uint16
	open     uint16
	pressure uint32
}

type stateTracker []stateInfo

type stateInfo struct {
	time     int
	pressure uint32
}

func (t *stateTracker) keep(time int, st state) bool {
	for _, e := range *t {
		if e.time <= time && e.pressure >= st.pressure {
			return false
		}
	}
	nt := (*t)[:0]
	for _, e := range *t {
		if e.time < time || e.pressure > st.pressure {
			nt = append(nt, e)
		}
	}
	*t = append(nt, stateInfo{time, st.pressure})
	return true
}

func ParseValveScan(line string) (vs ValveScan, err error) {
	tail, ok := util.CheckPrefix(line, "Valve ")
	if !ok {
		return ValveScan{}, fmt.Errorf("expected `Valve ` in %q", line)
	}
	vs.name, tail = util.NextWord(tail)
	tail, ok = util.CheckPrefix(tail, " has flow rate=")
	if !ok {
		return ValveScan{}, fmt.Errorf("expected ` has flow rate=` in %q", line)
	}
	vs.flowRate, ok, tail = util.NextInt(tail)
	if !ok {
		return ValveScan{}, fmt.Errorf("expected an integer in %q", line)
	}
	if tail, ok = util.CheckPrefix(tail, "; tunnel leads to valve "); !ok {
		tail, ok = util.CheckPrefix(tail, "; tunnels lead to valves ")
	}
	if !ok {
		return ValveScan{}, fmt.Errorf("expected `; tunnels? leads? to valves? in %q", line)
	}
	vs.tunnels = strings.Split(tail, ", ")
	return vs, nil
}

// plotting

var ExampleScan = []string{
	"Valve AA has flow rate=0; tunnels lead to valves DD, II, BB",
	"Valve BB has flow rate=13; tunnels lead to valves CC, AA",
	"Valve CC has flow rate=2; tunnels lead to valves DD, BB",
	"Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE",
	"Valve EE has flow rate=3; tunnels lead to valves FF, DD",
	"Valve FF has flow rate=0; tunnels lead to valves EE, GG",
	"Valve GG has flow rate=0; tunnels lead to valves FF, HH",
	"Valve HH has flow rate=22; tunnel leads to valve GG",
	"Valve II has flow rate=0; tunnels lead to valves AA, JJ",
	"Valve JJ has flow rate=21; tunnel leads to valve II",
}

func plot(lines []string, w io.Writer) error {
	const (
		startStyle = `,style="filled",fillcolor="#0f9d58",fontcolor="white"`
		stuckStyle = `,style="filled",fillcolor="#dddddd"`
	)
	scan, err := fn.MapE(lines, ParseValveScan)
	if err != nil {
		return err
	}
	fmt.Fprintln(w, "graph G {")
	for _, v := range scan {
		style := fn.If(v.name == "AA", startStyle, fn.If(v.flowRate == 0, stuckStyle, ""))
		fmt.Fprintf(w, "  v%s [label=\"%s\\n%d\"%s];\n", v.name, v.name, v.flowRate, style)
	}
	for _, v := range scan {
		for _, to := range v.tunnels {
			if to > v.name {
				fmt.Fprintf(w, "  v%s -- v%s;\n", v.name, to)
			}
		}
	}
	fmt.Fprintln(w, "}")
	return nil
}
