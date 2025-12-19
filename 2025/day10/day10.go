// Copyright 2025 Google LLC
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

// Package day10 solves AoC 2025 day 10.
package day10

import (
	"cmp"
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/fi"
	"github.com/fis/aoc/util/fn"
	"github.com/fis/aoc/util/ix"
)

func init() {
	glue.RegisterSolver(2025, 10, glue.LineSolver(glue.WithParser(parseMachine, solve)))
}

func solve(machines []machine) ([]string, error) {
	p1 := fn.SumF(machines, machine.configLights)
	p2 := fn.SumF(machines, machine.configJolts)
	return glue.Ints(p1, p2), nil
}

type machine struct {
	lights  uint32
	buttons []uint32
	jolts   []int
}

func (m machine) configLights() (totalPresses int) {
	for n, cfg := range fi.HammingOrder(len(m.buttons)) {
		lights := uint32(0)
		for i, b := range m.buttons {
			if cfg&(1<<i) != 0 {
				lights ^= b
			}
		}
		if lights == m.lights {
			return n
		}
	}
	return -1
}

func (m machine) configJolts() (totalPresses int) {
	A := matZero(len(m.jolts), len(m.buttons))
	for bi, b := range m.buttons {
		for ji := range m.jolts {
			if b&(1<<ji) != 0 {
				A.set(ji, bi, 1)
			}
		}
	}
	return solveEqs(A, m.jolts)
}

func solveEqs(A mat, b []int) (s int) {
	// solving this initial system of equations is based on:
	// https://www.epfl.ch/labs/disopt/wp-content/uploads/2018/09/hnf.pdf

	// find Hermite normal form of A: H = AU such that U is unimodular, H is lower triangular

	H, U := A, matI(A.n) // A is no longer used past this point, so okay to do this in-place

outer:
	for k := 0; k < min(H.m, H.n); k++ {
		// ensure H(k,j) >= 0 for all j >= k
		for j := k; j < H.n; j++ {
			if H.at(k, j) < 0 {
				H.scaleCol(j, -1)
				U.scaleCol(j, -1)
			}
		}
		// ensure H(k,k) > 0, H(k,j) == 0 for all j > k
		for {
			j1, j2 := -1, -1
			for j := k; j < H.n; j++ {
				if H.at(k, j) != 0 {
					if j1 == -1 {
						j1 = j
					} else {
						j2 = j
						break
					}
				}
			}
			if j1 == -1 { // this row is redundant, throw it away to achieve full row rank
				H.delRow(k)
				b = slices.Delete(b, k, k+1)
				k-- // reprocess the row
				continue outer
			}
			if j2 == -1 {
				if j1 != k {
					H.swapCol(k, j1)
					U.swapCol(k, j1)
				}
				break
			}
			d1, d2 := H.at(k, j1), H.at(k, j2)
			if d1 > d2 {
				H.subCol(j1, j2, d1/d2)
				U.subCol(j1, j2, d1/d2)
			} else {
				H.subCol(j2, j1, d2/d1)
				U.subCol(j2, j1, d2/d1)
			}
		}
		// ensure 0 <= H(k,j) < H(k,k) for all j < k
		dk := H.at(k, k)
		for j := range k {
			dj := H.at(k, j)
			if dj >= dk {
				H.subCol(j, k, dj/dk)
				U.subCol(j, k, dj/dk)
			} else if dj < 0 {
				H.addCol(j, k, (dk-dj-1)/dk)
				U.addCol(j, k, (dk-dj-1)/dk)
			}
		}
	}
	r := min(H.m, H.n)

	// find the solution vector y = [y1 ... yr 0 ... 0]^T for the system Hy = b

	y := make([]int, r)
	for i := range r {
		di, bi := H.at(i, i), b[i]
		for j := range i {
			bi -= y[j] * H.at(i, j)
		}
		if bi%di != 0 {
			panic("unsolvable: y not integral")
		}
		y[i] = bi / di
	}

	// find the lattice of integer solutions x = Uy' where y' y + arbitrary last "free" components

	x0 := vecZero(H.n)
	for i := range H.n {
		for j := range r {
			x0.a[i] += y[j] * U.at(i, j)
		}
		x0.s += x0.a[i]
	}

	basis := U.cols(r, H.n)
	for i := range basis {
		basis[i].flipNegative()
	}
	slices.SortFunc(basis, func(a, b vec) int { return cmp.Compare(b.s, a.s) })

	// find the (integer) values for the free variables that minimize sum(xi) subject to xi >= 0
	// this is done by finding the vertices that bound the region of valid solutions,
	// without worrying about integrality or the objective function, and then iterating over
	// the integer values within that bounding box

	switch len(basis) {
	case 0:
		return x0.s // unique solution, assume valid

	case 1:
		return solve1(math.MaxInt, unbounded, x0, basis[0])

	case 2:
		xt := vecZero(H.n)
		return solve2(math.MaxInt, unbounded, unbounded, x0, basis[0], basis[1], xt)

	case 3:
		xt0, xt1 := vecZero(H.n), vecZero(H.n)
		return solve3(x0, basis[0], basis[1], basis[2], xt0, xt1)

	default:
		panic("unsupported number of free variables")
	}
}

