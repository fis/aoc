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

// Package day19 solves AoC 2018 day 19.
package day19

import (
	"github.com/fis/aoc-go/2018/cpu"
	"github.com/fis/aoc-go/glue"
)

func init() {
	glue.RegisterSolver(2018, 19, glue.LineSolver(solve))
}

func solve(lines []string) ([]int, error) {
	prog, err := cpu.ParseProg(lines)
	if err != nil {
		return nil, err
	}

	s := cpu.State{}
	s.Run(prog)
	part1 := s.R[0]

	// For part 2, see the annotated assembly code below, as actually executing it would take too long.
	part2 := sumDiv(10551339)

	return []int{part1, part2}, nil
}

func sumDiv(n int) (sum int) {
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			sum += i
			sum += n / i
		}
	}
	return sum
}

/*

Annotated input program:

#ip 4
	00 addi 4 16 4  ; jmp .+16+1 -> init

part1:
	01 seti 1 3 3   ; r3 = 1
.oloop
	02 seti 1 4 2   ; r2 = 1
.iloop:
	03 mulr 3 2 1   ; r1 = r2*r3
	04 eqrr 1 5 1   ; r1 = r1==r5
	05 addr 1 4 4   ; jmp +r1+1 -> .pne or .peq
.pne:
	06 addi 4 1 4   ; jmp +1+1 -> .pne2
.peq:
	07 addr 3 0 0   ; r0 += r3
.pne:
	08 addi 2 1 2   ; r2++
	09 gtrr 2 5 1   ; r1 = r2>r5
	10 addr 4 1 4   ; jmp +r1+1 -> .r2le or .r2gt
.r2le:
	11 seti 2 2 4   ; jmp 2+1 -> .iloop
.r2gt:
	12 addi 3 1 3   ; r3++
	13 gtrr 3 5 1   ; r1 = r3>r5
	14 addr 1 4 4   ; jmp +r1+1 -> r3le or .r3gt
.r3le:
	15 seti 1 6 4   ; jmp 1+1 -> .oloop
.r3gt:
	16 mulr 4 4 4   ; halt

	;; part 1 pseudocode (assuming r5 == 939):
	;;
	;; r3 = 1
	;; do {
	;;   r2 = 1
	;;   do {
	;;     if (r2*r3 == 939) r0 += r3
	;;   } while (++r2 <= 939)
	;; } while (++r3 <= 939)
	;;
	;; IOW, sum all the divisors of 939: 1 + 3 + 313 + 939 = 1256

init:                            [? 0 ? ? * 0]
	17 addi 5 2 5   ; r5 += 2      -> r5:2
	18 mulr 5 5 5   ; r5 *= r5     -> r5:4
	19 mulr 4 5 5   ; r5 *= r4/IP  -> r5:4*19=76
	20 muli 5 11 5  ; r5 *= 11     -> r5:836
	21 addi 1 4 1   ; r1 += 4      -> r1:4
	22 mulr 1 4 1   ; r1 *= r4/IP  -> r1:4*22=88
	23 addi 1 15 1  ; r1 += 15     -> r1:103
	24 addr 5 1 5   ; r5 += r1     -> r5:939
	25 addr 4 0 4   ; jmp +r0+1 -> 26 for part 1, 27 for part 2
	26 seti 0 9 4   ; jmp 0+1 -> part1

part2:                           [? ? ? ? * 939]
	27 setr 4 2 1   ; r1 = 27      -> r1:27
	28 mulr 1 4 1   ; r1 *= 28     -> r1:756
	29 addr 4 1 1   ; r1 += 29     -> r1:785
	30 mulr 4 1 1   ; r1 *= 30     -> r1:23550
	31 muli 1 14 1  ; r1 *= 14     -> r1:329700
	32 mulr 1 4 1   ; r1 *= 32     -> r1:10550400
	33 addr 5 1 5   ; r5 += r1     -> r1:10551339
	34 seti 0 8 0   ; r0 = 0       -> r0:0
	35 seti 0 4 4   ; jmp 0+1 -> part1

*/
