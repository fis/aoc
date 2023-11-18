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

// Package day05 solves AoC 2021 day 5.
package day05

import (
	"strconv"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/ix"
)

const inputRegexp = `^(\d+),(\d+) -> (\d+),(\d+)$`

func init() {
	glue.RegisterSolver(2021, 5, glue.RegexpSolver{
		Solver: solve,
		Regexp: inputRegexp,
	})
}

func solve(input [][]string) ([]string, error) {
	lines := parseInput(input)
	canonicalize(lines)
	p1 := hvOverlapsArray(lines)
	p2 := hvdOverlapsTypewise(lines)
	return glue.Ints(p1, p2), nil
}

// Before you look below: yes, the "pairwise" function is entirely ridiculous, and you should have
// seen what the (incomplete) version with diagonals looked like. It did significantly outperform
// the initial simple "counting" method, but not after switching the method of how to track overlaps.
// Here's a representative benchmark output:
//   BenchmarkOverlaps/arrayHV-16       3240      330777 ns/op     966659 B/op       1 allocs/op
//   BenchmarkOverlaps/countingHV-16      86    13436727 ns/op    8209626 B/op    3875 allocs/op
//   BenchmarkOverlaps/pairwiseHV-16     843     1488492 ns/op     348414 B/op     164 allocs/op

// N.B. all these functions have an expected order for the endpoints of each line (see canonicalize).

func hvOverlapsArray(lines [][2]util.P) (overlaps int) {
	minX, maxX, minY, maxY := bounds(lines)
	W, H := maxX-minX+1, maxY-minY+1
	seen := make([]byte, W*H)
	for _, line := range lines {
		h, v := hv(line)
		switch {
		case h:
			y := line[0].Y
			x1, x2 := line[0].X, line[1].X
			for x := x1; x <= x2; x++ {
				i := (y-minY)*W + (x - minX)
				switch seen[i] {
				case 0:
					seen[i] = 1
				case 1:
					overlaps++
					seen[i] = 2
				}
			}
		case v:
			x := line[0].X
			y1, y2 := line[0].Y, line[1].Y
			for y := y1; y <= y2; y++ {
				i := (y-minY)*W + (x - minX)
				switch seen[i] {
				case 0:
					seen[i] = 1
				case 1:
					overlaps++
					seen[i] = 2
				}
			}
		}
	}
	return overlaps
}

func hvOverlapsCounting(lines [][2]util.P) (overlaps int) {
	freq := map[util.P]int{}
	for _, line := range lines {
		h, v := hv(line)
		switch {
		case h:
			y := line[0].Y
			x1, x2 := line[0].X, line[1].X
			for x := x1; x <= x2; x++ {
				freq[util.P{x, y}] = freq[util.P{x, y}] + 1
			}
		case v:
			x := line[0].X
			y1, y2 := line[0].Y, line[1].Y
			for y := y1; y <= y2; y++ {
				freq[util.P{x, y}] = freq[util.P{x, y}] + 1
			}
		}
	}
	for _, f := range freq {
		if f > 1 {
			overlaps++
		}
	}
	return overlaps
}

func hvOverlapsPairwise(lines [][2]util.P) int {
	overlaps := map[util.P]struct{}{}
	for i, N := 0, len(lines); i < N-1; i++ {
		for j := i + 1; j < N; j++ {
			l1, l2 := lines[i], lines[j]
			h1, v1 := hv(l1)
			h2, v2 := hv(l2)
			switch {
			case h1 && h2:
				if l1[0].Y == l2[0].Y {
					a1, a2 := l1[0].X, l1[1].X
					b1, b2 := l2[0].X, l2[1].X
					if a2 >= b1 && a1 <= b2 {
						x1, x2 := max(a1, b1), min(a2, b2)
						y := l1[0].Y
						for x := x1; x <= x2; x++ {
							overlaps[util.P{x, y}] = struct{}{}
						}
					}
				}
			case v1 && v2:
				if l1[0].X == l2[0].X {
					a1, a2 := l1[0].Y, l1[1].Y
					b1, b2 := l2[0].Y, l2[1].Y
					if a2 >= b1 && a1 <= b2 {
						y1, y2 := max(a1, b1), min(a2, b2)
						x := l1[0].X
						for y := y1; y <= y2; y++ {
							overlaps[util.P{x, y}] = struct{}{}
						}
					}
				}
			case h1 && v2:
				p := util.P{l2[0].X, l1[0].Y}
				x1, x2 := l1[0].X, l1[1].X
				y1, y2 := l2[0].Y, l2[1].Y
				if p.X >= x1 && p.X <= x2 && p.Y >= y1 && p.Y <= y2 {
					overlaps[p] = struct{}{}
				}
			case v1 && h2:
				p := util.P{l1[0].X, l2[0].Y}
				x1, x2 := l2[0].X, l2[1].X
				y1, y2 := l1[0].Y, l1[1].Y
				if p.X >= x1 && p.X <= x2 && p.Y >= y1 && p.Y <= y2 {
					overlaps[p] = struct{}{}
				}
			}
		}
	}
	return len(overlaps)
}

