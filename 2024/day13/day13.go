// Package day13 solves AoC 2024 day 13.
package day13

import (
	"io"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2024, 13, glue.GenericSolver(solve))
}

func solve(r io.Reader) ([]string, error) {
	claws, err := parseClaws(r)
	if err != nil {
		return nil, err
	}
	part1 := minCost(claws, 0)
	part2 := minCost(claws, 10000000000000)
	return glue.Ints(part1, part2), nil
}

func minCost(claws []clawMachine, offset int) (cost int) {
	for _, claw := range claws {
		if c, ok := claw.minCost(offset); ok {
			cost += c
		}
	}
	return cost
}

const (
	costA = 3
	costB = 1
)

type clawMachine struct {
	dA, dB util.P
	prize  util.P
}

func (claw *clawMachine) minCost(offset int) (cost int, ok bool) {
	dA, dB, p := claw.dA, claw.dB, claw.prize.AddXY(offset, offset)
	a := p.X*dB.Y - p.Y*dB.X
	d := dA.X*dB.Y - dA.Y*dB.X
	if a%d != 0 {
		return -1, false
	}
	b := p.Y*dA.X - p.X*dA.Y
	if b%d != 0 {
		return -1, false
	}
	return a/d*costA + b/d*costB, true
}

func parseClaws(r io.Reader) (claws []clawMachine, err error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	for {
		aX, ok, tail := nextInt(data)
		if !ok {
			return claws, nil
		}
		aY, _, tail := nextInt(tail)
		bX, _, tail := nextInt(tail)
		bY, _, tail := nextInt(tail)
		pX, _, tail := nextInt(tail)
		pY, _, tail := nextInt(tail)
		claws = append(claws, clawMachine{dA: util.P{aX, aY}, dB: util.P{bX, bY}, prize: util.P{pX, pY}})
		data = tail
	}
}

func nextInt(data []byte) (value int, ok bool, tail []byte) {
	neg := false
	for len(data) > 0 && (data[0] < '0' || data[0] > '9') {
		neg = data[0] == '-'
		data = data[1:]
	}
	for len(data) > 0 && data[0] >= '0' && data[0] <= '9' {
		ok = true
		value = 10*value + int(data[0]-'0')
		data = data[1:]
	}
	if neg {
		value = -value
	}
	return value, ok, data
}
