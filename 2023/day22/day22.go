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

// Package day22 solves AoC 2023 day 22.
package day22

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2023, 22, glue.LineSolver(glue.WithParser(parseBrick, solve)))
}

func solve(bricks []brick) ([]string, error) {
	p1, p2 := disintegrate(bricks)
	return glue.Ints(p1, p2), nil
}

const chuteSize = 10

func disintegrate(bricks []brick) (safe, fallen int) {
	util.SortBy(bricks, func(b brick) uint16 { return b[0].z })
	succ := drop(bricks)

	seen := util.MakeFixedBitmap1D(len(bricks) + 1)
	for i := 1; i <= len(bricks); i++ {
		seen.Clear()
		r := reachableWithout(succ, uint16(i), 0, seen)
		if r == len(bricks) {
			safe++
		} else {
			fallen += len(bricks) - r
		}
	}

	return safe, fallen
}

func reachableWithout(succ [][]uint16, removed uint16, at uint16, seen util.FixedBitmap1D) (reachable int) {
	reachable = 1
	for _, b := range succ[at] {
		if b != removed && !seen.Get(int(b)) {
			seen.Set(int(b))
			reachable += reachableWithout(succ, removed, b, seen)
		}
	}
	return reachable
}

func drop(bricks []brick) (succ [][]uint16) {
	succ = make([][]uint16, len(bricks)+1)
	var top [chuteSize][chuteSize]struct {
		z uint16
		b uint16
	}
	for i := range bricks {
		b := &bricks[i]
		switch {
		case b[0].x == b[1].x && b[0].y == b[1].y: // vertical brick or single cube
			x, y := b[0].x, b[0].y
			if d := b[0].z - top[y][x].z; d > 0 {
				b[0].z, b[1].z = b[0].z-d, b[1].z-d
			}
			succ[top[y][x].b] = append(succ[top[y][x].b], uint16(i+1))
			top[y][x].z = b[1].z + 1
			top[y][x].b = uint16(i + 1)
		case b[1].x > b[0].x: // horizontal brick with some X extent
			x, y, n := b[0].x, b[0].y, b[1].x-b[0].x+1
			minD := uint16(math.MaxUint16)
			for j := byte(0); j < n; j++ {
				if d := b[0].z - top[y][x+j].z; d < minD {
					minD = d
				}
			}
			if minD > 0 {
				b[0].z, b[1].z = b[0].z-minD, b[1].z-minD
			}
			for j, pb := byte(0), uint16(0xffff); j < n; j++ {
				if tz, tb := top[y][x+j].z, top[y][x+j].b; tz == b[0].z && tb != pb {
					pb = tb
					succ[tb] = append(succ[tb], uint16(i+1))
				}
				top[y][x+j].z = b[1].z + 1
				top[y][x+j].b = uint16(i + 1)
			}
		default: // horizontal brick with some Y extent
			x, y, n := b[0].x, b[0].y, b[1].y-b[0].y+1
			minD := uint16(math.MaxUint16)
			for j := byte(0); j < n; j++ {
				if d := b[0].z - top[y+j][x].z; d < minD {
					minD = d
				}
			}
			if minD > 0 {
				b[0].z, b[1].z = b[0].z-minD, b[1].z-minD
			}
			for j, pb := byte(0), uint16(0xffff); j < n; j++ {
				if tz, tb := top[y+j][x].z, top[y+j][x].b; tz == b[0].z && tb != pb {
					pb = tb
					succ[tb] = append(succ[tb], uint16(i+1))
				}
				top[y+j][x].z = b[1].z + 1
				top[y+j][x].b = uint16(i + 1)
			}
		}
	}
	return succ
}

type brick [2]p3

type p3 struct {
	x, y byte
	z    uint16
}

func parseBrick(line string) (b brick, err error) {
	lo, hi, ok := strings.Cut(line, "~")
	if !ok {
		return brick{}, fmt.Errorf("missing ~ in brick: %s", line)
	}
	for i, text := range []string{lo, hi} {
		s := util.Splitter(text)
		if s.Count(",") != 3 {
			return brick{}, fmt.Errorf("expected coordinate triple, got: %s", text)
		}
		x, y, z := s.Next(","), s.Next(","), s.Next(",")
		xi, errX := strconv.Atoi(x)
		yi, errY := strconv.Atoi(y)
		zi, errZ := strconv.Atoi(z)
		if errX != nil || errY != nil || errZ != nil {
			return brick{}, fmt.Errorf("bad coordinate: %w", errors.Join(errX, errY, errZ))
		}
		b[i].x, b[i].y, b[i].z = byte(xi), byte(yi), uint16(zi)
	}
	return b, nil
}
