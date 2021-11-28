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

// Package day12 solves AoC 2018 day 12.
package day12

import (
	"fmt"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2018, 12, glue.ChunkSolver(solve))
}

const initialPrefix = "initial state: "

func solve(chunks []string) ([]string, error) {
	if len(chunks) != 2 {
		return nil, fmt.Errorf("expected 2 chunks, got %d", len(chunks))
	}
	if !strings.HasPrefix(chunks[0], initialPrefix) {
		return nil, fmt.Errorf("invalid header: %s", chunks[0])
	}

	state := parseState(strings.TrimPrefix(chunks[0], initialPrefix))
	rules := parseRules(util.Lines(chunks[1]))

	state.evolve(20, rules)
	part1 := state.checksum()

	shift := state.findFixed(rules)
	state.offset += (50000000000 - state.gen) * shift
	part2 := state.checksum()

	return glue.Ints(part1, part2), nil
}

type stateVector struct {
	data   []byte
	offset int
	gen    int
}

func parseState(text string) stateVector {
	return stateVector{
		data:   asBits("...." + text + "...."),
		offset: -4,
	}
}

func (s *stateVector) evolve(generations int, r *ruleSet) {
	var next []byte
	for g := 0; g < generations; g++ {
		var disp int
		next, disp = r.step(s.data, next)
		s.data, s.offset, s.gen, next = next, s.offset+disp, s.gen+1, s.data
	}
}

func (s *stateVector) findFixed(r *ruleSet) (disp int) {
	var next []byte
	for {
		next, disp = r.step(s.data, next)
		if equalBytes(s.data, next) {
			return disp
		}
		s.data, s.offset, s.gen, next = next, s.offset+disp, s.gen+1, s.data
	}
}

func (s stateVector) checksum() (sum int) {
	for i, v := range s.data {
		if v != 0 {
			sum += s.offset + i
		}
	}
	return sum
}

type ruleSet [32]byte

func parseRules(lines []string) *ruleSet {
	r := new(ruleSet)
	for _, line := range lines {
		if len(line) != 10 || line[5:9] != " => " {
			continue
		}
		r[idx(asBits(line[0:5]))] = asBit(line[9])
	}
	return r
}

func (r *ruleSet) lookup(d []byte) byte {
	return r[idx(d)]
}

func (r *ruleSet) step(in, out []byte) ([]byte, int) {
	out = append(out[:0], 0, 0, 0, 0)
	disp := -2
	for x := 0; x+5 <= len(in); x++ {
		b := r.lookup(in[x : x+5])
		if len(out) == 4 && b == 0 {
			disp++
		} else {
			out = append(out, b)
		}
	}
	for out[len(out)-4] != 0 || out[len(out)-3] != 0 || out[len(out)-2] != 0 || out[len(out)-1] != 0 {
		out = append(out, 0)
	}
	return out, disp
}

func idx(d []byte) int {
	return int((d[0] << 4) | (d[1] << 3) | (d[2] << 2) | (d[3] << 1) | d[4])
}

func asBit(c byte) byte {
	if c == '#' {
		return 1
	}
	return 0
}

func asBits(s string) (out []byte) {
	for _, c := range []byte(s) {
		out = append(out, asBit(c))
	}
	return out
}

func equalBytes(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
