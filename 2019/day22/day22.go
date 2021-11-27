// Copyright 2019 Google LLC
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

// Package day22 solves AoC 2019 day 22.
package day22

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2019, 22, glue.LineSolver(solve))
}

func solve(lines []string) ([]int, error) {
	ops, err := parseShuffle(lines)
	if err != nil {
		return nil, err
	}
	p1 := shuffleForward(2019, 10007, ops)
	p2 := shuffleBackward(2020, 119315717514047, 101741582076661, ops)
	return []int{int(p1), int(p2)}, nil
}

type shuffleTrick int

const (
	shuffleDeal shuffleTrick = iota
	shuffleCut
	shuffleInterleave
)

type shuffleOp struct {
	trick shuffleTrick
	val   int64
}

func parseShuffle(lines []string) ([]shuffleOp, error) {
	ops := make([]shuffleOp, len(lines))
	for i, line := range lines {
		switch {
		case line == "deal into new stack":
			ops[i].trick = shuffleDeal
		case strings.HasPrefix(line, "cut "):
			val, err := strconv.ParseInt(line[4:], 10, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid number in line %q: %w", line, err)
			}
			ops[i].trick, ops[i].val = shuffleCut, val
		case strings.HasPrefix(line, "deal with increment "):
			val, err := strconv.ParseInt(line[20:], 10, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid number in line %q: %w", line, err)
			}
			ops[i].trick, ops[i].val = shuffleInterleave, val
		}
	}
	return ops, nil
}

func shuffleForward(pos, deckSize int64, ops []shuffleOp) int64 {
	for _, op := range ops {
		switch op.trick {
		case shuffleDeal:
			pos = deckSize - 1 - pos
		case shuffleCut:
			pos = (pos + deckSize - op.val) % deckSize
		case shuffleInterleave:
			pos = modmul(pos, op.val, deckSize)
		}
	}
	return pos
}

func shuffleBackward(pos, deckSize, reps int64, ops []shuffleOp) int64 {
	smul, soff := int64(1), int64(0)
	for i := len(ops) - 1; i >= 0; i-- {
		op := ops[i]
		switch op.trick {
		case shuffleDeal:
			smul, soff = deckSize-smul, deckSize-1-soff
		case shuffleCut:
			soff = (soff + deckSize + op.val) % deckSize
		case shuffleInterleave:
			inv := modinv(op.val, deckSize)
			smul, soff = modmul(smul, inv, deckSize), modmul(soff, inv, deckSize)
		}
	}
	mul, off := int64(1), int64(0)
	for ; reps > 0; reps >>= 1 {
		if reps&1 != 0 {
			mul, off = modmul(mul, smul, deckSize), (modmul(off, smul, deckSize)+soff)%deckSize
		}
		smul, soff = modmul(smul, smul, deckSize), modmul(soff, smul+1, deckSize)
	}
	return (modmul(pos, mul, deckSize) + off) % deckSize
}

func modmul(a, b, m int64) int64 {
	sum, sq := int64(0), b
	for ; a > 0; a >>= 1 {
		if a&1 != 0 {
			sum = (sum + sq) % m
		}
		sq = (sq + sq) % m
	}
	return sum
}

func modinv(a, m int64) int64 {
	g, x, _ := egcd(a, m)
	if g != 1 {
		return 0
	}
	return (x + m) % m
}

func egcd(a, b int64) (g, x, y int64) {
	if a == 0 {
		return b, 0, 1
	}
	g, y, x = egcd(b%a, a)
	return g, x - b/a*y, y
}
