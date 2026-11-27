// Package day06 solves AoC 2025 day 6.
package day06

import (
	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/fn"
)

func init() {
	glue.RegisterSolver(2025, 6, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	p1 := parseWorksheetWrong(lines).grandTotal()
	p2 := parseWorksheetRight(lines).grandTotal()
	return glue.Ints(p1, p2), nil
}

func parseWorksheetWrong(lines []string) (ws worksheet) {
	for _, line := range lines[:len(lines)-1] {
		nums := util.Ints(line)
		if n := len(nums) - len(ws); n > 0 {
			ws = append(ws, make(worksheet, n)...)
		}
		for i, n := range nums {
			ws[i].nums = append(ws[i].nums, n)
		}
	}
	for i, w := range util.Words(lines[len(lines)-1]) {
		ws[i].op = w[0]
	}
	return ws
}

func parseWorksheetRight(lines []string) (ws worksheet) {
	var p problem
	for i := len(lines[0]) - 1; i >= 0; i-- {
		num := 0
		for _, line := range lines[:len(lines)-1] {
			if line[i] == ' ' {
				continue
			}
			num = 10*num + int(line[i]-'0')
		}
		if num > 0 {
			p.nums = append(p.nums, num)
		}
		p.op = lines[len(lines)-1][i]
		if p.op != ' ' {
			ws = append(ws, p)
			p = problem{}
		}
	}
	return ws
}

type worksheet []problem

func (ws worksheet) grandTotal() int { return fn.SumF(ws, problem.solve) }

type problem struct {
	op   byte
	nums []int
}

func (p problem) solve() int {
	switch p.op {
	case '+':
		return fn.Sum(p.nums)
	case '*':
		return fn.Prod(p.nums)
	default:
		return -1
	}
}
