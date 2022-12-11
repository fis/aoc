// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package day09 solves AoC 2022 day 9.
package day09

import (
	"fmt"
	"strconv"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/ix"
)

const movePattern = `^([UDLR]) (\d+)$`

func init() {
	glue.RegisterSolver(2022, 9, glue.ParsableLineSolver[move]{
		Solver: solve,
		Parser: parseMove,
	})
}

func solve(moves []move) ([]string, error) {
	p1 := measureTail(moves)
	p2 := measureLongTail(moves)
	return glue.Ints(p1, p2), nil
}

// Solutions using util.Bitmap2D:

func measureTail(moves []move) int {
	head, tail := util.P{0, 0}, util.P{0, 0}
	var visited util.Bitmap2D
	visited.Set(0, 0)
	for _, m := range moves {
		d := dirStep[m.dir]
		for i := 0; i < m.steps; i++ {
			head = head.Add(d)
			tail = updateTail(head, tail)
			visited.Set(tail.X, tail.Y)
		}
	}
	return visited.Count()
}

func measureLongTail(moves []move) int {
	const ropeLength = 10
	var rope [ropeLength]util.P
	var visited util.Bitmap2D
	visited.Set(0, 0)
	for _, m := range moves {
		d := dirStep[m.dir]
		for i := 0; i < m.steps; i++ {
			rope[0] = rope[0].Add(d)
			for i := 1; i < ropeLength; i++ {
				rope[i] = updateTail(rope[i-1], rope[i])
			}
			tail := rope[ropeLength-1]
			visited.Set(tail.X, tail.Y)
		}
	}
	return visited.Count()
}

func updateTail(head, tail util.P) util.P {
	dx, dy := head.X-tail.X, head.Y-tail.Y
	sx, sy := ix.Sign(dx), ix.Sign(dy)
	if dx == sx && dy == sy {
		return tail
	}
	return util.P{tail.X + sx, tail.Y + sy}
}

type move struct {
	dir   direction
	steps int
}

type direction int

const (
	dirUp direction = iota
	dirDown
	dirLeft
	dirRight
)

func parseMove(line string) (move, error) {
	if len(line) < 3 || line[1] != ' ' {
		return move{}, fmt.Errorf("bad move format: %s", line)
	}
	dir, ok := dirMap[line[0]]
	if !ok {
		return move{}, fmt.Errorf("bad direction letter: %c", line[0])
	}
	steps, err := strconv.Atoi(line[2:])
	if err != nil {
		return move{}, fmt.Errorf("bad distance: %q: %w", line[2:], err)
	}
	return move{dir: dir, steps: steps}, nil
}

var dirMap = map[byte]direction{
	'U': dirUp,
	'D': dirDown,
	'L': dirLeft,
	'R': dirRight,
}

var dirStep = [...]util.P{
	dirUp:    {0, -1},
	dirDown:  {0, 1},
	dirLeft:  {-1, 0},
	dirRight: {1, 0},
}

// Solutions using a raw map[util.P]struct{}, for benchmarking:

func measureTailMap(moves []move) int {
	head, tail := util.P{0, 0}, util.P{0, 0}
	visited := map[util.P]struct{}{{0, 0}: {}}
	for _, m := range moves {
		d := dirStep[m.dir]
		for i := 0; i < m.steps; i++ {
			head = head.Add(d)
			tail = updateTail(head, tail)
			visited[tail] = struct{}{}
		}
	}
	return len(visited)
}

func measureLongTailMap(moves []move) int {
	const ropeLength = 10
	var rope [ropeLength]util.P
	visited := map[util.P]struct{}{{0, 0}: {}}
	for _, m := range moves {
		d := dirStep[m.dir]
		for i := 0; i < m.steps; i++ {
			rope[0] = rope[0].Add(d)
			for i := 1; i < ropeLength; i++ {
				rope[i] = updateTail(rope[i-1], rope[i])
			}
			visited[rope[ropeLength-1]] = struct{}{}
		}
	}
	return len(visited)
}
