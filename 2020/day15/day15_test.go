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

package day15

import "testing"

func TestSimulate(t *testing.T) {
	tests := []struct {
		initial []int
		want    int
	}{
		{initial: []int{0, 3, 6}, want: 436},
		{initial: []int{1, 3, 2}, want: 1},
		{initial: []int{2, 1, 3}, want: 10},
		{initial: []int{1, 2, 3}, want: 27},
		{initial: []int{2, 3, 1}, want: 78},
		{initial: []int{3, 2, 1}, want: 438},
		{initial: []int{3, 1, 2}, want: 1836},
	}
	for _, test := range tests {
		got := simulate(test.initial, 2020)
		if got != test.want {
			t.Fatalf("simulate(%v, 2020) = %d, want %d", test.initial, got, test.want)
		}
	}
}

func BenchmarkAlgos(b *testing.B) {
	tests := []struct {
		name string
		f    func([]int, int) int
	}{
		{name: "current", f: simulate},
	}
	initial, upTo := []int{15, 5, 1, 4, 7, 0}, 30000000
	want := 689
	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				got := test.f(initial, upTo)
				if got != want {
					b.Errorf("%s = %d, want %d", test.name, got, want)
				}
			}
		})
	}
}
