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

// Package day25 solves AoC 2017 day 25.
package day25

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2017, 25, glue.ChunkSolver(solve))
}

func solve(chunks []string) ([]string, error) {
	tm, steps, err := parseTM(chunks)
	if err != nil {
		return nil, err
	}

	tm.run(steps)
	p1 := tm.tape.popCount()

	return glue.Ints(p1), nil
}

type turingMachine struct {
	start  *turingState
	states []turingState
	tape   turingTape
	head   int
}

type turingState struct {
	label byte
	write [2]byte
	move  [2]int
	next  [2]*turingState
}

func (tm *turingMachine) run(steps int) {
	st := tm.start
	for step := 0; step < steps; step++ {
		old := tm.tape.read(tm.head)
		tm.tape.write(tm.head, st.write[old])
		tm.head += st.move[old]
		st = st.next[old]
	}
}

var regexpHead = regexp.MustCompile(`Begin in state ([A-Z])\.\nPerform a diagnostic checksum after (\d+) steps.`)
var regexpState = regexp.MustCompile(strings.Join([]string{
	`In state ([A-Z]):`,
	`  If the current value is 0:`,
	`    - Write the value ([01])\.`,
	`    - Move one slot to the (left|right)\.`,
	`    - Continue with state ([A-Z])\.`,
	`  If the current value is 1:`,
	`    - Write the value ([01])\.`,
	`    - Move one slot to the (left|right)\.`,
	`    - Continue with state ([A-Z])\.`,
}, `\n`))

func parseTM(blocks []string) (*turingMachine, int, error) {
	tm := &turingMachine{}
	moves := map[string]int{"left": -1, "right": +1}

	if len(blocks) < 2 {
		return nil, 0, fmt.Errorf("invalid TM: need at least 2 blocks, got %d", len(blocks))
	}
	tm.states = make([]turingState, len(blocks)-1)

	head := regexpHead.FindStringSubmatch(blocks[0])
	if head == nil {
		return nil, 0, fmt.Errorf("invalid TM: header %q does not match", blocks[0])
	}
	tm.start = &tm.states[head[1][0]-'A']
	steps, _ := strconv.Atoi(head[2])

	for _, block := range blocks[1:] {
		data := regexpState.FindStringSubmatch(block)
		if data == nil {
			return nil, 0, fmt.Errorf("invalid TM: block %q does not match", block)
		}
		s := &tm.states[data[1][0]-'A']
		s.label = data[1][0]
		s.write[0] = data[2][0] - '0'
		s.write[1] = data[5][0] - '0'
		s.move[0] = moves[data[3]]
		s.move[1] = moves[data[6]]
		s.next[0] = &tm.states[data[4][0]-'A']
		s.next[1] = &tm.states[data[7][0]-'A']
	}

	return tm, steps, nil
}

type turingTape struct {
	pos []byte
	neg []byte
}

func (t *turingTape) read(at int) byte {
	if at < 0 {
		at = -(at + 1)
		if at >= len(t.neg) {
			return 0
		}
		return t.neg[at]
	}
	if at >= len(t.pos) {
		return 0
	}
	return t.pos[at]
}

func (t *turingTape) write(at int, v byte) {
	if at < 0 {
		at = -(at + 1)
		if at < len(t.neg) {
			t.neg[at] = v
		} else if v != 0 {
			t.neg = append(t.neg, make([]byte, at+1-len(t.neg))...)
			t.neg[at] = v
		}
		return
	}
	if at < len(t.pos) {
		t.pos[at] = v
	} else if v != 0 {
		t.pos = append(t.pos, make([]byte, at+1-len(t.pos))...)
		t.pos[at] = v
	}
}

func (t *turingTape) popCount() (ones int) {
	for _, b := range t.pos {
		if b != 0 {
			ones++
		}
	}
	for _, b := range t.neg {
		if b != 0 {
			ones++
		}
	}
	return ones
}
