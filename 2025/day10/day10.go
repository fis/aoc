// Copyright 2025 Google LLC
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

// Package day10 solves AoC 2025 day 10.
package day10

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/fi"
	"github.com/fis/aoc/util/fn"
)

func init() {
	glue.RegisterSolver(2025, 10, glue.LineSolver(glue.WithParser(parseMachine, solve)))
}

func solve(machines []machine) ([]string, error) {
	p1 := fn.SumF(machines, machine.configLights)
	p2 := fn.SumF(machines, machine.configJolts)
	return glue.Ints(p1, p2), nil
}

type machine struct {
	lights  uint32
	buttons []uint32
	jolts   []int
}

func (m machine) configLights() (totalPresses int) {
	for n, cfg := range fi.HammingOrder(len(m.buttons)) {
		lights := uint32(0)
		for i, b := range m.buttons {
			if cfg&(1<<i) != 0 {
				lights ^= b
			}
		}
		if lights == m.lights {
			return n
		}
	}
	return -1
}

func (m machine) configJolts() (totalPresses int) {
	fixedPresses, bounds := m.simplify()
	if len(m.buttons) == 0 {
		return fixedPresses
	}
	solvedPresses := lpsolve(m.buttons, m.jolts, bounds)
	return fixedPresses + solvedPresses
}

// TODO: maybe try implementing this
func solveEqs(A [][]int, b []int) (x []int) {
	m, n := len(A), len(A[0])
	_, _ = m, n

	// build U such that U is unimodular and H = UA is upper triangular

	// find an integer solution for the system Hy = b, where y = U^-1 x

	// extract x = Uy

	return nil
}

/*
[....###..] (0,2,3,4,5,6,7) (0,2,3,4,7) (5,6) (2,7,8) (0,1,2,3,6,7,8) (0,1,2,3,4,5,6,7) (4,8) (2,4,5,6,7,8) (1,4) {30,31,60,30,46,40,51,60,42}

 p0 + p1 + p4 + p5 = 30
 p4 + p5 + p8 = 31
 p0 + p1 + p3 + p4 + p5 + p7 = 60
 p0 + p1 + p4 + p5 = 30
 p0 + p1 + p5 + p6 + p7 + p8 = 46
 p0 + p2 + p5 + p7 = 40
 p0 + p2 + p4 + p5 + p7 = 51
 p0 + p1 + p3 + p4 + p5 + p7 = 60
 p3 + p4 + p6 + p7 = 42

*/

func (m *machine) simplify() (fixedPresses int, bounds [][2]int) {
	bounds = make([][2]int, len(m.buttons))
	for i := range len(m.buttons) {
		bounds[i] = [2]int{0, math.MaxInt}
	}

	changed := true
	for changed {
		changed = false

		// remove joltages that are wired to same set of buttons (must have identical levels, otherwise unsolvable)

		wiring := m.wiring()
		wiringMap := make(map[uint32]int)
		for ji, w := range wiring {
			wiringMap[w] = m.jolts[ji]
		}
		if len(wiringMap) != len(wiring) {
			changed = true
			wiring, m.jolts = make([]uint32, 0, len(wiringMap)), make([]int, 0, len(wiringMap))
			for w, j := range wiringMap {
				wiring = append(wiring, w)
				m.jolts = append(m.jolts, j)
			}
			m.wireButtons(wiring)
		}

		// update upper bounds: button can't be pressed more times than any joltage if affects (minus lower bounds of other buttons affecting that joltage)

		for bi, b := range m.buttons {
			for ji, j := range m.jolts {
				if b&(1<<ji) == 0 {
					continue
				}
				for oi, o := range m.buttons {
					if oi == bi || o&(1<<ji) == 0 {
						continue
					}
					j -= bounds[oi][0]
				}
				if j < bounds[bi][1] {
					changed = true
					bounds[bi][1] = j
				}
			}
		}

		// update lower bounds: a button must be pressed at least as many times as required to bump up any joltage upper bounds of other buttons can't get high enough

		for bi, b := range m.buttons {
			for ji, j := range m.jolts {
				if b&(1<<ji) == 0 {
					continue
				}
				for oi, o := range m.buttons {
					if oi == bi || o&(1<<ji) == 0 {
						continue
					}
					j -= bounds[oi][1]
				}
				if j > bounds[bi][0] {
					changed = true
					bounds[bi][0] = j
				}
			}
		}

		// remove buttons that have equal bounds completely

		for bi := 0; bi < len(m.buttons); bi++ {
			if bounds[bi][0] == bounds[bi][1] {
				changed = true
				fixedPresses += bounds[bi][0]
				for ji := range len(m.jolts) {
					if m.buttons[bi]&(1<<ji) != 0 {
						m.jolts[ji] -= bounds[bi][0]
					}
				}
				m.buttons = slices.Delete(m.buttons, bi, bi+1)
				bounds = slices.Delete(bounds, bi, bi+1)
				bi--
			}
		}

		// remove joltages that have no buttons referring to them

		wiring = m.wiring()
		for ji := 0; ji < len(m.jolts); ji++ {
			if wiring[ji] == 0 {
				changed = true
				m.jolts = slices.Delete(m.jolts, ji, ji+1)
				wiring = slices.Delete(wiring, ji, ji+1)
				for bi, b := range m.buttons {
					mask := uint32(1)<<ji - 1
					m.buttons[bi] = (b & mask) | ((b >> 1) &^ mask)
				}
				ji--
			}
		}
	}

	return fixedPresses, bounds
}

