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

package day11

import "testing"

func TestPowerLevel(t *testing.T) {
	tests := []struct {
		x, y, s int
		want    int
	}{
		{x: 3, y: 5, s: 8, want: 4},
		{x: 122, y: 79, s: 57, want: -5},
		{x: 217, y: 196, s: 39, want: 0},
		{x: 101, y: 153, s: 71, want: 4},
	}
	for _, test := range tests {
		got := powerLevel(test.x, test.y, test.s)
		if got != test.want {
			t.Errorf("powerLevel(%d, %d, %d) = %d, want %d", test.x, test.y, test.s, got, test.want)
		}
	}
}

func TestSweetSpot(t *testing.T) {
	tests := []struct {
		s                     int
		wantX, wantY, wantPow int
	}{
		{s: 18, wantX: 33, wantY: 45, wantPow: 29},
		{s: 42, wantX: 21, wantY: 61, wantPow: 30},
	}
	for _, test := range tests {
		g := makeGrid(test.s)
		gotX, gotY, gotPow := sweetSpot(g)
		if gotX != test.wantX || gotY != test.wantY || gotPow != test.wantPow {
			t.Errorf("sweetSpot(%d) = (%d,%d,%d), want (%d,%d,%d)", test.s, gotX, gotY, gotPow, test.wantX, test.wantY, test.wantPow)
		}
	}
}

func TestSweetSquare(t *testing.T) {
	tests := []struct {
		s                               int
		wantX, wantY, wantSize, wantPow int
	}{
		{s: 18, wantX: 90, wantY: 269, wantSize: 16, wantPow: 113},
		{s: 42, wantX: 232, wantY: 251, wantSize: 12, wantPow: 119},
	}
	for _, test := range tests {
		g := makeGrid(test.s)
		gotX, gotY, gotSize, gotPow := sweetSquare(g)
		if gotX != test.wantX || gotY != test.wantY || gotSize != test.wantSize || gotPow != test.wantPow {
			t.Errorf("sweetSquare(%d) = (%d,%d,%d,%d), want (%d,%d,%d,%d)", test.s, gotX, gotY, gotSize, gotPow, test.wantX, test.wantY, test.wantSize, test.wantPow)
		}
	}
}
