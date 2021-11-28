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

package day10

import (
	"testing"
)

func TestPart1(t *testing.T) {
	N := 5
	lengths := []byte{3, 4, 1, 5}
	want := 12
	if got := part1(N, lengths); got != want {
		t.Errorf("part1(%d, %v) = %d, want %d", N, lengths, got, want)
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{input: ``, want: "a2582a3a0e66e6e86e3812dcb672a272"},
		{input: `AoC 2017`, want: "33efeb34ea91902bb2f59c9920caa6cd"},
		{input: `1,2,3`, want: "3efbe78a8d82f29979031a4aa0b16a9d"},
		{input: `1,2,4`, want: "63960835bcdc130f0b66d7ff4f6a5a8e"},
	}
	for _, test := range tests {
		if got := part2(256, 64, test.input); got != test.want {
			t.Errorf("part2(256, 64, %q) = %q, want %q", test.input, got, test.want)
		}
	}
}
