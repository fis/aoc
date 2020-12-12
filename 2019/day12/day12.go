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

// Package day12 solves AoC 2019 day 12.
package day12

import (
	"strconv"
	"strings"

	"github.com/fis/aoc-go/glue"
)

func init() {
	glue.RegisterSolver(2019, 12, glue.LineSolver(solve))
}

func solve(lines []string) ([]int, error) {
	initialState := parseState(lines)

	state := *initialState
	run(&state, 1000)
	part1 := totalEnergy(&state)

	part2 := cycle(initialState)

	return []int{part1, part2}, nil
}

const (
	moons = 4
	dims  = 3
)

type dimState struct {
	pos [moons]int
	vel [moons]int
}

func run(state *[dims]dimState, steps int) {
	for s := 0; s < steps; s++ {
		for d := 0; d < dims; d++ {
			step1(&state[d])
		}
	}
}

func step1(state *dimState) {
	// gravity
	for i := 0; i < moons-1; i++ {
		for j := i + 1; j < moons; j++ {
			if state.pos[i] < state.pos[j] {
				state.vel[i]++
				state.vel[j]--
			} else if state.pos[i] > state.pos[j] {
				state.vel[i]--
				state.vel[j]++
			}
		}
	}
	// velocity
	for i := 0; i < moons; i++ {
		state.pos[i] += state.vel[i]
	}
}

func totalEnergy(state *[dims]dimState) int {
	tot := 0
	for i := 0; i < moons; i++ {
		pot, kin := 0, 0
		for d := 0; d < dims; d++ {
			pot += abs(state[d].pos[i])
			kin += abs(state[d].vel[i])
		}
		tot += pot * kin
	}
	return tot
}

func cycle(state *[dims]dimState) int {
	var cycles [dims]int
	for d := 0; d < dims; d++ {
		c, s := 0, state[d]
		for {
			c++
			step1(&s)
			if s == state[d] {
				break
			}
		}
		cycles[d] = c
	}
	return lcm(cycles)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func lcm(cycles [dims]int) int {
	for d := dims - 2; d >= 0; d-- {
		div := gcd(cycles[d], cycles[d+1])
		cycles[d] = cycles[d] * cycles[d+1] / div
	}
	return cycles[0]
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func parseState(lines []string) *[dims]dimState {
	s := new([dims]dimState)
	if len(lines) != moons {
		panic("ill met by moonlight")
	}
	for i, line := range lines {
		parts := strings.Split(line, ",")
		if len(parts) != dims {
			panic("lost in a strange dimension")
		}
		for d, part := range parts {
			digits := strings.TrimFunc(part, func(r rune) bool { return r != '-' && (r < '0' || r > '9') })
			num, _ := strconv.Atoi(digits)
			s[d].pos[i] = num
		}
	}
	return s
}
