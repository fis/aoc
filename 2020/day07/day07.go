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

	"github.com/fis/aoc-go/util"
)

func Solve(path string) ([]string, error) {
	rules, err := util.ReadLines(path)
	if err != nil {
		return nil, err
	}
	g, err := parseRules(rules)
	if err != nil {
		return nil, err
	}

	bag := "shiny gold"
	part1 := countAncestors(g, bag)
	part2 := countDescendants(g, bag)

	return []string{strconv.Itoa(part1), strconv.Itoa(part2)}, nil
}

var (
	rulePattern = regexp.MustCompile(`^(\w+ \w+) bags contain ([^.]*)\.$`)
	bagPattern  = regexp.MustCompile(`^(\d+) (\w+ \w+) bags?$`)
)

func parseRules(rules []string) (g *util.Graph, err error) {
	g = &util.Graph{}

	for _, rule := range rules {
		m := rulePattern.FindStringSubmatch(rule)
		if m == nil {
			return nil, fmt.Errorf("invalid rule: %s", rule)
		}
		from, contents := m[1], m[2]
		fromV := g.V(from)
		if contents != "no other bags" {
			for _, content := range strings.Split(contents, ", ") {
				m = bagPattern.FindStringSubmatch(content)
				if m == nil {
					return nil, fmt.Errorf("invalid content: %s", content)
				}
				toW, _ := strconv.Atoi(m[1])
				toV := g.V(m[2])
				g.AddEdgeWV(fromV, toV, toW)
			}
		}
	}

	return g, nil
}

func countAncestors(g *util.Graph, node string) int {
	var (
		seen map[int]struct{}
		dfs  func(int) int
	)
	seen = make(map[int]struct{})
	dfs = func(at int) int {
		if _, ok := seen[at]; ok {
			return 0
		}
		seen[at] = struct{}{}
		count := 1
		g.RangePredV(at, func(next int) bool {
			count += dfs(next)
			return true
		})
		return count
	}
	return dfs(g.V(node)) - 1
}

func countDescendants(g *util.Graph, node string) int {
	var (
		memo map[int]int
		dfs  func(int) int
	)
	memo = make(map[int]int)
	dfs = func(at int) int {
		if c, ok := memo[at]; ok {
			return c
		}
		count := 0
		g.RangeSuccV(at, func(next int) bool {
			n := g.W(at, next)
			count += n * (1 + dfs(next))
			return true
		})
		memo[at] = count
		return count
	}
	return dfs(g.V(node))
}

func PrintRules(out io.Writer, rules []string, node string) error {
	g, err := parseRules(rules)
	if err != nil {
		return err
	}

	var (
		colors   map[int]string
		colorize func(int, func(*util.Graph, int, func(int) bool), string)
	)
	colors = make(map[int]string)
	colorize = func(at int, ranger func(*util.Graph, int, func(int) bool), color string) {
		colors[at] = color
		ranger(g, at, func(next int) bool {
			if _, ok := colors[next]; !ok {
				colorize(next, ranger, color)
			}
			return true
		})
	}

	nodeV := g.V(node)
	colorize(nodeV, (*util.Graph).RangePredV, "#1e8e3e")
	colorize(nodeV, (*util.Graph).RangeSuccV, "#1a73e8")
	colors[nodeV] = "#d93025"

	fmt.Fprint(out, "digraph bags {\n")
	g.RangeV(func(v int) {
		fg, bg := "black", "white"
		if c, ok := colors[v]; ok {
			fg, bg = "white", c
		}
		fmt.Fprintf(out, "  n%d [label=\"%s\", fillcolor=\"%s\", fontcolor=\"%s\", style=\"filled\"];\n", v, g.Name(v), bg, fg)
	})
	g.RangeV(func(v int) {
		g.RangeSuccV(v, func(v2 int) bool {
			fmt.Fprintf(out, "  n%d -> n%d [label=\"%d\"];\n", v, v2, g.W(v, v2))
			return true
		})
	})
	fmt.Fprint(out, "}\n")

	return nil
}
