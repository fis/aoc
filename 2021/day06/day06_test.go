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

package day06

import (
	"testing"
)

func TestCountFish(t *testing.T) {
	initial := []int{3, 4, 3, 1, 2}
	tests := []struct {
		days int
		want int
	}{
		{days: 1, want: 5},
		{days: 2, want: 6},
		{days: 3, want: 7},
		{days: 4, want: 9},
		{days: 5, want: 10},
		{days: 18, want: 26},
		{days: 80, want: 5934},
		{days: 256, want: 26984457539},
	}
	for _, test := range tests {
		if got := countFish(initial, test.days); got != test.want {
			t.Errorf("countFish(%v, %d) = %d, want %d", initial, test.days, got, test.want)
		}
	}
}
