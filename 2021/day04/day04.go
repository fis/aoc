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

// Package day04 solves AoC 2021 day 4.
package day04

import (
	"fmt"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2021, 4, glue.ChunkSolver(solve))
}

func solve(chunks []string) ([]string, error) {
	seq, boards, err := parseInput(chunks)
	if err != nil {
		return nil, err
	}
	p1, p2 := bingo(seq, boards)
	return glue.Ints(p1, p2), nil
}

func bingo(seq []int, boards []board) (firstWin, lastWin int) {
	won := make([]bool, len(boards))
	for _, n := range seq {
		allWon := true
		for i := range boards {
			if boards[i].mark(n) && !won[i] {
				won[i] = true
				score := n * boards[i].sum()
				if firstWin == 0 {
					firstWin = score
				}
				lastWin = score
			}
			allWon = allWon && won[i]
		}
		if allWon {
			break
		}
	}
	return firstWin, lastWin
}

const size = 5

type board struct {
	totalSum  int
	positions [][2]int
	marks     [2][size]int
}

func (b *board) initialize(numbers *[size * size]int) {
	for y, i := 0, 0; y < size; y++ {
		for x := 0; x < size; x, i = x+1, i+1 {
			n := numbers[i]
			b.totalSum += n
			for n >= len(b.positions) {
				b.positions = append(b.positions, [2]int{-1, -1})
			}
			b.positions[n] = [2]int{x, y}
		}
	}
}

func (b *board) mark(n int) (wins bool) {
	if n >= len(b.positions) {
		return false
	}
	p := b.positions[n]
	if p[0] < 0 {
		return false
	}
	b.totalSum -= n
	b.marks[0][p[0]]++
	b.marks[1][p[1]]++
	return b.marks[0][p[0]] >= size || b.marks[1][p[1]] >= size
}

func (b *board) sum() (s int) {
	return b.totalSum
}

func parseInput(chunks []string) (seq []int, boards []board, err error) {
	if len(chunks) < 2 {
		return nil, nil, fmt.Errorf("expected at least 2 chunks, got %d", len(chunks))
	}

	seq = util.Ints(chunks[0])
	chunks = chunks[1:]

	boards = make([]board, len(chunks))
	for i, chunk := range chunks {
		numbers := util.Ints(chunk)
		if len(numbers) != size*size {
			return nil, nil, fmt.Errorf("board %d has wrong size: got %d, want %d", i, len(numbers), size*size)
		}
		boards[i].initialize((*[size * size]int)(numbers))
	}

	return seq, boards, nil
}
