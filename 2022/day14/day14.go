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

// Package day14 solves AoC 2022 day 14.
package day14

import (
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/fn"
)

func init() {
	glue.RegisterSolver(2022, 14, glue.LineSolver(glue.WithParser(parsePath, solve)))
}

func solve(paths [][]util.P) ([]string, error) {
	bmp, sourceX := buildBitmap(paths)
	p1 := addSand(bmp, sourceX)
	p2 := p1 + addShadow(bmp, sourceX)
	return glue.Ints(p1, p2), nil
}

func buildBitmap(paths [][]util.P) (bmp util.FixedBitmap2D, sourceX int) {
	floor := fn.MaxF(paths, func(path []util.P) int {
		return fn.MaxF(path, func(p util.P) int { return p.Y })
	}) + 2
	minX, maxX := 500-floor, 500+floor
	bmp = util.MakeFixedBitmap2D(maxX-minX+1, floor)
	for _, path := range paths {
		from := path[0]
		if from.X >= minX && from.X <= maxX {
			bmp.Set(from.X-minX, from.Y)
		}
		for _, to := range path {
			d := fn.If(to.X == from.X,
				fn.If(to.Y < from.Y, util.P{0, -1}, util.P{0, 1}),
				fn.If(to.X < from.X, util.P{-1, 0}, util.P{1, 0}))
			for from != to {
				from = from.Add(d)
				if from.X >= minX && from.X <= maxX {
					bmp.Set(from.X-minX, from.Y)
				}
			}
		}
	}
	return bmp, 500 - minX
}

func addSand(bmp util.FixedBitmap2D, sourceX int) (grains int) {
	_, h := bmp.Size()
	for {
		x, y := sourceX, 0
	fall:
		for {
			switch {
			case y+1 == h:
				return grains
			case !bmp.Get(x, y+1):
				y = y + 1
			case !bmp.Get(x-1, y+1):
				x, y = x-1, y+1
			case !bmp.Get(x+1, y+1):
				x, y = x+1, y+1
			default:
				break fall
			}
		}
		grains++
		bmp.Set(x, y)
	}
}

func addShadow(bmp util.FixedBitmap2D, sourceX int) (grains int) {
	_, h := bmp.Size()
	shadow := 0
	for y := 1; y < h; y++ {
		x0, x1 := sourceX-y, sourceX+y
		for x := x0; x <= x1; x++ {
			if bmp.Get(x, y) {
				shadow++
			} else if bmp.Get(x-1, y-1) && bmp.Get(x, y-1) && bmp.Get(x+1, y-1) {
				bmp.Set(x, y)
				shadow++
			}
		}
	}
	return h*h - shadow
}

func parsePath(line string) (path []util.P, err error) {
	parts := strings.Split(line, " -> ")
	path = make([]util.P, len(parts))
	for i, part := range parts {
		p, err := util.ParseP(part)
		if err != nil {
			return nil, err
		}
		path[i] = p
	}
	return path, nil
}
