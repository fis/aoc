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

// Package day24 solves AoC 2023 day 24.
package day24

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
	glue.RegisterSolver(2023, 24, glue.LineSolver(glue.WithParser(parseHailstone, solve)))
}

func solve(stones []hailstone) ([]string, error) {
	p1 := countIntersectXY(stones, 2e14, 2e14, 4e14, 4e14)
	p := findCollider(stones)
	p2 := p.x + p.y + p.z
	return glue.Ints(p1, p2), nil
}

func countIntersectXY(stones []hailstone, minX, minY, maxX, maxY float64) (intersecting int) {
	for i := 0; i < len(stones)-1; i++ {
		a := stones[i]
		pa, va := a.p.asFloat(), a.v.asFloat()
		atx1, atx2 := (minX-pa.x)/va.x, (maxX-pa.x)/va.x
		aty1, aty2 := (minY-pa.y)/va.y, (maxY-pa.y)/va.y
		minT, maxT := max(min(atx1, atx2), min(aty1, aty2)), min(max(atx1, atx2), max(aty1, aty2))
		if minT > maxT || maxT < 0 {
			continue // does not cross the target area, or does so completely the past
		}
		for j := i + 1; j < len(stones); j++ {
			b := stones[j]
			pb, vb := b.p.asFloat(), b.v.asFloat()
			ta := (pb.x*vb.y + pa.y*vb.x - pa.x*vb.y - pb.y*vb.x) / (va.x*vb.y - va.y*vb.x)
			if ta < 0 || ta < minT || ta > maxT {
				continue
			}
			tb := (pa.y - pb.y + ta*va.y) / vb.y
			if tb >= 0 {
				intersecting++
			}
		}
	}
	return intersecting
}

func findCollider(stones []hailstone) p3 {
	s0, s1, s2 := pickThree(stones)
	p0, v0 := s0.p.asFloat(), s0.v.asFloat()
	p1, v1 := s1.p.asFloat(), s1.v.asFloat()
	p2, v2 := s2.p.asFloat(), s2.v.asFloat()
	dp1, dv1 := p1.Sub(p0), v1.Sub(v0)
	dp2, dv2 := p2.Sub(p0), v2.Sub(v0)
	A := [6][7]float64{
		{0, -dp1.z, dp1.y, 0, dv1.z, -dv1.y, p1.y*v1.z - p1.z*v1.y - p0.y*v0.z + p0.z*v0.y},
		{dp1.z, 0, -dp1.x, -dv1.z, 0, dv1.x, p1.z*v1.x - p1.x*v1.z - p0.z*v0.x + p0.x*v0.z},
		{-dp1.y, dp1.x, 0, dv1.y, -dv1.x, 0, p1.x*v1.y - p1.y*v1.x - p0.x*v0.y + p0.y*v0.x},
		{0, -dp2.z, dp2.y, 0, dv2.z, -dv2.y, p2.y*v2.z - p2.z*v2.y - p0.y*v0.z + p0.z*v0.y},
		{dp2.z, 0, -dp2.x, -dv2.z, 0, dv2.x, p2.z*v2.x - p2.x*v2.z - p0.z*v0.x + p0.x*v0.z},
		{-dp2.y, dp2.x, 0, dv2.y, -dv2.x, 0, p2.x*v2.y - p2.y*v2.x - p0.x*v0.y + p0.y*v0.x},
	}
	reduce(&A)
	pRz := ri(A[5][6] / A[5][5])
	pRy := ri((A[4][6] - A[4][5]*float64(pRz)) / A[4][4])
	pRx := ri((A[3][6] - A[3][5]*float64(pRz) - A[3][4]*float64(pRy)) / A[3][3])
	return p3{pRx, pRy, pRz}
}

func ri(x float64) int {
	return int(math.Round(x))
}

func reduce(A *[6][7]float64) {
	// standard Gaussian elimination with partial pivoting
	h, k := 0, 0
	for h < 6 && k < 7 {
		maxV, maxI := math.Abs(A[h][k]), h
		for i := h + 1; i < 6; i++ {
			if v := math.Abs(A[i][k]); v > maxV {
				maxV, maxI = v, i
			}
		}
		if A[maxI][k] == 0 {
			k++
			continue
		}
		if maxI > h {
			A[h], A[maxI] = A[maxI], A[h]
		}
		for i := h + 1; i < 6; i++ {
			f := A[i][k] / A[h][k]
			A[i][k] = 0
			for j := k + 1; j < 7; j++ {
				A[i][j] -= A[h][j] * f
			}
		}
		h, k = h+1, k+1
	}
}

func pickThree(stones []hailstone) (a, b, c hailstone) {
	const eps = 1e-6

	ai, bi, ci := 0, -1, -1
	a = stones[0]
	av := a.v.asFloat()

	for i, s := range stones[ai+1:] {
		sv := s.v.asFloat()
		rx := sv.x / av.x
		ry := sv.y / av.y
		rz := sv.z / av.z
		if math.Abs(rx-ry) >= eps || math.Abs(rx-rz) >= eps {
			b, bi = s, i
			break
		}
	}
	if bi < 0 {
		panic("no good b")
	}

	bv := b.v.asFloat()
	q1y, q1z := bv.y/bv.x, bv.z/bv.x
	q2y, q2z := av.y-av.x*q1y, av.z-av.x*q1z
	for i, s := range stones[ai+1+bi+1:] {
		sv := s.v.asFloat()
		k1 := (sv.y - sv.x*q1y) / q2y
		k2 := (sv.z - sv.x*q1z) / q2z
		if math.Abs(k1-k2) >= eps {
			c, ci = s, i
			break
		}
	}
	if ci < 0 {
		panic("no good c")
	}

	return a, b, c
}

type hailstone struct {
	p, v p3
}

type p3 struct {
	x, y, z int
}

func (p p3) asFloat() fp3 {
	return fp3{float64(p.x), float64(p.y), float64(p.z)}
}

type fp3 struct {
	x, y, z float64
}

func (p fp3) Sub(q fp3) fp3 {
	return fp3{p.x - q.x, p.y - q.y, p.z - q.z}
}

func parseHailstone(line string) (hs hailstone, err error) {
	pos, vel, ok := strings.Cut(line, "@")
	if !ok {
		return hailstone{}, fmt.Errorf("TODO")
	}
	hs.p, err = parseP3(pos)
	if err != nil {
		return hailstone{}, err
	}
	hs.v, err = parseP3(vel)
	if err != nil {
		return hailstone{}, err
	}
	return hs, nil
}

func parseP3(line string) (p p3, err error) {
	s := util.Splitter(line)
	if s.Count(",") != 3 {
		return p3{}, fmt.Errorf("TODO")
	}
	xs, ys, zs := s.Next(","), s.Next(","), s.Next(",")
	x, errX := strconv.Atoi(strings.TrimSpace(xs))
	y, errY := strconv.Atoi(strings.TrimSpace(ys))
	z, errZ := strconv.Atoi(strings.TrimSpace(zs))
	if errX != nil || errY != nil || errZ != nil {
		return p3{}, errors.Join(errX, errY, errZ)
	}
	return p3{x, y, z}, nil
}
