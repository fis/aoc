// Package day10 solves AoC 2021 day 10.
package day10

import (
	"fmt"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2021, 10, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	var (
		p1 int
		p2 []int
	)
	for _, line := range lines {
		if corrupted, score := check(line); corrupted {
			p1 += score
		} else {
			p2 = append(p2, score)
		}
	}
	if len(p2)%2 == 0 {
		return nil, fmt.Errorf("not an odd number of incomplete lines; got %d", len(p2))
	}
	return glue.Ints(p1, util.QuickSelect(p2, len(p2)/2)), nil
}

var delimPairs = map[byte]byte{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

var completionScores = map[byte]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

var errorScores = map[byte]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

func check(line string) (corrupted bool, score int) {
	var stack []byte
	for i := range len(line) {
		c := line[i]
		if pair, ok := delimPairs[c]; ok {
			stack = append(stack, pair)
			continue
		}
		if len(stack) == 0 {
			return true, 0
		}
		if c != stack[len(stack)-1] {
			return true, errorScores[c]
		}
		stack = stack[:len(stack)-1]
	}
	for i := len(stack) - 1; i >= 0; i-- {
		score *= 5
		score += completionScores[stack[i]]
	}
	return false, score
}
