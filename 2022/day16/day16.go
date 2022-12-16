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
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/fn"
	"golang.org/x/exp/maps"
)

func init() {
	glue.RegisterSolver(2022, 16, glue.LineSolver(glue.WithParser(parseValveScan, solve)))
	glue.RegisterPlotter(2022, 16, glue.LinePlotter(plot), map[string]string{"ex": strings.Join(ex, "\n")})
}

func solve(scan []valveScan) ([]string, error) {
	sum := preprocess(scan)
	p1 := releasePressure(sum, 30)
	p2 := releasePressure2(sum, 26)
	return glue.Ints(p1, p2), nil
}

type valveScan struct {
	name     string
	flowRate int
	tunnels  []string
}

type valveSummary struct {
	flowRates []int
	initDist  []int
	dist      [][]int
}

func preprocess(scan []valveScan) (sum valveSummary) {
	allValves, nonzeroValves := make(map[string]*valveScan), make(map[string]int)
	for i, v := range scan {
		allValves[v.name] = &scan[i]
		if v.flowRate != 0 {
			nonzeroValves[v.name] = len(sum.flowRates)
			sum.flowRates = append(sum.flowRates, v.flowRate)
		}
	}

	n := len(sum.flowRates)
	sum.initDist = make([]int, n)
	sum.dist = fn.MapRange(0, n, func(int) []int { return make([]int, n) })

	for _, from := range append(maps.Keys(nonzeroValves), "AA") {
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
					sum.dist[fromN][pn] = p.d
				} else {
					sum.initDist[pn] = p.d
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

func parseValveScan(line string) (vs valveScan, err error) {
	tail, ok := util.CheckPrefix(line, "Valve ")
	if !ok {
		return valveScan{}, fmt.Errorf("expected `Valve ` in %q", line)
	}
	vs.name, tail = util.NextWord(tail)
	tail, ok = util.CheckPrefix(tail, " has flow rate=")
	if !ok {
		return valveScan{}, fmt.Errorf("expected ` has flow rate=` in %q", line)
	}
	vs.flowRate, ok, tail = util.NextInt(tail)
	if !ok {
		return valveScan{}, fmt.Errorf("expected an integer in %q", line)
	}
	if tail, ok = util.CheckPrefix(tail, "; tunnel leads to valve "); !ok {
		tail, ok = util.CheckPrefix(tail, "; tunnels lead to valves ")
	}
	if !ok {
		return valveScan{}, fmt.Errorf("expected `; tunnels? leads? to valves? in %q", line)
	}
	vs.tunnels = strings.Split(tail, ", ")
	return vs, nil
}

// plotting

var ex = []string{
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
	scan, err := fn.MapE(lines, parseValveScan)
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
