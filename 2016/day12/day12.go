// Copyright 2023 Google LLC
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

// Package day12 solves AoC 2016 day 12.
package day12

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util/fn"
)

func init() {
	glue.RegisterSolver(2016, 12, glue.LineSolver(glue.WithParser(parseInst, solve)))

}

func solve(prog []inst) ([]string, error) {
	p1 := runProgram(prog, 0)
	p2 := runProgram(prog, 1)
	return glue.Ints(p1, p2), nil
}

func runProgram(prog []inst, initC int) int {
	regs := [4]int{2: initC}
	ip := 0
loop:
	for ip >= 0 && ip < len(prog) {
		switch i := prog[ip]; i.op {
		case opCpyIR:
			regs[i.y] = i.x
		case opCpyRR:
			regs[i.y] = regs[i.x]
		case opInc:
			regs[i.x]++
		case opDec:
			regs[i.x]--
		case opJnzR:
			if regs[i.x] != 0 {
				ip += i.y
				continue loop
			}
		case opJnzI:
			if i.x != 0 {
				ip += i.y
				continue loop
			}
		}
		ip++
	}
	return regs[0]
}

type inst struct {
	op   opcode
	x, y int
}

type opcode int

const (
	opCpyIR opcode = iota
	opCpyRR
	opInc
	opDec
	opJnzI
	opJnzR
)

var opArgs = map[string]int{
	"cpy": 2,
	"inc": 1,
	"dec": 1,
	"jnz": 2,
}

func parseInst(line string) (inst, error) {
	parts := strings.Split(line, " ")
	if wantArgs, ok := opArgs[parts[0]]; !ok {
		return inst{}, fmt.Errorf("invalid opcode: %q", parts[0])
	} else if len(parts) != 1+wantArgs {
		return inst{}, fmt.Errorf("unexpected number of arguments: %q, want %d", line, wantArgs)
	}
	switch parts[0] {
	case "cpy":
		y, err := parseReg(parts[2])
		if err != nil {
			return inst{}, err
		}
		if xr, err := parseReg(parts[1]); err == nil {
			return inst{op: opCpyRR, x: xr, y: y}, nil
		} else if xi, err := strconv.Atoi(parts[1]); err == nil {
			return inst{op: opCpyIR, x: xi, y: y}, nil
		}
		return inst{}, fmt.Errorf("not a register or an integer: %q", parts[1])
	case "inc", "dec":
		x, err := parseReg(parts[1])
		if err != nil {
			return inst{}, err
		}
		return inst{op: fn.If(parts[0] == "inc", opInc, opDec), x: x}, nil
	case "jnz":
		y, err := strconv.Atoi(parts[2])
		if err != nil {
			return inst{}, err
		}
		if xr, err := parseReg(parts[1]); err == nil {
			return inst{op: opJnzR, x: xr, y: y}, nil
		} else if xi, err := strconv.Atoi(parts[1]); err == nil {
			return inst{op: opJnzI, x: xi, y: y}, nil
		}
		return inst{}, fmt.Errorf("not a register or an integer: %q", parts[1])
	}
	return inst{}, fmt.Errorf("impossible")
}

func parseReg(name string) (int, error) {
	if len(name) != 1 {
		return 0, fmt.Errorf("not a register name: %q", name)
	}
	r := int(name[0] - 'a')
	if r < 0 || r >= 4 {
		return 0, fmt.Errorf("not a register name: %q", name)
	}
	return r, nil
}
