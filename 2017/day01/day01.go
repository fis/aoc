// Package day01 solves AoC 2017 day 1.
package day01

import (
	"fmt"

	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2017, 1, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	if len(lines) != 1 {
		return nil, fmt.Errorf("expecting only 1 line of input, got %d", len(lines))
	}
	digits := []byte(lines[0])
	p1 := checksum(digits, 1)
	p2 := checksum(digits, len(digits)/2)
	return glue.Ints(p1, p2), nil
}

func checksum(digits []byte, offset int) (sum int) {
	for i, d := range digits {
		if d == digits[(i+offset)%len(digits)] {
			sum += int(d - '0')
		}
	}
	return sum
}
