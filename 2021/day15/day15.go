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

// Package day15 solves AoC 2021 day 15.
package day15

import (
	"container/heap"
	"fmt"
	"math"

	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2021, 15, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	w, h, level, err := readLevel(lines)
	if err != nil {
		return nil, err
	}
	p1 := shortestPathDijkstraBQ(w, h, level, 1)
	p2 := shortestPathDijkstraBQ(w, h, level, 5)
	return glue.Ints(int(p1), int(p2)), nil
}

type coord struct {
	x, y int32
}

type path struct {
	at coord
	d  int32
}

type pathq []path

func shortestPathDijkstra(w, h int32, level [][]byte, scale int32) int32 {
	W, H := w*scale, h*scale
	from := coord{0, 0}
	to := coord{W - 1, H - 1}
	dist := make([]int32, W*H)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[0] = 0
	fringe := pathq{{at: from, d: 0}}
	for len(fringe) > 0 {
		p := heap.Pop(&fringe).(path)
		if p.at == to {
			return p.d
		}
		if od := dist[p.at.y*W+p.at.x]; od < p.d {
			continue // path no longer relevant
		}
		for _, q := range [4]coord{{p.at.x, p.at.y - 1}, {p.at.x, p.at.y + 1}, {p.at.x - 1, p.at.y}, {p.at.x + 1, p.at.y}} {
			if q.x < 0 || q.x >= W || q.y < 0 || q.y >= H {
				continue
			}
			qr := 1 + (int32(level[q.y%h][q.x%w])+q.x/w+q.y/h-1)%9
			qp := path{at: q, d: p.d + qr}
			if od := dist[q.y*W+q.x]; od <= qp.d {
				continue
			}
			dist[q.y*W+q.x] = qp.d
			heap.Push(&fringe, qp)
		}
	}
	return -1 // can't get there from here
}

func shortestPathDijkstraBQ(w, h int32, level [][]byte, scale int32) int32 {
	W, H := w*scale, h*scale
	from := coord{0, 0}
	to := coord{W - 1, H - 1}
	dist := make([]int32, W*H)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[0] = 0
	var fringe bucketq
	fringe.push(0, from)
	for {
		pd, p := fringe.pop()
		if p == to {
			return pd
		}
		if od := dist[p.y*W+p.x]; od < pd {
			continue // path no longer relevant
		}
		for _, q := range [4]coord{{p.x, p.y - 1}, {p.x, p.y + 1}, {p.x - 1, p.y}, {p.x + 1, p.y}} {
			if q.x < 0 || q.x >= W || q.y < 0 || q.y >= H {
				continue
			}
			qd := pd + 1 + (int32(level[q.y%h][q.x%w])+q.x/w+q.y/h-1)%9
			if od := dist[q.y*W+q.x]; od <= qd {
				continue
			}
			dist[q.y*W+q.x] = qd
			fringe.push(qd, q)
		}
	}
}

const bucketSpan = 16 // a power of 2 that's > than the maximum edge length

type bucketq struct {
	at      int32
	buckets [bucketSpan]struct {
		prio   int32
		coords []coord
	}
}

func (bq *bucketq) pop() (p int32, c coord) {
	for len(bq.buckets[bq.at].coords) == 0 {
		bq.at = (bq.at + 1) & (bucketSpan - 1)
	}
	coords := bq.buckets[bq.at].coords
	c = coords[len(coords)-1]
	bq.buckets[bq.at].coords = coords[:len(coords)-1]
	return bq.buckets[bq.at].prio, c
}

func (bq *bucketq) push(p int32, c coord) {
	i := p & (bucketSpan - 1)
	bq.buckets[i].prio = p
	bq.buckets[i].coords = append(bq.buckets[i].coords, c)
}

type hpath struct {
	at coord
	d  int32
	hd int32
}

type hpathq []hpath

func shortestPathAStar(w, h int32, level [][]byte, scale int32) int32 {
	W, H := w*scale, h*scale
	from := coord{0, 0}
	to := coord{W - 1, H - 1}
	dist := make([]int32, W*H)
	for i := range dist {
		dist[i] = -1
	}
	dist[0] = 0
	fringe := hpathq{{at: from, d: 0, hd: distM(from, to)}}
	for len(fringe) > 0 {
		p := heap.Pop(&fringe).(hpath)
		if p.at == to {
			return p.d
		}
		if od := dist[p.at.y*W+p.at.x]; od < p.d {
			continue // path no longer relevant
		}
		for _, q := range [4]coord{{p.at.x, p.at.y - 1}, {p.at.x, p.at.y + 1}, {p.at.x - 1, p.at.y}, {p.at.x + 1, p.at.y}} {
			if q.x < 0 || q.x >= W || q.y < 0 || q.y >= H {
				continue
			}
			qr := 1 + (int32(level[q.y%h][q.x%w])+q.x/w+q.y/h-1)%9
			qp := hpath{at: q, d: p.d + qr, hd: p.d + qr + distM(q, to)}
			if od := dist[q.y*W+q.x]; od >= 0 && od <= qp.d {
				continue
			}
			dist[q.y*W+q.x] = qp.d
			heap.Push(&fringe, qp)
		}
	}
	return -1 // can't get there from here
}

func distM(a, b coord) int32 {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

func readLevel(lines []string) (w, h int32, data [][]byte, err error) {
	if len(lines) == 0 || len(lines[0]) == 0 {
		return 0, 0, nil, fmt.Errorf("empty level")
	}
	w, h = int32(len(lines[0])), int32(len(lines))
	block := make([]byte, w*h)
	for y, line := range lines {
		if int32(len(line)) != w {
			return 0, 0, nil, fmt.Errorf("misshapen level: expected %d columns, got %d", w, len(line))
		}
		row := block[y*int(w) : (y+1)*int(w)]
		for x := int32(0); x < w; x++ {
			c := line[x]
			if c < '0' || c > '9' {
				return 0, 0, nil, fmt.Errorf("bad level data: expected '0'..'9', got '%c' (%v)", c, c)
			}
			row[x] = c - '0'
		}
		data = append(data, row)
	}
	return w, h, data, nil
}

func (q pathq) Len() int           { return len(q) }
func (q pathq) Less(i, j int) bool { return q[i].d < q[j].d }
func (q pathq) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }

func (q *pathq) Push(x interface{}) {
	*q = append(*q, x.(path))
}

func (q *pathq) Pop() interface{} {
	old, n := *q, len(*q)
	path := old[n-1]
	*q = old[0 : n-1]
	return path
}

func (q hpathq) Len() int           { return len(q) }
func (q hpathq) Less(i, j int) bool { return q[i].hd < q[j].hd }
func (q hpathq) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }

func (q *hpathq) Push(x interface{}) {
	*q = append(*q, x.(hpath))
}

func (q *hpathq) Pop() interface{} {
	old, n := *q, len(*q)
	path := old[n-1]
	*q = old[0 : n-1]
	return path
}
