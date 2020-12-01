// Package day13 solves AoC 2019 day 13.
package day13

import (
	"strconv"

	"github.com/fis/aoc-go/intcode"
)

func Solve(path string) ([]string, error) {
	prog, err := intcode.Load(path)
	if err != nil {
		return nil, err
	}
	p1 := part1(prog)
	prog[0] = 2
	p2 := part2(prog)
	return []string{strconv.Itoa(p1), strconv.Itoa(p2)}, nil
}

func part1(prog []int64) int {
	out, _ := intcode.Run(prog, nil)
	blocks := 0
	for i := 2; i < len(out); i += 3 {
		if out[i] == 2 {
			blocks++
		}
	}
	return blocks
}

func part2(prog []int64) int {
	var (
		vm  intcode.VM
		tok intcode.WalkToken
	)
	var score, ball, paddle int64
	vm.Load(prog)
	for vm.Walk(&tok) {
		if tok.IsInput() {
			switch {
			case paddle < ball:
				tok.ProvideInput(1)
			case paddle > ball:
				tok.ProvideInput(-1)
			default:
				tok.ProvideInput(0)
			}
			continue
		}
		x := tok.ReadOutput()
		// ignore middle output (Y position)
		vm.Walk(&tok)
		vm.Walk(&tok)
		out := tok.ReadOutput()
		switch {
		case x == -1:
			score = out
		case x >= 0 && out == 3:
			paddle = x
		case x >= 0 && out == 4:
			ball = x
		}
	}
	return int(score)
}
