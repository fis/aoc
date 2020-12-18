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

// Package day16 solves AoC 2018 day 16.
package day16

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"math/bits"
	"strconv"

	"github.com/fis/aoc-go/glue"
	"github.com/fis/aoc-go/util"
)

func init() {
	glue.RegisterSolver(2018, 16, glue.GenericSolver(solve))
}

func solve(r io.Reader) ([]string, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	parts := bytes.SplitN(data, []byte{'\n', '\n', '\n', '\n'}, 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("expected two major chunks, got %d", len(parts))
	}
	chunks, err := util.ScanAll(bytes.NewReader(parts[0]), util.ScanChunks)
	if err != nil {
		return nil, err
	}
	samples, err := parseSamples(chunks)
	if err != nil {
		return nil, err
	}
	prog, err := parseInsts(util.Lines(string(parts[1])))
	if err != nil {
		return nil, err
	}

	opMap, p1, err := findOpMap(samples)
	if err != nil {
		return nil, err
	}
	p2 := executeProg(prog, opMap)

	return []string{strconv.Itoa(p1), strconv.Itoa(p2)}, nil
}

type sample struct {
	inst          instruction
	before, after regState
}

func parseSamples(chunks []string) (out []sample, err error) {
	for _, chunk := range chunks {
		s, err := parseSample(util.Lines(chunk))
		if err != nil {
			return nil, err
		}
		out = append(out, s)
	}
	return out, nil
}

func parseSample(lines []string) (s sample, err error) {
	if len(lines) != 3 {
		return sample{}, fmt.Errorf("invalid sample: %v: expected 3 lines", lines)
	}
	if s.inst, err = parseInst(lines[1]); err != nil {
		return sample{}, err
	}
	if _, err := fmt.Sscanf(lines[0], "Before: [%d, %d, %d, %d]", &s.before[0], &s.before[1], &s.before[2], &s.before[3]); err != nil {
		return sample{}, fmt.Errorf("invalid before state: %s: %w", lines[0], err)
	}
	if _, err := fmt.Sscanf(lines[2], "After:  [%d, %d, %d, %d]", &s.after[0], &s.after[1], &s.after[2], &s.after[3]); err != nil {
		return sample{}, fmt.Errorf("invalid after state: %s: %w", lines[2], err)
	}
	return s, nil
}

type instruction struct {
	op      int
	a, b, c int
}

func parseInsts(lines []string) (prog []instruction, err error) {
	for _, line := range lines {
		inst, err := parseInst(line)
		if err != nil {
			return nil, err
		}
		prog = append(prog, inst)
	}
	return prog, nil
}

func parseInst(line string) (inst instruction, err error) {
	if _, err := fmt.Sscanf(line, "%d %d %d %d", &inst.op, &inst.a, &inst.b, &inst.c); err != nil {
		return instruction{}, fmt.Errorf("invalid instruction: %s: %w", line, err)
	}
	return inst, nil
}

type regState [4]int

func (r regState) equal(s regState) bool {
	return r[0] == s[0] && r[1] == s[1] && r[2] == s[2] && r[3] == s[3]
}

type opcode int

const (
	addr opcode = iota
	addi
	mulr
	muli
	banr
	bani
	borr
	bori
	setr
	seti
	gtir
	gtri
	gtrr
	eqir
	eqri
	eqrr
)

const (
	firstOp = addr
	lastOp  = eqrr
)

func findOpMap(samples []sample) (opMap [16]opcode, multiChoice int, err error) {
	possible := [16]uint{}
	for i := range possible {
		possible[i] = (uint(1) << (lastOp + 1)) - 1
	}
	for _, s := range samples {
		if s.inst.op < 0 || s.inst.op >= 16 {
			return [16]opcode{}, 0, fmt.Errorf("invalid opcode: %d", s.inst.op)
		}
		valid := validOps(s)
		if bits.OnesCount(valid) >= 3 {
			multiChoice++
		}
		possible[s.inst.op] &= valid
	}
	unassigned := (uint(1) << (lastOp + 1)) - 1
next:
	for unassigned > 0 {
		for i, p := range possible {
			if unassigned&(uint(1)<<i) == 0 {
				continue
			}
			if n := bits.OnesCount(p); n == 0 {
				return [16]opcode{}, 0, fmt.Errorf("impossible: no choices for %d", i)
			} else if n == 1 {
				opMap[i] = opcode(bits.TrailingZeros(p))
				unassigned &= ^(uint(1) << i)
				for i := range possible {
					possible[i] &= ^p
				}
				continue next
			}
		}
		return [16]opcode{}, 0, fmt.Errorf("impossible: ambiguous: %b", possible)
	}
	return opMap, multiChoice, nil
}

func validOps(s sample) (mask uint) {
	a, b, c := s.inst.a, s.inst.b, s.inst.c
	for op := firstOp; op <= lastOp; op++ {
		if s.after.equal(execute(op, a, b, c, s.before)) {
			mask |= uint(1) << op
		}
	}
	return mask
}

func executeProg(prog []instruction, opMap [16]opcode) int {
	var s regState
	for _, inst := range prog {
		s = execute(opMap[inst.op], inst.a, inst.b, inst.c, s)
	}
	return s[0]
}

func execute(op opcode, a, b, c int, r regState) regState {
	switch op {
	case addr:
		r[c] = r[a] + r[b]
	case addi:
		r[c] = r[a] + b
	case mulr:
		r[c] = r[a] * r[b]
	case muli:
		r[c] = r[a] * b
	case banr:
		r[c] = r[a] & r[b]
	case bani:
		r[c] = r[a] & b
	case borr:
		r[c] = r[a] | r[b]
	case bori:
		r[c] = r[a] | b
	case setr:
		r[c] = r[a]
	case seti:
		r[c] = a
	case gtir:
		r[c] = asInt(a > r[b])
	case gtri:
		r[c] = asInt(r[a] > b)
	case gtrr:
		r[c] = asInt(r[a] > r[b])
	case eqir:
		r[c] = asInt(a == r[b])
	case eqri:
		r[c] = asInt(r[a] == b)
	case eqrr:
		r[c] = asInt(r[a] == r[b])
	}
	return r
}

func asInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
