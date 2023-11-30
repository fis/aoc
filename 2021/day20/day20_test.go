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

package day20

import (
	"fmt"
	"strings"
	"testing"

	"github.com/fis/aoc/util"
)

var ex = []string{
	strings.Join([]string{
		"..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..",
		"#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####",
		".#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..",
		"#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#",
	}, ""),
	"#..#.",
	"#....",
	"##..#",
	"..#..",
	"..###",
}

var algos = []struct {
	name string
	f    func(algoLine string, imgLines []string) (lit2, lit50 int)
}{
	{name: "enhanceBytes", f: enhanceBytes},
	{name: "enhanceBits", f: enhanceBits},
	{name: "enhanceBitsPar-1", f: func(a string, i []string) (int, int) { return enhanceBitsPar(a, i, 1) }},
	{name: "enhanceBitsPar-2", f: func(a string, i []string) (int, int) { return enhanceBitsPar(a, i, 2) }},
	{name: "enhanceBitsPar-4", f: func(a string, i []string) (int, int) { return enhanceBitsPar(a, i, 4) }},
	{name: "enhanceBitsPar-8", f: func(a string, i []string) (int, int) { return enhanceBitsPar(a, i, 8) }},
	{name: "enhanceBitsPar-16", f: func(a string, i []string) (int, int) { return enhanceBitsPar(a, i, 16) }},
}

func TestEnhance(t *testing.T) {
	tests := []struct {
		name          string
		algo          string
		img           []string
		want2, want50 int
	}{
		{name: "ex", algo: ex[0], img: ex[1:], want2: 35, want50: 3351},
		{name: "!ex", algo: "#" + ex[0][1:511] + ".", img: ex[1:], want2: 24, want50: 3352},
	}
	for _, alg := range algos {
		for _, test := range tests {
			got2, got50 := alg.f(test.algo, test.img)
			if got2 != test.want2 || got50 != test.want50 {
				t.Errorf("%s(%s) = (%d, %d), want (%d, %d)", alg.name, test.name, got2, got50, test.want2, test.want50)
			}
		}
	}
}

func BenchmarkEnhance(b *testing.B) {
	day20, err := util.ReadChunks("../../testdata/2021/day20.txt")
	if err != nil {
		b.Fatal(err)
	}
	tests := []struct {
		name          string
		algo          string
		img           []string
		want2, want50 int
	}{
		{name: "ex", algo: ex[0], img: ex[1:], want2: 35, want50: 3351},
		{name: "!ex", algo: "#" + ex[0][1:511] + ".", img: ex[1:], want2: 24, want50: 3352},
		{name: "day20", algo: day20[0], img: util.Lines(day20[1]), want2: 5359, want50: 12333},
	}
	for _, alg := range algos {
		for _, test := range tests {
			name := fmt.Sprintf("%s/%s", alg.name, test.name)
			b.Run(name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					alg.f(test.algo, test.img)
				}
			})
		}
	}

}
