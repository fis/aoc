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

package day05

import (
	"testing"

	"github.com/fis/aoc/util"
)

var ex = [][2]util.P{
	{{0, 9}, {5, 9}},
	{{8, 0}, {0, 8}},
	{{9, 4}, {3, 4}},
	{{2, 2}, {2, 1}},
	{{7, 0}, {7, 4}},
	{{6, 4}, {2, 0}},
	{{0, 9}, {2, 9}},
	{{3, 4}, {1, 4}},
	{{0, 0}, {8, 8}},
	{{5, 5}, {8, 2}},
}

const (
	wantHV  = 5
	wantHVD = 12
)

func TestHVOverlaps(t *testing.T) {
	tests := []struct {
		name string
		f    func([][2]util.P) int
	}{
		{name: "array", f: hvOverlapsArray},
		{name: "counting", f: hvOverlapsCounting},
		{name: "pairwise", f: hvOverlapsPairwise},
	}
	canonicalize(ex)
	for _, test := range tests {
		if got := test.f(ex); got != wantHV {
			t.Errorf("%s = %d, want %d", test.name, got, wantHV)
		}
	}
}

func TestHVDOverlaps(t *testing.T) {
	tests := []struct {
		name string
		f    func([][2]util.P) int
	}{
		{name: "array", f: hvdOverlapsArray},
		{name: "typewise", f: hvdOverlapsTypewise},
	}
	canonicalize(ex)
	for _, test := range tests {
		if got := test.f(ex); got != wantHVD {
			t.Errorf("%s = %d, want %d", test.name, got, wantHVD)
		}
	}
}

func BenchmarkOverlaps(b *testing.B) {
	input, err := util.ReadRegexp("../../testdata/2021/day05.txt", inputRegexp)
	if err != nil {
		b.Fatal(err)
	}
	lines := parseInput(input)
	canonicalize(lines)
	tests := []struct {
		name string
		f    func([][2]util.P) int
		want int
	}{
		{name: "arrayHV", f: hvOverlapsArray, want: 5576},
		{name: "countingHV", f: hvOverlapsCounting, want: 5576},
		{name: "pairwiseHV", f: hvOverlapsPairwise, want: 5576},
		{name: "arrayHVD", f: hvdOverlapsArray, want: 18144},
		{name: "typewiseHVD", f: hvdOverlapsTypewise, want: 18144},
	}
	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				if got := test.f(lines); got != test.want {
					b.Errorf("%s = %d, want %d", test.name, got, test.want)
				}
			}
		})
	}
}
