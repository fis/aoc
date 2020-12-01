// Package day07 solves AoC 2019 day 7.
package day07

import (
	"strconv"

	"github.com/fis/aoc-go/intcode"
)

func Solve(path string) ([]string, error) {
	prog, err := intcode.Load(path)
	if err != nil {
		return nil, err
	}
	return []string{
		strconv.FormatInt(part1(prog), 10),
		strconv.FormatInt(part2(prog), 10),
	}, nil
}

const ampCount = 5

func part1(prog []int64) int64 {
	return findBest(prog, &[5]int64{0, 1, 2, 3, 4}, run)
}

func part2(prog []int64) int64 {
	return findBest(prog, &[5]int64{5, 6, 7, 8, 9}, runFeedback)
}

func findBest(prog []int64, phases *[ampCount]int64, runner func([]int64, *[ampCount]int64) int64) int64 {
	best := int64(0)
	permutations(phases, 5, func(phases *[ampCount]int64) {
		sig := runner(prog, phases)
		if sig > best {
			best = sig
		}
	})
	return best
}

func run(prog []int64, phases *[ampCount]int64) int64 {
	sig := int64(0)
	for i := 0; i < ampCount; i++ {
		out, _ := intcode.Run(prog, []int64{phases[i], sig})
		sig = out[0]
	}
	return sig
}

func runFeedback(prog []int64, phases *[ampCount]int64) int64 {
	amps, tokens := [ampCount]intcode.VM{}, [ampCount]intcode.WalkToken{}
	for i := 0; i < ampCount; i++ {
		amps[i].Load(prog)
		amps[i].Walk(&tokens[i])
		tokens[i].ProvideInput(phases[i])
		amps[i].Walk(&tokens[i])
	}
	sig, done := int64(0), false
	for !done {
		for i := 0; i < ampCount; i++ {
			tokens[i].ProvideInput(sig)
			amps[i].Walk(&tokens[i])
			sig = tokens[i].ReadOutput()
			amps[i].Walk(&tokens[i])
			done = done || tokens[i].IsEmpty()
		}
	}
	return sig
}

func permutations(vals *[ampCount]int64, k int, cb func(*[ampCount]int64)) {
	if k == 1 {
		cb(vals)
		return
	}
	permutations(vals, k-1, cb)
	even := k&1 == 0
	for i := 0; i < k-1; i++ {
		if even {
			vals[i], vals[k-1] = vals[k-1], vals[i]
		} else {
			vals[0], vals[k-1] = vals[k-1], vals[0]
		}
		permutations(vals, k-1, cb)
	}
}
