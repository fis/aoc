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

// Package day19 solves AoC 2021 day 19.
package day19

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2021, 19, glue.ChunkSolver(solve))
}

func solve(chunks []string) ([]string, error) {
	scanners, err := parseInput(chunks)
	if err != nil {
		return nil, err
	}
	scannerPos, beaconPos := scanners.buildMap()
	p1 := len(beaconPos)
	p2 := maxDist(scannerPos)
	return glue.Ints(p1, p2), nil
}

func maxDist(points []p3) (maxD int) {
	for i, p := range points[:len(points)-1] {
		for _, q := range points[i+1:] {
			d := int(abs(q.x-p.x) + abs(q.y-p.y) + abs(q.z-p.z))
			if d > maxD {
				maxD = d
			}
		}
	}
	return maxD
}

type scannerSet struct {
	scanners []scanner
}

func (ss *scannerSet) buildMap() (scannerPos []p3, beaconPos map[p3]struct{}) {
	beaconPos = make(map[p3]struct{})
	used := make(map[int]struct{})
	scannerPos = ss.mergeToMap(scannerPos, beaconPos, 0, identityTr, used)
	if len(used) != len(ss.scanners) {
		panic("could not merge all scanners")
	}
	return scannerPos, beaconPos
}

func (ss *scannerSet) mergeToMap(scannerPos []p3, beaconPos map[p3]struct{}, scanner int, tr transform, used map[int]struct{}) []p3 {
	scannerPos = append(scannerPos, tr.apply(p3{0, 0, 0}))
	for _, p := range ss.scanners[scanner].beacons {
		beaconPos[tr.apply(p)] = struct{}{}
	}

	used[scanner] = struct{}{}
	for next := 0; next < len(ss.scanners); next++ {
		if _, ok := used[next]; ok {
			continue
		}
		overlap := ss.overlap(scanner, next)
		if len(overlap) == 0 {
			continue
		}
		nextTr := tr.combine(ss.align(scanner, next, overlap))
		scannerPos = ss.mergeToMap(scannerPos, beaconPos, next, nextTr, used)
	}

	return scannerPos
}

func (ss *scannerSet) align(ref, next int, overlap [][2]int32) (tr transform) {
	r1 := ss.scanners[ref].beacons[overlap[0][0]]
	r2 := ss.scanners[ref].beacons[overlap[1][0]]
	rd := p3{r2.x - r1.x, r2.y - r1.y, r2.z - r1.z}
	n1 := ss.scanners[next].beacons[overlap[0][1]]
	n2 := ss.scanners[next].beacons[overlap[1][1]]
	nd := p3{n2.x - n1.x, n2.y - n1.y, n2.z - n1.z}
	for rot := 0; ; rot++ {
		tr = rotationTr[rot]
		if tr.apply(nd) == rd {
			break
		}
	}
	n1r := tr.apply(n1)
	tr[0][3], tr[1][3], tr[2][3] = r1.x-n1r.x, r1.y-n1r.y, r1.z-n1r.z
	return tr
}

func (ss *scannerSet) overlap(id1, id2 int) (match [][2]int32) {
	const (
		wantPoints    = 12
		wantDistances = wantPoints * (wantPoints - 1) / 2
	)
	var distPairs [wantDistances][2]distance
	if !ss.matchDistances(id1, id2, distPairs[:]) {
		return nil
	}

	points1 := make(map[int32][][2]distance)
	for _, dp := range distPairs {
		points1[dp[0].p1] = append(points1[dp[0].p1], dp)
		points1[dp[0].p2] = append(points1[dp[0].p2], dp)
	}
	if len(points1) != wantPoints {
		panic("impossible: got right amount of distances but wrong amount of points")
	}

	match = make([][2]int32, wantPoints)
	at := 0
	for p1, dps := range points1 {
		p2a, p2b := dps[0][1].p1, dps[0][1].p2
		if p2a == dps[1][1].p1 || p2a == dps[1][1].p2 {
			match[at] = [2]int32{p1, p2a}
		} else if p2b == dps[1][1].p1 || p2b == dps[1][1].p2 {
			match[at] = [2]int32{p1, p2b}
		} else {
			panic("impossible: no match found after all")
		}
		at++
	}
	return match
}

func (ss *scannerSet) matchDistances(id1, id2 int, out [][2]distance) bool {
	s1, s2 := &ss.scanners[id1], &ss.scanners[id2]
	d1, d2 := s1.dists, s2.dists
	at := 0
	for len(d1) > 0 && len(d2) > 0 {
		if d1[0].d == d2[0].d {
			out[at] = [2]distance{d1[0], d2[0]}
			at++
			if at == len(out) {
				return true
			}
			d1, d2 = d1[1:], d2[1:]
		} else if distLess(d1[0].d, d2[0].d) {
			d1 = d1[1:]
		} else {
			d2 = d2[1:]
		}
	}
	return false
}

type scanner struct {
	beacons []p3       // beacon positions
	dists   []distance // pairwise distances between all beacons, in order
}

