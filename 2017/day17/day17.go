// Package day17 solves AoC 2017 day 17.
package day17

import (
	"fmt"

	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2017, 17, glue.IntSolver(solve))
}

func solve(input []int) ([]string, error) {
	if len(input) != 1 {
		return nil, fmt.Errorf("expected 1 int as input, got %d", len(input))
	}
	p1 := part1(input[0])
	p2 := part2(input[0])
	return glue.Ints(p1, p2), nil
}

func part1(skip int) int {
	r := newRing()
	for i := 1; i <= 2017; i++ {
		r.move(skip)
		r.insert(i)
	}
	return r.pos.next.val
}

func part2(skip int) int {
	pos, size, after0 := 0, 1, -1
	for i := 1; i <= 50000000; i++ {
		pos = (pos + skip) % size
		if pos == 0 {
			after0 = i
		}
		pos++
		size++
	}
	return after0
}

type ring struct {
	size int
	pos  *node
}

type node struct {
	val  int
	next *node
}

func newRing() *ring {
	n := &node{val: 0}
	n.next = n
	return &ring{size: 1, pos: n}
}

func (r *ring) move(steps int) {
	if steps >= r.size {
		steps %= r.size
	}
	for range steps {
		r.pos = r.pos.next
	}
}

func (r *ring) insert(val int) {
	n := &node{val: val, next: r.pos.next}
	r.pos.next = n
	r.pos = n
	r.size++
}
