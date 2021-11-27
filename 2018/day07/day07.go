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

// Package day07 solves AoC 2018 day 7.
package day07

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2018, 7, glue.GenericSolver(solve))
	glue.RegisterPlotter(2018, 7, glue.LinePlotter(plotDeps), map[string]string{"ex": example})
}

func solve(input io.Reader) ([]string, error) {
	lines, err := util.ScanAll(input, bufio.ScanLines)
	if err != nil {
		return nil, err
	}

	part1 := strings.Join(toplexSort(parseRules(lines)), "")
	_, part2 := timedSort(parseRules(lines), 5, 60)

	return []string{part1, strconv.Itoa(part2)}, nil
}

func parseRules(lines []string) *util.Graph {
	g := &util.Graph{}
	for _, line := range lines {
		var from, to string
		if _, err := fmt.Sscanf(line, "Step %s must be finished before step %s can begin.", &from, &to); err != nil {
			continue
		}
		g.AddEdge(from, to)
	}
	return g
}

func toplexSort(g *util.Graph) (order []string) {
	avail := labelHeap{}
	g.RangeV(func(v int) {
		if g.NumPredV(v) == 0 {
			heap.Push(&avail, g.Name(v))
		}
	})
	for len(avail) > 0 {
		from := heap.Pop(&avail).(string)
		order = append(order, from)
		fromV := g.V(from)
		g.RangeSuccV(fromV, func(toV int) bool {
			g.DelEdgeV(fromV, toV)
			if g.NumPredV(toV) == 0 {
				heap.Push(&avail, g.Name(toV))
			}
			return true
		})
	}
	return order
}

func timedSort(g *util.Graph, workers, baseTime int) (order []string, totalTime int) {
	avail := labelHeap{}
	g.RangeV(func(v int) {
		if g.NumPredV(v) == 0 {
			heap.Push(&avail, g.Name(v))
		}
	})
	busy := workHeap{}
	now := 0
	for len(avail) > 0 || len(busy) > 0 {
		for len(busy) > 0 && now >= busy[0].readyAt {
			wi := heap.Pop(&busy).(workItem)
			order = append(order, wi.label)
			fromV := g.V(wi.label)
			g.RangeSuccV(fromV, func(toV int) bool {
				g.DelEdgeV(fromV, toV)
				if g.NumPredV(toV) == 0 {
					heap.Push(&avail, g.Name(toV))
				}
				return true
			})
		}
		for len(avail) > 0 && workers-len(busy) > 0 {
			from := heap.Pop(&avail).(string)
			dur := baseTime + int(from[0]-'A'+1)
			heap.Push(&busy, workItem{readyAt: now + dur, label: from})
		}
		if len(busy) > 0 {
			now = busy[0].readyAt
		}
	}
	return order, now
}

type labelHeap []string

func (h labelHeap) Len() int           { return len(h) }
func (h labelHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h labelHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h *labelHeap) Push(x interface{}) {
	*h = append(*h, x.(string))
}

func (h *labelHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type workItem struct {
	readyAt int
	label   string
}

type workHeap []workItem

func (h workHeap) Len() int      { return len(h) }
func (h workHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h workHeap) Less(i, j int) bool {
	if h[i].readyAt != h[j].readyAt {
		return h[i].readyAt < h[j].readyAt
	}
	return h[i].label < h[j].label
}

func (h *workHeap) Push(x interface{}) {
	*h = append(*h, x.(workItem))
}

func (h *workHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

var example = `Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.
`

func plotDeps(lines []string, w io.Writer) error {
	g := parseRules(lines)
	return g.WriteDOT(w, "deps", nil, nil)
}
