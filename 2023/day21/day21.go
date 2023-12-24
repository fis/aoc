// Copyright 2023 Google LLC
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

// Package day21 solves AoC 2023 day 21.
package day21

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"os"
	"slices"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/fn"
	"golang.org/x/sys/unix"
)

func init() {
	glue.RegisterSolver(2023, 21, glue.FixedLevelSolver(solve))
	glue.RegisterPlotter(2023, 21, "", plotter{}, nil)
}

const target = 26501365 // 589 // 26501365

func solve(l *util.FixedLevel) ([]string, error) {
	const (
		smallSteps = 64
		bigSteps   = 26501365
	)
	mod := bigSteps % l.W
	reachable := countReachable(l, []int{smallSteps, mod, mod + l.W, mod + 2*l.W})
	p1 := reachable[0]
	p2 := interp(reachable[1:], bigSteps/l.W)
	return glue.Ints(p1, p2), nil
}

func interp(s []int, n int) int {
	a := (s[2] - 2*s[1] + s[0]) / 2
	b := s[1] - s[0] - a
	c := s[0]
	return a*n*n + b*n + c
}

func countReachable(l *util.FixedLevel, steps []int) (reachable []int) {
	size := l.W
	if l.H != size {
		panic("it's not hip to be unsquare")
	}
	if size&1 == 0 {
		panic("map isn't odd enough")
	}

	maxSteps := fn.Max(steps)
	reachable = make([]int, len(steps))
	for i, s := range steps {
		reachable[i] = fn.If(s > 0 && s&1 == 0, 1, 0)
	}

	rings := (maxSteps + size/2) / size
	rsize := 2*rings + 1
	if rsize*rsize > 64 {
		panic("that's a step too far")
	}

	seen := make([]uint64, size*size)
	type state struct {
		p util.P
		d int
	}
	fringe := util.MakeQueue[state](2048)

	startX, startY, ok := l.Find('S')
	if !ok {
		panic("where are you coming from with this?")
	}
	l.Set(startX, startY, '.')
	seen[startY*size+startX] |= 1 << (rings*rsize + rings)
	startX, startY = startX+rings*size, startY+rings*size

	fringe.Push(state{p: util.P{startX, startY}, d: 0})
	for !fringe.Empty() {
		p := fringe.Pop()
		for _, n := range p.p.Neigh() {
			rx, ry := n.X/size, n.Y/size
			ri := ry*rsize + rx
			gx, gy := n.X%size, n.Y%size
			gi := gy*size + gx
			if l.Data[gi] != '.' || seen[gi]&(1<<ri) != 0 {
				continue
			}
			seen[gi] |= 1 << ri
			nd := p.d + 1
			for i := len(steps) - 1; i >= 0; i-- {
				s := steps[i]
				if nd > s {
					break
				}
				if nd&1 == s&1 {
					reachable[i]++
				}
			}
			if nd < maxSteps {
				fringe.Push(state{p: n, d: nd})
			}
		}
	}

	return reachable
}

func countReachableDynamic(l *util.FixedLevel, steps []int) (reachable []int) {
	// This is effectively `countReachable`, except using a list of repetition indices
	// instead of a bitmap in order to avoid having any upper bound on steps.

	size := l.W

	maxSteps := fn.Max(steps)
	reachable = make([]int, len(steps))
	for i, s := range steps {
		reachable[i] = fn.If(s > 0 && s&1 == 0, 1, 0)
	}

	seen := make([][]util.P, size*size)
	type state struct {
		p util.P
		d int
	}
	fringe := util.MakeQueue[state](2048)

	startX, startY, _ := l.Find('S')
	l.Set(startX, startY, '.')
	seen[startY*size+startX] = append(seen[startY*size+startX], util.P{0, 0})

	fringe.Push(state{p: util.P{startX, startY}, d: 0})
	for !fringe.Empty() {
		p := fringe.Pop()
		for _, n := range p.p.Neigh() {
			gx, gy := (n.X%size+size)%size, (n.Y%size+size)%size
			rx, ry := (n.X-gx)/size, (n.Y-gy)/size
			gi := gy*size + gx
			if l.Data[gi] != '.' || slices.Contains(seen[gi], util.P{rx, ry}) {
				continue
			}
			seen[gi] = append(seen[gi], util.P{rx, ry})
			nd := p.d + 1
			for i, s := range steps {
				if nd <= s && nd&1 == s&1 {
					reachable[i]++
				}
			}
			if nd < maxSteps {
				fringe.Push(state{p: n, d: nd})
			}
		}
	}

	return reachable
}

// plotting

type plotter struct{}

func (plotter) Plot(r io.Reader, w io.Writer) error {
	const steps = 64

	if f, ok := w.(*os.File); ok {
		_, err := unix.IoctlGetTermios(int(f.Fd()), unix.TCGETS)
		if err == nil {
			return errors.New("refusing to output PNG to a terminal")
		}
	}

	input, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	l := util.ParseFixedLevel(input)
	x, y, _ := l.Find('S')
	type state struct {
		p util.P
		d int
	}
	if x < steps || y < steps || x+steps >= l.W || y+steps >= l.H {
		return fmt.Errorf("level too small: %dx%d", l.W, l.H)
	}
	q := util.MakeQueue[state](64)
	q.Push(state{p: util.P{x, y}, d: 0})
	for !q.Empty() {
		p := q.Pop()
		for _, n := range p.p.Neigh() {
			if l.At(n.X, n.Y) == '.' {
				l.Set(n.X, n.Y, 'O')
				if p.d+1 < steps {
					q.Push(state{p: n, d: p.d + 1})
				}
			}
		}
	}

	var (
		colorGarden  = color.NRGBA{R: 15, G: 157, B: 88, A: 255}
		colorRock    = color.NRGBA{R: 219, G: 68, B: 55, A: 255}
		colorVisited = color.NRGBA{R: 66, G: 133, B: 244, A: 255}
	)
	img := image.NewNRGBA(image.Rect(0, 0, 2*l.W, 2*l.H))
	for y := 0; y < l.H; y++ {
		for x := 0; x < l.W; x++ {
			var c color.NRGBA
			switch l.At(x, y) {
			case '.':
				c = colorGarden
			case '#':
				c = colorRock
			default:
				c = colorVisited
			}
			img.SetNRGBA(2*x, 2*y, c)
			img.SetNRGBA(2*x+1, 2*y, c)
			img.SetNRGBA(2*x, 2*y+1, c)
			img.SetNRGBA(2*x+1, 2*y+1, c)
		}
	}
	return png.Encode(w, img)
}
