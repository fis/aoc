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

// Package day21 solves AoC 2018 day 21.
package day21

import (
	"github.com/fis/aoc/2018/cpu"
	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2018, 21, glue.LineSolver(solve))
}

func solve(lines []string) ([]int, error) {
	prog, err := cpu.ParseProg(lines)
	if err != nil {
		return nil, err
	}
	p1 := part1(prog)
	p2 := part2()
	return []int{p1, p2}, nil
}

func part1(prog cpu.Prog) int {
	s := cpu.State{}
	s.IPBound, s.IPR = prog.IPBound, prog.IPR
	for s.IP != 28 {
		i := prog.Code[s.IP]
		s.Step(i.Op, i.A, i.B, i.C)
	}
	return s.R[3]
}

func part2() int {
	prev := iter(0)
	seen := map[int]struct{}{prev: {}}
	for {
		next := iter(prev)
		if _, ok := seen[next]; ok {
			return prev
		}
		seen[next] = struct{}{}
		prev = next
	}
}

func iter(r3 int) int {
	r4 := r3 | 0x10000
	r3 = 2176960
	for {
		r3 = (r3 + (r4 & 0xff)) & 0xffffff
		r3 = (r3 * 65899) & 0xffffff
		if r4 < 256 {
			break
		}
		r4 >>= 8
	}
	return r3
}

/*

Annotated input program:

#ip 2
check:
	00 seti 123 0 3        ; r3 = 123
.loop
	01 bani 3 456 3        ; r3 &= 456
	02 eqri 3 72 3         ; r3 = r3 == 72
	03 addr 3 2 2          ; jmp .+r3+1 -> .fail | main
.fail:
	04 seti 0 0 2          ; jmp 0+1 -> .loop
main:
	05 seti 0 6 3          ; r3 = 0
.L0:
	06 bori 3 65536 4      ; r4 = r3|0x10000
	07 seti 2176960 8 3    ; r3 = 2176960
.L1:
	08 bani 4 255 1        ; r1 = r4&0xff
	09 addr 3 1 3          ; r3 += r1
	10 bani 3 16777215 3   ; r3 &= 0xffffff
	11 muli 3 65899 3      ; r3 *= 65899
	12 bani 3 16777215 3   ; r3 &= 0xffffff
	13 gtir 256 4 1        ; r1 = 256>r4
	14 addr 1 2 2          ; jmp .+r1+1 -> .L2 | .L3
.L2:
	15 addi 2 1 2          ; jmp .+1+1 -> .L4
.L3:
	16 seti 27 7 2         ; jmp 27+1 -> .L10
.L4:
	17 seti 0 9 1          ; r1 = 0
.L5:
	18 addi 1 1 5          ; r5 = r1+1
	19 muli 5 256 5        ; r5 *= 256
	20 gtrr 5 4 5          ; r5 = r5>r4
	21 addr 5 2 2          ; jmp .+r5+1 -> .L6 | .L7
.L6:
	22 addi 2 1 2          ; jmp.+1+1 -> .L8
.L7:
	23 seti 25 7 2         ; jmp 25+1 -> .L9
.L8:
	24 addi 1 1 1          ; r1++
	25 seti 17 2 2         ; jmp 17+1 -> .L5
.L9:
	26 setr 1 7 4          ; r4 = r1
	27 seti 7 9 2          ; jmp 7+1 -> .L1
.L10:
	28 eqrr 3 0 1          ; r1 = r0==r3
	29 addr 1 2 2          ; halt if r1
	30 seti 5 9 2          ; jmp 5+1 -> .L0

Main loop, simplified by merging instructions and unravelling conditional jumps:

L0:
	r4 = r3 | 0x10000
	r3 = 2176960
L1:
	r3 = (r3 + (r4 & 0xff)) * 65899  // 24-bit addition and multiplication
	if (r4 < 256) goto L10
	r1 = 0
L5:
	if ((r1+1) * 256 > r4) goto L9
	r1++
	goto L5
L9:
	r4 = r1
	goto L1
L10:
	if (r0 == r3) halt
	goto L0

With jump instructions replaced by structured control flow:

	do {
		r4 = r3 | 0x10000
		r3 = 2176960
		for {
			r3 = (r3 + (r4 & 0xff)) * 65899  // 24-bit addition and multiplication
			if (r4 < 256) break
			for r1 = 0; (r1+1)*256 <= r4; r1++ {}
			r4 = r1
		}
	} while (r3 != r0)

And with the innermost loop replaced with the equivalent mathematical operation:

	do {
		r4 = r3 | 0x10000
		r3 = 2176960
		for {
			r3 = (r3 + (r4 & 0xff)) * 65899  // 24-bit addition and multiplication
			if (r4 < 256) break
			r4 >>= 8
		}
	} while (r3 != r0)

To minimize the number of instructions executed, we just need to set r0 whatever the value of r3 will be on the first run.
To maximize (while still halting), we'll need to set r0 to the last value it gets before hitting some previous number again.

*/
