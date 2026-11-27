// Package day10 solves AoC 2020 day 10.
package day10

import (
	"sort"

	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2020, 10, glue.IntSolver(solve))
}

func solve(adapters []int) ([]string, error) {
	d1, d3 := deltas(adapters)
	arr := arrangements(adapters)
	return glue.Ints(d1*d3, arr), nil
}

func chain(adapters []int) []int {
	chain := append([]int{0}, adapters...)
	sort.Ints(chain)
	return append(chain, chain[len(chain)-1]+3)
}

func deltas(adapters []int) (d1, d3 int) {
	ch := chain(adapters)
	for i, l := range ch[:len(ch)-1] {
		r := ch[i+1]
		if r-l == 1 {
			d1++
		} else if r-l == 3 {
			d3++
		}
	}

	return d1, d3
}

func arrangements(adapters []int) int {
	ch := chain(adapters)
	ways := make([]int, len(ch))
	ways[0] = 1
	for at := 1; at < len(ways); at++ {
		n := ways[at-1]
		if at-2 >= 0 && ch[at]-ch[at-2] <= 3 {
			n += ways[at-2]
		}
		if at-3 >= 0 && ch[at]-ch[at-3] <= 3 {
			n += ways[at-3]
		}
		ways[at] = n
	}
	return ways[len(ways)-1]
}
