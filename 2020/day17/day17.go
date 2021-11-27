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

// Package day17 solves AoC 2020 day 17.
package day17

import (
	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2020, 17, glue.LevelSolver{Solver: solve, Empty: '.'})
}

func solve(level *util.Level) ([]int, error) {
	state3, min3, max3 := loadLevel3(level)
	for i := 0; i < 6; i++ {
		state3, min3, max3 = cycle3(state3, min3, max3)
	}
	state4, min4, max4 := loadLevel4(level)
	for i := 0; i < 6; i++ {
		state4, min4, max4 = cycle4(state4, min4, max4)
	}
	return []int{len(state3), len(state4)}, nil
}

// P3 is a simple three-dimensional coordinate point.
type P3 struct{ X, Y, Z int }

// P4 is a simple four-dimensional coordinate point.
type P4 struct{ X, Y, Z, W int }

func loadLevel3(level *util.Level) (out map[P3]struct{}, outMin, outMax P3) {
	out = make(map[P3]struct{})
	min, max := level.Bounds()
	for y := min.Y; y <= max.Y; y++ {
		for x := min.X; x <= max.X; x++ {
			if level.At(x, y) == '#' {
				out[P3{x, y, 0}] = struct{}{}
			}
		}
	}
	return out, P3{min.X, min.Y, 0}, P3{max.X, max.Y, 0}
}

func loadLevel4(level *util.Level) (out map[P4]struct{}, outMin, outMax P4) {
	out = make(map[P4]struct{})
	min, max := level.Bounds()
	for y := min.Y; y <= max.Y; y++ {
		for x := min.X; x <= max.X; x++ {
			if level.At(x, y) == '#' {
				out[P4{x, y, 0, 0}] = struct{}{}
			}
		}
	}
	return out, P4{min.X, min.Y, 0, 0}, P4{max.X, max.Y, 0, 0}
}

func cycle3(in map[P3]struct{}, min, max P3) (out map[P3]struct{}, newMin, newMax P3) {
	out = make(map[P3]struct{})
	newMin, newMax = min, max
	for z := min.Z - 1; z <= max.Z+1; z++ {
		for y := min.Y - 1; y <= max.Y+1; y++ {
			for x := min.X - 1; x <= max.X+1; x++ {
				p := P3{x, y, z}
				active, count := false, 0
				if _, ok := in[p]; ok {
					active, count = true, -1
				}
				for nz := z - 1; nz <= z+1; nz++ {
					for ny := y - 1; ny <= y+1; ny++ {
						for nx := x - 1; nx <= x+1; nx++ {
							if _, ok := in[P3{nx, ny, nz}]; ok {
								count++
							}
						}
					}
				}
				if count == 3 || (active && count == 2) {
					out[p] = struct{}{}
					newMin.X = imin(newMin.X, x)
					newMax.X = imax(newMax.X, x)
					newMin.Y = imin(newMin.Y, y)
					newMax.Y = imax(newMax.Y, y)
					newMin.Z = imin(newMin.Z, z)
					newMax.Z = imax(newMax.Z, z)
				}
			}
		}
	}
	return out, newMin, newMax
}

func cycle4(in map[P4]struct{}, min, max P4) (out map[P4]struct{}, newMin, newMax P4) {
	out = make(map[P4]struct{})
	newMin, newMax = min, max
	for w := min.W - 1; w <= max.W+1; w++ {
		for z := min.Z - 1; z <= max.Z+1; z++ {
			for y := min.Y - 1; y <= max.Y+1; y++ {
				for x := min.X - 1; x <= max.X+1; x++ {
					p := P4{x, y, z, w}
					active, count := false, 0
					if _, ok := in[p]; ok {
						active, count = true, -1
					}
					for nw := w - 1; nw <= w+1; nw++ {
						for nz := z - 1; nz <= z+1; nz++ {
							for ny := y - 1; ny <= y+1; ny++ {
								for nx := x - 1; nx <= x+1; nx++ {
									if _, ok := in[P4{nx, ny, nz, nw}]; ok {
										count++
									}
								}
							}
						}
					}
					if count == 3 || (active && count == 2) {
						out[p] = struct{}{}
						newMin.X = imin(newMin.X, x)
						newMax.X = imax(newMax.X, x)
						newMin.Y = imin(newMin.Y, y)
						newMax.Y = imax(newMax.Y, y)
						newMin.Z = imin(newMin.Z, z)
						newMax.Z = imax(newMax.Z, z)
						newMin.W = imin(newMin.W, w)
						newMax.W = imax(newMax.W, w)
					}
				}
			}
		}
	}
	return out, newMin, newMax
}

func imin(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func imax(a, b int) int {
	if b > a {
		return b
	}
	return a
}
