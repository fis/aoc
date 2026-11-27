// Package day11 solves AoC 2020 day 11.
package day11

import (
	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2020, 11, glue.LevelSolver{Solver: solve, Empty: '.'})
}

func solve(level *util.Level) ([]string, error) {
	fp1 := fixedPoint(level, nearMap, 4)
	fp2 := fixedPoint(level, farMap, 5)
	return glue.Ints(fp1, fp2), nil
}

func fixedPoint(level *util.Level, mapper func(*util.Level) [][8]int, tolerance int) (occupied int) {
	neigh := mapper(level)
	cur := make([]bool, len(neigh))
	next := make([]bool, len(neigh))
	for simulate(cur, next, neigh, tolerance) {
		cur, next = next, cur
	}
	for _, seat := range next {
		if seat {
			occupied++
		}
	}
	return occupied
}

func nearMap(level *util.Level) [][8]int {
	offsets := make(map[util.P]int)
	for p := range level.All() {
		offsets[p] = len(offsets)
	}
	neigh := make([][8]int, len(offsets))
	for p := range level.All() {
		n, ni := &neigh[offsets[p]], 0
		for _, np := range p.Neigh8() {
			if level.At(np.X, np.Y) == 'L' {
				n[ni] = offsets[np]
				ni++
			}
		}
		for ; ni < 8; ni++ {
			n[ni] = -1
		}
	}
	return neigh
}

func farMap(level *util.Level) [][8]int {
	offsets := make(map[util.P]int)
	for p := range level.All() {
		offsets[p] = len(offsets)
	}
	neigh := make([][8]int, len(offsets))
	for p := range level.All() {
		n, ni := &neigh[offsets[p]], 0
		for _, d := range (util.P{0, 0}).Neigh8() {
			for i := 1; level.InBounds(p.X+i*d.X, p.Y+i*d.Y); i++ {
				if level.At(p.X+i*d.X, p.Y+i*d.Y) == 'L' {
					n[ni] = offsets[util.P{p.X + i*d.X, p.Y + i*d.Y}]
					ni++
					break
				}
			}
		}
		for ; ni < 8; ni++ {
			n[ni] = -1
		}
	}
	return neigh
}

func simulate(in, out []bool, neigh [][8]int, tolerance int) (changed bool) {
	for i, ni := range neigh {
		nc := 0
		for j := 0; j < 8 && ni[j] >= 0; j++ {
			if in[ni[j]] {
				nc++
			}
		}
		c := in[i]
		if !c && nc == 0 {
			c = true
			changed = true
		} else if c && nc >= tolerance {
			c = false
			changed = true
		}
		out[i] = c
	}
	return changed
}
