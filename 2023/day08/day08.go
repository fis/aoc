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
	"regexp"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/fn"
	"github.com/fis/aoc/util/ix"
)

func init() {
	glue.RegisterSolver(2023, 8, glue.ChunkSolver(solve))
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
	start := fn.MaxF(cycles, func(c cycle) int { return c.start })
	for i := range cycles {
		cycles[i].end -= start
	}
	for len(cycles) > 1 {
		i := len(cycles) - 2
		cycles[i] = mergeCycles(cycles[i], cycles[i+1])
		cycles = cycles[:i+1]
	}
	return start + cycles[0].end
}

type cycle struct {
	start int
	size  int
	end   int
}

func findCycle(g *graph, from int, dirs []direction) cycle {
	n, steps, d := from, 0, 0
	seen := make([][]int, len(g.nodes))
	for i := range seen {
		seen[i] = make([]int, len(dirs))
		for j := range seen[i] {
			seen[i][j] = -1
		}
	}
	end := -1 // assumes there's just one suitable end to it
	for {
		seen[n][d] = steps
		steps++
		n = g.edges[n][dirs[d]]
		d++
		if d == len(dirs) {
			d = 0
		}
		if label := g.labels[n]; label[len(label)-1] == 'Z' {
			end = steps
		}
		if start := seen[n][d]; start >= 0 {
			return cycle{start: start, size: steps - start, end: end}
		}
	}
}

func mergeCycles(ca, cb cycle) cycle {
	if ca.size < cb.size {
		ca, cb = cb, ca
	}
	for a := ca.end; ; a += ca.size {
		if a%cb.size == cb.end {
			return cycle{size: ix.LCM(ca.size, cb.size), end: a}
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