func (m machine) wiring() (wiring []uint32) {
	wiring = make([]uint32, len(m.jolts))
	for bi, b := range m.buttons {
		for ji := range len(m.jolts) {
			if b&(1<<ji) != 0 {
				wiring[ji] |= 1 << bi
			}
		}
	}
	return wiring
}

func (m machine) wireButtons(wiring []uint32) {
	for bi := range m.buttons {
		m.buttons[bi] = 0
	}
	for ji, w := range wiring {
		for bi := range len(m.buttons) {
			if w&(1<<bi) != 0 {
				m.buttons[bi] |= 1 << ji
			}
		}
	}
}

func parseMachine(line string) (m machine, err error) {
	words := util.Words(line)
	if len(words) < 3 {
		return machine{}, fmt.Errorf("expected at least 3 words, got: %s", line)
	}
	lights, buttons, jolts := words[0], words[1:len(words)-1], words[len(words)-1]
	if len(lights) < 3 || lights[0] != '[' || lights[len(lights)-1] != ']' {
		return machine{}, fmt.Errorf("invalid lights: expected [...], got: %s", lights)
	}
	for i, c := range []byte(lights[1 : len(lights)-1]) {
		switch c {
		case '.':
			// zero by default
		case '#':
			m.lights |= 1 << i
		default:
			return machine{}, fmt.Errorf("invalid lights: expected ./#, got: %c", c)
		}
	}
	m.buttons = make([]uint32, len(buttons))
	for i, bs := range buttons {
		for _, n := range util.Ints(bs) {
			m.buttons[i] |= 1 << n
		}
	}
	if len(jolts) < 3 || jolts[0] != '{' || jolts[len(jolts)-1] != '}' {
		return machine{}, fmt.Errorf("invalid jolts: expected {...}, got: %s", jolts)
	}
	m.jolts = util.Ints(jolts)
	return m, nil
}

func (m machine) String() string {
	var s strings.Builder
	s.WriteByte('[')
	for i := range len(m.jolts) {
		if m.lights&(1<<i) != 0 {
			s.WriteByte('#')
		} else {
			s.WriteByte('.')
		}
	}
	s.WriteByte(']')
	for _, b := range m.buttons {
		s.WriteByte(' ')
		s.WriteByte('(')
		first := true
		for i := range 32 {
			if b&(1<<i) != 0 {
				if first {
					first = false
				} else {
					s.WriteByte(',')
				}
				fmt.Fprintf(&s, "%d", i)
			}
		}
		s.WriteByte(')')
	}
	s.WriteByte(' ')
	s.WriteByte('{')
	for i, j := range m.jolts {
		if i > 0 {
			s.WriteByte(',')
		}
		fmt.Fprintf(&s, "%d", j)
	}
	s.WriteByte('}')
	return s.String()
}
