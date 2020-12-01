// Package day15 solves AoC 2019 day 15.
package day15

import (
	"fmt"
	"strconv"

	"github.com/fis/aoc-go/intcode"
	"github.com/fis/aoc-go/util"
)

func Solve(path string) ([]string, error) {
	prog, err := intcode.Load(path)
	if err != nil {
		return nil, err
	}

	dr := realDroid{}
	dr.vm.Load(prog)
	level := explore(&dr)

	p1, target := distance(level, util.P{0, 0}, 'O')
	p2, _ := distance(level, target, ' ')

	return []string{strconv.Itoa(p1), strconv.Itoa(p2)}, nil
}

type tile byte

type droid interface {
	tryMove(to util.P) byte
	mustMove(to util.P)
}

func explore(dr droid) *util.Level {
	level := util.ParseLevelString(".", ' ')

	var dfs func(util.P)
	dfs = func(pos util.P) {
		for _, step := range pos.Neigh() {
			if level.At(step.X, step.Y) != ' ' {
				continue // already visited
			}
			tile := dr.tryMove(step)
			level.Set(step.X, step.Y, tile)
			if tile == '#' {
				continue // wall, didn't move
			}
			dfs(step)
			dr.mustMove(pos)
		}
	}

	dfs(util.P{0, 0})
	return level
}

func distance(level *util.Level, from util.P, toTile byte) (int, util.P) {
	seen := make(map[util.P]struct{})
	fringe := []util.P{from}
	d := 0
	for len(fringe) > 0 {
		d++
		var newFringe []util.P
		for _, pos := range fringe {
			seen[pos] = struct{}{}
			for _, step := range pos.Neigh() {
				tile := level.At(step.X, step.Y)
				if tile == toTile {
					return d, step
				} else if tile == '#' {
					continue // wall
				} else if _, ok := seen[step]; ok {
					continue // visited
				}
				newFringe = append(newFringe, step)
			}
		}
		fringe = newFringe
	}
	return d - 1, from
}

type realDroid struct {
	vm  intcode.VM
	tok intcode.WalkToken
	pos util.P
}

func (dr *realDroid) tryMove(to util.P) byte {
	dx, dy, dir := to.X-dr.pos.X, to.Y-dr.pos.Y, int64(0)
	switch {
	case dx == 0 && dy == -1:
		dir = 1 // north
	case dx == 0 && dy == 1:
		dir = 2 // south
	case dx == -1 && dy == 0:
		dir = 3 // west
	case dx == 1 && dy == 0:
		dir = 4 // east
	default:
		panic(fmt.Sprintf("invalid move: %v to %v", dr.pos, to))
	}
	dr.vm.Walk(&dr.tok)
	dr.tok.ProvideInput(dir)
	dr.vm.Walk(&dr.tok)
	switch out := dr.tok.ReadOutput(); out {
	case 0:
		return '#' // hit a wall
	case 1:
		dr.pos = to
		return '.' // regular corridor
	case 2:
		dr.pos = to
		return 'O' // found the oxygen system
	default:
		panic(fmt.Sprintf("invalid result: %d", out))
	}
}

func (dr *realDroid) mustMove(to util.P) {
	tile := dr.tryMove(to)
	if tile == '#' {
		panic(fmt.Sprintf("invalid move: %v to %v: must succeed", dr.pos, to))
	}
}
