// Package day19 solves AoC 2019 day 19.
package day19

import (
	"strconv"

	"github.com/fis/aoc2019-go/intcode"
	"github.com/fis/aoc2019-go/util"
)

const N = 50
const M = 100

func Solve(path string) ([]string, error) {
	prog, err := intcode.Load(path)
	if err != nil {
		return nil, err
	}
	probe := prober(prog)
	return []string{
		strconv.Itoa(part1(50, probe)),
		strconv.Itoa(part2(50, 100, probe)),
	}, nil
}

func part1(size int, probe func(x, y int) bool) int {
	count := 0
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
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

func part2(start, size int, probe func(x, y int) bool) int {
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
			return 10000*bx + by
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
