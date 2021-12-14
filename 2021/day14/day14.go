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

// Package day14 solves AoC 2021 day 14.
package day14

import (
	"fmt"
	"strings"

	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2021, 14, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	rb, initial, ends, err := parsePairRules(lines)
	if err != nil {
		return nil, err
	}
	c1 := rb.updateCountsN(initial, 10)
	c2 := rb.updateCountsN(c1, 30)
	q1 := rb.countElements(c1, ends)
	q2 := rb.countElements(c2, ends)
	return glue.Ints(maxDiff(q1), maxDiff(q2)), nil
}

// Approach A: keep track of the count of pairs.

type pairRuleBook struct {
	numElements int
	pairs       [][2]byte
	rules       [][2]int
}

func (rb *pairRuleBook) countElements(pairCounts []int, ends [2]byte) (quants []int) {
	quants = make([]int, rb.numElements)
	for p, c := range pairCounts {
		pe := rb.pairs[p]
		quants[pe[0]] += c
		quants[pe[1]] += c
	}
	quants[ends[0]]++
	quants[ends[1]]++
	for i := range quants {
		quants[i] /= 2
	}
	return quants
}

func (rb *pairRuleBook) updateCountsN(in []int, steps int) (out []int) {
	cur, next := append([]int(nil), in...), []int(nil)
	for i := 0; i < steps; i++ {
		next = rb.updateCounts(cur, next)
		cur, next = next, cur
	}
	return cur
}

func (rb *pairRuleBook) updateCounts(in, buf []int) (out []int) {
	buf = append(buf[:0], make([]int, len(rb.pairs))...)
	for p, c := range in {
		r := rb.rules[p]
		buf[r[0]] += c
		buf[r[1]] += c
	}
	return buf
}

func parsePairRules(lines []string) (rb *pairRuleBook, initial []int, ends [2]byte, err error) {
	if len(lines) < 3 || len(lines[0]) < 2 || lines[1] != "" {
		return nil, nil, [2]byte{}, fmt.Errorf("invalid rulebook: expected initial polymer, blank line, list of rules")
	}

	elementMap := map[byte]byte{}
	mapElement := func(k byte) byte {
		if v, ok := elementMap[k]; ok {
			return v
		}
		v := byte(len(elementMap))
		elementMap[k] = v
		return v
	}

	pairs := [][2]byte(nil)
	pairMap := map[[2]byte]int{}
	mapPair := func(a, b byte) int {
		ab := [2]byte{a, b}
		if p, ok := pairMap[ab]; ok {
			return p
		}
		p := len(pairs)
		pairs = append(pairs, ab)
		pairMap[ab] = p
		return p
	}

	rules := [][2]int(nil)
	for _, line := range lines[2:] {
		parts := strings.Split(line, " -> ")
		if len(parts) != 2 || len(parts[0]) != 2 || len(parts[1]) != 1 {
			return nil, nil, [2]byte{}, fmt.Errorf("invalid rule, expected \"XY -> Z\", got %q", line)
		}
		a, b, c := mapElement(parts[0][0]), mapElement(parts[0][1]), mapElement(parts[1][0])
		p, p1, p2 := mapPair(a, b), mapPair(a, c), mapPair(c, b)
		if p >= len(rules) {
			rules = append(rules, make([][2]int, p-len(rules)+1)...)
		}
		rules[p] = [2]int{p1, p2}
	}

	initial = make([]int, len(pairs))
	for i := 0; i < len(lines[0])-1; i++ {
		a, b := mapElement(lines[0][i]), mapElement(lines[0][i+1])
		p, ok := pairMap[[2]byte{a, b}]
		if !ok {
			return nil, nil, [2]byte{}, fmt.Errorf("unknown pair %s in initial polymer", lines[0][i:i+2])
		}
		initial[p]++
	}
	start, end := mapElement(lines[0][0]), mapElement(lines[0][len(lines[0])-1])

	return &pairRuleBook{numElements: len(elementMap), pairs: pairs, rules: rules}, initial, [2]byte{start, end}, nil
}

// Approach B: recurse on pairs with memoization.

type elementCounter struct {
	pairCache [][][][]int
}

func (ec *elementCounter) countPolymer(polymer []byte, steps int, rb *ruleBook) []int {
	quants := make([]int, len(rb.elements))
	for i := 0; i < len(polymer)-1; i++ {
		qp := ec.countPair(polymer[i], polymer[i+1], steps, rb)
		for j := range quants {
			quants[j] += qp[j]
		}
	}
	if len(polymer) >= 2 {
		for _, e := range polymer[1 : len(polymer)-1] {
			quants[e]--
		}
	}
	return quants
}

