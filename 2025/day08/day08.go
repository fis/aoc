// Package day08 solves AoC 2025 day 8.
package day08

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util/heap"
)

func init() {
	glue.RegisterSolver(2025, 8, glue.LineSolver(glue.WithParser(parseP3, solve)))
}

func solve(points []p3) ([]string, error) {
	p1, p2 := makeConnections(points, 1000)
	return glue.Ints(p1, p2), nil
}

func makeConnections(points []p3, topSizeCutoff int) (p1, p2 int) {
	N := len(points)

	type pairDist struct{ left, right, d2 int }
	dists := make(heap.Heap[[2]int], 0, N*(N-1)/2)
	for left := 0; left < N-1; left++ {
		for right := left + 1; right < N; right++ {
			l, r := points[left], points[right]
			dx, dy, dz := r.x-l.x, r.y-l.y, r.z-l.z
			dists = append(dists, heap.Item[[2]int]{P: dx*dx + dy*dy + dz*dz, Value: [2]int{left, right}})
		}
	}
	dists.Heapify()

	circuits, circuitMap, numCircuits := make([][]int, N), make([]int, N), N
	for i := range N {
		circuits[i] = []int{i}
		circuitMap[i] = i
	}
	numConns := 0
	for numCircuits > 1 {
		_, pd := dists.Pop()
		l, r := pd[0], pd[1]
		numConns++

		if lc, rc := circuitMap[l], circuitMap[r]; lc != rc {
			if len(circuits[rc]) > len(circuits[lc]) {
				l, r, lc, rc = r, l, rc, lc // faster to relabel the smaller component
			}
			for _, i := range circuits[rc] {
				circuitMap[i] = lc
			}
			circuits[lc] = append(circuits[lc], circuits[rc]...)
			circuits[rc] = nil
			numCircuits--
			if numCircuits == 1 {
				p2 = points[l].x * points[r].x
			}
		}

		if numConns == topSizeCutoff {
			top := top3{}
			for _, c := range circuits {
				top.push(len(c))
			}
			p1 = top[0] * top[1] * top[2]
		}
	}

	return p1, p2
}

type p3 struct{ x, y, z int }

func parseP3(s string) (p p3, err error) {
	xs, yzs, ok1 := strings.Cut(s, ",")
	ys, zs, ok2 := strings.Cut(yzs, ",")
	if !ok1 || !ok2 {
		return p3{}, fmt.Errorf("point not in x,y,z form: %q", s)
	}
	x, err1 := strconv.Atoi(xs)
	y, err2 := strconv.Atoi(ys)
	z, err3 := strconv.Atoi(zs)
	if err1 != nil || err2 != nil || err3 != nil {
		return p3{}, fmt.Errorf("point components not all integers: %q: %w", s, errors.Join(err1, err2, err3))
	}
	return p3{x, y, z}, nil
}

type top3 [3]int

func (t *top3) push(n int) {
	switch {
	case n > t[0]:
		t[0], t[1], t[2] = n, t[0], t[1]
	case n > t[1]:
		t[1], t[2] = n, t[1]
	case n > t[2]:
		t[2] = n
	}
}
