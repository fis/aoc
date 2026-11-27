// Package day19 solves AoC 2019 day 19.
package day19

import (
	"github.com/fis/aoc/2019/intcode"
	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2019, 19, intcode.Solver(solve))
}

func solve(prog []int64) ([]int64, error) {
	probe := prober(prog)
	return []int64{
		part1(50, probe),
		part2(50, 100, probe),
	}, nil
}

func part1(size int, probe func(x, y int) bool) int64 {
	var count int64
	for y := range size {
		for x := range size {
			c := " "
			if probe(x, y) {
				count++
				c = "#"
			}
			util.Diag(c)
		}
		util.Diag("\n")
	}
	return count
}

func part2(start, size int, probe func(x, y int) bool) int64 {
	left := 0
	for !probe(left, start) {
		left++
	}
	right := left
	for probe(right+1, start) {
		right++
	}

	history := make([]beam, size)
	history[start%size] = beam{left, right}

	for y := start + 1; ; /* ever */ y++ {
		for !probe(left, y) {
			left++
		}
		for probe(right+1, y) {
			right++
		}
		history[y%size] = beam{left, right}
		if y < start+size {
			continue
		}
		prev := history[(y-size+1)%size]
		if left >= prev.left && left+size-1 <= prev.right {
			bx, by := left, y-size+1
			return int64(10000*bx + by)
		}
	}
}

func prober(prog []int64) func(x, y int) bool {
	return func(x, y int) bool {
		out, _ := intcode.Run(prog, []int64{int64(x), int64(y)})
		return out[0] != 0
	}
}

type beam struct {
	left  int
	right int
}
