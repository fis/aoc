// Copyright 2022 Google LLC
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

// Package day17 solves AoC 2022 day 17.
package day17

import (
	"fmt"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util/fn"
)

func init() {
	glue.RegisterSolver(2022, 17, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	if len(lines) != 1 {
		return nil, fmt.Errorf("expecting one line, got %d", len(lines))
	}
	jets := lines[0]
	p1 := dropRocks(jets, 2022)
	p2 := analyzeRocks(jets, 1000000000000)
	return glue.Ints(p1, p2), nil
}

func dropRocks(jets string, rocks int) int {
	sim, rock := newSimulation(jets), 0
	for rock < rocks {
		if sim.step() {
			rock++
		}
	}
	return sim.height()
}

func analyzeRocks(jets string, rocks int) int {
	sim, rock := newSimulation(jets), 0
	for sim.height() < fingerprintSize {
		for i := 0; i < len(jets); i++ {
			if sim.step() {
				rock++
			}
		}
	}

	type state struct {
		rock   int
		height int
	}
	states := make(map[fingerprint]state)
	discarded := 0

outerLoop:
	for {
		discarded += sim.trimHeight(fingerprintSize)
		fp := sim.fingerprint()
		if prev, ok := states[fp]; ok {
			skipRocks := rock - prev.rock
			skipHeight := discarded + sim.height() - prev.height
			n := (rocks - rock) / skipRocks
			rock += n * skipRocks
			discarded += n * skipHeight
			break outerLoop
		}
		states[fp] = state{rock: rock, height: discarded + sim.height()}
		for i := 0; i < len(jets); i++ {
			if sim.step() {
				rock++
				if rock == rocks {
					break outerLoop
				}
			}
		}
	}

	for rock < rocks {
		if sim.step() {
			rock++
		}
	}

	return discarded + sim.height()
}

type simulation struct {
	chute        []byte
	airGap       int
	jets         string
	jetI, shapeI int
	rock         int
	rockX, rockY int
}

func newSimulation(jets string) *simulation {
	sim := &simulation{jets: jets, shapeI: len(rockShapes) - 1}
	sim.createRock()
	return sim
}

func (sim *simulation) height() int {
	return len(sim.chute) - sim.airGap
}

func (sim *simulation) trimHeight(newHeight int) int {
	if drop := sim.height() - newHeight; drop > 0 {
		copy(sim.chute, sim.chute[drop:])
		sim.chute = sim.chute[:len(sim.chute)-drop]
		sim.rockY -= drop
		return drop
	}
	return 0
}

func (sim *simulation) createRock() {
	sim.shapeI++
	if sim.shapeI == len(rockShapes) {
		sim.shapeI = 0
	}
	rs := rockShapes[sim.shapeI]
	if addH := 3 + rs.h - sim.airGap; addH > 0 {
		sim.chute = append(sim.chute, make([]byte, addH)...)
		sim.airGap += addH
	}
	sim.rockX, sim.rockY = 2, len(sim.chute)-sim.airGap+3
}

func (sim *simulation) step() bool {
	jet := sim.jets[sim.jetI]
	sim.jetI++
	if sim.jetI == len(sim.jets) {
		sim.jetI = 0
	}

	rs := rockShapes[sim.shapeI]

	newX := fn.If(jet == '<', sim.rockX-1, sim.rockX+1)
	if newX >= 0 && newX+rs.w <= 7 {
		blocked := false
		for y := 0; y < rs.h; y++ {
			if sim.chute[sim.rockY+y]&(rs.shape[y]<<(7-rs.w-newX)) != 0 {
				blocked = true
				break
			}
		}
		if !blocked {
			sim.rockX = newX
		}
	}

	newY := sim.rockY - 1
	blocked := newY < 0
	if !blocked {
		for y := 0; y < rs.h; y++ {
			if sim.chute[newY+y]&(rs.shape[y]<<(7-rs.w-sim.rockX)) != 0 {
				blocked = true
				break
			}
		}
	}
	if !blocked {
		sim.rockY = newY
		return false
	}

	for y := 0; y < rs.h; y++ {
		sim.chute[sim.rockY+y] |= rs.shape[y] << (7 - rs.w - sim.rockX)
	}
	topY := max(len(sim.chute)-sim.airGap, sim.rockY+rs.h)
	sim.airGap = len(sim.chute) - topY

	sim.createRock()
	return true
}

const fingerprintSize = 64 // must be large enough that discarding history past this won't affect the results

type fingerprint struct {
	chute        [fingerprintSize]byte
	shapeI       int
	rockX, rockY int
}

func (sim *simulation) fingerprint() fingerprint {
	top := sim.height()
	return fingerprint{
		chute:  *(*[fingerprintSize]byte)(sim.chute[top-fingerprintSize:]),
		shapeI: sim.shapeI,
		rockX:  sim.rockX - top, rockY: sim.rockY - top,
	}
}

type rockShape struct {
	w, h  int
	shape [4]byte
}

var rockShapes = [5]rockShape{
	{w: 4, h: 1, shape: [4]byte{0b1111}},
	{w: 3, h: 3, shape: [4]byte{0b010, 0b111, 0b010}},
	{w: 3, h: 3, shape: [4]byte{0b111, 0b001, 0b001}},
	{w: 1, h: 4, shape: [4]byte{0b1, 0b1, 0b1, 0b1}},
	{w: 2, h: 2, shape: [4]byte{0b11, 0b11}},
}
