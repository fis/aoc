// Package day09 solves AoC 2021 day 9.
package day09

import (
	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2021, 9, glue.LevelSolver{
		Solver: solve,
		Empty:  ' ',
	})
}

func solve(level *util.Level) ([]string, error) {
	p1 := riskLevels(level)
	p2 := basinSizes(level)
	return glue.Ints(p1, p2), nil
}

func riskLevels(level *util.Level) (totalRisk int) {
outer:
	for p, b := range level.All() {
		for _, n := range p.Neigh() {
			nb := level.At(n.X, n.Y)
			if nb != ' ' && nb <= b {
				continue outer
			}
		}
		totalRisk += int(b-'0') + 1
	}
	return totalRisk
}

func basinSizes(level *util.Level) (score int) {
	min, max := level.Bounds()
	var sizes []int
	for y := min.Y; y <= max.Y; y++ {
		for x := min.X; x <= max.X; x++ {
			if b := level.At(x, y); b < '0' || b > '8' {
				continue // not in a basin
			}
			sizes = append(sizes, basinSize(level, x, y))
		}
	}
	a, b, c := top3(sizes)
	return a * b * c
}

func basinSize(level *util.Level, x, y int) (size int) {
	edge := []util.P{{x, y}}
	for len(edge) > 0 {
		p := edge[len(edge)-1]
		edge = edge[:len(edge)-1]
		if b := level.At(p.X, p.Y); b < '0' || b > '8' {
			continue
		}
		level.Set(p.X, p.Y, ' ')
		size++
		for _, n := range p.Neigh() {
			if nb := level.At(n.X, n.Y); nb >= '0' && nb <= '8' {
				edge = append(edge, n)
			}
		}
	}
	return size
}

func top3(data []int) (a, b, c int) {
	c, b, a = util.Sort3(data[0], data[1], data[2])
	for _, x := range data[3:] {
		if x >= a {
			a, b, c = x, a, b
		} else if x >= b {
			b, c = x, b
		} else if x > c {
			c = x
		}
	}
	return a, b, c
}
