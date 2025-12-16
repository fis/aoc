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
	fixedPresses, _ := m.simplify()
	if len(m.buttons) == 0 {
		return fixedPresses
	}

	A := matZero(len(m.jolts), len(m.buttons))
	for bi, b := range m.buttons {
		for ji := range m.jolts {
			if b&(1<<ji) != 0 {
				A.set(ji, bi, 1)
			}
		}
	}
	solvedPresses := solveEqs(A, m.jolts)

	return fixedPresses + solvedPresses
}

func solveEqs(A mat, b []int) (s int) {
	// solving the initial system of equation following the lead of:
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
				// A.delRow(k)
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

	x0 := make([]int, H.n)
	for i := range H.n {
		for j := range r {
			x0[i] += y[j] * U.at(i, j)
		}
	}

	basis := U.cols(r, H.n)
	free := len(basis)

	// find the (integer) values for the free variables that minimize sum(xi) subject to xi >= 0
	// this is done by finding the vertices that bound the region of valid solutions,
	// without worrying about integrality or the objective function, and then iterating over
	// the integer values within that bounding box

	switch free {
	case 0:
		return fn.Sum(x0) // unique solution, assume valid

	case 1:
		x, xc := make([]int, H.n), make([]int, H.n)
		return solve1(x, math.MaxInt, unbounded, x0, basis[0], xc)

	case 2:
		x, xc0, xc1 := make([]int, H.n), make([]int, H.n), make([]int, H.n)
		return solve2(x, math.MaxInt, unbounded, unbounded, x0, basis[0], basis[1], xc0, xc1)

	case 3:
		x, xc0, xc1, xc2 := make([]int, H.n), make([]int, H.n), make([]int, H.n), make([]int, H.n)
		return solve3(x, x0, basis[0], basis[1], basis[2], xc0, xc1, xc2)

	default:
		panic("unsupported number of free variables")
	}
}

func solve1(x []int, minS int, zb ival, x0, basis, xc []int) int {
	for i, bi := range basis {
		switch {
		case bi < 0:
			z := ix.RatOf(-x0[i], bi)
			zb = zb.updateMax(z.Floor())
		case bi > 0:
			z := ix.RatOf(-x0[i], bi)
			zb = zb.updateMin(z.Ceil())
		}
	}
	if !zb.isFinite() {
		return minS
	}
	for z := zb.min; z <= zb.max; z++ {
		if s := fn.Sum(combine(xc, x0, basis, z)); fn.Min(xc) >= 0 && s < minS {
			copy(x, xc)
			minS = s
		}
	}
	return minS
}

func solve2(x []int, minS int, zb0, zb1 ival, x0, basis0, basis1, xc0, xc1 []int) int {
	n := len(x)
	zr0, zr1 := empty, empty
	for i := 0; i < n-1; i++ {
	nextIntersect:
		for j := i + 1; j < n; j++ {
			a1, b1, c1 := basis0[i], basis1[i], -x0[i]
			a2, b2, c2 := basis0[j], basis1[j], -x0[j]
			det := a1*b2 - b1*a2
			if det == 0 {
				continue
			}
			z0, z1 := ix.RatOf(c1*b2-b1*c2, det), ix.RatOf(a1*c2-c1*a2, det)
			for k := range n {
				if k == i || k == j {
					continue // 0 by construction
				}
				xk := ix.RatI(x0[k]).Add(z0.MulI(basis0[k])).Add(z1.MulI(basis1[k]))
				if xk.P < 0 {
					continue nextIntersect // in wrong half, can't be a vertex of valid region
				}
			}
			zr0 = zr0.extend(z0.Ceil(), z0.Floor())
			zr1 = zr1.extend(z1.Ceil(), z1.Floor())
		}
	}
	zb0, zb1 = zb0.isect(zr0), zb1.isect(zr1)
	if !zb0.isFinite() || !zb1.isFinite() {
		panic("failed to bound 2 free variables")
	}
	if zb1.size() < zb0.size() {
		zb0, zb1, basis0, basis1 = zb1, zb0, basis1, basis0
	}
	for z0 := zb0.min; z0 <= zb0.max; z0++ {
		combine(xc0, x0, basis0, z0)
		minS = solve1(x, minS, zb1, xc0, basis1, xc1)
	}
	return minS
}

