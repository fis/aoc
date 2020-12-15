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

// Package day14 solves AoC 2020 day 14.
package day14

import (
	"fmt"

	"github.com/fis/aoc-go/glue"
)

func init() {
	glue.RegisterSolver(2020, 14, glue.LineSolver(solve))
}

func solve(lines []string) ([]int, error) {
	code, err := parseCode(lines)
	if err != nil {
		return nil, err
	}

	p1 := sumValues(evaluate1(code))
	p2 := sumValues(evaluate2(code))

	return []int{int(p1), int(p2)}, nil
}

type opcode int

const (
	opMask opcode = iota
	opSet
)

type instruction struct {
	op         opcode
	a1, a2, a3 uint
}

func parseCode(lines []string) (out []instruction, err error) {
	for _, line := range lines {
		var (
			mask       string
			a1, a2, a3 uint
		)
		if _, err := fmt.Sscanf(line, "mask = %s", &mask); err == nil {
			a1, a2, a3 = 0, 0, 0
			for i, c := range mask {
				bit := uint(1) << (35 - i)
				switch c {
				case '0':
					a1 |= bit
				case '1':
					a2 |= bit
				case 'X':
					a3 |= bit
				}
			}
			out = append(out, instruction{op: opMask, a1: a1, a2: a2, a3: a3})
		} else if _, err := fmt.Sscanf(line, "mem[%d] = %d", &a1, &a2); err == nil {
			out = append(out, instruction{op: opSet, a1: a1, a2: a2, a3: 0})
		} else {
			return nil, fmt.Errorf("invalid instruction: %s", line)
		}
	}
	return out, nil
}

func evaluate1(code []instruction) map[uint]uint {
	mem := make(map[uint]uint)
	orMask, andMask := uint(0), ^uint(0)
	for _, inst := range code {
		switch inst.op {
		case opMask:
			orMask, andMask = inst.a2, ^inst.a1
		case opSet:
			mem[inst.a1] = (inst.a2 | orMask) & andMask
		}
	}
	return mem
}

func evaluate2(code []instruction) map[uint]uint {
	mem := make(map[uint]uint)
	orMask, floatMask, floatBits := uint(0), ^uint(0), []uint{}
	for _, inst := range code {
		switch inst.op {
		case opMask:
			orMask, floatMask = inst.a2, ^inst.a3
			floatBits = makeFloatBits(inst.a3, floatBits)
		case opSet:
			baseAddr := (inst.a1 | orMask) & floatMask
			for _, f := range floatBits {
				mem[baseAddr|f] = inst.a2
			}
		}
	}
	return mem
}

func makeFloatBits(mask uint, buf []uint) []uint {
	ones := make([]uint, 0, 10)
	for mask != 0 {
		lowBits := mask ^ (mask - 1)
		lowBit := lowBits ^ (lowBits >> 1)
		ones = append(ones, lowBit)
		mask ^= lowBit
	}
	buf = buf[:0]
	for i, n := 0, 1<<len(ones); i < n; i++ {
		v := uint(0)
		for j, b := range ones {
			if i&(1<<j) != 0 {
				v |= b
			}
		}
		buf = append(buf, v)
	}
	return buf
}

func sumValues(mem map[uint]uint) (sum uint) {
	for _, v := range mem {
		sum += v
	}
	return sum
}
