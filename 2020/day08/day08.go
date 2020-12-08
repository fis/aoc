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

// Package day08 solves AoC 2020 day 8.
package day08

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fis/aoc-go/util"
)

func Solve(path string) ([]string, error) {
	lines, err := util.ReadLines(path)
	if err != nil {
		return nil, err
	}
	code, err := parseCode(lines)
	if err != nil {
		return nil, err
	}

	_, part1 := loopCheck(code)
	part2 := repair(code)

	return []string{strconv.Itoa(part1), strconv.Itoa(part2)}, nil
}

type opcode int

const (
	opAcc opcode = iota
	opJmp
	opNop
)

type instruction struct {
	op  opcode
	arg int
}

func parseCode(lines []string) (out []instruction, err error) {
	var mnemonics = map[string]opcode{"acc": opAcc, "jmp": opJmp, "nop": opNop}

	for _, line := range lines {
		parts := strings.SplitN(line, " ", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid instruction: %s", line)
		}
		op, ok := mnemonics[parts[0]]
		if !ok {
			return nil, fmt.Errorf("invalid opcode: %s", parts[0])
		}
		arg, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("invalid argument: %s", parts[1])
		}
		out = append(out, instruction{op: op, arg: arg})
	}

	return out, nil
}

func loopCheck(code []instruction) (loop bool, acc int) {
	seen := make([]bool, len(code))
	for at := 0; at < len(code); {
		if seen[at] {
			return true, acc
		}
		seen[at] = true
		switch code[at].op {
		case opAcc:
			acc += code[at].arg
			at++
		case opJmp:
			at += code[at].arg
		case opNop:
			at++
		}
	}
	return false, acc
}

func repair(code []instruction) int {
	var flip = [...]opcode{opJmp: opNop, opNop: opJmp}
	for at := range code {
		if code[at].op == opAcc {
			continue
		}
		code[at].op = flip[code[at].op]
		if loop, acc := loopCheck(code); !loop {
			return acc
		}
		code[at].op = flip[code[at].op]
	}
	panic("this code is unfixable")
}
