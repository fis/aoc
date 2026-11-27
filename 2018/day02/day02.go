// Package day02 solves AoC 2018 day 2.
package day02

import (
	"strconv"

	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2018, 2, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	cs := checksum(lines)
	box := findBox(lines)

	return []string{strconv.Itoa(cs), box}, nil
}

func checksum(ids []string) int {
	var twos, threes int
	for _, id := range ids {
		var freqs [26]int
		for _, c := range id {
			if c >= 'a' && c <= 'z' {
				freqs[c-'a']++
			}
		}
		has2, has3 := false, false
		for _, f := range freqs {
			has2 = has2 || f == 2
			has3 = has3 || f == 3
		}
		if has2 {
			twos++
		}
		if has3 {
			threes++
		}
	}
	return twos * threes
}

func diff1(left, right string) (bool, string) {
	pos := -1
	for i := range left {
		if left[i] != right[i] {
			if pos >= 0 {
				return false, ""
			}
			pos = i
		}
	}
	if pos < 0 {
		return false, ""
	}
	return true, left[:pos] + left[pos+1:]
}

func findBox(ids []string) string {
	for i, l := range ids[:len(ids)-1] {
		for _, r := range ids[i+1:] {
			if ok, s := diff1(l, r); ok {
				return s
			}
		}
	}
	return ""
}
