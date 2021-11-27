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

package day02

import "testing"

func TestChecksumSheet(t *testing.T) {
	sheet := [][]int{{5, 1, 9, 5}, {7, 5, 3}, {2, 4, 6, 8}}
	want := 18
	if got := checksumSheet(sheet); got != want {
		t.Errorf("checksumSheet(%v) = %d, want %d", sheet, got, want)
	}
}

func TestChecksumRow(t *testing.T) {
	tests := []struct {
		row  []int
		want int
	}{
		{row: []int{5, 1, 9, 5}, want: 8},
		{row: []int{7, 5, 3}, want: 4},
		{row: []int{2, 4, 6, 8}, want: 6},
	}
	for _, test := range tests {
		if got := checksumRow(test.row); got != test.want {
			t.Errorf("checksumRow(%v) = %d, want %d", test.row, got, test.want)
		}
	}
}

func TestDivisibleSheet(t *testing.T) {
	sheet := [][]int{{5, 9, 2, 8}, {9, 4, 7, 3}, {3, 8, 6, 5}}
	want := 9
	if got, err := divisibleSheet(sheet); err != nil {
		t.Errorf("divisibleSheet(%v): %v", sheet, err)
	} else if got != want {
		t.Errorf("divisibleSheet(%v) = %d, want %d", sheet, got, want)
	}
}
func TestDivisibleRow(t *testing.T) {
	tests := []struct {
		row  []int
		want int
	}{
		{row: []int{5, 9, 2, 8}, want: 4},
		{row: []int{9, 4, 7, 3}, want: 3},
		{row: []int{3, 8, 6, 5}, want: 2},
	}
	for _, test := range tests {
		if got, err := divisibleRow(test.row); err != nil {
			t.Errorf("divisibleRow(%v): %v", test.row, err)
		} else if got != test.want {
			t.Errorf("divisibleRow(%v) = %d, want %d", test.row, got, test.want)
		}
	}
}
