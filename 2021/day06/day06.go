// Package day06 solves AoC 2021 day 6.
package day06

import "github.com/fis/aoc/glue"

func init() {
	glue.RegisterSolver(2021, 6, glue.IntSolver(solve))
}

func solve(input []int) ([]string, error) {
	p1 := countFish(input, 80)
	p2 := countFish(input, 256)
	return glue.Ints(p1, p2), nil
}

func countFish(initial []int, days int) (total int) {
	counts := [9]int{}
	for _, i := range initial {
		counts[i]++
	}
	offset := 0
	for range days {
		counts[(offset+7)%9] += counts[offset]
		offset = (offset + 1) % 9
	}
	for _, c := range counts {
		total += c
	}
	return total
}
