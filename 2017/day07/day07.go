// Copyright 2021 Google LLC
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

// Package day07 solves AoC 2017 day 7.
package day07

import (
	"strconv"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util/graph"
)

func init() {
	glue.RegisterSolver(2017, 7, glue.RegexpSolver{
		Solver: solve,
		Regexp: `^(\w+) \((\d+)\)(?: -> (\w+(?:, \w+)*))?$`,
	})
}

func solve(lines [][]string) ([]string, error) {
	progs := parseLines(lines)
	g := buildGraph(progs)
	p1 := findRoot(g)
	_, p2 := fixWeight(p1, progs)
	return []string{p1, strconv.Itoa(p2)}, nil
}

type program struct {
	name     string
	weight   int
	subNames []string
}

func findRoot(g *graph.Sparse) string {
	numPred := make([]int, g.Len())
	for u := 0; u < g.Len(); u++ {
		for it := g.Succ(u); it.Valid(); it = g.Next(it) {
			numPred[it.Head()]++
		}
	}
	for u, n := range numPred {
		if n == 0 {
			return g.Label(u)
		}
	}
	return ""
}

func fixWeight(root string, progs map[string]*program) (treeW, fixedW int) {
	prog := progs[root]
	subN := len(prog.subNames)
	if subN == 0 {
		return prog.weight, -1
	}
	subW, subWdist := make([]int, subN), make(map[int]int)
	for i, sub := range prog.subNames {
		subW[i], fixedW = fixWeight(sub, progs)
		if fixedW >= 0 {
			return -1, fixedW
		}
		subWdist[subW[i]] = subWdist[subW[i]] + 1
	}
	if len(subWdist) == 1 { // all children are balanced
		return prog.weight + subN*subW[0], -1
	}
	if len(subWdist) > 2 {
		panic("impossible: len(subWdist) > 2")
	}
	badW, goodW := -1, -1
	for w, count := range subWdist {
		if count == 1 {
			badW = w
		} else {
			goodW = w
		}
	}
	if badW == -1 || goodW == -1 {
		panic("impossible: badW == -1 || goodW == -1")
	}
	for i, sub := range prog.subNames {
		if subW[i] == badW {
			return -1, progs[sub].weight + (goodW - badW)
		}
	}
	panic("impossible: bad weight not found")
}

func buildGraph(progs map[string]*program) *graph.Sparse {
	g := graph.NewBuilder()
	for _, prog := range progs {
		u := g.V(prog.name)
		for _, sub := range prog.subNames {
			v := g.V(sub)
			g.AddEdge(u, v)
		}
	}
	return g.SparseDigraph()
}

func parseLines(lines [][]string) (progs map[string]*program) {
	progs = make(map[string]*program)
	for _, line := range lines {
		prog := parseLine(line)
		progs[prog.name] = &prog
	}
	return progs
}

func parseLine(line []string) (prog program) {
	prog.name = line[0]
	prog.weight, _ = strconv.Atoi(line[1])
	if line[2] != "" {
		prog.subNames = strings.Split(line[2], ", ")
	}
	return prog
}
