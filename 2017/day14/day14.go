// Package day14 solves AoC 2017 day 14.
package day14

import (
	"fmt"
	"math/bits"

	"github.com/fis/aoc/2017/knot"
	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2017, 14, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	if len(lines) != 1 {
		return nil, fmt.Errorf("expected 1 line of input, got %d", len(lines))
	}
	hashes, p1 := hash(lines[0])
	lvl := buildLevel(hashes)
	p2 := countRegions(lvl)
	return glue.Ints(p1, p2), nil
}

func hash(key string) (hashes [][]byte, used int) {
	for i := range 128 {
		rowKey := fmt.Sprintf("%s-%d", key, i)
		hash := knot.Hash(knot.N, knot.Rounds, rowKey)
		hashes = append(hashes, hash)
		for _, b := range hash {
			used += bits.OnesCount8(b)
		}
	}
	return hashes, used
}

func buildLevel(hashes [][]byte) *util.Level {
	lvl := util.EmptyLevel(util.P{0, 0}, util.P{knot.N - 1, len(hashes) - 1}, '.')
	for y, row := range hashes {
		for x1, b := range row {
			for x2 := range 8 {
				if ((b << x2) & 0x80) != 0 {
					lvl.Set(8*x1+x2, y, '#')
				}
			}
		}
	}
	return lvl
}

func countRegions(lvl *util.Level) (regions int) {
	seen := map[util.P]struct{}{}
	for start := range lvl.All() {
		if _, found := seen[start]; found {
			continue
		}
		regions++
		edge := []util.P{start}
		for len(edge) > 0 {
			at := edge[len(edge)-1]
			edge = edge[:len(edge)-1]
			if _, found := seen[at]; found {
				continue
			}
			seen[at] = struct{}{}
			for _, n := range at.Neigh() {
				if _, found := seen[n]; !found && lvl.At(n.X, n.Y) == '#' {
					edge = append(edge, n)
				}
			}
		}
	}
	return regions
}