func solve3(x, x0, basis0, basis1, basis2, xc0, xc1, xc2 []int) int {
	minS := math.MaxInt
	zb0, zb1, zb2 := empty, empty, empty
	n := len(x)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
		nextIntersect:
			for k := j + 1; k < n; k++ {
				a1, b1, c1, d1 := basis0[i], basis1[i], basis2[i], -x0[i]
				a2, b2, c2, d2 := basis0[j], basis1[j], basis2[j], -x0[j]
				a3, b3, c3, d3 := basis0[k], basis1[k], basis2[k], -x0[k]
				det := det3(a1, b1, c1, a2, b2, c2, a3, b3, c3)
				if det == 0 {
					continue
				}
				z0 := ix.RatOf(det3(d1, b1, c1, d2, b2, c2, d3, b3, c3), det)
				z1 := ix.RatOf(det3(a1, d1, c1, a2, d2, c2, a3, d3, c3), det)
				z2 := ix.RatOf(det3(a1, b1, d1, a2, b2, d2, a3, b3, d3), det)
				for l := range n {
					if l == i || l == j || l == k {
						continue // 0 by construction
					}
					xl := ix.RatI(x0[l]).Add(z0.MulI(basis0[l])).Add(z1.MulI(basis1[l])).Add(z2.MulI(basis2[l]))
					if xl.P < 0 {
						continue nextIntersect // in wrong half, can't be a vertex of valid region
					}
				}
				zb0 = zb0.extend(z0.Ceil(), z0.Floor())
				zb1 = zb1.extend(z1.Ceil(), z1.Floor())
				zb2 = zb2.extend(z2.Ceil(), z2.Floor())
			}
		}
	}
	if !zb0.isFinite() || !zb1.isFinite() || !zb2.isFinite() {
		panic("failed to bound 3 free variables")
	}
	if zb2.size() < zb1.size() {
		zb1, zb2, basis1, basis2 = zb2, zb1, basis2, basis1
	}
	if zb1.size() < zb0.size() {
		zb0, zb1, basis0, basis1 = zb1, zb0, basis1, basis0
	}
	for z0 := zb0.min; z0 <= zb0.max; z0++ {
		combine(xc0, x0, basis0, z0)
		minS = solve2(x, minS, zb1, zb2, xc0, basis1, basis2, xc1, xc2)
	}
	return minS
}

func det3(a, b, c, d, e, f, g, h, i int) int {
	return a*e*i + b*f*g + c*d*h - c*e*g - b*d*i - a*f*h
}

func combine(xc, x0 []int, b []int, z int) []int {
	for i, x0i := range x0 {
		xc[i] = x0i + z*b[i]
	}
	return xc
}

type ival struct {
	min, max int
}

var (
	unbounded = ival{min: math.MinInt, max: math.MaxInt}
	empty     = ival{min: math.MaxInt, max: math.MinInt}
)

func (iv ival) isFinite() bool         { return iv.min != math.MinInt && iv.max != math.MaxInt }
func (iv ival) size() int              { return iv.max - iv.min + 1 }
func (iv ival) updateMin(x int) ival   { return ival{max(iv.min, x), iv.max} }
func (iv ival) updateMax(x int) ival   { return ival{iv.min, min(iv.max, x)} }
func (iv ival) extend(lo, hi int) ival { return ival{min(iv.min, lo), max(iv.max, hi)} }
func (iv ival) isect(iv2 ival) ival    { return ival{max(iv.min, iv2.min), min(iv.max, iv2.max)} }

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

