// Package day04 solves AoC 2024 day 4.
package day04

import (
	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2024, 4, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	part1 := countMatches(lines, "XMAS")
	part2 := countXmas(lines)
	return glue.Ints(part1, part2), nil
}

func countMatches(lines []string, s string) (count int) {
	W, H := len(lines[0]), len(lines)
	neigh := util.P{0, 0}.Neigh8()
	for y := range H {
		for x := range W {
			if lines[y][x] != s[0] {
				continue
			}
			for _, d := range neigh {
				if isMatch(lines, x+d.X, y+d.Y, d.X, d.Y, s[1:]) {
					count++
				}
			}
		}
	}
	return count
}

func isMatch(lines []string, x, y, dx, dy int, s string) bool {
	W, H := len(lines[0]), len(lines)
	for _, c := range []byte(s) {
		if x < 0 || y < 0 || x >= W || y >= H || lines[y][x] != c {
			return false
		}
		x, y = x+dx, y+dy
	}
	return true
}

func countXmas(lines []string) (count int) {
	W, H := len(lines[0]), len(lines)
	for y := 1; y < H-1; y++ {
		for x := 1; x < W-1; x++ {
			if lines[y][x] != 'A' {
				continue
			}
			c0, c1, c2, c3 := lines[y-1][x-1], lines[y-1][x+1], lines[y+1][x-1], lines[y+1][x+1]
			if (c0 == 'M' && c3 == 'S') || (c0 == 'S' && c3 == 'M') {
				if (c1 == 'M' && c2 == 'S') || (c1 == 'S' && c2 == 'M') {
					count++
				}
			}
		}
	}
	return count
}
