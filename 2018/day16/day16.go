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

	"github.com/fis/aoc-go/2018/cpu"
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
	before, after cpu.State
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
	if _, err := fmt.Sscanf(lines[0], "Before: [%d, %d, %d, %d]", &s.before.R[0], &s.before.R[1], &s.before.R[2], &s.before.R[3]); err != nil {
		return sample{}, fmt.Errorf("invalid before state: %s: %w", lines[0], err)
	}
	if _, err := fmt.Sscanf(lines[2], "After:  [%d, %d, %d, %d]", &s.after.R[0], &s.after.R[1], &s.after.R[2], &s.after.R[3]); err != nil {
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

func findOpMap(samples []sample) (opMap [16]cpu.Op, multiChoice int, err error) {
	possible := [16]uint{}
	for i := range possible {
		possible[i] = (uint(1) << (cpu.LastOp + 1)) - 1
	}
	for _, s := range samples {
		if s.inst.op < 0 || s.inst.op >= 16 {
			return [16]cpu.Op{}, 0, fmt.Errorf("invalid opcode: %d", s.inst.op)
		}
		valid := validOps(s)
		if bits.OnesCount(valid) >= 3 {
			multiChoice++
		}
		possible[s.inst.op] &= valid
	}
	unassigned := (uint(1) << (cpu.LastOp + 1)) - 1
next:
	for unassigned > 0 {
		for i, p := range possible {
			if unassigned&(uint(1)<<i) == 0 {
				continue
			}
			if n := bits.OnesCount(p); n == 0 {
				return [16]cpu.Op{}, 0, fmt.Errorf("impossible: no choices for %d", i)
			} else if n == 1 {
				opMap[i] = cpu.Op(bits.TrailingZeros(p))
				unassigned &^= uint(1) << i
				for i := range possible {
					possible[i] &^= p
				}
				continue next
			}
		}
		return [16]cpu.Op{}, 0, fmt.Errorf("impossible: ambiguous: %b", possible)
	}
	return opMap, multiChoice, nil
}

func validOps(s sample) (mask uint) {
	a, b, c := s.inst.a, s.inst.b, s.inst.c
	for op := cpu.FirstOp; op <= cpu.LastOp; op++ {
		xs := s.before
		xs.Step(op, a, b, c)
		if xs.R == s.after.R {
			mask |= uint(1) << op
		}
	}
	return mask
}

func executeProg(prog []instruction, opMap [16]cpu.Op) int {
	var s cpu.State
	for _, inst := range prog {
		s.Step(opMap[inst.op], inst.a, inst.b, inst.c)
	}
	return s.R[0]
}
