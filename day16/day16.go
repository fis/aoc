// Package day16 solves AoC 2019 day 16.
package day16

import (
	"github.com/fis/aoc2019-go/util"
)

func Solve(path string) ([]string, error) {
	lines, err := util.ReadLines(path)
	if err != nil {
		return nil, err
	}

	sig := digits(lines[0])
	fft(sig, 100)
	p1 := undigits(sig[:8])

	sig = digits(lines[0])
	p2 := undigits(rfft(sig, 100, 10000))

	return []string{p1, p2}, nil
}

func fft(sig []int, phases int) {
	work := make([]int, len(sig))
	for p := 0; p < phases; p++ {
		for n := range work {
			r := 0
			for i := n; i < len(sig); {
				for j := 0; j <= n && i < len(sig); i, j = i+1, j+1 {
					r += sig[i]
				}
				i += n + 1
				for j := 0; j <= n && i < len(sig); i, j = i+1, j+1 {
					r -= sig[i]
				}
				i += n + 1
			}
			work[n] = abs(r % 10)
		}
		copy(sig, work)
	}
}

func rfft(sig []int, phases, reps int) []int {
	off := offset(sig[:7])
	if off < reps*len(sig)/2 {
		panic("rfft offset error")
	}
	// TODO: think of a less computationally intensive way to do this
	work := make([]int, reps*len(sig)-off)
	for i := range work {
		work[i] = sig[(off+i)%len(sig)]
	}
	for p := 0; p < phases; p++ {
		c := 0
		for n := len(work) - 1; n >= 0; n-- {
			c = (c + work[n]) % 10
			work[n] = c
		}
	}
	return work[:8]
}

func digits(str string) []int {
	sig := make([]int, len(str))
	for i, r := range str {
		sig[i] = int(r - '0')
	}
	return sig
}

func undigits(sig []int) string {
	bytes := make([]byte, len(sig))
	for i, v := range sig {
		bytes[i] = byte('0' + abs(v%10))
	}
	return string(bytes)
}

func offset(digits []int) int {
	off := 0
	for _, d := range digits {
		off = 10*off + d
	}
	return off
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
