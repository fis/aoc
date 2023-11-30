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

package day15

import (
	"fmt"
	"testing"

	"github.com/fis/aoc/util"
)

var ex = []string{
	"1163751742",
	"1381373672",
	"2136511328",
	"3694931569",
	"7463417111",
	"1319128137",
	"1359912421",
	"3125421639",
	"1293138521",
	"2311944581",
}

var algos = []struct {
	name string
	f    func(w, h int32, level [][]byte, scale int32) int32
}{
	{"Dijkstra", shortestPathDijkstra},
	{"DijkstraBQ", shortestPathDijkstraBQ},
	{"AStar", shortestPathAStar},
}

func TestShortestPath(t *testing.T) {
	w, h, level, err := readLevel(ex)
	if err != nil {
		t.Fatalf("readLevel: %v", err)
	}
	tests := []struct {
		scale int32
		want  int32
	}{
		{scale: 1, want: 40},
		{scale: 5, want: 315},
	}
	for _, alg := range algos {
		for _, test := range tests {
			if got := alg.f(w, h, level, test.scale); got != test.want {
				t.Errorf("%s, scale %d = %d, want %d", alg.name, test.scale, got, test.want)
			}
		}
	}
}

func BenchmarkShortestPath(b *testing.B) {
	exW, exH, exLevel, err := readLevel(ex)
	if err != nil {
		b.Fatalf("readLevel(ex): %v", err)
	}

	dayLines, err := util.ReadLines("../../testdata/2021/day15.txt")
	if err != nil {
		b.Fatalf("ReadLines(day15): %v", err)
	}
	dayW, dayH, dayLevel, err := readLevel(dayLines)
	if err != nil {
		b.Fatalf("readLevel(day15): %v", err)
	}

	samples := []struct {
		name  string
		w, h  int32
		level [][]byte
	}{
		{"ex", exW, exH, exLevel},
		{"day15", dayW, dayH, dayLevel},
	}

	for _, sample := range samples {
		for _, scale := range []int32{1, 5} {
			for _, alg := range algos {
				name := fmt.Sprintf("%s/scale=%d/%s", sample.name, scale, alg.name)
				b.Run(name, func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						alg.f(sample.w, sample.h, sample.level, scale)
					}
				})
			}
		}
	}
}
