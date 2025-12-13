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

// Package day12 solves AoC 2025 day 12.
package day12

import (
	"math/bits"
	"strconv"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2025, 12, glue.ChunkSolver(solve))
}

func solve(chunks []string) ([]string, error) {
	shapes, regions := parseGifts(chunks)
	p1 := countFitting(shapes, regions) // such a scam
	return glue.Ints(p1), nil
}

func countFitting(shapes []shape, regions []region) (fitting int) {
	shapeAreas := make([]int, len(shapes))
	for i, s := range shapes {
		shapeAreas[i] = s.area()
	}

	for _, r := range regions {
		reqShapes, reqArea := 0, 0
		for i, req := range r.reqs {
			reqShapes += req
			reqArea += req * shapeAreas[i]
		}
		switch {
		case reqArea > r.w*r.h:
			// obviously impossible: needs more total area for gifts than in region
		case reqShapes <= ((r.w+2)/3)*((r.h+2)/3):
			fitting++ // obviously possible: can fit each gift in its own 3x3 box
		default:
			panic("oh no, there's a real puzzle here")
		}
	}

	return fitting
}

type (
	shape  uint16
	region struct {
		w, h int
		reqs []int
	}
)

func (s shape) area() int { return bits.OnesCount16(uint16(s)) }

func parseGifts(chunks []string) (shapes []shape, regions []region) {
	shapes = make([]shape, len(chunks)-1)
	for i, chunk := range chunks[:len(chunks)-1] {
		var s shape
		lines := util.Lines(chunk)
		for y, line := range lines[1:] {
			for x, c := range []byte(line) {
				if c == '#' {
					s |= 1 << (3*y + x)
				}
			}
		}
		shapes[i] = s
	}

	regionLines := util.Lines(chunks[len(chunks)-1])
	regions = make([]region, len(regionLines))
	for i, line := range regionLines {
		size, reqs, _ := strings.Cut(line, ": ")
		ws, hs, _ := strings.Cut(size, "x")
		w, _ := strconv.Atoi(ws)
		h, _ := strconv.Atoi(hs)
		regions[i] = region{w: w, h: h, reqs: util.Ints(reqs)}
	}

	return shapes, regions
}
