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
	"io"
	"strconv"
	"strings"

	"github.com/fis/aoc-go/util"
)

func init() {
	util.RegisterSolver(8, util.LineSolver(solve))
}

func solve(lines []string) ([]int, error) {
	code, err := parseCode(lines)
	if err != nil {
		return nil, err
	}

	_, part1 := loopCheck(code)
	part2 := repair(code)

	return []int{part1, part2}, nil
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
	type branch struct{ to, acc int }
	var branches []branch

	seen := make([]bool, len(code))

	for at, acc := 0, 0; !seen[at]; {
		seen[at] = true
		switch code[at].op {
		case opAcc:
			acc += code[at].arg
			at++
		case opJmp:
			branches = append(branches, branch{to: at + 1, acc: acc})
			at += code[at].arg
		case opNop:
			branches = append(branches, branch{to: at + code[at].arg, acc: acc})
			at++
		}
	}

	for _, branch := range branches {
		at, acc := branch.to, branch.acc
		for {
			if at >= len(code) {
				return acc
			}
			if seen[at] {
				break
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
	}

	panic("this code is unfixable")
}

func PrintGraph(out io.Writer, lines []string) error {
	var mnemonics = map[opcode]string{opAcc: "acc", opJmp: "jmp", opNop: "nop"}

	code, err := parseCode(lines)
	if err != nil {
		return err
	}

	g := &util.Graph{}
	verts := make([]int, len(code)+1)
	for i, inst := range code {
		verts[i] = g.V(fmt.Sprintf("%d: %s %+d", i, mnemonics[inst.op], inst.arg))
	}
	verts[len(code)] = g.V("halt")

	for i, inst := range code {
		switch inst.op {
		case opAcc:
			g.AddEdgeWV(verts[i], verts[i+1], 0)
		case opJmp:
			g.AddEdgeWV(verts[i], verts[i+inst.arg], 0)
			g.AddEdgeWV(verts[i], verts[i+1], 1)
		case opNop:
			g.AddEdgeWV(verts[i], verts[i+1], 0)
			g.AddEdgeWV(verts[i], verts[i+inst.arg], 1)
		}
	}

	fmt.Fprint(out, "digraph prog {\n")
	g.RangeV(func(v int) {
		per := ""
		if v == verts[0] || v == verts[len(verts)-1] {
			per = ", peripheries=2"
		}
		fmt.Fprintf(out, "  n%d [label=\"%s\"%s];\n", v, g.Name(v), per)
	})
	g.RangeV(func(v int) {
		g.RangeSuccV(v, func(v2 int) bool {
			color := ""
			if g.W(v, v2) == 1 {
				color = ` [color="red"]`
			}
			fmt.Fprintf(out, "  n%d -> n%d%s;\n", v, v2, color)
			return true
		})
	})
	fmt.Fprint(out, "}\n")

	return nil
}
