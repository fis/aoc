// Package day07 solves AoC 2024 day 7.
package day07

import (
	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/ix"
)

func init() {
	glue.RegisterSolver(2024, 7, glue.LineSolver(glue.WithParser(parseEquation, solve)))
}

func solve(eqs []equation) ([]string, error) {
	part1 := sumPossible(eqs)
	part2 := sumPossible3(eqs)
	return glue.Ints(part1, part2), nil
}

func sumPossible(eqs []equation) (sum int) {
	for _, eq := range eqs {
		if isPossible(eq.result, eq.nums[0], eq.nums[1:]) {
			sum += eq.result
		}
	}
	return sum
}

func isPossible(result, value int, nums []int) bool {
	if len(nums) == 0 {
		return result == value
	}
	if nv := value * nums[0]; nv <= result && isPossible(result, nv, nums[1:]) {
		return true
	}
	if nv := value + nums[0]; nv <= result && isPossible(result, nv, nums[1:]) {
		return true
	}
	return false
}

func sumPossible3(eqs []equation) (sum int) {
	for _, eq := range eqs {
		if isPossible3(eq.result, eq.nums[0], eq.nums[1:]) {
			sum += eq.result
		}
	}
	return sum
}

func isPossible3(result, value int, nums []int) bool {
	if len(nums) == 0 {
		return result == value
	}
	if nv := value * nums[0]; nv <= result && isPossible3(result, nv, nums[1:]) {
		return true
	}
	if nv := value + nums[0]; nv <= result && isPossible3(result, nv, nums[1:]) {
		return true
	}
	if nv := ix.Concat(value, nums[0]); nv <= result && isPossible3(result, nv, nums[1:]) {
		return true
	}
	return false
}

type equation struct {
	result int
	nums   []int
}

func parseEquation(line string) (equation, error) {
	nums := util.Ints(line)
	return equation{result: nums[0], nums: nums[1:]}, nil
}
