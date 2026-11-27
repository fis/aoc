// Package day02 solves AoC 2017 day 2.
package day02

import (
	"fmt"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2017, 2, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	sheet := make([][]int, len(lines))
	for i, line := range lines {
		sheet[i] = util.Ints(line)
	}
	p1 := checksumSheet(sheet)
	p2, err := divisibleSheet(sheet)
	if err != nil {
		return nil, err
	}
	return glue.Ints(p1, p2), nil
}

func checksumSheet(rows [][]int) (sum int) {
	for _, row := range rows {
		sum += checksumRow(row)
	}
	return sum
}

func checksumRow(row []int) int {
	min, max := row[0], row[0]
	for _, n := range row[1:] {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return max - min
}

func divisibleSheet(rows [][]int) (sum int, err error) {
	for _, row := range rows {
		d, err := divisibleRow(row)
		if err != nil {
			return 0, err
		}
		sum += d
	}
	return sum, nil
}

func divisibleRow(row []int) (int, error) {
	for _, i := range row {
		for _, j := range row {
			if i > j && i%j == 0 {
				return i / j, nil
			}
		}
	}
	return 0, fmt.Errorf("invalid input: %v: no evenly divisible pair", row)
}
