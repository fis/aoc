// Copyright 2020 Google LLC
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

// Package day13 solves AoC 2018 day 13.
package day13

import (
	"fmt"
	"io"
	"io/ioutil"
	"sort"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2018, 13, glue.GenericSolver(solve))
}

func solve(r io.Reader) ([]string, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	level := util.ParseLevel(data, ' ')

	crashX, crashY := simulate(level.Copy(), false)
	lastX, lastY := simulate(level.Copy(), true)

	return []string{
		fmt.Sprintf("%d,%d", crashX, crashY),
		fmt.Sprintf("%d,%d", lastX, lastY),
	}, nil
}

type cart struct {
	pos, dir util.P
	state    int
}

func simulate(level *util.Level, keepGoing bool) (outX, outY int) {
	carts, occupied := []*cart(nil), map[util.P]struct{}{}
	level.Range(func(x, y int, b byte) {
		switch b {
		case '^':
			carts = append(carts, &cart{pos: util.P{x, y}, dir: util.P{0, -1}})
		case 'v':
			carts = append(carts, &cart{pos: util.P{x, y}, dir: util.P{0, 1}})
		case '<':
			carts = append(carts, &cart{pos: util.P{x, y}, dir: util.P{-1, 0}})
		case '>':
			carts = append(carts, &cart{pos: util.P{x, y}, dir: util.P{1, 0}})
		}
	})
	for _, c := range carts {
		track := byte('-')
		if c.dir.Y != 0 {
			track = '|'
		}
		level.Set(c.pos.X, c.pos.Y, track)
		occupied[c.pos] = struct{}{}
	}

	for {
		sort.Slice(carts, func(i, j int) bool { return lessC(carts[i], carts[j]) })
		for len(carts) > 0 && carts[len(carts)-1] == nil {
			carts = carts[:len(carts)-1]
		}

		if len(carts) == 0 {
			panic("all outta carts")
		} else if len(carts) == 1 {
			return carts[0].pos.X, carts[0].pos.Y
		}

		for _, c := range carts {
			if c == nil {
				continue
			}

			delete(occupied, c.pos)
			c.pos = c.pos.Add(c.dir)
			if _, ok := occupied[c.pos]; ok {
				if !keepGoing {
					return c.pos.X, c.pos.Y
				}
				delete(occupied, c.pos)
				for i := range carts {
					if carts[i] != nil && carts[i].pos.X == c.pos.X && carts[i].pos.Y == c.pos.Y {
						carts[i] = nil
					}
				}
				continue
			}
			occupied[c.pos] = struct{}{}

			switch level.At(c.pos.X, c.pos.Y) {
			case '+':
				switch c.state {
				case 0:
					c.dir.X, c.dir.Y = c.dir.Y, -c.dir.X
				case 1:
					// no action
				case 2:
					c.dir.X, c.dir.Y = -c.dir.Y, c.dir.X
				}
				c.state = (c.state + 1) % 3
			case '/':
				c.dir.X, c.dir.Y = -c.dir.Y, -c.dir.X
			case '\\':
				c.dir.X, c.dir.Y = c.dir.Y, c.dir.X
			case '-', '|':
				// no action
			default:
				panic("cart derailed")
			}
		}
	}
}

func lessC(a, b *cart) bool {
	if a == nil || b == nil {
		return b == nil
	}
	return a.pos.Y < b.pos.Y || (a.pos.Y == b.pos.Y && a.pos.X < b.pos.X)
}
