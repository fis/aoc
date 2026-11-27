// Package day08 solves AoC 2018 day 8.
package day08

import (
	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2018, 8, glue.IntSolver(solve))
}

func solve(input []int) ([]string, error) {
	n, _ := parseLicense(input)
	cs, v := n.checksum(), n.value()
	return glue.Ints(cs, v), nil
}

type node struct {
	child []node
	data  []int
}

func parseLicense(input []int) (n node, consumed int) {
	nc, nd, input := input[0], input[1], input[2:]
	consumed += 2
	for range nc {
		c, cl := parseLicense(input)
		n.child = append(n.child, c)
		input = input[cl:]
		consumed += cl
	}
	n.data = input[:nd]
	consumed += nd
	return n, consumed
}

func (n *node) checksum() (cs int) {
	for _, c := range n.child {
		cs += c.checksum()
	}
	for _, d := range n.data {
		cs += d
	}
	return cs
}

func (n *node) value() (v int) {
	if len(n.child) == 0 {
		for _, d := range n.data {
			v += d
		}
		return v
	}

	var cv []int
	for _, c := range n.child {
		cv = append(cv, c.value())
	}
	for _, d := range n.data {
		if d >= 1 && d <= len(cv) {
			v += cv[d-1]
		}
	}
	return v
}
