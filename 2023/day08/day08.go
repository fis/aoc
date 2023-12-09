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

// Package day08 solves AoC 2023 day 8.
package day08

import (
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/fn"
	"github.com/fis/aoc/util/ix"
)

func init() {
	glue.RegisterSolver(2023, 8, glue.ChunkSolver(solve))
	glue.RegisterPlotter(2023, 8, glue.ChunkPlotter(plot), map[string]string{"ex1": ex1, "ex2": ex2, "ex3": ex3})
}

func solve(chunks []string) ([]string, error) {
	g, dirs, err := parseMaps(chunks)
	if err != nil {
		return nil, err
	}
	p1 := countSteps(g, g.v("AAA"), g.v("ZZZ"), dirs)
	p2 := countGhostSteps(g, dirs)
	return glue.Ints(p1, p2), nil
}

func countSteps(g *graph, from, to int, dirs []direction) (steps int) {
	d := 0
	for from != to {
		steps++
		from = g.edges[from][dirs[d]]
		d++
		if d == len(dirs) {
			d = 0
		}
	}
	return steps
}

func countGhostSteps(g *graph, dirs []direction) (steps int) {
	var cycles []cycle
	for label, node := range g.nodes {
		if label[len(label)-1] == 'A' {
			cycles = append(cycles, findCycle(g, node, dirs))
		}
	}
	for len(cycles) > 1 {
		i := len(cycles) - 2
		cycles[i] = mergeCycles(cycles[i], cycles[i+1])
		cycles = cycles[:i+1]
	}
	return cycles[0].end
}

type cycle struct {
	start int
	size  int
	end   int
}

func findCycle(g *graph, from int, dirs []direction) cycle {
	seen := fn.MapRange(0, len(g.nodes), func(int) int { return -1 })
	n, steps := from, 0
	end := -1
	for {
		seen[n] = steps
		for i, d := range dirs {
			n = g.edges[n][dirs[d]]
			if label := g.labels[n]; label[len(label)-1] == 'Z' {
				end = steps + i + 1
			}
		}
		steps += len(dirs)
		if start := seen[n]; start >= 0 {
			size := steps - start
			return cycle{start: start, size: size, end: end % size}
		}
	}
}

func mergeCycles(c1, c2 cycle) cycle {
	if c1.size < c2.size {
		c1, c2 = c2, c1
	}
	for a := c1.end; ; a += c1.size {
		if a < c1.start || a < c2.start {
			continue
		}
		if a%c2.size == c2.end {
			return cycle{size: ix.LCM(c1.size, c2.size), end: a}
		}
	}
}

type direction byte

const (
	dirL direction = iota
	dirR
)

type graph struct {
	nodes  map[string]int
	labels []string
	edges  [][2]int // array indices: dirL, dirR
}

func (g *graph) v(label string) int {
	if n, ok := g.nodes[label]; ok {
		return n
	}
	n := len(g.labels)
	g.nodes[label] = n
	g.labels = append(g.labels, label)
	g.edges = append(g.edges, [2]int{-1, -1})
	return n
}

var nodePattern = regexp.MustCompile(`(\w{3}) = \((\w{3}), (\w{3})\)`)

func parseMaps(chunks []string) (g *graph, dirs []direction, err error) {
	if len(chunks) != 2 {
		return nil, nil, fmt.Errorf("expected 2 chunks, got %d", len(chunks))
	}
	dirs = make([]direction, 0, len(chunks[0]))
	for _, d := range chunks[0] {
		switch d {
		case 'L':
			dirs = append(dirs, dirL)
		case 'R':
			dirs = append(dirs, dirR)
		default:
			return nil, nil, fmt.Errorf("bad direction: %q", d)
		}
	}
	g = &graph{nodes: make(map[string]int)}
	for _, line := range util.Lines(chunks[1]) {
		parts := nodePattern.FindStringSubmatch(line)
		if len(parts) == 0 {
			return nil, nil, fmt.Errorf("bad edge: %s", line)
		}
		from, left, right := g.v(parts[1]), g.v(parts[2]), g.v(parts[3])
		g.edges[from] = [2]int{left, right}
	}
	return g, dirs, err
}

// plotting

var (
	ex1 = strings.TrimPrefix(`
RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)
`, "\n")
	ex2 = strings.TrimPrefix(`
LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)
`, "\n")
	ex3 = strings.TrimPrefix(`
LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)
`, "\n")
)

func plot(chunks []string, w io.Writer) error {
	g, _, err := parseMaps(chunks)
	if err != nil {
		return err
	}
	fmt.Fprintln(w, "digraph G {")
	for i, label := range g.labels {
		fmt.Fprintf(w, "  n%d [label=\"%s\"];\n", i, label)
	}
	for i, e := range g.edges {
		for j, label := range []string{"L", "R"} {
			fmt.Fprintf(w, "  n%d -> n%d [label=\"%s\"];\n", i, e[j], label)
		}
	}
	fmt.Fprintln(w, "}")
	return nil
}
