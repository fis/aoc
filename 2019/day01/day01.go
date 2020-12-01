// Package day01 solves AoC 2019 day 1.
package day01

import (
	"strconv"

	"github.com/fis/aoc-go/util"
)

func Solve(path string) ([]string, error) {
	modules, err := util.ReadIntRows(path)
	if err != nil {
		return nil, err
	}
	return []string{
		strconv.Itoa(part1(modules)),
		strconv.Itoa(part2(modules)),
	}, nil
}

func part1(modules []int) int {
	sum := 0
	for _, m := range modules {
		sum += moduleFuel(m)
	}
	return sum
}

func moduleFuel(w int) int {
	return w/3 - 2
}

func part2(modules []int) int {
	sum := 0
	for _, m := range modules {
		sum += moduleFuelComplete(m)
	}
	return sum
}

func moduleFuelComplete(w int) int {
	total := 0
	for {
		f := moduleFuel(w)
		if f <= 0 {
			return total
		}
		total += f
		w = f
	}
}
