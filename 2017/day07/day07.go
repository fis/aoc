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
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2017, 7, glue.GenericSolver(solve))
}

func solve(input io.Reader) ([]string, error) {
	lines, err := util.ScanAll(input, bufio.ScanLines)
	if err != nil {
		return nil, err
	}
	progs, err := parseLines(lines)
	if err != nil {
		return nil, err
	}
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

func findRoot(g *util.Graph) string {
	root := ""
	g.Range(func(name string) {
		if g.NumPred(name) == 0 {
			root = name
		}
	})
	return root
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

func buildGraph(progs map[string]*program) *util.Graph {
	g := &util.Graph{}
	for _, prog := range progs {
		fromV := g.V(prog.name)
		for _, sub := range prog.subNames {
			toV := g.V(sub)
			g.AddEdgeV(fromV, toV)
		}
	}
	return g
}

func parseLines(lines []string) (progs map[string]*program, err error) {
	progs = make(map[string]*program)
	for _, line := range lines {
		prog, err := parseLine(line)
		if err != nil {
			return nil, err
		}
		progs[prog.name] = &prog
	}
	return progs, nil
}

func parseLine(line string) (prog program, err error) {
	parts := strings.SplitN(line, " ", 4)
	if n := len(parts); n != 2 && n != 4 {
		return program{}, fmt.Errorf("unexpected number of parts: want 2 or 4, got %d", n)
	}
	prog.name = parts[0]
	if _, err = fmt.Sscanf(parts[1], "(%d)", &prog.weight); err != nil {
		return program{}, fmt.Errorf("unexpected weight: %q: wanted (N)", parts[1])
	}
	if len(parts) == 2 {
		return prog, nil
	}
	if parts[2] != "->" {
		return program{}, fmt.Errorf("unexpected separator: %q: wanted ->", parts[3])
	}
	prog.subNames = strings.Split(parts[3], ", ")
	return prog, nil
}
