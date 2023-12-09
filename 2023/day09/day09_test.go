// Copyright 2023 Google LLC
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

import "testing"

var tests = []struct {
	seq        []int
	next, prev int
}{
	{seq: []int{0, 3, 6, 9, 12, 15}, next: 18, prev: -3},
	{seq: []int{1, 3, 6, 10, 15, 21}, next: 28, prev: 0},
	{seq: []int{10, 13, 16, 21, 30, 45}, next: 68, prev: 5},
}

func TestPredict(t *testing.T) {
	for _, test := range tests {
		if got := predict(test.seq); got != test.next {
			t.Errorf("predict(%v) = %d, want %d", test.seq, got, test.next)
		}
	}
}

func TestUnpredict(t *testing.T) {
	for _, test := range tests {
		if got := unpredict(test.seq); got != test.prev {
			t.Errorf("unpredict(%v) = %d, want %d", test.seq, got, test.prev)
		}
	}
}
