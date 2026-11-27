// Package day08 solves AoC 2024 day 8.
package day08

import (
	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2024, 8, glue.FixedLevelSolver(solve))
}

func solve(level *util.FixedLevel) ([]string, error) {
	antennae := findAntennae(level)
	part1 := countAntinodes(level, antennae)
	part2 := countAllAntinodes(level, antennae)
	return glue.Ints(part1, part2), nil
}

func countAntinodes(level *util.FixedLevel, antennae [][]util.P) int {
	W, H := level.W, level.H
	antinodes := util.MakeFixedBitmap1D(W * H)
	for _, fants := range antennae {
		for i, ip := range fants {
			for j, jp := range fants {
				if i == j {
					continue
				}
				ax, ay := 2*ip.X-jp.X, 2*ip.Y-jp.Y
				if ax >= 0 && ax < W && ay >= 0 && ay < H {
					antinodes.Set(ay*W + ax)
				}
			}
		}
	}
	return antinodes.Count()
}

func countAllAntinodes(level *util.FixedLevel, antennae [][]util.P) int {
	W, H := level.W, level.H
	antinodes := util.MakeFixedBitmap1D(W * H)
	for _, fants := range antennae {
		for i, ip := range fants {
			for j, jp := range fants {
				if i == j {
					continue
				}
				dx, dy := jp.X-ip.X, jp.Y-ip.Y
				for x, y := jp.X, jp.Y; x >= 0 && x < W && y >= 0 && y < H; x, y = x+dx, y+dy {
					antinodes.Set(y*W + x)
				}
			}
		}
	}
	return antinodes.Count()
}

func findAntennae(level *util.FixedLevel) [][]util.P {
	antennae := make([][]util.P, maxFreq)
	for y := range level.H {
		for x := range level.W {
			if f := freqIdx(level.At(x, y)); f >= 0 {
				antennae[f] = append(antennae[f], util.P{x, y})
			}
		}
	}
	return antennae
}

const maxFreq = 2*26 + 10

func freqIdx(c byte) int {
	if c >= '0' && c <= '9' {
		return int(c - '0')
	}
	if c >= 'A' && c <= 'Z' {
		return 10 + int(c-'A')
	}
	if c >= 'a' && c <= 'z' {
		return 10 + 26 + int(c-'a')
	}
	return -1
}
