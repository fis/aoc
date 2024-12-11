// Copyright 2024 Google LLC
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

package day11

import (
	"testing"
)

var ex = []int{125, 17}

func TestCountAll(t *testing.T) {
	var c counter
	want := 55312
	if got := c.countAll(ex, 25); got != want {
		t.Errorf("countAll(ex, 25) = %d, want %d", got, want)
	}
}

func TestSplit(t *testing.T) {
	tests := []struct {
		input     int
		wantLeft  int
		wantRight int
		wantOK    bool
	}{
		{125, 0, 0, false},
		{17, 1, 7, true},
		{253000, 253, 0, true},
		{14168, 0, 0, false},
		{512072, 512, 72, true},
	}
	for _, test := range tests {
		left, right, ok := split(test.input)
		if left != test.wantLeft || right != test.wantRight || ok != test.wantOK {
			t.Errorf("split(%d) = (%d, %d, %t), want (%d, %d, %t)", test.input, left, right, ok, test.wantLeft, test.wantRight, test.wantOK)
		}
	}
}

func BenchmarkSplit(b *testing.B) {
	for range b.N {
		for _, s := range []int{125, 17, 253000, 14168, 512072, 123456789, 1234567890, 12345678901, 123456789012} {
			_, _, _ = split(s)
		}
	}
}
