// Package day09 solves AoC 2023 day 9.
package day09

import (
	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/fn"
)

func init() {
	glue.RegisterSolver(2023, 9, glue.LineSolver(glue.WithParser(func(line string) ([]int, error) { return util.Ints(line), nil }, solve)))
}

func solve(seqs [][]int) ([]string, error) {
	p1 := fn.SumF(seqs, predict)
	p2 := fn.SumF(seqs, unpredict)
	return glue.Ints(p1, p2), nil
}

func predict(seq []int) int {
	buf := append([]int(nil), seq...)
	var n int
	for n = len(buf) - 1; n > 0; n-- {
		nz := false
		for i := range n {
			buf[i] = buf[i+1] - buf[i]
			nz = nz || buf[i] != 0
		}
		if !nz {
			break
		}
	}
	return fn.Sum(buf[n:])
}

func unpredict(seq []int) int {
	buf := append([]int(nil), seq...)
	var s int
	for s = 1; s < len(buf); s++ {
		nz := false
		for i := len(buf) - 1; i >= s; i-- {
			buf[i] = buf[i-1] - buf[i]
			nz = nz || buf[i] != 0
		}
		if !nz {
			break
		}
	}
	return fn.Sum(buf[:s])
}
