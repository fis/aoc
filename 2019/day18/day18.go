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

// Package day18 solves AoC 2019 day 18.
package day18

import (
	"container/heap"
	"errors"
	"fmt"
	"strings"

	"github.com/fis/aoc-go/util"
)

func init() {
	util.RegisterSolver(18, util.LevelSolver{Solver: solve, Empty: '#'})
}

func solve(level *util.Level) ([]int, error) {
	part1 := solveLevel(level)

	x, y, ok := level.Find('@')
	if !ok {
		return nil, errors.New("could not find unique @ in input level")
	}
	for _, d := range [][2]int{{-1, 0}, {0, -1}, {0, 0}, {0, 1}, {1, 0}} {
		level.Set(x+d[0], y+d[1], '#')
	}
	for _, d := range [][2]int{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}} {
		level.Set(x+d[0], y+d[1], '@')
	}

	part2 := solveLevel(level)

	return []int{part1, part2}, nil
}

func solveLevel(level *util.Level) int {
	g := buildGraph(level)
	g.dump()
	switch len(g.starts) {
	case 1:
		return g.collect1()
	case 4:
		return g.collect4()
	}
	panic("odd number of entrances")
}

type graph struct {
	verts  []*graphV
	starts []*graphV
	keys   labelSet
}

type graphV struct {
	id    int
	key   label
	edges []graphE
	best  map[labelSet]int
}

type graphE struct {
	v           *graphV
	d           int
	keys, doors labelSet
}

func (g *graph) collect1() int {
	if len(g.starts) != 1 {
		panic("collect1: unexpected number of starting points")
	}

	for _, v := range g.verts {
		v.best = make(map[labelSet]int)
	}

	front := frontier{
		&frontierNode{
			data: g.starts[0],
			keys: emptyLabelSet,
			d:    0,
		},
	}
	maxD := -1
	for len(front) > 0 {
		f := heap.Pop(&front).(*frontierNode)
		if f.d > maxD {
			maxD = f.d
			util.Diagf("exploring at %d, frontier size %d\n", maxD, len(front))
		}
		if f.keys == g.keys {
			return f.d
		}
		fv := f.data.(*graphV)
		for _, e := range fv.edges {
			if !e.doors.isSubsetOf(f.keys) {
				continue // this path is unavailable
			}
			keys := f.keys | e.keys
			d := f.d + e.d
			if b, ok := e.v.best[keys]; ok && d >= b {
				continue // this edge is not worth it
			}
			e.v.best[keys] = d
			heap.Push(&front, &frontierNode{data: e.v, keys: keys, d: d})
		}
	}
	panic("mission: impossible")
}

func (g *graph) collect4() int {
	if len(g.starts) != 4 {
		panic("collect4: unexpected number of starting points")
	}

	var start [4]int
	for i, v := range g.starts {
		start[i] = v.id
	}

	allBest := make(map[[4]int]map[labelSet]int)

	front := frontier{
		&frontierNode{
			data: start,
			keys: emptyLabelSet,
			d:    0,
		},
	}
	maxD := -1
	for len(front) > 0 {
		f := heap.Pop(&front).(*frontierNode)
		if f.d > maxD {
			maxD = f.d
			util.Diagf("exploring at %d, frontier size %d\n", maxD, len(front))
		}
		if f.keys == g.keys {
			return f.d
		}
		fvis := f.data.([4]int)
		for ri, fvi := range fvis {
			fv := g.verts[fvi]
			for _, e := range fv.edges {
				if !e.doors.isSubsetOf(f.keys) {
					continue // this path is unavailable
				}
				keys := f.keys | e.keys
				d := f.d + e.d
				nvis := fvis
				nvis[ri] = e.v.id
				best, ok := allBest[nvis]
				if ok {
					if b, ok := best[keys]; ok && d >= b {
						continue // this edge is not worth it
					}
				} else {
					best = make(map[labelSet]int)
					allBest[nvis] = best
				}
				best[keys] = d
				heap.Push(&front, &frontierNode{data: nvis, keys: keys, d: d})
			}
		}
	}
	panic("mission: impossible")
}

type frontier []*frontierNode

type frontierNode struct {
	data interface{}
	keys labelSet
	d    int
}

