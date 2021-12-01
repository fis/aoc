// Copyright 2021 Google LLC
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

// Package day22 solves AoC 2017 day 22.
package day22

import (
	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2017, 22, glue.LevelSolver{
		Solver: solve,
		Empty:  '.',
	})
}

func solve(l *util.Level) ([]string, error) {
	p1 := simulateSimple(l.Copy(), 10000)
	p2 := simulateEvolvedArray(l, 10000000)
	return glue.Ints(p1, p2), nil
}

func simulateSimple(l *util.Level, rounds int) (infected int) {
	minP, maxP := l.Bounds()
	at, d := util.P{(minP.X + maxP.X) / 2, (minP.Y + maxP.Y) / 2}, util.P{0, -1}
	for round := 0; round < rounds; round++ {
		if l.At(at.X, at.Y) == '#' {
			d = util.P{-d.Y, d.X}
			l.Set(at.X, at.Y, '.')
		} else {
			d = util.P{d.Y, -d.X}
			l.Set(at.X, at.Y, '#')
			infected++
		}
		at = at.Add(d)
	}
	return infected
}

func simulateEvolvedLevel(l *util.Level, rounds int) (infected int) {
	minP, maxP := l.Bounds()
	at, d := util.P{(minP.X + maxP.X) / 2, (minP.Y + maxP.Y) / 2}, util.P{0, -1}
	for round := 0; round < rounds; round++ {
		switch c := l.At(at.X, at.Y); c {
		case '.':
			d = util.P{d.Y, -d.X}
			l.Set(at.X, at.Y, 'W')
		case 'W':
			l.Set(at.X, at.Y, '#')
			infected++
		case '#':
			d = util.P{-d.Y, d.X}
			l.Set(at.X, at.Y, 'F')
		case 'F':
			d = util.P{-d.X, -d.Y}
			l.Set(at.X, at.Y, '.')
		}
		at = at.Add(d)
	}
	return infected
}

func simulateEvolvedArray(l *util.Level, rounds int) (infected int) {
	const (
		X0    = -256
		Y0    = -185
		W     = 512
		Wbits = 9
		H     = 400
	)
	arr := [W * H]byte{}

	minP, maxP := l.Bounds()
	for y := minP.Y; y <= maxP.Y; y++ {
		for x := minP.X; x <= maxP.X; x++ {
			if l.At(x, y) == '#' {
				arr[((y-Y0)<<Wbits)+(x-X0)] = 2
			}
		}
	}
	atX, atY, dX, dY := (minP.X+maxP.X)/2, (minP.Y+maxP.Y)/2, 0, -1

	for round := 0; round < rounds; round++ {
		atI := ((atY - Y0) << Wbits) + (atX - X0)
		switch c := arr[atI]; c {
		case 0:
			dX, dY = dY, -dX
			arr[atI] = 1
		case 1:
			arr[atI] = 2
			infected++
		case 2:
			dX, dY = -dY, dX
			arr[atI] = 3
		case 3:
			dX, dY = -dX, -dY
			arr[atI] = 0
		}
		atX += dX
		atY += dY
	}
	return infected
}
