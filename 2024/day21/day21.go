// Package day21 solves AoC 2024 day 21.
package day21

import (
	"math"
	"strconv"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/fn"
)

func init() {
	glue.RegisterSolver(2024, 21, glue.LineSolver(glue.WithParser(parseCode, solve)))
}

func solve(codes []entryCode) ([]string, error) {
	cc := makeCalculator()
	part1 := fn.SumF(codes, func(code entryCode) int { return cc.codeComplexity(code, 2) })
	part2 := fn.SumF(codes, func(code entryCode) int { return cc.codeComplexity(code, 25) })
	return glue.Ints(part1, part2), nil
}

type costCalculator struct {
	cache map[uint64]int
}

func makeCalculator() costCalculator {
	return costCalculator{cache: make(map[uint64]int)}
}

func (cc costCalculator) codeComplexity(code entryCode, keypads int) int {
	return cc.numCost(numKeyPos[10], code.keys, keypads, make([]dirKey, 0, 32)) * code.value
}

func (cc costCalculator) numCost(from util.P, keys []byte, keypads int, pressed []dirKey) int {
	if len(keys) == 0 {
		return cc.dirComplexity(pressed, keypads)
	}
	to := numKeyPos[keys[0]]
	if from == to {
		return cc.numCost(from, keys[1:], keypads, append(pressed, dirA))
	}
	costH, costV := math.MaxInt, math.MaxInt
	if to.X > from.X {
		costH = cc.numCost(from.AddXY(1, 0), keys, keypads, append(pressed, dirRight))
	} else if to.X < from.X {
		if from.X-1 != numBadX || from.Y != numBadY {
			costH = cc.numCost(from.AddXY(-1, 0), keys, keypads, append(pressed, dirLeft))
		}
	}
	if to.Y > from.Y {
		if from.X != numBadX || from.Y+1 != numBadY {
			costV = cc.numCost(from.AddXY(0, 1), keys, keypads, append(pressed, dirDown))
		}
	} else if to.Y < from.Y {
		costV = cc.numCost(from.AddXY(0, -1), keys, keypads, append(pressed, dirUp))
	}
	return min(costH, costV)
}

func (cc costCalculator) dirComplexity(keys []dirKey, keypads int) (cost int) {
	stack := make([]dirKey, keypads)
	for _, key := range keys {
		cost += cc.dirCost(stack, key)
		stack[0] = key
	}
	return cost
}

func (cc costCalculator) dirCost(at []dirKey, to dirKey) int {
	at0 := at[0]
	if len(at) == 1 {
		return util.DistM(dirKeyPos[at0], dirKeyPos[to]) + 1
	}
	if at0 == to {
		return cc.dirCost(at[1:], dirA)
	}
	ck := cacheKey(at, to)
	if cost, ok := cc.cache[ck]; ok {
		return cost
	}
	at1 := at[1]
	costH, costV := math.MaxInt, math.MaxInt
	atP, toP := dirKeyPos[at0], dirKeyPos[to]
	if toP.X > atP.X {
		costH = cc.dirCost(at[1:], dirRight)
		at[0], at[1] = dirKeyAt[atP.Y][atP.X+1], dirRight
		costH += cc.dirCost(at, to)
	} else if toP.X < atP.X {
		if atP.X-1 != dirBadX || atP.Y != dirBadY {
			costH = cc.dirCost(at[1:], dirLeft)
			at[0], at[1] = dirKeyAt[atP.Y][atP.X-1], dirLeft
			costH += cc.dirCost(at, to)
		}
	}
	at[0], at[1] = at0, at1
	if toP.Y > atP.Y {
		if atP.X != dirBadX || atP.Y+1 != dirBadY {
			costV = cc.dirCost(at[1:], dirDown)
			at[0], at[1] = dirKeyAt[atP.Y+1][atP.X], dirDown
			costV += cc.dirCost(at, to)
		}
	} else if toP.Y < atP.Y {
		if atP.X != dirBadX || atP.Y-1 != dirBadY {
			costV = cc.dirCost(at[1:], dirUp)
			at[0], at[1] = dirKeyAt[atP.Y-1][atP.X], dirUp
			costV += cc.dirCost(at, to)
		}
	}
	at[0], at[1] = at0, at1
	cost := min(costH, costV)
	cc.cache[ck] = cost
	return cost
}

func cacheKey(at []dirKey, to dirKey) (cacheKey uint64) {
	cacheKey = uint64(len(at))
	cacheKey = 5*cacheKey + uint64(to)
	for _, key := range at {
		cacheKey = 5*cacheKey + uint64(key)
	}
	return cacheKey
}

type dirKey byte

const (
	dirA dirKey = iota
	dirRight
	dirDown
	dirLeft
	dirUp
)

const (
	numBadX = 0
	numBadY = 3
	dirBadX = 0
	dirBadY = 0
)

var (
	numKeyPos = [11]util.P{
		// 0 1 2 3 4 5 6 7 8 9 A
		{1, 3}, {0, 2}, {1, 2}, {2, 2}, {0, 1}, {1, 1}, {2, 1}, {0, 0}, {1, 0}, {2, 0}, {2, 3},
	}
	dirKeyPos = [5]util.P{
		// A > v < ^
		{2, 0}, {2, 1}, {1, 1}, {0, 1}, {1, 0},
	}
	dirKeyAt = [2][3]dirKey{
		{0, dirUp, dirA},
		{dirLeft, dirDown, dirRight},
	}
)

type entryCode struct {
	keys  []byte
	value int
}

func parseCode(line string) (entryCode, error) {
	keys := make([]byte, len(line))
	for i := range len(line) - 1 {
		keys[i] = line[i] - '0'
	}
	keys[len(line)-1] = 10
	value, _ := strconv.Atoi(line[:len(line)-1])
	return entryCode{keys: keys, value: value}, nil
}
