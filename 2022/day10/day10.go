// Package day10 solves AoC 2022 day 10.
package day10

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util/fn"
)

func init() {
	glue.RegisterSolver(2022, 10, glue.LineSolver(glue.WithParser(parseInstruction, solve)))
}

func solve(prog []instruction) ([]string, error) {
	p1 := sigStrength(prog)
	img := render(prog)
	return append(glue.Ints(p1), img...), nil
}

func sigStrength(prog []instruction) (strength int) {
	const (
		sampleCycle = 19
		sampleMod   = 40
	)
	cycle, x := 0, 1
	for _, inst := range prog {
		c := cycle % sampleMod
		if c <= sampleCycle && c+inst.cycles > sampleCycle {
			strength += (cycle - c + sampleCycle + 1) * x
		}
		cycle, x = cycle+inst.cycles, x+inst.addend
	}
	return strength
}

func render(prog []instruction) []string {
	const (
		W = 40
		H = 6
	)
	var screen [H][W]byte
	cycle, x := 0, 1
	for _, inst := range prog {
		for c := range inst.cycles {
			px, py := (cycle+c)%W, (cycle+c)/W
			screen[py][px] = fn.If[byte](px >= x-1 && px <= x+1, '#', ' ')
		}
		cycle, x = cycle+inst.cycles, x+inst.addend
	}
	return fn.Map(screen[:], func(row [W]byte) string { return string(row[:]) })
}

type instruction struct {
	cycles int
	addend int
}

func parseInstruction(line string) (instruction, error) {
	if line == "noop" {
		return instruction{cycles: 1, addend: 0}, nil
	}
	if strings.HasPrefix(line, "addx ") {
		addend, err := strconv.Atoi(line[5:])
		if err != nil {
			return instruction{}, fmt.Errorf("invalid addend: %q: %w", line[5:], err)
		}
		return instruction{cycles: 2, addend: addend}, nil
	}
	return instruction{}, fmt.Errorf("invalid instruction: %q", line)
}
