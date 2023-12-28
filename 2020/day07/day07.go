// Copyright 2020 Google LLC
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

// Package day07 solves AoC 2020 day 7.
package day07

import (
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util/graph"
)

func init() {
	glue.RegisterSolver(2020, 7, glue.LineSolver(solve))
	glue.RegisterPlotter(2020, 7, "", glue.LinePlotter(plotRules), map[string]string{"ex1": ex1, "ex2": ex2})
}

func solve(rules []string) ([]string, error) {
	g, err := parseRules(rules)
	if err != nil {
		return nil, err
	}

	bag := "shiny gold"
	part1 := countAncestors(g, bag)
	part2 := countDescendants(g, bag)

	return glue.Ints(part1, part2), nil
}

var (
	rulePattern = regexp.MustCompile(`^(\w+ \w+) bags contain ([^.]*)\.$`)
	bagPattern  = regexp.MustCompile(`^(\d+) (\w+ \w+) bags?$`)
)

func parseRules(rules []string) (*graph.DenseW, error) {
	g := graph.NewBuilder()

	for _, rule := range rules {
		m := rulePattern.FindStringSubmatch(rule)
		if m == nil {
			return nil, fmt.Errorf("invalid rule: %s", rule)
		}
		from, contents := m[1], m[2]
		u := g.V(from)
		if contents != "no other bags" {
			for _, content := range strings.Split(contents, ", ") {
				m = bagPattern.FindStringSubmatch(content)
				if m == nil {
					return nil, fmt.Errorf("invalid content: %s", content)
				}
				w, _ := strconv.Atoi(m[1])
				v := g.V(m[2])
				g.AddEdgeW(u, v, w)
			}
		}
	}

	return g.DenseDigraphW(), nil
}

func countAncestors(g *graph.DenseW, node string) int {
	seen := make([]bool, g.Len())
	var dfs func(int) int
	dfs = func(at int) int {
		if seen[at] {
			return 0
		}
		seen[at] = true
		count := 1
		for it := g.Pred(at); it.Valid(); it = g.Next(it) {
			count += dfs(it.Tail())
		}
		return count
	}
	v, _ := g.V(node)
	return dfs(v) - 1
}

func countDescendants(g *graph.DenseW, node string) int {
	memo := make([]int, g.Len())
	var dfs func(int) int
	dfs = func(at int) int {
		if c := memo[at]; c > 0 {
			return c
		}
		count := 0
		for it := g.Succ(at); it.Valid(); it = g.Next(it) {
			next := it.Head()
			n := g.W(at, next)
			count += n * (1 + dfs(next))
		}
		memo[at] = count
		return count
	}
	v, _ := g.V(node)
	return dfs(v)
}

var (
	ex1 = `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.
`
	ex2 = `shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
dark yellow bags contain 2 dark green bags.
dark green bags contain 2 dark blue bags.
dark blue bags contain 2 dark violet bags.
dark violet bags contain no other bags.
`
)

func plotRules(rules []string, out io.Writer) error {
	g, err := parseRules(rules)
	if err != nil {
		return err
	}

	var (
		colors   map[int]string
		colorize func(int, func(int) graph.DenseIt, func(graph.DenseIt) int, string)
	)
	colors = make(map[int]string)
	colorize = func(u int, next func(int) graph.DenseIt, end func(graph.DenseIt) int, color string) {
		colors[u] = color
		for it := next(u); it.Valid(); it = g.Next(it) {
			v := end(it)
			if _, ok := colors[v]; !ok {
				colorize(v, next, end, color)
			}
		}
	}

	nodeV, _ := g.V("shiny gold")
	colorize(nodeV, g.Pred, graph.DenseIt.Tail, `"#1e8e3e"`)
	colorize(nodeV, g.Succ, graph.DenseIt.Head, `"#1a73e8"`)
	colors[nodeV] = `"#d93025"`

	return graph.WriteDOT(g, out, "bags", true, func(v int) map[string]string {
		fg, bg := `"black"`, `"white"`
		if c, ok := colors[v]; ok {
			fg, bg = `"white"`, c
		}
		return map[string]string{"fillcolor": bg, "fontcolor": fg, "style": `"filled"`}
	}, nil)
}
