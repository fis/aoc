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

// Package day23 solves AoC 2018 day 23.
package day23

import (
	"fmt"
	"math"
	"sort"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/ix"
)

func init() {
	glue.RegisterSolver(2018, 23, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	bots := []nanobot(nil)
	for _, line := range lines {
		bot, err := parseBot(line)
		if err != nil {
			return nil, err
		}
		bots = append(bots, bot)
	}

	part1 := inRange(bots)
	part2 := bestPos(bots)

	return glue.Ints(part1, part2), nil
}

type p3 struct {
	x, y, z int
}

type nanobot struct {
	p p3
	r int
}

func inRange(bots []nanobot) (count int) {
	max := nanobot{r: -1}
	for _, bot := range bots {
		if bot.r > max.r {
			max = bot
		}
	}
	for _, bot := range bots {
		if dist(max.p, bot.p) <= max.r {
			count++
		}
	}
	return count
}

func bestPos(bots []nanobot) (d int) {
	minP, maxP := bots[0].p, bots[0].p
	for _, bot := range bots[1:] {
		minP.x = min(minP.x, bot.p.x)
		maxP.x = max(maxP.x, bot.p.x)
		minP.y = min(minP.y, bot.p.y)
		maxP.y = max(maxP.y, bot.p.y)
		minP.z = min(minP.z, bot.p.z)
		maxP.z = max(maxP.z, bot.p.z)
	}
	d, _ = findBest(bots, minP, maxP, 0, -1, -1, 0)
	return d
}

const linearSearch = 4

type subCube struct {
	min, max p3
	bots     []nanobot
	baseN    int
}

var treeCache = [32]struct {
	bots  [8][]nanobot
	cubes [8]subCube
}{}

func findBest(bots []nanobot, min, max p3, baseN, boundD, boundN, level int) (bestD, bestN int) {
	if max.x-min.x < linearSearch && max.y-min.y < linearSearch && max.z-min.z < linearSearch {
		bestD, bestN = boundD, boundN
		for z := min.z; z <= max.z; z++ {
			for y := min.y; y <= max.y; y++ {
				for x := min.x; x <= max.x; x++ {
					n := baseN
					for _, bot := range bots {
						if dist(bot.p, p3{x, y, z}) <= bot.r {
							n++
						}
					}
					if n < bestN {
						continue
					}
					d := dist0(p3{x, y, z})
					if n > bestN || d < bestD {
						bestD, bestN = d, n
					}
				}
			}
		}
		return bestD, bestN
	}
	cubes := treeCache[level].cubes[:0]
	if max.x-min.x >= linearSearch {
		midX := min.x + (max.x-min.x)/2
		max1, min2 := p3{midX - 1, max.y, max.z}, p3{midX, min.y, min.z}
		cubes = append(cubes, subCube{min: min, max: max1}, subCube{min: min2, max: max})
	} else {
		cubes = append(cubes, subCube{min: min, max: max})
	}
	if max.y-min.y >= linearSearch {
		midY := min.y + (max.y-min.y)/2
		for i, l := 0, len(cubes); i < l; i++ {
			cmin, cmax := cubes[i].min, cubes[i].max
			max1, min2 := p3{cmax.x, midY - 1, cmax.z}, p3{cmin.x, midY, cmin.z}
			cubes[i] = subCube{min: cmin, max: max1}
			cubes = append(cubes, subCube{min: min2, max: cmax})
		}
	}
	if max.z-min.z >= linearSearch {
		midZ := min.z + (max.z-min.z)/2
		for i, l := 0, len(cubes); i < l; i++ {
			cmin, cmax := cubes[i].min, cubes[i].max
			max1, min2 := p3{cmax.x, cmax.y, midZ - 1}, p3{cmin.x, cmin.y, midZ}
			cubes[i] = subCube{min: cmin, max: max1}
			cubes = append(cubes, subCube{min: min2, max: cmax})
		}
	}
	for i := range cubes {
		cubes[i].bots, cubes[i].baseN = filterBots(bots, cubes[i].min, cubes[i].max, &treeCache[level].bots[i])
		cubes[i].baseN += baseN
	}
	sort.Slice(cubes, func(i, j int) bool {
		li, lj := cubes[i].baseN+len(cubes[i].bots), cubes[j].baseN+len(cubes[j].bots)
		if li == lj {
			di := minimumD(cubes[i].min, cubes[i].max, p3{0, 0, 0})
			dj := minimumD(cubes[j].min, cubes[j].max, p3{0, 0, 0})
			return di < dj
		}
		return li > lj
	})
	bestD, bestN = boundD, boundN
	for _, cube := range cubes {
		if cube.baseN+len(cube.bots) < bestN {
			break
		}
		minD := minimumD(cube.min, cube.max, p3{0, 0, 0})
		if cube.baseN+len(cube.bots) == bestN && minD >= bestD {
			break
		}
		bestD, bestN = findBest(cube.bots, cube.min, cube.max, cube.baseN, bestD, bestN, level+1)
	}
	return bestD, bestN
}

func filterBots(bots []nanobot, min, max p3, cache *[]nanobot) (possible []nanobot, known int) {
	possible = (*cache)[:0]
	for _, bot := range bots {
		if maximumD(min, max, bot.p) <= bot.r {
			known++
		} else if minimumD(min, max, bot.p) <= bot.r {
			possible = append(possible, bot)
		}
	}
	*cache = possible
	return possible, known
}

func minimumD(min, max, p p3) int {
	dx, dy, dz := 0, 0, 0
	if max.x < p.x {
		dx = p.x - max.x
	} else if min.x > p.x {
		dx = min.x - p.x
	}
	if max.y < p.y {
		dy = p.y - max.y
	} else if min.y > p.y {
		dy = min.y - p.y
	}
	if max.z < p.z {
		dz = p.z - max.z
	} else if min.z > p.z {
		dz = min.z - p.z
	}
	return dx + dy + dz
}

func maximumD(min, max, p p3) int {
	dx, dy, dz := ix.Abs(min.x-p.x), ix.Abs(min.y-p.y), ix.Abs(min.z-p.z)
	if d := ix.Abs(max.x - p.x); d > dx {
		dx = d
	}
	if d := ix.Abs(max.y - p.y); d > dy {
		dy = d
	}
	if d := ix.Abs(max.z - p.z); d > dz {
		dz = d
	}
	return dx + dy + dz
}

func parseBot(line string) (bot nanobot, err error) {
	if _, err := fmt.Sscanf(line, "pos=<%d,%d,%d>, r=%d", &bot.p.x, &bot.p.y, &bot.p.z, &bot.r); err != nil {
		return nanobot{}, err
	}
	return bot, nil
}

func dist(a, b p3) int {
	return ix.Abs(a.x-b.x) + ix.Abs(a.y-b.y) + ix.Abs(a.z-b.z)
}

func dist0(p p3) int {
	return ix.Abs(p.x) + ix.Abs(p.y) + ix.Abs(p.z)
}

// This code would solve part 2 of the puzzle using Chebyshev distance, accidentally.
// It's also very inefficient. It's left here for posterity and/or as a warning to others.

type overlapX struct {
	min, max int
	n        int
}

type overlapXY struct {
	min, max int
	laps     overlapSetX
}

type overlapXYZ struct {
	min, max int
	laps     overlapSetXY
}

type overlapSetX []overlapX
type overlapSetXY []overlapXY
type overlapSetXYZ []overlapXYZ

func emptySetX() overlapSetX {
	return overlapSetX{{min: math.MinInt64, max: math.MaxInt64, n: 0}}
}

func emptySetXY() overlapSetXY {
	return overlapSetXY{{min: math.MinInt64, max: math.MaxInt64, laps: emptySetX()}}
}

func emptySetXYZ() overlapSetXYZ {
	return overlapSetXYZ{{min: math.MinInt64, max: math.MaxInt64, laps: emptySetXY()}}
}

func (lap overlapXY) copy() overlapXY {
	lap.laps = append(overlapSetX(nil), lap.laps...)
	return lap
}

func (lap overlapXYZ) copy() overlapXYZ {
	old := lap.laps
	lap.laps = overlapSetXY(nil)
	for _, o := range old {
		lap.laps = append(lap.laps, o.copy())
	}
	return lap
}

func (os *overlapSetX) increment(min, max int) {
	hi := sort.Search(len(*os), func(i int) bool {
		return (*os)[i].min > max
	}) - 1
	lo := sort.Search(len(*os), func(i int) bool {
		return (*os)[i].max >= min
	})
	if (*os)[lo].min < min {
		lap := (*os)[lo]
		(*os)[lo].max, lap.min = min-1, min
		lo, hi = lo+1, hi+1
		os.insert(lo, lap)
	}
	if (*os)[hi].max > max {
		lap := (*os)[hi]
		(*os)[hi].max, lap.min = max, max+1
		os.insert(hi+1, lap)
	}
	for i := lo; i <= hi; i++ {
		(*os)[i].n++
	}
}

func (os *overlapSetXY) increment(min, max util.P) {
	hi := sort.Search(len(*os), func(i int) bool {
		return (*os)[i].min > max.Y
	}) - 1
	lo := sort.Search(len(*os), func(i int) bool {
		return (*os)[i].max >= min.Y
	})
	if (*os)[lo].min < min.Y {
		lap := (*os)[lo].copy()
		(*os)[lo].max, lap.min = min.Y-1, min.Y
		lo, hi = lo+1, hi+1
		os.insert(lo, lap)
	}
	if (*os)[hi].max > max.Y {
		lap := (*os)[hi].copy()
		(*os)[hi].max, lap.min = max.Y, max.Y+1
		os.insert(hi+1, lap)
	}
	for i := lo; i <= hi; i++ {
		(*os)[i].laps.increment(min.X, max.X)
	}
}

func (os *overlapSetXYZ) increment(min, max p3) {
	hi := sort.Search(len(*os), func(i int) bool {
		return (*os)[i].min > max.z
	}) - 1
	lo := sort.Search(len(*os), func(i int) bool {
		return (*os)[i].max >= min.z
	})
	if (*os)[lo].min < min.z {
		lap := (*os)[lo].copy()
		(*os)[lo].max, lap.min = min.z-1, min.z
		lo, hi = lo+1, hi+1
		os.insert(lo, lap)
	}
	if (*os)[hi].max > max.z {
		lap := (*os)[hi].copy()
		(*os)[hi].max, lap.min = max.z, max.z+1
		os.insert(hi+1, lap)
	}
	for i := lo; i <= hi; i++ {
		(*os)[i].laps.increment(util.P{min.x, min.y}, util.P{max.x, max.y})
	}
}

func (os *overlapSetX) insert(at int, lap overlapX) {
	if at == len(*os) {
		*os = append(*os, lap)
	} else {
		n := len(*os)
		*os = append(*os, (*os)[n-1])
		if end := n - 1; end > at {
			copy((*os)[at+1:end+1], (*os)[at:end])
		}
		(*os)[at] = lap
	}
}

func (os *overlapSetXY) insert(at int, lap overlapXY) {
	if at == len(*os) {
		*os = append(*os, lap)
	} else {
		n := len(*os)
		*os = append(*os, (*os)[n-1])
		if end := n - 1; end > at {
			copy((*os)[at+1:end+1], (*os)[at:end])
		}
		(*os)[at] = lap
	}
}

func (os *overlapSetXYZ) insert(at int, lap overlapXYZ) {
	if at == len(*os) {
		*os = append(*os, lap)
	} else {
		n := len(*os)
		*os = append(*os, (*os)[n-1])
		if end := n - 1; end > at {
			copy((*os)[at+1:end+1], (*os)[at:end])
		}
		(*os)[at] = lap
	}
}

func (os overlapSetX) findBest() (d, n int) {
	bestD, bestN := -1, -1
	for _, lap := range os {
		if lap.n < bestN {
			continue
		}
		d := closest(lap.min, lap.max)
		if lap.n > bestN || d < bestD {
			bestD, bestN = d, lap.n
		}
	}
	return bestD, bestN
}

func (os overlapSetXY) findBest() (d, n int) {
	bestD, bestN := -1, -1
	for _, lap := range os {
		xd, n := lap.laps.findBest()
		if n < bestN {
			continue
		}
		yd := closest(lap.min, lap.max)
		if n > bestN || xd+yd < bestD {
			bestD, bestN = xd+yd, n
		}
	}
	return bestD, bestN
}

func (os overlapSetXYZ) findBest() (d, n int) {
	bestD, bestN := -1, -1
	for _, lap := range os {
		xyd, n := lap.laps.findBest()
		if n < bestN {
			continue
		}
		zd := closest(lap.min, lap.max)
		if n > bestN || xyd+zd < bestD {
			bestD, bestN = xyd+zd, n
		}
	}
	return bestD, bestN
}

func closest(min, max int) int {
	if min > 0 {
		return min
	}
	if max < 0 {
		return -max
	}
	return 0
}
