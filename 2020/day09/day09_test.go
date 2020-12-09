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

package day09

import (
	"testing"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		data []int
		win  int
		want int
	}{
		// Simple cases: 1-25, followed by one valid/invalid number.
		{data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26}, win: 25, want: -1},
		{data: []int{25, 24, 23, 22, 21, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 49}, win: 25, want: -1},
		{data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 100}, win: 25, want: 100},
		{data: []int{25, 24, 23, 22, 21, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 50}, win: 25, want: 50},
		// One step further: 20, 1-25 excluding 20, 45, followed by one valid/invalid number.
		{data: []int{20, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 1, 21, 22, 23, 24, 25, 45, 26}, win: 25, want: -1},
		{data: []int{20, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 1, 21, 22, 23, 24, 25, 45, 65}, win: 25, want: 65},
		{data: []int{20, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 1, 21, 22, 23, 24, 25, 45, 64}, win: 25, want: -1},
		{data: []int{20, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 1, 21, 22, 23, 24, 25, 45, 66}, win: 25, want: -1},
		// The so-called larger example, with a window size of 5.
		{data: []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}, win: 5, want: 127},
	}
	for _, test := range tests {
		got := validate(test.data, test.win)
		if got != test.want {
			t.Errorf("validate(%v, %d) = %d, want %d", test.data, test.win, got, test.want)
		}
	}
}

func TestFindSum(t *testing.T) {
	data := []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}
	key := 127
	wantMin, wantMax := 15, 47
	min, max := findSum(data, key)
	if min != wantMin || max != wantMax {
		t.Errorf("findSum(%v, %d) = (%d, %d), want (%d, %d)", data, key, min, max, wantMin, wantMax)
	}
}
