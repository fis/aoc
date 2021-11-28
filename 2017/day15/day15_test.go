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

package day15

import (
	"testing"
)

const (
	xA = 65
	xB = 8921
)

func TestJudge(t *testing.T) {
	tests := []struct {
		N    int
		want int
	}{
		{N: 5, want: 1},
		{N: 40000000, want: 588},
	}
	for _, test := range tests {
		if got := judge(xA, xB, test.N); got != test.want {
			t.Errorf("judge(..., %d) = %d, want %d", test.N, got, test.want)
		}
	}
}

func TestJudge2(t *testing.T) {
	tests := []struct {
		N    int
		want int
	}{
		{N: 5, want: 0},
		{N: 1055, want: 0},
		{N: 1056, want: 1},
		{N: 5000000, want: 309},
	}
	for _, test := range tests {
		if got := judge2(xA, xB, test.N); got != test.want {
			t.Errorf("judge2(..., %d) = %d, want %d", test.N, got, test.want)
		}
	}
}
