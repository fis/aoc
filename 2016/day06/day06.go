// Package day06 solves AoC 2016 day 6.
package day06

import (
	"math"

	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2016, 6, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	p1, p2 := decode(lines)
	return []string{p1, p2}, nil
}

func decode(lines []string) (string, string) {
	N := len(lines[0])
	m1, m2 := make([]byte, N), make([]byte, N)
	for i := range N {
		var freqs [26]int
		for _, line := range lines {
			if b := line[i]; b >= 'a' && b <= 'z' {
				freqs[b-'a']++
			}
		}
		maxB, maxF, minB, minF := byte(0), math.MinInt, byte(0), math.MaxInt
		for b := range byte(26) {
			f := freqs[b]
			if f == 0 {
				continue
			}
			if f > maxF {
				maxB, maxF = b, f
			}
			if f < minF {
				minB, minF = b, f
			}
		}
		m1[i] = 'a' + maxB
		m2[i] = 'a' + minB
	}
	return string(m1), string(m2)
}
