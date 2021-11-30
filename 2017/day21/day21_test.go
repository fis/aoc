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

package day21

import (
	"testing"
)

func TestIterate(t *testing.T) {
	book, err := parseBook([][]string{
		{"../.#", "##./#../..."},
		{".#./..#/###", "#..#/..../..../#..#"},
	})
	if err != nil {
		t.Fatal(err)
	}

	bmp := newBitmap(3)
	bmp.put3(0, 0, rootTile)
	bmp = iterate(bmp, book, 2)

	want := 12
	if got := bmp.popCount(); got != want {
		t.Errorf("iterate(2) -> %d pixels, want %d", got, want)
	}
}

func TestSplitBits(t *testing.T) {
	t3 := tile3(0b100_010_001)

	// scenario 1:
	// #.|.
	// .#|.
	// ..|#

	bmp := newBitmap(64)
	bmp.put3(30, 0, t3)
	if got := bmp.get3(30, 0); got != t3 {
		t.Errorf("scenario 1 get3 = %x, want %x", got, t3)
	}
	for _, test := range []struct {
		x, y int
		t    tile2
	}{
		{x: 30, y: 0, t: 0b10_01},
		{x: 30, y: 2, t: 0},
		{x: 32, y: 0, t: 0},
		{x: 32, y: 2, t: 0b10_00},
	} {
		if got := bmp.get2(test.x, test.y); got != test.t {
			t.Errorf("scenario 1 get2(%d,%d) = %x, want %x", test.x, test.y, got, test.t)
		}
	}

	// scenario 2:
	// #|..
	// .|#.
	// .|.#

	bmp = newBitmap(64)
	bmp.put3(31, 0, t3)
	if got := bmp.get3(31, 0); got != t3 {
		t.Errorf("scenario 2 get3 = %x, want %x", got, t3)
	}
	for _, test := range []struct {
		x, y int
		t    tile2
	}{
		{x: 30, y: 0, t: 0b01_00},
		{x: 30, y: 2, t: 0},
		{x: 32, y: 0, t: 0b00_10},
		{x: 32, y: 2, t: 0b01_00},
	} {
		if got := bmp.get2(test.x, test.y); got != test.t {
			t.Errorf("scenario 2 get2(%d,%d) = %x, want %x", test.x, test.y, got, test.t)
		}
	}
}
