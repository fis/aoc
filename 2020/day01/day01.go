// Package day01 solves AoC 2020 day 1.
package day01

import (
	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2020, 1, glue.IntSolver(solve))
}

func solve(expenses []int) ([]string, error) {
	return glue.Ints(part1(expenses), part2(expenses)), nil
}

func part1(expenses []int) int {
	for i, l := range expenses[0 : len(expenses)-1] {
		for _, r := range expenses[i+1:] {
			if l+r == 2020 {
				return l * r
			}
		}
	}
	return -1
}

func part2(expenses []int) int {
	for i, a := range expenses[0 : len(expenses)-2] {
		for j, b := range expenses[i+1 : len(expenses)-1] {
			for _, c := range expenses[i+j+2:] {
				if a+b+c == 2020 {
					return a * b * c
				}
			}
		}
	}
	return -1
}
