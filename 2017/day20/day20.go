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

// Package day20 solves AoC 2017 day 20.
package day20

import (
	"math"
	"sort"
	"strconv"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util/ix"
)

const inputRegexp = `^p=<(-?\d+),(-?\d+),(-?\d+)>, v=<(-?\d+),(-?\d+),(-?\d+)>, a=<(-?\d+),(-?\d+),(-?\d+)>$`

func init() {
	glue.RegisterSolver(2017, 20, glue.RegexpSolver{
		Solver: solve,
		Regexp: inputRegexp,
	})
}

func solve(input [][]string) ([]string, error) {
	ps := parseInput(input)
	p1 := closest(ps)
	p2 := collideSim(ps, 40)
	return glue.Ints(p1, p2), nil
}

type p3 struct {
	x, y, z int
}

type particle struct {
	p, v, a p3
}

func closest(ps []particle) (minI int) {
	minP, minV, minA := math.MaxInt, math.MaxInt, math.MaxInt
	for i, p := range ps {
		dP, dV, dA := p.p.dist(), p.v.dist(), p.a.dist()
		if dA < minA || (dA == minA && dV < minV) || (dA == minA && dV == minV && dP < minP) {
			minI, minP, minV, minA = i, dP, dV, dA
		}
	}
	return minI
}

func collideSim(ps []particle, rounds int) int {
	for t := 0; t < rounds; t++ {
		positions := map[p3][]int{}
		for i, p := range ps {
			if p == (particle{}) {
				continue
			}
			at := p.pos(t)
			positions[at] = append(positions[at], i)
		}
		for _, list := range positions {
			if len(list) < 2 {
				continue
			}
			for _, i := range list {
				ps[i] = particle{}
			}
		}
	}
	surviving := 0
	for _, p := range ps {
		if p != (particle{}) {
			surviving++
		}
	}
	return surviving
}

func collideCalc(ps []particle) int {
	type collisionKey struct {
		t int
		p p3
	}

	collisions := map[collisionKey][]int{}
	collisionKeys := []collisionKey(nil)
	for i, N := 0, len(ps); i < N-1; i++ {
		for j := i + 1; j < N; j++ {
			if t := collide3D(ps[i], ps[j]); t >= 0 {
				key := collisionKey{t: t, p: ps[i].pos(t)}
				if _, old := collisions[key]; !old {
					collisionKeys = append(collisionKeys, key)
				}
				collisions[key] = insert(collisions[key], i, j)
			}
		}
	}

	sort.Slice(collisionKeys, func(i, j int) bool {
		return collisionKeys[i].t < collisionKeys[j].t
	})
	collided := make([]bool, len(ps))
	survived := len(ps)
	for _, key := range collisionKeys {
		size := 0
		for _, p := range collisions[key] {
			if !collided[p] {
				size++
			}
		}
		if size >= 2 {
			for _, p := range collisions[key] {
				if !collided[p] {
					collided[p] = true
					survived--
				}
			}
		}
	}

	return survived
}

func insert(list []int, a, b int) []int {
	foundA, foundB := false, false
	for _, i := range list {
		if i == a {
			foundA = true
		}
		if i == b {
			foundB = true
		}
	}
	if !foundA {
		list = append(list, a)
	}
	if !foundB {
		list = append(list, b)
	}
	return list
}

func collide3D(p1, p2 particle) (t int) {
	t1x, t2x, allX := collide1D(p1.p.x, p1.v.x, p1.a.x, p2.p.x, p2.v.x, p2.a.x)
	if t1x < 0 && !allX {
		return -1
	}
	t1y, t2y, allY := collide1D(p1.p.y, p1.v.y, p1.a.y, p2.p.y, p2.v.y, p2.a.y)
	if t1y < 0 && !allY {
		return -1
	}
	t1z, t2z, allZ := collide1D(p1.p.z, p1.v.z, p1.a.z, p2.p.z, p2.v.z, p2.a.z)
	if t1z < 0 && !allZ {
		return -1
	}
	t1, t2, all := combine(t1x, t2x, t1y, t2y, allX, allY)
	t1, t2, _ = combine(t1, t2, t1z, t2z, all, allZ)
	if t1 >= 0 && t2 >= 0 {
		if t1 < t2 {
			return t1
		} else {
			return t2
		}
	} else if t1 >= 0 {
		return t1
	} else if t2 >= 0 {
		return t2
	}
	return -1
}

func collide1D(p1, v1, a1, p2, v2, a2 int) (t1, t2 int, all bool) {
	pd, vd, ad := p1-p2, v1-v2, a1-a2
	A, B, C := ad, 2*vd+ad, 2*pd
	if A == 0 && B == 0 {
		return -1, -1, C == 0
	} else if A == 0 {
		if C%B == 0 && -C/B >= 0 {
			return -C / B, -1, false
		} else {
			return -1, -1, false
		}
	}
	disc := B*B - 4*A*C
	if disc < 0 {
		return -1, -1, false
	} else if disc == 0 {
		if B%(2*A) == 0 && -B/(2*A) >= 0 {
			return -B / (2 * A), -1, false
		} else {
			return -1, -1, false
		}
	}
	sqDisc := ix.Sqrt(disc)
	if sqDisc*sqDisc != disc {
		return -1, -1, false
	}
	i1 := (-B-sqDisc)%(2*A) == 0 && (-B-sqDisc)/(2*A) >= 0
	i2 := (-B+sqDisc)%(2*A) == 0 && (-B+sqDisc)/(2*A) >= 0
	if i1 && i2 {
		return (-B - sqDisc) / (2 * A), (-B + sqDisc) / (2 * A), false
	} else if i1 {
		return (-B - sqDisc) / (2 * A), -1, false
	} else if i2 {
		return (-B + sqDisc) / (2 * A), -1, false
	}
	return -1, -1, false
}

func combine(t1a, t2a, t1b, t2b int, allA, allB bool) (t1, t2 int, all bool) {
	if allA && allB {
		return -1, -1, true
	} else if allA {
		return t1b, t2b, false
	} else if allB {
		return t1a, t2a, false
	}
	if (t1a == t1b && t2a == t2b) || (t1a == t2b && t2a == t1b) {
		return t1a, t2a, false
	}
	if t1a == t1b || t1a == t2b {
		return t1a, -1, false
	}
	if t2a == t1b || t2a == t2b {
		return t2a, -1, false
	}
	return -1, -1, false
}

func parseInput(input [][]string) []particle {
	ps := make([]particle, len(input))
	for i, row := range input {
		ps[i].p.x, _ = strconv.Atoi(row[0])
		ps[i].p.y, _ = strconv.Atoi(row[1])
		ps[i].p.z, _ = strconv.Atoi(row[2])
		ps[i].v.x, _ = strconv.Atoi(row[3])
		ps[i].v.y, _ = strconv.Atoi(row[4])
		ps[i].v.z, _ = strconv.Atoi(row[5])
		ps[i].a.x, _ = strconv.Atoi(row[6])
		ps[i].a.y, _ = strconv.Atoi(row[7])
		ps[i].a.z, _ = strconv.Atoi(row[8])
	}
	return ps
}

func (p particle) pos(t int) p3 {
	// p(t) = p0 + t*v0 + (t+1)*t/2 * a0
	tt := (t + 1) * t / 2
	return p3{
		x: p.p.x + t*p.v.x + tt*p.a.x,
		y: p.p.y + t*p.v.y + tt*p.a.y,
		z: p.p.z + t*p.v.z + tt*p.a.z,
	}
}

func (p p3) dist() int {
	return ix.Abs(p.x) + ix.Abs(p.y) + ix.Abs(p.z)
}
