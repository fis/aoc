// Package day17 solves AoC 2019 day 17.
package day17

import (
	"strconv"

	"github.com/fis/aoc2019-go/intcode"
	"github.com/fis/aoc2019-go/util"
)

var input = []string{
	"A,A,B,C,B,C,B,C,C,A",
	"R,8,L,4,R,4,R,10,R,8",
	"L,12,L,12,R,8,R,8",
	"R,10,R,4,R,4",
	"n",
}

func Solve(path string) ([]string, error) {
	prog, err := intcode.Load(path)
	if err != nil {
		return nil, err
	}

	level := capture(prog)
	p1 := crosses(level)

	prog[0] = 2
	vm := intcode.VM{}
	vm.Load(prog)
	out := vm.Run(unlines(input))

	return []string{strconv.Itoa(p1), strconv.FormatInt(out[len(out)-1], 10)}, nil
}

func capture(prog []int64) *util.Level {
	vm := intcode.VM{}
	vm.Load(prog)
	out := vm.Run([]int64{})
	level := util.ParseLevelString("", '.')
	x, y := 0, 0
	for _, v := range out {
		switch v {
		case 10:
			x, y = 0, y+1
		case '#':
			level.Set(x, y, '#')
			fallthrough
		default:
			x++
		}
	}
	return level
}

func crosses(level *util.Level) int {
	sum := 0
	level.Range(func(x, y int, _ byte) {
		for _, n := range (util.P{x, y}).Neigh() {
			if level.At(n.X, n.Y) != '#' {
				return
			}
		}
		sum += x * y
	})
	return sum
}

func unlines(lines []string) []int64 {
	var out []int64
	for _, line := range lines {
		for _, r := range line {
			out = append(out, int64(r))
		}
		out = append(out, '\n')
	}
	return out
}