func (m mat) cols(j1, j2 int) [][]int {
	cols := make([][]int, j2-j1)
	for c, j := 0, j1; j < j2; c, j = c+1, j+1 {
		cols[c] = make([]int, m.m)
		for i := range m.m {
			cols[c][i] = m.at(i, j)
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

func (m *machine) simplify() (fixedPresses int, bounds [][2]int) {
	bounds = make([][2]int, len(m.buttons))
	for i := range len(m.buttons) {
		bounds[i] = [2]int{0, math.MaxInt}
	}

	changed := true
	for changed {
		changed = false

		// remove joltages that are wired to same set of buttons (must have identical levels, otherwise unsolvable)

		wiring := m.wiring()
		wiringMap := make(map[uint32]int)
		for ji, w := range wiring {
			wiringMap[w] = m.jolts[ji]
		}
		if len(wiringMap) != len(wiring) {
			changed = true
			wiring, m.jolts = make([]uint32, 0, len(wiringMap)), make([]int, 0, len(wiringMap))
			for w, j := range wiringMap {
				wiring = append(wiring, w)
				m.jolts = append(m.jolts, j)
			}
			m.wireButtons(wiring)
		}

		// update upper bounds: button can't be pressed more times than any joltage if affects (minus lower bounds of other buttons affecting that joltage)

		for bi, b := range m.buttons {
			for ji, j := range m.jolts {
				if b&(1<<ji) == 0 {
					continue
				}
				for oi, o := range m.buttons {
					if oi == bi || o&(1<<ji) == 0 {
						continue
					}
					j -= bounds[oi][0]
				}
				if j < bounds[bi][1] {
					changed = true
					bounds[bi][1] = j
				}
			}
		}

		// update lower bounds: a button must be pressed at least as many times as required to bump up any joltage upper bounds of other buttons can't get high enough

		for bi, b := range m.buttons {
			for ji, j := range m.jolts {
				if b&(1<<ji) == 0 {
					continue
				}
				for oi, o := range m.buttons {
					if oi == bi || o&(1<<ji) == 0 {
						continue
					}
					j -= bounds[oi][1]
				}
				if j > bounds[bi][0] {
					changed = true
					bounds[bi][0] = j
				}
			}
		}

		// remove buttons that have equal bounds completely

		for bi := 0; bi < len(m.buttons); bi++ {
			if bounds[bi][0] == bounds[bi][1] {
				changed = true
				fixedPresses += bounds[bi][0]
				for ji := range len(m.jolts) {
					if m.buttons[bi]&(1<<ji) != 0 {
						m.jolts[ji] -= bounds[bi][0]
					}
				}
				m.buttons = slices.Delete(m.buttons, bi, bi+1)
				bounds = slices.Delete(bounds, bi, bi+1)
				bi--
			}
		}

		// remove joltages that have no buttons referring to them

		wiring = m.wiring()
		for ji := 0; ji < len(m.jolts); ji++ {
			if wiring[ji] == 0 {
				changed = true
				m.jolts = slices.Delete(m.jolts, ji, ji+1)
				wiring = slices.Delete(wiring, ji, ji+1)
				for bi, b := range m.buttons {
					mask := uint32(1)<<ji - 1
					m.buttons[bi] = (b & mask) | ((b >> 1) &^ mask)
				}
				ji--
			}
		}
	}

	return fixedPresses, bounds
}

func (m machine) wiring() (wiring []uint32) {
	wiring = make([]uint32, len(m.jolts))
	for bi, b := range m.buttons {
		for ji := range len(m.jolts) {
			if b&(1<<ji) != 0 {
				wiring[ji] |= 1 << bi
			}
		}
	}
	return wiring
}

func (m machine) wireButtons(wiring []uint32) {
	for bi := range m.buttons {
		m.buttons[bi] = 0
	}
	for ji, w := range wiring {
		for bi := range len(m.buttons) {
			if w&(1<<bi) != 0 {
				m.buttons[bi] |= 1 << ji
			}
		}
	}
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
