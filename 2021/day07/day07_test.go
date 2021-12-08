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

package day07

import (
	"testing"

	"github.com/fis/aoc/util"
)

var ex = []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

func TestAlign(t *testing.T) {
	tests := []struct {
		name     string
		cost     func(n, x int) int
		wantX    int
		wantCost int
	}{
		{name: "cost1", cost: cost1, wantX: 2, wantCost: 37},
		{name: "cost2", cost: cost2, wantX: 5, wantCost: 168},
	}
	for _, test := range tests {
		if x, cost := align(ex, test.cost); x != test.wantX || cost != test.wantCost {
			t.Errorf("align(%v, %s) = (%d, %d), want (%d, %d)", ex, test.name, x, cost, test.wantX, test.wantCost)
		}
	}
}

func TestAlgos(t *testing.T) {
	input, err := util.ReadInts("../days/testdata/day07.txt")
	if err != nil {
		t.Fatal(err)
	}
	inputs := [][]int{ex, input}
	want1 := [][2]int{{2, 37}, {317, 352997}}
	want2 := [][2]int{{5, 168}, {466, 101571302}}

	tests := []struct {
		name  string
		f     func(input []int) (x, cost int)
		wants [][2]int
	}{
		{
			name:  "align/cost1",
			f:     func(input []int) (x, cost int) { return align(input, cost1) },
			wants: want1,
		},
		{name: "align1Points", f: align1Points, wants: want1},
		{name: "align1MedianSort", f: align1MedianSort, wants: want1},
		{name: "align1MedianQS", f: align1MedianQS, wants: want1},
		{
			name:  "align/cost2",
			f:     func(input []int) (x, cost int) { return align(input, cost2) },
			wants: want2,
		},
		{name: "align2Mean", f: align2Mean, wants: want2},
	}
	for _, test := range tests {
		for i, input := range inputs {
			want := test.wants[i]
			if gotX, gotCost := test.f(input); gotX != want[0] || gotCost != want[1] {
				t.Errorf("%s = (%d, %d), want (%d, %d)", test.name, gotX, gotCost, want[0], want[1])
			}
		}
	}
}

func BenchmarkAlgos(b *testing.B) {
	input, err := util.ReadInts("../days/testdata/day07.txt")
	if err != nil {
		b.Fatal(err)
	}

	tests := []struct {
		name string
		f    func(input []int) (x, cost int)
		want [2]int
	}{
		{
			name: "align/cost1",
			f:    func(input []int) (x, cost int) { return align(input, cost1) },
			want: [2]int{317, 352997},
		},
		{name: "align1Points", f: align1Points, want: [2]int{317, 352997}},
		{name: "align1MedianSort", f: align1MedianSort, want: [2]int{317, 352997}},
		{name: "align1MedianQS", f: align1MedianQS, want: [2]int{317, 352997}},
		{
			name: "align/cost2",
			f:    func(input []int) (x, cost int) { return align(input, cost2) },
			want: [2]int{466, 101571302},
		},
		{name: "align2Mean", f: align2Mean, want: [2]int{466, 101571302}},
	}
	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				if gotX, gotCost := test.f(input); gotX != test.want[0] || gotCost != test.want[1] {
					b.Errorf("%s = (%d, %d), want (%d, %d)", test.name, gotX, gotCost, test.want[0], test.want[1])
				}
			}
		})
	}
}
