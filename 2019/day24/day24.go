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

// Package day24 solves AoC 2019 day 24.
package day24

import (
	"fmt"

	"github.com/fis/aoc-go/glue"
)

func init() {
	glue.RegisterSolver(2019, 24, glue.LineSolver(solve))
}

func solve(lines []string) ([]int, error) {
	initial := parseState(lines)
	p1 := findRepeating(initial)
	p2 := countBugs(initial, 200)
	return []int{int(p1), p2}, nil
}

type state uint32

const (
	stateW    = 5
	stateH    = 5
	maxStates = 256
)

func findRepeating(s state) state {
	neigh := neighbors()
	seen := map[state]struct{}{s: {}}
	for {
		ps := s
		for i, bits := range neigh {
			count := state(0)
			for _, bit := range bits {
				count += (ps >> bit) & 1
			}
			live := ps&(1<<i) != 0
			if !live && (count == 1 || count == 2) {
				s |= 1 << i
			} else if live && count != 1 {
				s &= ^(1 << i)
			}
		}
		if _, ok := seen[s]; ok {
			return s
		}
		seen[s] = struct{}{}
	}
}

func countBugs(initial state, steps int) int {
	neigh := recursiveNeighbors()

	var s, ps [maxStates]state
	low, high := maxStates/2, maxStates/2
	s[low] = initial

	bugs := 0
	for t := initial; t > 0; t >>= 1 {
		if t&1 != 0 {
			bugs++
		}
	}

	for t := 0; t < steps; t++ {
		copy(ps[low-2:high+3], s[low-2:high+3])
		for d := low - 1; d <= high+1; d++ {
			for i, bitsets := range neigh {
				count := state(0)
				for _, rb := range bitsets {
					psd := ps[d+rb.depth]
					for _, bit := range rb.bits {
						count += (psd >> bit) & 1
					}
				}
				live := ps[d]&(1<<i) != 0
				if !live && (count == 1 || count == 2) {
					bugs++
					s[d] |= 1 << i
				} else if live && count != 1 {
					bugs--
					s[d] &= ^(1 << i)
				}
			}
		}
		if s[low-1] != 0 {
			low--
		}
		if s[high+1] != 0 {
			high++
		}
	}

	return bugs
}

func neighbors() [stateW * stateH][]state {
	var neigh [stateW * stateH][]state
	for i := range neigh {
		x, y := i%stateW, i/stateW
		for _, n := range [4][2]int{{x - 1, y}, {x + 1, y}, {x, y - 1}, {x, y + 1}} {
			if n[0] >= 0 && n[0] < stateW && n[1] >= 0 && n[1] < stateH {
				neigh[i] = append(neigh[i], state(stateW*n[1]+n[0]))
			}
		}
	}
	return neigh
}

type relbits struct {
	depth int
	bits  []state
}

func recursiveNeighbors() [stateW * stateH][]relbits {
	var rneigh [stateW * stateH][]relbits
	for i, bits := range neighbors() {
		rneigh[i] = append(rneigh[i], relbits{depth: 0, bits: bits})
	}
	mid := state(stateH/2*stateW + stateW/2)
	rneigh[mid] = nil
	{
		var top, bottom []state
		for x := 0; x < stateW; x++ {
			top = append(top, state(x))
			bottom = append(bottom, state((stateH-1)*stateW+x))
		}
		rneigh[mid-stateW] = append(rneigh[mid-stateW], relbits{depth: 1, bits: top})
		rneigh[mid+stateW] = append(rneigh[mid+stateW], relbits{depth: 1, bits: bottom})
	}
	{
		var left, right []state
		for y := 0; y < stateH; y++ {
			left = append(left, state(y*stateW))
			right = append(right, state((y+1)*stateW-1))
		}
		rneigh[mid-1] = append(rneigh[mid-1], relbits{depth: 1, bits: left})
		rneigh[mid+1] = append(rneigh[mid+1], relbits{depth: 1, bits: right})
	}
	rneigh[0] = append(rneigh[0], relbits{depth: -1, bits: []state{mid - 1, mid - stateW}})
	rneigh[stateW-1] = append(rneigh[stateW-1], relbits{depth: -1, bits: []state{mid + 1, mid - stateW}})
	rneigh[(stateH-1)*stateW] = append(rneigh[(stateH-1)*stateW], relbits{depth: -1, bits: []state{mid - 1, mid + stateW}})
	rneigh[stateH*stateW-1] = append(rneigh[stateH*stateW-1], relbits{depth: -1, bits: []state{mid + 1, mid + stateW}})
	for x := 1; x < stateW-1; x++ {
		rneigh[x] = append(rneigh[x], relbits{depth: -1, bits: []state{mid - stateW}})
		rneigh[(stateH-1)*stateW+x] = append(rneigh[(stateH-1)*stateW+x], relbits{depth: -1, bits: []state{mid + stateW}})
	}
	for y := 1; y < stateH-1; y++ {
		rneigh[y*stateW] = append(rneigh[y*stateW], relbits{depth: -1, bits: []state{mid - 1}})
		rneigh[(y+1)*stateW-1] = append(rneigh[(y+1)*stateW-1], relbits{depth: -1, bits: []state{mid + 1}})
	}
	return rneigh
}

func parseState(lines []string) state {
	s := state(0)
	for y, line := range lines {
		for x, r := range line {
			if r == '#' {
				s |= 1 << (stateW*y + x)
			}
		}
	}
	return s
}

func printState(s state) {
	for y := 0; y < stateH; y++ {
		for x := 0; x < stateW; x++ {
			c := '.'
			if s&(1<<(stateW*y+x)) != 0 {
				c = '#'
			}
			fmt.Printf("%c", c)
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}
