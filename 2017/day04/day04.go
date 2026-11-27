// Package day04 solves AoC 2017 day 4.
package day04

import (
	"sort"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2017, 4, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	p1, p2 := 0, 0
	for _, line := range lines {
		words := util.Words(line)
		if part1(words) {
			p1++
		}
		if part2(words) {
			p2++
		}
	}
	return glue.Ints(p1, p2), nil
}

func part1(passphrase []string) bool {
	seen := make(map[string]struct{})
	for _, word := range passphrase {
		if _, found := seen[word]; found {
			return false
		}
		seen[word] = struct{}{}
	}
	return true
}

func part2(passphrase []string) bool {
	seen := make(map[string]struct{})
	for _, word := range passphrase {
		chars := []byte(word)
		sort.Slice(chars, func(i, j int) bool { return chars[i] < chars[j] })
		word = string(chars)
		if _, found := seen[word]; found {
			return false
		}
		seen[word] = struct{}{}
	}
	return true
}
