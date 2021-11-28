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
	for i := 0; i < nc; i++ {
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