func solve1(minS int, zb ival, x0, basis vec) int {
	for i, bi := range basis.a {
		switch {
		case bi < 0:
			z := ix.Rat{P: x0.a[i], Q: -bi} // z = -x[0]/bi
			zb = zb.updateMax(z.Floor())
		case bi > 0:
			z := ix.Rat{P: -x0.a[i], Q: bi} // z = -x[0]/bi
			zb = zb.updateMin(z.Ceil())
		}
	}
	if zb.empty() {
		return minS
	}
	return min(minS, x0.s+zb.min*basis.s)
}

func solve2(minS int, zb0, zb1 ival, x0, basis0, basis1, xt vec) int {
	n := x0.len()
	zr0, zr1 := empty, empty
	var bestZ0, bestZ1, bestS ix.Rat
	for i := 0; i < n-1; i++ {
	nextIntersect:
		for j := i + 1; j < n; j++ {
			a1, b1, c1 := basis0.a[i], basis1.a[i], -x0.a[i]
			a2, b2, c2 := basis0.a[j], basis1.a[j], -x0.a[j]
			det := a1*b2 - b1*a2
			if det == 0 {
				continue
			}
			z0p, z1p := c1*b2-b1*c2, a1*c2-c1*a2
			for k := range n {
				if k == i || k == j {
					continue // 0 by construction
				}
				if ix.RatSign(x0.a[k]*det+z0p*basis0.a[k]+z1p*basis1.a[k], det) < 0 {
					continue nextIntersect // in wrong half, can't be a vertex of valid region
				}
			}
			z0, z1 := ix.RatU(z0p, det), ix.RatU(z1p, det)
			zr0 = zr0.extend(z0.Ceil(), z0.Floor())
			zr1 = zr1.extend(z1.Ceil(), z1.Floor())
			s := ix.RatU(x0.s*det+z0p*basis0.s+z1p*basis1.s, det)
			if !bestS.IsValid() || s.Compare(bestS) < 0 {
				bestZ0, bestZ1, bestS = z0, z1, s
			}
		}
	}
	if bestS.IsValid() && bestZ0.IsInt() && bestZ1.IsInt() {
		return min(minS, bestS.P/bestS.Q)
	}
	zb0, zb1 = zb0.isect(zr0), zb1.isect(zr1)
	if !zb0.isFinite() || !zb1.isFinite() {
		panic("failed to bound 2 free variables")
	}
	if zb0.empty() || zb1.empty() {
		return minS
	}
	if basis0.s == 0 {
		return min(minS, x0.s) // objective function is the same for all these solutions
	}
	margin := zb1.min * basis1.s // best-case impact to solution cost from adjusting z1
	for z0 := zb0.min; z0 <= zb0.max; z0++ {
		combine(&xt, x0, basis0, z0)
		if xt.s+margin >= minS {
			break // can't find a better solution
		}
		_ = margin
		minS = solve1(minS, zb1, xt, basis1)
	}
	return minS
}

