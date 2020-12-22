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

// Package day22 solves AoC 2020 day 22.
package day22

import (
	"encoding/binary"
	"fmt"
	"hash/fnv"
	"strconv"

	"github.com/fis/aoc-go/glue"
	"github.com/fis/aoc-go/util"
)

func init() {
	glue.RegisterSolver(2020, 22, glue.ChunkSolver(solve))
}

func solve(chunks []string) ([]int, error) {
	if len(chunks) != 2 {
		return nil, fmt.Errorf("expected 2 chunks, got %d", len(chunks))
	}
	p1, err := parseDeck(1, util.Lines(chunks[0]))
	if err != nil {
		return nil, err
	}
	p2, err := parseDeck(2, util.Lines(chunks[1]))
	if err != nil {
		return nil, err
	}

	part1 := game(p1.copy(p1.len(), p2.len()), p2.copy(p2.len(), p1.len()))
	part2 := recursive(p1, p2, map[uint64]int{})
	if part2 < 0 {
		part2 = -part2
	}

	return []int{part1, part2}, nil
}

func game(p1, p2 *deck) (score int) {
	for p1.len() > 0 && p2.len() > 0 {
		c1, c2 := p1.pop(), p2.pop()
		if c1 > c2 {
			p1.push2(c1, c2)
		} else {
			p2.push2(c2, c1)
		}
	}
	if p1.len() == 0 {
		p1 = p2
	}
	n := p1.len()
	p1.rangeCards(func(i, c int) {
		score += (n - i) * c
	})
	return score
}

func recursive(p1, p2 *deck, scores map[uint64]int) (score int) {
	i1, i2 := p1.hash(false), p2.hash(true)
	if score, ok := scores[i1^i2]; ok {
		return score
	}
	seen := map[uint64]struct{}{}
	for p1.len() > 0 && p2.len() > 0 {
		h1, h2 := p1.hash(false), p2.hash(true)
		if _, ok := seen[h1^h2]; ok {
			break
		}
		seen[h1^h2] = struct{}{}
		c1, c2 := p1.pop(), p2.pop()
		if p1.len() >= c1 && p2.len() >= c2 {
			s := recursive(p1.copy(c1, c2), p2.copy(c2, c1), scores)
			if s > 0 {
				p1.push2(c1, c2)
			} else {
				p2.push2(c2, c1)
			}
		} else if c1 > c2 {
			p1.push2(c1, c2)
		} else {
			p2.push2(c2, c1)
		}
	}
	sign := 1
	if p1.len() == 0 {
		p1, sign = p2, -1
	}
	n := p1.len()
	p1.rangeCards(func(i, c int) {
		score += (n - i) * c
	})
	score *= sign
	scores[i1^i2] = score
	return score
}

type deck struct {
	c           []int
	start, size int
}

func (d *deck) copy(cards, space int) *deck {
	d2 := &deck{c: make([]int, cards+space), start: 0, size: cards}
	if d.start+cards > len(d.c) {
		n := len(d.c) - d.start
		copy(d2.c[0:n], d.c[d.start:])
		copy(d2.c[n:cards], d.c[0:cards-n])
	} else {
		copy(d2.c[0:cards], d.c[d.start:d.start+cards])
	}
	return d2
}

func (d *deck) len() int {
	return d.size
}

func (d *deck) pop() int {
	if d.size == 0 {
		panic("pop on empty deck")
	}
	c := d.c[d.start]
	d.start, d.size = (d.start+1)%len(d.c), d.size-1
	return c
}

func (d *deck) push(c int) {
	if d.size == len(d.c) {
		panic("push on full deck")
	}
	d.c[(d.start+d.size)%len(d.c)] = c
	d.size++
}

func (d *deck) push2(c1, c2 int) {
	d.push(c1)
	d.push(c2)
}

func (d *deck) rangeCards(cb func(i, c int)) {
	for i := 0; i < d.size; i++ {
		cb(i, d.c[(d.start+i)%len(d.c)])
	}
}

func (d *deck) hash(flip bool) uint64 {
	h := fnv.New64a()
	buf := [4]byte{}
	d.rangeCards(func(_, c int) {
		binary.LittleEndian.PutUint32(buf[:], uint32(c))
		h.Write(buf[:])
	})
	s := h.Sum64()
	if flip {
		s = (s << 32) | (s >> 32)
	}
	return s
}

func parseDeck(player int, lines []string) (d *deck, err error) {
	hdr := fmt.Sprintf("Player %d:", player)
	if len(lines) < 1 || lines[0] != hdr {
		return nil, fmt.Errorf("bad/missing header in deck")
	}
	lines = lines[1:]
	c := make([]int, 2*len(lines))
	for i, line := range lines {
		c[i], err = strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
	}
	return &deck{c: c, start: 0, size: len(lines)}, nil
}
