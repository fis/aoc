// Package day22 solves AoC 2024 day 22.
package day22

import (
	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/fn"
)

func init() {
	glue.RegisterSolver(2024, 22, glue.IntSolver(solve))
}

func solve(secrets []int) ([]string, error) {
	part1, part2 := analyze(secrets)
	return glue.Ints(part1, part2), nil
}

const (
	numSteps = 2000
	numSeq   = 19 * 19 * 19 * 19
)

func analyze(secrets []int) (part1, part2 int) {
	seen := util.MakeFixedBitmap1D(numSeq)
	bananas := make([]uint32, numSeq)
	for _, seed := range secrets {
		seen.Clear()
		n := uint32(seed)
		a := n % 10
		n = step(n)
		b := n % 10
		n = step(n)
		c := n % 10
		n = step(n)
		d := n % 10
		for range numSteps - 3 {
			n = step(n)
			e := n % 10
			seq := 9 + b - a
			seq = 19*seq + 9 + c - b
			seq = 19*seq + 9 + d - c
			seq = 19*seq + 9 + e - d
			if !seen.Get(int(seq)) {
				seen.Set(int(seq))
				bananas[seq] += e
			}
			a, b, c, d = b, c, d, e
		}
		part1 += int(n)
	}
	part2 = int(fn.Max(bananas))
	return part1, part2
}

func stepN(n, t int) int {
	v := uint32(n)
	for range t {
		v = step(v)
	}
	return int(v)
}

func step(v uint32) uint32 {
	v = (v ^ (v << 6)) & 0xffffff
	v = (v ^ (v >> 5)) & 0xffffff
	v = (v ^ (v << 11)) & 0xffffff
	return v
}
