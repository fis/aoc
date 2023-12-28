// Copyright 2019 Google LLC
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

// Package day14 solves AoC 2019 day 14.
package day14

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util/fn"
	"github.com/fis/aoc/util/graph"
)

func init() {
	glue.RegisterSolver(2019, 14, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	reactions := parseReactions(lines)

	part1 := ore(1, reactions)
	part2 := maxFuel(1000000000000, reactions)

	return glue.Ints(part1, part2), nil
}

type pile struct {
	name string
	q    int
}

type reaction struct {
	out pile
	in  []pile
}

func ore(wantFuel int, reactions map[string]reaction) int {
	order := reactionOrder(reactions)
	want := map[string]int{"FUEL": wantFuel}
	return oreFor(want, order, reactions)
}

func maxFuel(ore int, reactions map[string]reaction) int {
	order := reactionOrder(reactions)
	start, end := 1, ore+1
	for end-start >= 2 {
		mid := start + (end-start)/2
		got := oreFor(map[string]int{"FUEL": mid}, order, reactions)
		if got > ore {
			end = mid
		} else {
			start = mid
		}
	}
	return start
}

func reactionOrder(reactions map[string]reaction) []string {
	gb := graph.NewBuilder()
	for out, r := range reactions {
		for _, in := range r.in {
			gb.AddEdgeL(out, in.name)
		}
	}
	g := gb.DenseDigraph()
	order := g.TopoSort(false)
	return fn.Map(order, g.Label)
}

func oreFor(want map[string]int, order []string, reactions map[string]reaction) int {
	for _, ch := range order {
		n, ok := want[ch]
		if !ok {
			continue // not needed
		}
		if ch == "ORE" {
			return n
		}
		delete(want, ch)
		r := reactions[ch]
		k := (n + r.out.q - 1) / r.out.q
		for _, in := range r.in {
			want[in.name] += k * in.q
		}
	}
	panic("no ore required")
}

func parseReactions(lines []string) map[string]reaction {
	reactions := make(map[string]reaction)
	for _, line := range lines {
		parts := strings.Split(line, " => ")
		if len(parts) != 2 {
			panic(fmt.Sprintf("invalid reaction: %s", line))
		}
		r := reaction{out: parsePile(parts[1])}
		for _, spec := range strings.Split(parts[0], ", ") {
			r.in = append(r.in, parsePile(spec))
		}
		reactions[r.out.name] = r
	}
	return reactions
}

func parsePile(spec string) pile {
	parts := strings.Split(spec, " ")
	if len(parts) != 2 {
		panic(fmt.Sprintf("invalid pile: %s", spec))
	}
	q, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(fmt.Sprintf("invalid pile (not a number): %s", spec))
	}
	return pile{name: parts[1], q: q}
}
