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
	g, w, err := parseRules(rules)
	if err != nil {
		return nil, err
	}

	bag := "shiny gold"
	part1 := countAncestors(g, bag)
	part2 := countDescendants(g, w, bag)

	return []string{strconv.Itoa(part1), strconv.Itoa(part2)}, nil
}

var (
	rulePattern = regexp.MustCompile(`^(\w+ \w+) bags contain ([^.]*)\.$`)
	bagPattern  = regexp.MustCompile(`^(\d+) (\w+ \w+) bags?$`)
)

func parseRules(rules []string) (g *util.Graph, w map[[2]int]int, err error) {
	g = &util.Graph{}
	w = make(map[[2]int]int)

	for _, rule := range rules {
		m := rulePattern.FindStringSubmatch(rule)
		if m == nil {
			return nil, nil, fmt.Errorf("invalid rule: %s", rule)
		}
		from, contents := m[1], m[2]
		fromV := g.V(from)
		if contents != "no other bags" {
			for _, content := range strings.Split(contents, ", ") {
				m = bagPattern.FindStringSubmatch(content)
				if m == nil {
					return nil, nil, fmt.Errorf("invalid content: %s", content)
				}
				toW, _ := strconv.Atoi(m[1])
				toV := g.V(m[2])
				g.AddEdgeV(fromV, toV)
				w[[2]int{fromV, toV}] = toW
			}
		}
	}

	return g, w, nil
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

func countDescendants(g *util.Graph, w map[[2]int]int, node string) int {
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
			n := w[[2]int{at, next}]
			count += n * (1 + dfs(next))
			return true
		})
		memo[at] = count
		return count
	}
	return dfs(g.V(node))
}
