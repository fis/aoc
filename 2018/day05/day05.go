// Package day05 solves AoC 2018 day 5.
package day05

import (
	"fmt"

	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2018, 5, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	if len(lines) != 1 {
		return nil, fmt.Errorf("expecting only 1 line of input, got %d", len(lines))
	}
	p1 := reducedLen(lines[0])
	p2 := improve(lines[0])
	return glue.Ints(p1, p2), nil
}

func reducedLen(polymer string) int {
	buf := reduce([]byte(polymer))
	return len(buf)
}

func reduce(polymer []byte) []byte {
	for {
		var i, o int
		for i < len(polymer) {
			if i+1 < len(polymer) && polymer[i] == unitPairs[polymer[i+1]] {
				i += 2
			} else {
				if o < i {
					polymer[o] = polymer[i]
				}
				i++
				o++
			}
		}
		if o < i {
			polymer = polymer[:o]
		} else {
			return polymer
		}
	}
}

func improve(polymer string) int {
	min := len(polymer)
	for i := range byte(26) {
		buf := reduce(prune([]byte(polymer), 'a'+i, 'A'+i))
		if len(buf) < min {
			min = len(buf)
		}
	}
	return min
}

func prune(polymer []byte, u1, u2 byte) []byte {
	var i, o int
	for i < len(polymer) {
		if polymer[i] != u1 && polymer[i] != u2 {
			if o < i {
				polymer[o] = polymer[i]
			}
			o++
		}
		i++
	}
	return polymer[:o]
}

var unitPairs [256]byte

func init() {
	for i := range byte(26) {
		unitPairs['a'+i] = 'A' + i
		unitPairs['A'+i] = 'a' + i
	}
}
