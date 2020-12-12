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

// Package day20 solves AoC 2019 day 20.
package day20

import (
	"container/heap"
	"unicode"

	"github.com/fis/aoc-go/glue"
	"github.com/fis/aoc-go/util"
)

func init() {
	glue.RegisterSolver(2019, 20, glue.LevelSolver{Solver: solve, Empty: ' '})
}

func solve(level *util.Level) ([]int, error) {
	dist := distances(level)
	p1 := shortest(label{name: "AA", outer: true}, label{name: "ZZ", outer: true}, dist)
	p2 := recursive(label{name: "AA", outer: true}, label{name: "ZZ", outer: true}, dist)
	return []int{p1, p2}, nil
}

type label struct {
	name  string
	outer bool
}

type distance struct {
	d     int
	depth int
}

type path struct {
	at    label
	depth int
	d     int
}

type pathq []path

func shortest(from, to label, edges map[label]map[label]distance) int {
	dist := map[label]int{from: 0}
	fringe := pathq{{at: from, d: 0}}
	for len(fringe) > 0 {
		p := heap.Pop(&fringe).(path)
		if p.at == to {
			return p.d
		}
		if od := dist[p.at]; od < p.d {
			continue // obsolete path
		}
		for to, e := range edges[p.at] {
			ed := p.d + e.d
			if od, ok := dist[to]; ok && od <= ed {
				continue // seen better
			}
			dist[to] = ed
			heap.Push(&fringe, path{at: to, d: ed})
		}
	}
	return -1
}

func recursive(from, to label, edges map[label]map[label]distance) int {
	type node struct {
		at    label
		depth int
	}
	dist := map[node]int{{at: from, depth: 0}: 0}
	fringe := pathq{{at: from, depth: 0, d: 0}}
	for len(fringe) > 0 {
		p := heap.Pop(&fringe).(path)
		if p.at == to && p.depth == 0 {
			return p.d
		}
		if od := dist[node{at: p.at, depth: p.depth}]; od < p.d {
			continue // obsolete path
		}
		for to, e := range edges[p.at] {
			ed, edepth := p.d+e.d, p.depth+e.depth
			if edepth < 0 {
				continue // cannot ascend
			}
			if od, ok := dist[node{at: to, depth: edepth}]; ok && od <= ed {
				continue // seen better
			}
			dist[node{at: to, depth: edepth}] = ed
			heap.Push(&fringe, path{at: to, depth: edepth, d: ed})
		}
	}
	return -1
}

func distances(level *util.Level) map[label]map[label]distance {
	labels := make(map[util.P]label)

	level.Range(func(x, y int, c byte) {
		if !unicode.IsUpper(rune(c)) {
			return
		}
		for _, d := range (util.P{}).Neigh() {
			c2 := level.At(x+d.X, y+d.Y)
			if !unicode.IsUpper(rune(c2)) || level.At(x+2*d.X, y+2*d.Y) != '.' {
				continue
			}
			var name string
			if d.X > 0 || d.Y > 0 {
				name = string([]byte{c, c2})
			} else {
				name = string([]byte{c2, c})
			}
			labels[util.P{x + 2*d.X, y + 2*d.Y}] = label{
				name:  name,
				outer: !level.InBounds(x-d.X, y-d.Y),
			}
		}
	})

	allDist := make(map[label]map[label]distance)

	for start, from := range labels {
		dist := make(map[label]distance)
		seen := make(map[util.P]struct{})
		fringe := []util.P{start}
		d := 0
		for len(fringe) > 0 {
			d++
			var newFringe []util.P
			for _, p := range fringe {
				seen[p] = struct{}{}
				for _, step := range p.Neigh() {
					if _, ok := seen[step]; ok {
						continue
					}
					if to, ok := labels[step]; ok {
						if _, ok := dist[to]; !ok {
							dist[to] = distance{d: d, depth: 0} // best path from -> to
						}
						continue
					}
					if level.At(step.X, step.Y) == '.' {
						newFringe = append(newFringe, step)
					}
				}
			}
			fringe = newFringe
		}
		allDist[from] = dist
	}

	for in := range allDist {
		if in.outer {
			continue
		}
		out := label{name: in.name, outer: true}
		if _, ok := allDist[out]; !ok {
			continue
		}
		allDist[in][out] = distance{d: 1, depth: 1}
		allDist[out][in] = distance{d: 1, depth: -1}
	}

	return allDist
}

func (q pathq) Len() int {
	return len(q)
}

func (q pathq) Less(i, j int) bool {
	return q[i].d < q[j].d
}

func (q pathq) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *pathq) Push(x interface{}) {
	*q = append(*q, x.(path))
}

func (q *pathq) Pop() interface{} {
	old, n := *q, len(*q)
	path := old[n-1]
	*q = old[0 : n-1]
	return path
}
