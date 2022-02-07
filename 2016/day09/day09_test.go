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

package day09

import (
	"testing"
)

func TestDecompressLen(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{in: "ADVENT", want: 6},
		{in: "A(1x5)BC", want: 7},
		{in: "(3x3)XYZ", want: 9},
		{in: "A(2x2)BCD(2x2)EFG", want: 11},
		{in: "(6x1)(1x3)A", want: 6},
		{in: "X(8x2)(3x3)ABCY", want: 18},
	}
	for _, test := range tests {
		if got, err := decompressLen(test.in); err != nil {
			t.Errorf("decompressLen(%s): %v", test.in, err)
		} else if got != test.want {
			t.Errorf("decompressLen(%s) = %d, want %d", test.in, got, test.want)
		}
	}
}

func TestRecursiveLen(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{in: "(3x3)XYZ", want: 9},
		{in: "X(8x2)(3x3)ABCY", want: 20},
		{in: "(27x12)(20x12)(13x14)(7x10)(1x12)A", want: 241920},
		{in: "(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN", want: 445},
	}
	for _, test := range tests {
		if got, err := recursiveLen(test.in); err != nil {
			t.Errorf("recursiveLen(%s): %v", test.in, err)
		} else if got != test.want {
			t.Errorf("recursiveLen(%s) = %d, want %d", test.in, got, test.want)
		}
	}
}