func hv(line [2]util.P) (h, v bool) {
	h = line[0].Y == line[1].Y
	v = line[0].X == line[1].X
	return h, v
}

func hvdOverlapsArray(lines [][2]util.P) (overlaps int) {
	minX, maxX, minY, maxY := bounds(lines)
	W, H := maxX-minX+1, maxY-minY+1
	seen := make([]byte, W*H)
	for _, line := range lines {
		len := max(ix.Abs(line[1].X-line[0].X), ix.Abs(line[1].Y-line[0].Y))
		dx, dy := (line[1].X-line[0].X)/len, (line[1].Y-line[0].Y)/len
		for i, x, y := 0, line[0].X-minX, line[0].Y-minY; i <= len; i, x, y = i+1, x+dx, y+dy {
			i := y*W + x
			switch seen[i] {
			case 0:
				seen[i] = 1
			case 1:
				overlaps++
				seen[i] = 2
			}
		}
	}
	return overlaps
}

func hvdOverlapsTypewise(lines [][2]util.P) (overlaps int) {
	minX, maxX, minY, maxY := bounds(lines)
	W, H := maxX-minX+1, maxY-minY+1
	seen := make([]byte, W*H)
	for _, line := range lines {
		h, v, dA, dD := hvd(line)
		switch {
		case h:
			y := line[0].Y
			x1, x2 := line[0].X, line[1].X
			for x := x1; x <= x2; x++ {
				i := (y-minY)*W + (x - minX)
				switch seen[i] {
				case 0:
					seen[i] = 1
				case 1:
					overlaps++
					seen[i] = 2
				}
			}
		case v:
			x := line[0].X
			y1, y2 := line[0].Y, line[1].Y
			for y := y1; y <= y2; y++ {
				i := (y-minY)*W + (x - minX)
				switch seen[i] {
				case 0:
					seen[i] = 1
				case 1:
					overlaps++
					seen[i] = 2
				}
			}
		case dA || dD:
			p1, p2 := line[0], line[1]
			if p1.X > p2.X {
				p1, p2 = p2, p1
			}
			yd := 1
			if dD {
				yd = -1
			}
			for x, y := p1.X, p1.Y; x <= p2.X; x, y = x+1, y+yd {
				i := (y-minY)*W + (x - minX)
				switch seen[i] {
				case 0:
					seen[i] = 1
				case 1:
					overlaps++
					seen[i] = 2
				}
			}
		}
	}
	return overlaps
}

func hvd(line [2]util.P) (h, v, dA, dD bool) {
	h = line[0].Y == line[1].Y
	v = line[0].X == line[1].X
	dA = line[1].Y-line[0].Y == ix.Abs(line[1].X-line[0].X)
	dD = line[1].Y-line[0].Y == -ix.Abs(line[1].X-line[0].X)
	return h, v, dA, dD
}

func bounds(lines [][2]util.P) (minX, maxX, minY, maxY int) {
	minX, maxX, minY, maxY = lines[0][0].X, lines[0][0].X, lines[0][0].Y, lines[0][0].Y
	for _, line := range lines {
		for d := 0; d < 2; d++ {
			p := line[d]
			if p.X < minX {
				minX = p.X
			}
			if p.X > maxX {
				maxX = p.X
			}
			if p.Y < minY {
				minY = p.Y
			}
			if p.Y > maxY {
				maxY = p.Y
			}
		}
	}
	return minX, maxX, minY, maxY
}

func parseInput(input [][]string) (lines [][2]util.P) {
	lines = make([][2]util.P, len(input))
	for i, parts := range input {
		lines[i][0].X, _ = strconv.Atoi(parts[0])
		lines[i][0].Y, _ = strconv.Atoi(parts[1])
		lines[i][1].X, _ = strconv.Atoi(parts[2])
		lines[i][1].Y, _ = strconv.Atoi(parts[3])
	}
	return lines
}

func canonicalize(lines [][2]util.P) {
	for i := range lines {
		p1, p2 := lines[i][0], lines[i][1]
		if p1.X > p2.X || (p1.X == p2.X && p1.Y > p2.Y) {
			lines[i][0], lines[i][1] = p2, p1
		}
	}
}
