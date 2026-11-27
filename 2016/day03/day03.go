// Package day03 solves AoC 2016 day 3.
package day03

import "github.com/fis/aoc/glue"

func init() {
	glue.RegisterSolver(2016, 3, glue.IntSolver(solve))
}

func solve(nums []int) ([]string, error) {
	p1 := part1(nums)
	p2 := part2(nums)
	return glue.Ints(p1, p2), nil
}

func part1(nums []int) (valid int) {
	for i := 0; i+2 < len(nums); i += 3 {
		if isValid(nums[i], nums[i+1], nums[i+2]) {
			valid++
		}
	}
	return valid
}

func part2(nums []int) (valid int) {
	for base := 0; base+8 < len(nums); base += 9 {
		for i := range 3 {
			if isValid(nums[base+i], nums[base+i+3], nums[base+i+6]) {
				valid++
			}
		}
	}
	return valid
}

func isValid(a, b, c int) bool {
	return a+b > c && a+c > b && b+c > a
}
