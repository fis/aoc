// Package day06 solves AoC 2017 day 6.
package day06

import "github.com/fis/aoc/glue"

func init() {
	glue.RegisterSolver(2017, 6, glue.IntSolver(solve))
}

func solve(input []int) ([]string, error) {
	banks := make([]byte, len(input))
	for i, n := range input {
		banks[i] = byte(n)
	}
	p1 := loopLength(banks)
	p2 := loopLength(banks)
	return glue.Ints(p1, p2), nil
}

func loopLength(banks []byte) int {
	seen := map[string]struct{}{string(banks): {}}
	steps := 0
	for {
		balance(banks)
		steps++
		state := string(banks)
		if _, found := seen[state]; found {
			return steps
		}
		seen[state] = struct{}{}
	}
}

func balance(banks []byte) {
	N := byte(len(banks))
	src, blocks := argmax(banks)
	banks[src] = 0
	if blocks >= N {
		all := blocks / N
		blocks = blocks % N
		for i := range banks {
			banks[i] += all
		}
	}
	for i := byte(1); i <= blocks; i++ {
		banks[(byte(src)+i)%N]++
	}
}

func argmax(xs []byte) (maxI int, maxX byte) {
	maxI, maxX = -1, 0
	for i, x := range xs {
		if maxI == -1 || x > maxX {
			maxI, maxX = i, x
		}
	}
	return maxI, maxX
}
