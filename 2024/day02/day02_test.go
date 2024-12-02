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

package day02

import (
	"testing"
)

var ex = [][]int{
	{7, 6, 4, 2, 1},
	{1, 2, 7, 8, 9},
	{9, 7, 6, 2, 1},
	{1, 3, 2, 4, 5},
	{8, 6, 4, 4, 1},
	{1, 3, 6, 7, 9},
}

func TestIsSafe(t *testing.T) {
	tests := []struct {
		report []int
		dampen bool
		want   bool
	}{
		{report: ex[0], want: true},
		{report: ex[1], want: false},
		{report: ex[2], want: false},
		{report: ex[3], want: false},
		{report: ex[4], want: false},
		{report: ex[5], want: true},
	}
	for _, test := range tests {
		if got := isSafe(test.report); got != test.want {
			t.Errorf("isSafe(%v) = %t, want %t", test.report, got, test.want)
		}
	}
}

func TestIsSafeDampened(t *testing.T) {
	tests := []struct {
		report []int
		dampen bool
		want   bool
	}{
		{report: ex[0], want: true},
		{report: ex[1], want: false},
		{report: ex[2], want: false},
		{report: ex[3], want: true},
		{report: ex[4], want: true},
		{report: ex[5], want: true},
	}
	for _, test := range tests {
		if got := isSafeDampened(test.report); got != test.want {
			t.Errorf("isSafe(%v) = %t, want %t", test.report, got, test.want)
		}
	}
}