func (ec *elementCounter) countPair(a, b byte, steps int, rb *ruleBook) []int {
	switch steps {
	case 0:
		quants := make([]int, len(rb.elements))
		quants[a]++
		quants[b]++
		return quants
	case 1:
		quants := make([]int, len(rb.elements))
		quants[a]++
		quants[b]++
		quants[rb.rules[a][b]]++
		return quants
	default:
		if quants := ec.cacheGet(a, b, steps); quants != nil {
			return quants
		}
		c := rb.rules[a][b]
		qa, qb := ec.countPair(a, c, steps-1, rb), ec.countPair(c, b, steps-1, rb)
		quants := make([]int, len(rb.elements))
		for i := range quants {
			quants[i] = qa[i] + qb[i]
		}
		quants[c]-- // double-counted
		ec.cachePut(a, b, steps, quants)
		return quants
	}
}

func (ec *elementCounter) cacheGet(a, b byte, steps int) []int {
	if int(a) >= len(ec.pairCache) {
		return nil
	}
	if int(b) >= len(ec.pairCache[a]) {
		return nil
	}
	if steps >= len(ec.pairCache[a][b]) {
		return nil
	}
	return ec.pairCache[a][b][steps]
}

func (ec *elementCounter) cachePut(a, b byte, steps int, quants []int) {
	if ia := int(a); ia >= len(ec.pairCache) {
		ec.pairCache = append(ec.pairCache, make([][][][]int, ia-len(ec.pairCache)+1)...)
	}
	if ib := int(b); ib >= len(ec.pairCache[a]) {
		ec.pairCache[a] = append(ec.pairCache[a], make([][][]int, ib-len(ec.pairCache[a])+1)...)
	}
	if steps >= len(ec.pairCache[a][b]) {
		ec.pairCache[a][b] = append(ec.pairCache[a][b], make([][]int, steps-len(ec.pairCache[a][b])+1)...)
	}
	ec.pairCache[a][b][steps] = quants
}

type ruleBook struct {
	elements []byte
	rules    [][]byte
}

func (rb *ruleBook) expandN(in []byte, steps int) (out []byte) {
	cur, next := append([]byte(nil), in...), []byte(nil)
	for i := 0; i < steps; i++ {
		next = rb.expand(cur, next)
		cur, next = next, cur
	}
	return cur
}

func (rb *ruleBook) expand(in, buf []byte) (out []byte) {
	out = buf[:0]
	for i := 0; i < len(in)-1; i++ {
		a, b := in[i], in[i+1]
		if i == 0 {
			out = append(out, a)
		}
		out = append(out, rb.rules[a][b])
		out = append(out, b)
	}
	return out
}

func (rb *ruleBook) asString(polymer []byte) string {
	s := strings.Builder{}
	for _, e := range polymer {
		s.WriteByte(rb.elements[e])
	}
	return s.String()
}

func parseRules(lines []string) (rb *ruleBook, initial []byte, err error) {
	if len(lines) < 3 || lines[1] != "" {
		return nil, nil, fmt.Errorf("invalid rulebook: expected initial polymer, blank line, list of rules")
	}

	elementMap := [256]byte{}
	elements := []byte(nil)
	mapElement := func(k byte) byte {
		if v := elementMap[k]; v != 0 {
			return v - 1
		}
		v := byte(len(elements))
		elements = append(elements, k)
		elementMap[k] = v + 1
		return v
	}

	rules := [][]byte(nil)
	for _, line := range lines[2:] {
		parts := strings.Split(line, " -> ")
		if len(parts) != 2 || len(parts[0]) != 2 || len(parts[1]) != 1 {
			return nil, nil, fmt.Errorf("invalid rule, expected \"XY -> Z\", got %q", line)
		}
		a, b, c := mapElement(parts[0][0]), mapElement(parts[0][1]), mapElement(parts[1][0])
		if ia := int(a); ia >= len(rules) {
			rules = append(rules, make([][]byte, ia-len(rules)+1)...)
		}
		if ib := int(b); ib >= len(rules[a]) {
			rules[a] = append(rules[a], make([]byte, ib-len(rules[a])+1)...)
		}
		rules[a][b] = c
	}

	for i := 0; i < len(lines[0]); i++ {
		e := elementMap[lines[0][i]]
		if e == 0 {
			return nil, nil, fmt.Errorf("unknown element %c in initial polymer", e)
		}
		initial = append(initial, e-1)
	}

	return &ruleBook{elements: elements, rules: rules}, initial, nil
}

// Common utilities

func maxDiff(quants []int) int {
	min, max := quants[0], quants[0]
	for _, q := range quants[1:] {
		if q < min {
			min = q
		}
		if q > max {
			max = q
		}
	}
	return max - min
}