func solve3(x0, basis0, basis1, basis2, xt0, xt1 vec) int {
	minS := math.MaxInt
	zb0, zb1, zb2 := empty, empty, empty
	n := x0.len()
	var bestZ0, bestZ1, bestZ2, bestS ix.Rat
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
		nextIntersect:
			for k := j + 1; k < n; k++ {
				a1, b1, c1, d1 := basis0.a[i], basis1.a[i], basis2.a[i], -x0.a[i]
				a2, b2, c2, d2 := basis0.a[j], basis1.a[j], basis2.a[j], -x0.a[j]
				a3, b3, c3, d3 := basis0.a[k], basis1.a[k], basis2.a[k], -x0.a[k]
				det := det3(a1, b1, c1, a2, b2, c2, a3, b3, c3)
				if det == 0 {
					continue
				}
				z0p := det3(d1, b1, c1, d2, b2, c2, d3, b3, c3)
				z1p := det3(a1, d1, c1, a2, d2, c2, a3, d3, c3)
				z2p := det3(a1, b1, d1, a2, b2, d2, a3, b3, d3)
				for l := range n {
					if l == i || l == j || l == k {
						continue // 0 by construction
					}
					if ix.RatSign(x0.a[l]*det+z0p*basis0.a[l]+z1p*basis1.a[l]+z2p*basis2.a[l], det) < 0 {
						continue nextIntersect // in wrong half, can't be a vertex of valid region
					}
				}
				z0, z1, z2 := ix.RatU(z0p, det), ix.RatU(z1p, det), ix.RatU(z2p, det)
				zb0 = zb0.extend(z0.Ceil(), z0.Floor())
				zb1 = zb1.extend(z1.Ceil(), z1.Floor())
				zb2 = zb2.extend(z2.Ceil(), z2.Floor())
				s := ix.RatU(x0.s*det+z0p*basis0.s+z1p*basis1.s+z2p*basis2.s, det)
				if !bestS.IsValid() || s.Compare(bestS) < 0 {
					bestZ0, bestZ1, bestZ2, bestS = z0, z1, z2, s
				}
			}
		}
	}
	if bestS.IsValid() && bestZ0.IsInt() && bestZ1.IsInt() && bestZ2.IsInt() {
		return min(minS, bestS.P/bestS.Q)
	}
	if !zb0.isFinite() || !zb1.isFinite() || !zb2.isFinite() {
		panic("failed to bound 3 free variables")
	}
	if basis0.s == 0 {
		return min(minS, x0.s) // objective function is the same for all these solutions
	}
	margin := zb1.min*basis1.s + zb2.min*basis2.s // best-case impact to solution cost from adjusting z1, z2
	for z0 := zb0.min; z0 <= zb0.max; z0++ {
		combine(&xt0, x0, basis0, z0)
		if xt0.s+margin >= minS {
			break // can't find a better solution
		}
		minS = solve2(minS, zb1, zb2, xt0, basis1, basis2, xt1)
	}
	return minS
}

func det3(a, b, c, d, e, f, g, h, i int) int {
	return a*e*i + b*f*g + c*d*h - c*e*g - b*d*i - a*f*h
}

func combine(xc *vec, x0, b vec, z int) {
	xc.s = 0
	for i, x0i := range x0.a {
		xc.a[i] = x0i + z*b.a[i]
		xc.s += xc.a[i]
	}
}

type ival struct {
	min, max int
}

var (
	unbounded = ival{min: math.MinInt, max: math.MaxInt}
	empty     = ival{min: math.MaxInt, max: math.MinInt}
)

func (iv ival) isFinite() bool         { return iv.min != math.MinInt && iv.max != math.MaxInt }
func (iv ival) empty() bool            { return iv.max < iv.min }
func (iv ival) updateMin(x int) ival   { return ival{max(iv.min, x), iv.max} }
func (iv ival) updateMax(x int) ival   { return ival{iv.min, min(iv.max, x)} }
func (iv ival) extend(lo, hi int) ival { return ival{min(iv.min, lo), max(iv.max, hi)} }
func (iv ival) isect(iv2 ival) ival    { return ival{max(iv.min, iv2.min), min(iv.max, iv2.max)} }

type vec struct {
	a []int
	s int
}

func vecZero(n int) vec { return vec{a: make([]int, n), s: 0} }
func (v vec) len() int  { return len(v.a) }

func (v *vec) set(o vec) {
	copy(v.a, o.a)
	v.s = o.s
}

func (v *vec) flipNegative() {
	if v.s >= 0 {
		return
	}
	for i, x := range v.a {
		v.a[i] = -x
	}
	v.s = -v.s
}

type mat struct {
	m, n int
	a    []int
}

func matZero(m, n int) mat {
	return mat{m: m, n: n, a: make([]int, m*n)}
}