func buildGraph(level *util.Level) *graph {
	g := &graph{}

	tiles := make(map[util.P]label)
	verts := make(map[util.P]*graphV)
	var sources []util.P

	level.Range(func(x, y int, b byte) {
		p := util.P{x, y}
		l := noLabel
		switch {
		case b >= 'a' && b <= 'z':
			l = keyLabel(b - 'a')
			verts[p] = &graphV{key: l}
			sources = append(sources, p)
			g.keys |= l.asSet()
		case b >= 'A' && b <= 'Z':
			l = doorLabel(b - 'A')
		case b == '@':
			// as a tile, treat like '.'
			verts[p] = &graphV{key: noLabel}
			sources = append(sources, p)
		case b != '.':
			panic(fmt.Sprintf("unexpected feature: %c", b))
		}
		tiles[util.P{x, y}] = l
	})

	for _, v := range verts {
		v.id = len(g.verts)
		g.verts = append(g.verts, v)
		if v.key == noLabel {
			g.starts = append(g.starts, v)
		}
	}

	for _, from := range sources {
		fromV := verts[from]
		type path struct {
			at    util.P
			keys  labelSet
			doors labelSet
		}
		fringe := []path{{at: from, keys: emptyLabelSet, doors: emptyLabelSet}}
		seen := make(map[util.P]struct{})
		d := 0
		for len(fringe) > 0 {
			d++
			var newFringe []path
			for _, p := range fringe {
				seen[p.at] = struct{}{}
				for _, step := range p.at.Neigh() {
					l, hasTile := tiles[step]
					_, isSeen := seen[step]
					if !hasTile || isSeen {
						continue // been there, done that; or nothing to see
					}
					keys, doors := p.keys, p.doors
					if l.isKey() {
						keys |= l.asSet()
						fromV.edges = append(fromV.edges, graphE{v: verts[step], d: d, keys: keys, doors: doors})
					} else if l.isDoor() {
						doors |= l.asKey().asSet()
					}
					newFringe = append(newFringe, path{at: step, keys: keys, doors: doors})
				}
			}
			fringe = newFringe
		}
	}

	return g
}

// heap.Interface implementation for the priority queue

func (f frontier) Len() int {
	return len(f)
}

func (f frontier) Less(i, j int) bool {
	return f[i].d < f[j].d
}

func (f frontier) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func (f *frontier) Push(x interface{}) {
	node := x.(*frontierNode)
	*f = append(*f, node)
}

func (f *frontier) Pop() interface{} {
	old, n := *f, len(*f)
	node := old[n-1]
	*f = old[0 : n-1]
	return node
}

// utilities for labels and bitset-based label sets

type label uint64

const (
	noLabel      = label(0)
	startLabel   = label(1)
	minKeyLabel  = label(uint64(1) << 1)
	maxKeyLabel  = label(uint64(1) << 26)
	minDoorLabel = label(uint64(1) << 27)
	maxDoorLabel = label(uint64(1) << 52)
)

func keyLabel(key byte) label {
	return label(uint64(1) << (1 + key))
}

func doorLabel(door byte) label {
	return label(uint64(1) << (27 + door))
}

func (l label) isKey() bool {
	return l >= minKeyLabel && l <= maxKeyLabel
}

func (l label) isDoor() bool {
	return l >= minDoorLabel && l <= maxDoorLabel
}

func (l label) asKey() label {
	if l.isDoor() {
		return l >> 26
	}
	return l
}

func (l label) asDoor() label {
	if l.isKey() {
		return l << 26
	}
	return l
}

func (l label) asSet() labelSet {
	return labelSet(l)
}

func (l label) String() string {
	var (
		c byte
		v label
	)
	switch {
	case l == noLabel:
		return "."
	case l == startLabel:
		return "@"
	case l.isKey():
		c, v = 'a', minKeyLabel
	case l.isDoor():
		c, v = 'A', minDoorLabel
	}
	for l != v && v != 0 {
		c, v = c+1, v<<1
	}
	return string([]byte{c})
}

type labelSet uint64

const (
	emptyLabelSet = labelSet(0)
)

func (ls labelSet) isSubsetOf(other labelSet) bool {
	return (ls & other) == ls
}

func (ls labelSet) diff(other labelSet) labelSet {
	return ls & ^other
}

func (ls labelSet) single() bool {
	return (ls & (ls - 1)) == emptyLabelSet
}

func (ls labelSet) contains(l label) bool {
	return (ls & labelSet(l)) != emptyLabelSet
}

func (ls labelSet) without(l label) labelSet {
	return ls & ^labelSet(l)
}

func (ls *labelSet) next(l *label) bool {
	if *ls == 0 {
		return false
	}
	old := *ls
	*ls = old & (old - 1)
	*l = label(old ^ *ls)
	return true
}

func (ls labelSet) String() string {
	out := strings.Builder{}
	for l := noLabel; ls.next(&l); {
		out.WriteString(l.String())
	}
	return out.String()
}

// debugging utilities

func (v *graphV) String() string {
	if v.key.isKey() {
		return v.key.String()
	}
	return fmt.Sprintf("@%d", v.id)
}

func (g *graph) dump() {
	for _, v := range g.verts {
		util.Diagf("%v:", v)
		for _, e := range v.edges {
			util.Diagf(" %s/%v/%v/%d", e.v, e.keys, e.doors, e.d)
		}
		util.Diagln()
	}
}
