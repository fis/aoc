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

// Package cpu provides an emulator for the unnamed time travel control computer/wearable from AoC 2018.
package cpu

import (
	"fmt"
)

// Op represents one of the device's 16 opcodes.
type Op int

const (
	// AddR (add register), rC = rA + rB.
	AddR Op = iota
	// AddI (add immediate), rC = rA + #B.
	AddI
	// MulR (multiply register), rC = rA * rB.
	MulR
	// MulI (multiply immediate), rC = rA * #B.
	MulI
	// BanR (bitwise AND register), rC = rA & rB.
	BanR
	// BanI (bitwise AND immediate), rC = rA & #B.
	BanI
	// BorR (bitwise OR register), rC = rA | rB.
	BorR
	// BorI (bitwise OR immediate), rC = rA | #B.
	BorI
	// SetR (set register), rC = rA, B ignored.
	SetR
	// SetI (set immediate), rC = #A, B ignored.
	SetI
	// GtIR (greater-than immediate/register), rC = #A > rB.
	GtIR
	// GtRI (greater-than register/immediate), rC = rA > #B.
	GtRI
	// GtRR (greater-than register/register), rC = rA > rB.
	GtRR
	// EqIR (equal immediate/register), rC = #A == rB.
	EqIR
	// EqRI (equal register/immediate), rC = rA == #B.
	EqRI
	// EqRR (equal register/register), rC = rA == rB.
	EqRR
	// FirstOp is the numerically first opcode; you can iterate from it to LastOp.
	FirstOp = AddR
	// LastOp is the numerically last opcode; you can iterate to it from FirstOp.
	LastOp = EqRR
)

var (
	opToName = [...]string{
		AddR: "addr", AddI: "addi", MulR: "mulr", MulI: "muli", BanR: "banr", BanI: "bani", BorR: "borr", BorI: "bori",
		SetR: "setr", SetI: "seti", GtIR: "gtir", GtRI: "gtri", GtRR: "gtrr", EqIR: "eqir", EqRI: "eqri", EqRR: "eqrr",
	}
	nameToOp map[string]Op
)

func init() {
	nameToOp = make(map[string]Op)
	for op := FirstOp; op <= LastOp; op++ {
		nameToOp[opToName[op]] = op
	}
}

func (op Op) String() string {
	return opToName[op]
}

// OpNamed returns the opcode with the given mnemonic, or false as the second argument if the mnemonic is not valid.
func OpNamed(s string) (op Op, ok bool) {
	if op, ok = nameToOp[s]; ok {
		return op, ok
	}
	return 0, false
}

// Inst represents a single instruction: an opcode, and the A, B and C operands.
type Inst struct {
	Op      Op
	A, B, C int
}

// ParseInst converts the textual format of an instruction (e.g., "seti 5 0 1") to an instruction.
func ParseInst(s string) (i Inst, ok bool) {
	var op string
	if _, err := fmt.Sscanf(s, "%s %d %d %d", &op, &i.A, &i.B, &i.C); err != nil {
		return Inst{}, false
	} else if i.Op, ok = OpNamed(op); !ok {
		return Inst{}, false
	}
	return i, true
}

// Prog represents an entire program: a sequence of instructions, together with IP binding instructions.
type Prog struct {
	Code    []Inst
	IPBound bool
	IPR     int
}

// ParseProg reads the text of a program (one symbolic instruction per line, plus an optional IP binding header).
func ParseProg(lines []string) (p Prog, err error) {
	if len(lines) == 0 {
		return Prog{}, nil
	}
	if _, err := fmt.Sscanf(lines[0], "#ip %d", &p.IPR); err == nil {
		p.IPBound = true
		lines = lines[1:]
	}
	p.Code = make([]Inst, len(lines))
	for i, line := range lines {
		var ok bool
		p.Code[i], ok = ParseInst(line)
		if !ok {
			return Prog{}, fmt.Errorf("invalid instruction: %s", line)
		}
	}
	return p, nil
}

// State holds the entire CPU state: 6 registers and the instruction pointer,
// as well as the current IP/register binding state.
type State struct {
	R       [6]int
	IP      int
	IPBound bool
	IPR     int
}

// Run executes an entire program until it halts. The IP binding state is reset to be that of the program,
// but the IP (and other registers) are itself not reset to zero.
func (s *State) Run(p Prog) {
	s.IPBound, s.IPR = p.IPBound, p.IPR
	for s.IP >= 0 && s.IP < len(p.Code) {
		i := p.Code[s.IP]
		//fmt.Printf("ip=%d %v %v %d %d %d", s.IP, s.R, i.Op, i.A, i.B, i.C)
		s.Step(i.Op, i.A, i.B, i.C)
		//fmt.Printf(" %v\n", s.R)
	}
}

// Step executes one CPU cycle, given the operation to execute.
func (s *State) Step(op Op, a, b, c int) {
	if s.IPBound {
		s.R[s.IPR] = s.IP
	}
	switch op {
	case AddR:
		s.R[c] = s.R[a] + s.R[b]
	case AddI:
		s.R[c] = s.R[a] + b
	case MulR:
		s.R[c] = s.R[a] * s.R[b]
	case MulI:
		s.R[c] = s.R[a] * b
	case BanR:
		s.R[c] = s.R[a] & s.R[b]
	case BanI:
		s.R[c] = s.R[a] & b
	case BorR:
		s.R[c] = s.R[a] | s.R[b]
	case BorI:
		s.R[c] = s.R[a] | b
	case SetR:
		s.R[c] = s.R[a]
	case SetI:
		s.R[c] = a
	case GtIR:
		s.R[c] = asInt(a > s.R[b])
	case GtRI:
		s.R[c] = asInt(s.R[a] > b)
	case GtRR:
		s.R[c] = asInt(s.R[a] > s.R[b])
	case EqIR:
		s.R[c] = asInt(a == s.R[b])
	case EqRI:
		s.R[c] = asInt(s.R[a] == b)
	case EqRR:
		s.R[c] = asInt(s.R[a] == s.R[b])
	}
	if s.IPBound {
		s.IP = s.R[s.IPR]
	}
	s.IP++
}

func asInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
