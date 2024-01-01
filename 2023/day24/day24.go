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
	"math/big"
	"strconv"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/ix"
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
			ta := ((pb.x-pa.x)*vb.y - (pb.y-pa.y)*vb.x) / (va.x*vb.y - va.y*vb.x)
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
	p0, v0, p1, v1, p2, v2 := s0.p, s0.v, s1.p, s1.v, s2.p, s2.v
	dp1, dv1, dp2, dv2 := p1.Sub(p0), v1.Sub(v0), p2.Sub(p0), v2.Sub(v0)
	var A [6][7]*big.Rat
	setRow(&A[0], 0, -dp1.z, dp1.y, 0, dv1.z, -dv1.y, p1.y*v1.z-p1.z*v1.y-p0.y*v0.z+p0.z*v0.y)
	setRow(&A[1], dp1.z, 0, -dp1.x, -dv1.z, 0, dv1.x, p1.z*v1.x-p1.x*v1.z-p0.z*v0.x+p0.x*v0.z)
	setRow(&A[2], -dp1.y, dp1.x, 0, dv1.y, -dv1.x, 0, p1.x*v1.y-p1.y*v1.x-p0.x*v0.y+p0.y*v0.x)
	setRow(&A[3], 0, -dp2.z, dp2.y, 0, dv2.z, -dv2.y, p2.y*v2.z-p2.z*v2.y-p0.y*v0.z+p0.z*v0.y)
	setRow(&A[4], dp2.z, 0, -dp2.x, -dv2.z, 0, dv2.x, p2.z*v2.x-p2.x*v2.z-p0.z*v0.x+p0.x*v0.z)
	setRow(&A[5], -dp2.y, dp2.x, 0, dv2.y, -dv2.x, 0, p2.x*v2.y-p2.y*v2.x-p0.x*v0.y+p0.y*v0.x)
	reduce(&A)
	var pRx, pRy, pRz big.Rat
	pRz.Inv(A[5][5]).Mul(&pRz, A[5][6])
	pRy.Inv(A[4][4]).Mul(&pRy, A[4][6].Sub(A[4][6], A[4][5].Mul(A[4][5], &pRz)))
	pRx.Inv(A[3][3]).Mul(&pRx, A[3][6].Sub(A[3][6], A[3][5].Mul(A[3][5], &pRz)).Sub(A[3][6], A[3][4].Mul(A[3][4], &pRy)))
	if !pRx.IsInt() || !pRy.IsInt() || !pRz.IsInt() {
		panic("expected an integer solution")
	}
	return p3{int(pRx.Num().Int64()), int(pRy.Num().Int64()), int(pRz.Num().Int64())}
}

func setRow(row *[7]*big.Rat, xs ...int) {
	for i := 0; i < 7; i++ {
		row[i] = big.NewRat(int64(xs[i]), 1)
	}
}

func reduce(A *[6][7]*big.Rat) {
	// standard Gaussian elimination with partial pivoting
	h, k := 0, 0
	var maxV, v, f, fj, zero big.Rat
	for h < 6 && k < 7 {
		maxV.Abs(A[h][k])
		maxI := h
		for i := h + 1; i < 6; i++ {
			v.Abs(A[i][k])
			if v.Cmp(&maxV) > 0 {
				maxV, maxI = v, i
			}
		}
		if A[maxI][k].Cmp(&zero) == 0 {
			k++
			continue
		}
		if maxI > h {
			A[h], A[maxI] = A[maxI], A[h]
		}
		for i := h + 1; i < 6; i++ {
			f.Inv(A[h][k]).Mul(&f, A[i][k])
			A[i][k].SetInt64(0)
			for j := k + 1; j < 7; j++ {
				A[i][j].Sub(A[i][j], fj.Mul(&f, A[h][j]))
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

// alternative solution

func altFindCollider(stones []hailstone) p3 {
	vX, vY, vZ := make(map[int][]int), make(map[int][]int), make(map[int][]int)
	for _, s := range stones {
		vX[s.v.x] = append(vX[s.v.x], s.p.x)
		vY[s.v.y] = append(vY[s.v.y], s.p.y)
		vZ[s.v.z] = append(vZ[s.v.z], s.p.z)
	}
	vRx, vRy, vRz := constrainV(vX), constrainV(vY), constrainV(vZ)
	a, b := stones[0], stones[1]
	nAx, nAy, nAz := a.v.y*vRz-a.v.z*vRy, a.v.z*vRx-a.v.x*vRz, a.v.x*vRy-a.v.y*vRx
	pAB := a.p.Sub(b.p)
	pABnA := pAB.x*nAx + pAB.y*nAy + pAB.z*nAz
	vBnA := b.v.x*nAx + b.v.y*nAy + b.v.z*nAz
	tB := pABnA / vBnA
	return p3{b.p.x + tB*(b.v.x-vRx), b.p.y + tB*(b.v.y-vRy), b.p.z + tB*(b.v.z-vRz)}
}

func constrainV(m map[int][]int) int {
	const (
		minV = -2000
		maxV = 2000
	)
	possible := []int(nil)
	for v, ps := range m {
		if len(ps) < 2 {
			continue
		}
		for i := 0; i < len(ps)-1; i++ {
			for j := i + 1; j < len(ps); j++ {
				a, b := ps[i], ps[j]
				gap := ix.Abs(a - b)
				if len(possible) == 0 {
					for rv := minV; rv <= maxV; rv++ {
						if rv != 0 && gap%rv == 0 {
							possible = append(possible, rv+v)
						}
					}
				} else {
					filtered := possible[:0]
					for _, ov := range possible {
						if rv := ov - v; rv != 0 && gap%rv == 0 {
							filtered = append(filtered, ov)
						}
					}
					possible = filtered
				}
				if len(possible) < 1 {
					panic("no choice")
				} else if len(possible) == 1 {
					return possible[0]
				}
			}
		}
	}
	panic("too much choice")
}

// parsing

type hailstone struct {
	p, v p3
}

type p3 struct {
	x, y, z int
}

func (p p3) Sub(q p3) p3 {
	return p3{p.x - q.x, p.y - q.y, p.z - q.z}
}

func (p p3) asFloat() fp3 {
	return fp3{float64(p.x), float64(p.y), float64(p.z)}
}

type fp3 struct {
	x, y, z float64
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