type distance struct {
	d      p3    // |x2-x1|, |y2-y1|, |x2-z1| in order of magnitude
	p1, p2 int32 // endpoints the distance is for
}

func (s *scanner) calcDists() {
	distMap := make(map[p3][2]int32)
	for i := 0; i < len(s.beacons)-1; i++ {
		p := s.beacons[i]
		for j := i + 1; j < len(s.beacons); j++ {
			q := s.beacons[j]
			dx, dy, dz := abs(q.x-p.x), abs(q.y-p.y), abs(q.z-p.z)
			da, db, dc := sort3(dx, dy, dz)
			d := p3{da, db, dc}
			if _, seen := distMap[d]; seen {
				panic("non-unique distance!") // could happen in reality, doesn't happen in puzzle-land
			}
			distMap[d] = [2]int32{int32(i), int32(j)}
		}
	}
	for d, p := range distMap {
		s.dists = append(s.dists, distance{d: d, p1: p[0], p2: p[1]})
	}
	sort.Slice(s.dists, func(i, j int) bool { return distLess(s.dists[i].d, s.dists[j].d) })
}

func distLess(a, b p3) bool {
	if a.x < b.x {
		return true
	}
	if a.x == b.x && a.y < b.y {
		return true
	}
	if a.x == b.x && a.y == b.y && a.z < b.z {
		return true
	}
	return false
}

func parseInput(chunks []string) (ss *scannerSet, err error) {
	ss = &scannerSet{scanners: make([]scanner, len(chunks))}
	for i, chunk := range chunks {
		header := fmt.Sprintf("--- scanner %d ---\n", i)
		if len(chunk) < len(header) {
			return nil, fmt.Errorf("short chunk: %q", chunk)
		} else if !strings.HasPrefix(chunk, header) {
			return nil, fmt.Errorf("expected header %q, got %q", header, chunk[:len(header)])
		}
		lines, err := util.ScanAllRegexp(strings.NewReader(chunk[len(header):]), `^(-?\d+),(-?\d+),(-?\d+)$`)
		if err != nil {
			return nil, err
		}
		ss.scanners[i].beacons = make([]p3, len(lines))
		for j, line := range lines {
			x, _ := strconv.Atoi(line[0])
			y, _ := strconv.Atoi(line[1])
			z, _ := strconv.Atoi(line[2])
			ss.scanners[i].beacons[j] = p3{int32(x), int32(y), int32(z)}
		}
		ss.scanners[i].calcDists()
	}
	return ss, nil
}

type transform [4][4]int32

func (tr transform) apply(p p3) (q p3) {
	q.x = tr[0][0]*p.x + tr[0][1]*p.y + tr[0][2]*p.z + tr[0][3]
	q.y = tr[1][0]*p.x + tr[1][1]*p.y + tr[1][2]*p.z + tr[1][3]
	q.z = tr[2][0]*p.x + tr[2][1]*p.y + tr[2][2]*p.z + tr[2][3]
	return q
}

func (tr transform) combine(next transform) (out transform) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				out[i][j] += tr[i][k] * next[k][j]
			}
		}
	}
	return out
}

var identityTr = transform{{1, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}}

var rotationTr [24]transform

func init() {
	x, y, z, w := [4]int32{1, 0, 0, 0}, [4]int32{0, 1, 0, 0}, [4]int32{0, 0, 1, 0}, [4]int32{0, 0, 0, 1}
	X, Y, Z := [4]int32{-1, 0, 0, 0}, [4]int32{0, -1, 0, 0}, [4]int32{0, 0, -1, 0}
	rotationTr = [24]transform{
		{x, y, z, w}, {y, X, z, w}, {X, Y, z, w}, {Y, x, z, w},
		{Z, y, x, w}, {y, z, x, w}, {z, Y, x, w}, {Y, Z, x, w},
		{X, y, Z, w}, {y, x, Z, w}, {x, Y, Z, w}, {Y, X, Z, w},
		{z, y, X, w}, {y, Z, X, w}, {Z, Y, X, w}, {Y, z, X, w},
		{x, Z, y, w}, {Z, X, y, w}, {X, z, y, w}, {z, x, y, w},
		{x, z, Y, w}, {Z, x, Y, w}, {X, Z, Y, w}, {z, X, Y, w},
	}
}

type p3 struct {
	x, y, z int32
}

func sort3(x, y, z int32) (a, b, c int32) {
	if x <= y { // x <= y
		if x <= z { // x <= y, x <= z
			if y <= z { // x <= y, x <= z, y <= z
				return x, y, z
			} else { // x <= y, x <= z, y > z
				return x, z, y
			}
		} else { // x <= y, x > z
			return z, x, y
		}
	} else { // x > y
		if y <= z { // x > y, y <= z
			if x <= z { // x > y, y <= z, x <= z
				return y, x, z
			} else { // x > y, y <= z, x > z
				return y, z, x
			}
		} else { // x > y, y > z
			return z, y, x
		}
	}
}

func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}
