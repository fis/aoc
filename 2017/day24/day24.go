// Package day24 solves AoC 2017 day 24.
package day24

import (
	"strconv"

	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2017, 24, glue.RegexpSolver{
		Solver: solve,
		Regexp: `^(\d+)/(\d+)$`,
	})
}

func solve(lines [][]string) ([]string, error) {
	inv := &inventory{}
	for _, line := range lines {
		c := &component{}
		c.p1, _ = strconv.Atoi(line[0])
		c.p2, _ = strconv.Atoi(line[1])
		inv.give(c)
	}

	p1, p2, _ := bestBridge(0, inv)

	return glue.Ints(p1, p2), nil
}

func bestBridge(from int, inv *inventory) (strongest, longest, length int) {
	next := inv.find(from)
	if len(next) == 0 {
		return 0, 0, 0
	}
	for _, c := range next {
		var nextStrongest, nextLongest, nextLength int

		inv.take(c)
		if from == c.p1 {
			nextStrongest, nextLongest, nextLength = bestBridge(c.p2, inv)
		} else {
			nextStrongest, nextLongest, nextLength = bestBridge(c.p1, inv)
		}
		inv.give(c)

		if s := c.p1 + c.p2 + nextStrongest; s > strongest {
			strongest = s
		}
		if s, l := c.p1+c.p2+nextLongest, 1+nextLength; l > length || (l == length && s > longest) {
			longest, length = s, l
		}
	}
	return strongest, longest, length
}

type component struct {
	p1, p2 int
}

type inventory struct {
	chains []*chain
}

type chain struct {
	c    *component
	next *chain
}

func (i *inventory) find(p int) (comps []*component) {
	if p >= len(i.chains) {
		return nil
	}
	for at := i.chains[p]; at != nil; at = at.next {
		comps = append(comps, at.c)
	}
	return comps
}

func (i *inventory) take(c *component) {
	i.removeFromChain(c.p1, c)
	if c.p2 != c.p1 {
		i.removeFromChain(c.p2, c)
	}
}

func (i *inventory) give(c *component) {
	i.insertToChain(c.p1, c)
	if c.p2 != c.p1 {
		i.insertToChain(c.p2, c)
	}
}

func (i *inventory) insertToChain(p int, c *component) {
	if p >= len(i.chains) {
		i.chains = append(i.chains, make([]*chain, p-len(i.chains)+1)...)
	}
	i.chains[p] = &chain{c: c, next: i.chains[p]}
}

func (i *inventory) removeFromChain(p int, c *component) {
	prev, at := &i.chains[p], i.chains[p]
	for at.c != c {
		prev, at = &at.next, at.next
	}
	*prev = at.next
}
