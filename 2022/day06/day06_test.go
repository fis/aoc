// Copyright 2022 Google LLC
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

package day06

import (
	"fmt"
	"testing"

	"github.com/fis/aoc/util"
)

var algos = []struct {
	name string
	f    func(string, int) int
}{
	{"bitset", findMarkerBitset},
	{"windowed", findMarkerWindowed},
}

func TestFindMarker(t *testing.T) {
	tests := []struct {
		sig  string
		size int
		want int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 4, 7},
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 14, 19},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 4, 5},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 14, 23},
		{"nppdvjthqldpwncqszvftbrmjlhg", 4, 6},
		{"nppdvjthqldpwncqszvftbrmjlhg", 14, 23},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 4, 10},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 14, 29},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 4, 11},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 14, 26},
	}
	for _, test := range tests {
		for _, algo := range algos {
			got := algo.f(test.sig, test.size)
			if got != test.want {
				t.Errorf("findMarker[%s](%s, %d) = %d, want %d", algo.name, test.sig, test.size, got, test.want)
			}
		}
	}
}

func BenchmarkFindMarker(b *testing.B) {
	lines, err := util.ReadLines("../days/testdata/day06.txt")
	if err != nil || len(lines) != 1 {
		b.Fatal(err)
	}
	sig := lines[0]
	wants := map[int]int{4: 1655, 14: 2665}
	for _, size := range []int{4, 14} {
		want := wants[size]
		for _, algo := range algos {
			testCase := fmt.Sprintf("size=%d/algo=%s", size, algo.name)
			b.Run(testCase, func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					got := algo.f(sig, size)
					if got != want {
						b.Errorf("findMarker[%s](sig, %d) = %d, want %d", algo.name, size, got, want)
					}
				}
			})
		}
	}
}
