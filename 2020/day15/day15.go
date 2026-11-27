// Package day15 solves AoC 2020 day 15.
package day15

import (
	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2020, 15, glue.IntSolver(solve))
}

func solve(numbers []int) ([]string, error) {
	p1 := simulate(numbers, 2020)
	p2 := simulate(numbers, 30000000)
	return glue.Ints(p1, p2), nil
}

func simulate(initial []int, upTo int) (num int) {
	lastTurn := make([]int, upTo)
	for turn, num := range initial[:len(initial)-1] {
		lastTurn[num] = turn + 1
	}
	num = initial[len(initial)-1]
	for turn := len(initial); turn < upTo; turn++ {
		var next int
		if last := lastTurn[num]; last == 0 {
			next = 0
		} else {
			next = turn - last
		}
		lastTurn[num], num = turn, next
	}
	return num
}
