// Copyright 2021 Google LLC
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

// Package day16 solves AoC 2017 day 16.
package day16

import (
	"fmt"
	"strings"

	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2017, 16, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	if len(lines) != 1 {
		return nil, fmt.Errorf("expected 1 line of input, got %d", len(lines))
	}
	dance, err := parseMoves(strings.Split(lines[0], ","))
	if err != nil {
		return nil, err
	}
	p1 := perform(16, dance)
	p2 := encore(16, dance, 1000000000)
	return []string{p1, p2}, nil
}

type moveType int

const (
	moveSpin moveType = iota
	moveExchange
	movePartner
)

type danceMove struct {
	typ    moveType
	a1, a2 int
}

func perform(N int, moves []danceMove) string {
	line, progPos := make([]byte, N), make([]int, N)
	for i := 0; i < N; i++ {
		line[i] = byte('a' + i)
		progPos[i] = i
	}

	start := 0
	for _, move := range moves {
		switch move.typ {
		case moveSpin:
			start = (start + N - move.a1) % N
		case moveExchange:
			pA, pB := (start+move.a1)%N, (start+move.a2)%N
			line[pA], line[pB] = line[pB], line[pA]
			progPos[line[pA]-'a'], progPos[line[pB]-'a'] = pA, pB
		case movePartner:
			pA, pB := progPos[move.a1], progPos[move.a2]
			line[pA], line[pB] = line[pB], line[pA]
			progPos[move.a1], progPos[move.a2] = pB, pA
		}
	}

	if start == 0 {
		return string(line)
	} else {
		return string(line[start:]) + string(line[:start])
	}
}

func encore(N int, moves []danceMove, times int) string {
	line1, line2 := make([]byte, N), make([]byte, N)
	perm1, perm2 := make([]byte, N), make([]byte, N)
	renamed := make([]byte, N)
	for i := 0; i < N; i++ {
		line1[i] = byte('a' + i)
		perm1[i] = byte(i)
		renamed[i] = byte(i)
	}

	poff := 0
	for _, move := range moves {
		switch move.typ {
		case moveSpin:
			poff = (poff + N - move.a1) % N
		case moveExchange:
			pA, pB := (poff+move.a1)%N, (poff+move.a2)%N
			perm1[pA], perm1[pB] = perm1[pB], perm1[pA]
		case movePartner:
			nA, nB := renamed[move.a1], renamed[move.a2]
			renamed[nA], renamed[nB] = renamed[nB], renamed[nA]
		}
	}
	if poff != 0 {
		copy(perm2[:N-poff], perm1[poff:])
		copy(perm2[N-poff:], perm1[:poff])
		perm1, perm2 = perm2, perm1
	}

	applyRenames := (times & 1) == 1
	for times > 0 {
		if (times & 1) == 1 {
			for i := 0; i < N; i++ {
				line2[i] = line1[perm1[i]]
			}
			line1, line2 = line2, line1
		}
		for i := 0; i < N; i++ {
			perm2[i] = perm1[perm1[i]]
		}
		perm1, perm2 = perm2, perm1
		times >>= 1
	}
	if applyRenames {
		for i := 0; i < N; i++ {
			line2[i] = byte('a' + renamed[line1[i]-'a'])
		}
		line1 = line2
	}

	return string(line1)
}

func parseMoves(specs []string) (moves []danceMove, err error) {
	moves = make([]danceMove, len(specs))
	for i, spec := range specs {
		moves[i], err = parseMove(spec)
		if err != nil {
			return nil, err
		}
	}
	return moves, nil
}

func parseMove(spec string) (danceMove, error) {
	var (
		a1, a2 int
		r1, r2 rune
	)
	if _, err := fmt.Sscanf(spec, "s%d", &a1); err == nil {
		return danceMove{typ: moveSpin, a1: a1}, nil
	}
	if _, err := fmt.Sscanf(spec, "x%d/%d", &a1, &a2); err == nil {
		return danceMove{typ: moveExchange, a1: a1, a2: a2}, nil
	}
	if _, err := fmt.Sscanf(spec, "p%c/%c", &r1, &r2); err == nil {
		return danceMove{typ: movePartner, a1: int(r1 - 'a'), a2: int(r2 - 'a')}, nil
	}
	return danceMove{}, fmt.Errorf("invalid move: %q", spec)
}
