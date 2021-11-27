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

// Package day20 solves AoC 2018 day 20.
package day20

import (
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2018, 20, glue.LineSolver(solve))
}

func solve(lines []string) ([]int, error) {
	if len(lines) != 1 {
		return nil, fmt.Errorf("expected one line, got %d", len(lines))
	}
	ex, err := parseDirex(lines[0])
	if err != nil {
		return nil, err
	}

	l := make(layout)
	ex.trace([]util.P{{0, 0}}, l)
	part1, part2 := l.radius()

	return []int{part1, part2}, nil
}

type direction int

const (
	north direction = iota
	west
	south
	east
)

var dirV = [...]util.P{north: {0, -1}, west: {-1, 0}, south: {0, 1}, east: {1, 0}}

func (d direction) rev() direction       { return (d + 2) % 4 }
func (d direction) move(p util.P) util.P { return p.Add(dirV[d]) }

// layout representation: a point will exist in the map iff a room exists, and the value of the entry
// will be a bitmap of directions, with the bit set iff there is a door to that direction from the room.
type layout map[util.P]uint

func (l layout) String() string {
	sb := strings.Builder{}
	l.Write(&sb)
	return sb.String()
}

func (l layout) Write(w io.Writer) {
	level := l.asLevel()
	min, max := level.Bounds()
	level.WriteRect(w, util.P{min.X - 1, min.Y - 1}, util.P{max.X + 1, max.Y + 1})
}

func (l layout) asLevel() (level *util.Level) {
	level = util.ParseLevelString("X", '#')
	for p, doors := range l {
		pp := p.Scale(2)
		level.Set(pp.X, pp.Y, '.')
		for d := north; d <= east; d++ {
			if (doors & (uint(1) << d)) != 0 {
				dp := d.move(pp)
				level.Set(dp.X, dp.Y, "-|"[d%2])
			}
		}
	}
	level.Set(0, 0, 'X')
	return level
}

func (l layout) trace(from util.P, d direction) {
	l[from] |= uint(1) << d
	l[d.move(from)] |= uint(1) << d.rev()
}

func (l layout) radius() (r int, far int) {
	type distP struct {
		p util.P
		d int
	}
	fringe, seen := []distP{{p: util.P{0, 0}, d: 0}}, map[util.P]struct{}{{0, 0}: {}}
	for len(fringe) > 0 {
		at := fringe[0]
		fringe = fringe[1:]
		seen[at.p] = struct{}{}
		if at.d > r {
			r = at.d
		}
		if at.d >= 1000 {
			far++
		}
		for doors, d := l[at.p], north; d <= east; d++ {
			if (doors & (uint(1) << d)) != 0 {
				to := d.move(at.p)
				if _, ok := seen[to]; !ok {
					fringe = append(fringe, distP{p: to, d: at.d + 1})
				}
			}
		}
	}
	return r, far
}

type direxOp int

const (
	opLit direxOp = iota
	opCat
	opAlt
)

type direx struct {
	op  direxOp
	lit []direction
	sub []*direx
}

func (ex *direx) trace(from []util.P, l layout) (end []util.P) {
	switch ex.op {
	case opLit:
		for _, at := range from {
			for _, d := range ex.lit {
				l.trace(at, d)
				at = d.move(at)
			}
			end = append(end, at)
		}
		return end
	case opCat:
		end = from
		for _, s := range ex.sub {
			end = s.trace(end, l)
		}
		return end
	case opAlt:
		for _, s := range ex.sub {
			end = append(end, s.trace(from, l)...)
		}
		return dedup(end)
	}
	return nil
}

func dedup(ps []util.P) []util.P {
	if len(ps) == 0 {
		return nil
	}
	sort.Slice(ps, func(i, j int) bool {
		return ps[i].Y < ps[j].Y || (ps[i].Y == ps[j].Y && ps[i].X < ps[j].X)
	})
	at := 0
	for i, p := range ps {
		if at == 0 || ps[at-1].X != p.X || ps[at-1].Y != p.Y {
			if at < i {
				ps[at] = p
			}
			at++
		}
	}
	return ps[:at]
}

func (ex *direx) String() string {
	if ex.op == opLit {
		return fmt.Sprintf("lit:%v", ex.lit)
	}
	sub := []string(nil)
	for _, s := range ex.sub {
		sub = append(sub, s.String())
	}
	switch ex.op {
	case opCat:
		return fmt.Sprintf("cat:%v", sub)
	case opAlt:
		return fmt.Sprintf("alt:%v", sub)
	default:
		return "bad"
	}
}

func parseDirex(in string) (*direx, error) {
	if !strings.HasPrefix(in, "^") || !strings.HasSuffix(in, "$") {
		return nil, fmt.Errorf("expected ^...$")
	}
	ex, tail, err := parseDirexAlt(in[1 : len(in)-1])
	if err != nil {
		return nil, err
	} else if tail != "" {
		return nil, fmt.Errorf("junk after expression")
	}
	return ex, nil
}

func parseDirexAlt(in string) (ex *direx, tail string, err error) {
	ex, in, err = parseDirexCat(in)
	if err != nil {
		return nil, "", err
	}
	if len(in) == 0 || in[0] != '|' {
		return ex, in, nil
	}
	sub := []*direx{ex}
	for len(in) > 0 && in[0] == '|' {
		ex, in, err = parseDirexCat(in[1:])
		if err != nil {
			return nil, "", err
		}
		sub = append(sub, ex)
	}
	return &direx{op: opAlt, sub: sub}, in, nil
}

func parseDirexCat(in string) (ex *direx, tail string, err error) {
	sub := []*direx(nil)
	for {
		ex, in, err = parseDirexAtom(in)
		if err != nil {
			return nil, "", err
		} else if ex == nil {
			break
		}
		sub = append(sub, ex)
	}
	switch len(sub) {
	case 0:
		return &direx{op: opLit}, in, nil
	case 1:
		return sub[0], in, nil
	default:
		return &direx{op: opCat, sub: sub}, in, nil
	}
}

func parseDirexAtom(in string) (ex *direx, tail string, err error) {
	switch {
	case len(in) == 0 || in[0] == ')' || in[0] == '|':
		return nil, in, nil
	case isDir(in[0]):
		lit := []direction(nil)
		for len(in) > 0 && isDir(in[0]) {
			lit = append(lit, asDir(in[0]))
			in = in[1:]
		}
		return &direx{op: opLit, lit: lit}, in, nil
	case in[0] == '(':
		ex, in, err = parseDirexAlt(in[1:])
		if err != nil {
			return nil, "", err
		}
		if len(in) == 0 || in[0] != ')' {
			return nil, "", fmt.Errorf("expected ), got %q", in)
		}
		return ex, in[1:], nil
	default:
		return nil, "", fmt.Errorf("expected direction or ( or ) or |")
	}
}

func isDir(b byte) bool {
	return b == 'N' || b == 'W' || b == 'S' || b == 'E'
}

func asDir(b byte) direction {
	switch b {
	case 'N':
		return north
	case 'W':
		return west
	case 'S':
		return south
	case 'E':
		return east
	}
	return direction(-1)
}
