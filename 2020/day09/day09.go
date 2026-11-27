// Package day09 solves AoC 2020 day 9.
package day09

import (
	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2020, 9, glue.IntSolver(solve))
}

func solve(data []int) ([]string, error) {
	invalid := validate(data, 25)
	min, max := findSum(data, invalid)
	return glue.Ints(invalid, min+max), nil
}

func validate(data []int, win int) int {
	sums := make(map[int]int)
	for i, x := range data[:win-1] {
		for _, y := range data[i+1 : win] {
			sums[x+y] = sums[x+y] + 1
		}
	}

	for i := win; i < len(data); i++ {
		n := data[i]
		if sums[n] == 0 {
			return n
		}
		o := data[i-win]
		for _, x := range data[i-win+1 : i] {
			sums[o+x] = sums[o+x] - 1
			sums[n+x] = sums[n+x] + 1
		}
	}

	return -1
}

func findSum(data []int, key int) (min, max int) {
	l, r, sum := 0, 0, data[0]
	for {
		if sum < key {
			r++
			sum += data[r]
		} else if sum > key {
			sum -= data[l]
			l++
		} else {
			return minMax(data[l : r+1])
		}
	}
}

func minMax(data []int) (min, max int) {
	min, max = data[0], data[0]
	for _, n := range data[1:] {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min, max
}