func matI(n int) mat {
	m := matZero(n, n)
	for i := range n {
		m.set(i, i, 1)
	}
	return m
}

func (m mat) at(i, j int) int { return m.a[i*m.n+j] }
func (m mat) row(i int) []int { return m.a[i*m.n : (i+1)*m.n] }
func (m mat) set(i, j, v int) { m.a[i*m.n+j] = v }

func (m mat) cols(j1, j2 int) []vec {
	cols := make([]vec, j2-j1)
	for c, j := 0, j1; j < j2; c, j = c+1, j+1 {
		cols[c] = vecZero(m.m)
		for i := range m.m {
			v := m.at(i, j)
			cols[c].a[i] = v
			cols[c].s += v
		}
	}
	return cols
}

func (m mat) swapCol(j1, j2 int) {
	for i := range m.m {
		p1, p2 := i*m.n+j1, i*m.n+j2
		m.a[p1], m.a[p2] = m.a[p2], m.a[p1]
	}
}

func (m mat) scaleCol(j, k int) {
	for i := range m.m {
		m.a[i*m.n+j] *= k
	}
}

func (m mat) addCol(jd, js, s int) {
	for i := range m.m {
		m.a[i*m.n+jd] += s * m.a[i*m.n+js]
	}
}

func (m mat) subCol(jd, js, s int) {
	for i := range m.m {
		m.a[i*m.n+jd] -= s * m.a[i*m.n+js]
	}
}

func (m *mat) delRow(i int) {
	copy(m.a[i*m.n:], m.a[(i+1)*m.n:])
	m.m--
	m.a = m.a[:m.m*m.n]
}

func (m mat) String() string {
	cw := 1
	for _, v := range m.a {
		cw = max(cw, ix.SLen(v))
	}
	var s strings.Builder
	for i := range m.m {
		for j, v := range m.row(i) {
			if j > 0 {
				s.WriteByte(' ')
			}
			w := max(1, ix.SLen(v))
			for range cw - w {
				s.WriteByte(' ')
			}
			fmt.Fprintf(&s, "%d", v)
		}
		if i < m.m-1 {
			s.WriteByte('\n')
		}
	}
	return s.String()
}

func parseMachine(line string) (m machine, err error) {
	words := util.Words(line)
	if len(words) < 3 {
		return machine{}, fmt.Errorf("expected at least 3 words, got: %s", line)
	}
	lights, buttons, jolts := words[0], words[1:len(words)-1], words[len(words)-1]
	if len(lights) < 3 || lights[0] != '[' || lights[len(lights)-1] != ']' {
		return machine{}, fmt.Errorf("invalid lights: expected [...], got: %s", lights)
	}
	for i, c := range []byte(lights[1 : len(lights)-1]) {
		switch c {
		case '.':
			// zero by default
		case '#':
			m.lights |= 1 << i
		default:
			return machine{}, fmt.Errorf("invalid lights: expected ./#, got: %c", c)
		}
	}
	m.buttons = make([]uint32, len(buttons))
	for i, bs := range buttons {
		for _, n := range util.Ints(bs) {
			m.buttons[i] |= 1 << n
		}
	}
	if len(jolts) < 3 || jolts[0] != '{' || jolts[len(jolts)-1] != '}' {
		return machine{}, fmt.Errorf("invalid jolts: expected {...}, got: %s", jolts)
	}
	m.jolts = util.Ints(jolts)
	return m, nil
}

func (m machine) String() string {
	var s strings.Builder
	s.WriteByte('[')
	for i := range len(m.jolts) {
		if m.lights&(1<<i) != 0 {
			s.WriteByte('#')
		} else {
			s.WriteByte('.')
		}
	}
	s.WriteByte(']')
	for _, b := range m.buttons {
		s.WriteByte(' ')
		s.WriteByte('(')
		first := true
		for i := range 32 {
			if b&(1<<i) != 0 {
				if first {
					first = false
				} else {
					s.WriteByte(',')
				}
				fmt.Fprintf(&s, "%d", i)
			}
		}
		s.WriteByte(')')
	}
	s.WriteByte(' ')
	s.WriteByte('{')
	for i, j := range m.jolts {
		if i > 0 {
			s.WriteByte(',')
		}
		fmt.Fprintf(&s, "%d", j)
	}
	s.WriteByte('}')
	return s.String()
}
