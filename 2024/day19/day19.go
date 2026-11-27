// Package day19 solves AoC 2024 day 19.
package day19

import (
	"fmt"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2024, 19, glue.ChunkSolver(solve))
}

func solve(chunks []string) ([]string, error) {
	if len(chunks) != 2 {
		return nil, fmt.Errorf("expected 2 chunks, got %d", len(chunks))
	}
	towels := strings.Split(chunks[0], ", ")
	designs := util.Lines(chunks[1])
	part1, part2 := countPossible(designs, towels)
	return glue.Ints(part1, part2), nil
}

func countPossible(designs, towels []string) (part1, part2 int) {
	tree := buildTowelTree(towels)

	var possible []int

	for _, design := range designs {
		possible = util.EnsureSize(possible, len(design)+1)
		possible[len(design)] = 1
		for start := len(design) - 1; start >= 0; start-- {
			possible[start] = 0
			for end, t := start+1, tree; end <= len(design); end++ {
				t = t.child[stripeID(design[end-1])]
				if t == nil {
					break
				}
				if t.valid {
					possible[start] += possible[end]
				}
			}
		}
		if possible[0] > 0 {
			part1++
		}
		part2 += possible[0]
	}

	return part1, part2
}

type towelTree struct {
	valid bool
	child [8]*towelTree
}

func buildTowelTree(towels []string) *towelTree {
	root := &towelTree{}
	for _, towel := range towels {
		at := root
		for _, c := range []byte(towel) {
			s := stripeID(c)
			if at.child[s] == nil {
				at.child[s] = &towelTree{}
			}
			at = at.child[s]
		}
		at.valid = true
	}
	return root
}

func stripeID(c byte) int {
	// selected bits result in differentiating "wubrg"
	return int(c&3 | (c>>2)&4)
}
