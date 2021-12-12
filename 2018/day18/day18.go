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

// Package day18 solves AoC 2018 day 18.
package day18

import (
	"hash/fnv"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2018, 18, glue.LevelSolver{Solver: solve, Empty: ' '})
}

func solve(level *util.Level) ([]string, error) {
	data, w, h := convertLevel(level)

	part1 := value(evolve(append([]state(nil), data...), w, h, 10))

	target := 1000000000
	data, at, period := findCycle(data, w, h)
	at += (target - at) / period * period
	if at < target {
		data = evolve(data, w, h, target-at)
	}
	part2 := value(data)

	return glue.Ints(part1, part2), nil
}

type state = byte

const (
	stateOpen state = iota
	stateTrees
	stateLumber
)

func convertLevel(level *util.Level) (data []state, w, h int) {
	min, max := level.Bounds()
	w, h = max.X-min.X+1, max.Y-min.Y+1
	data = make([]state, w*h)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			switch level.At(min.X+x, min.Y+y) {
			case '.':
				data[y*w+x] = stateOpen
			case '|':
				data[y*w+x] = stateTrees
			case '#':
				data[y*w+x] = stateLumber
			}
		}
	}
	return data, w, h
}

func evolve(data []state, w, h, generations int) []state {
	next := make([]state, len(data))
	for g := 0; g < generations; g++ {
		step(data, next, w, h)
		data, next = next, data
	}
	return data
}

func findCycle(data []state, w, h int) (out []state, at, period int) {
	seen := map[uint64]int{}
	next := make([]state, len(data))
	for at = 0; ; at++ {
		hash := hashLevel(data)
		if prev, ok := seen[hash]; ok {
			return data, at, at - prev
		}
		seen[hash] = at
		step(data, next, w, h)
		data, next = next, data
	}
}

func step(in, out []state, w, h int) {
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			nt, nl := 0, 0
			for _, n := range (util.P{x, y}).Neigh8() {
				if n.X < 0 || n.X >= w || n.Y < 0 || n.Y >= h {
					continue
				}
				switch in[n.Y*w+n.X] {
				case stateTrees:
					nt++
				case stateLumber:
					nl++
				}
			}
			v := in[y*w+x]
			switch {
			case v == stateOpen && nt >= 3:
				v = stateTrees
			case v == stateTrees && nl >= 3:
				v = stateLumber
			case v == stateLumber && (nt == 0 || nl == 0):
				v = stateOpen
			}
			out[y*w+x] = v
		}
	}
}

func value(data []state) int {
	nt, nl := 0, 0
	for _, v := range data {
		switch v {
		case stateTrees:
			nt++
		case stateLumber:
			nl++
		}
	}
	return nt * nl
}

func hashLevel(data []state) uint64 {
	h := fnv.New64a()
	h.Write(data)
	return h.Sum64()
}
