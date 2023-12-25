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

package day24

import (
	"testing"

	"github.com/fis/aoc/util/fn"
)

var ex = []string{
	"19, 13, 30 @ -2,  1, -2",
	"18, 19, 22 @ -1, -1, -2",
	"20, 25, 34 @ -2, -2, -4",
	"12, 31, 28 @ -1, -2, -1",
	"20, 19, 15 @  1, -5, -3",
}

func TestCountIntersectXY(t *testing.T) {
	stones, err := fn.MapE(ex, parseHailstone)
	if err != nil {
		t.Fatal(err)
	}
	want := 2
	if got := countIntersectXY(stones, 7, 7, 27, 27); got != want {
		t.Errorf("countIntersectXY(ex, 7, 7, 27, 27) = %d, want %d", got, want)
	}
}

func TestFindCollider(t *testing.T) {
	stones, err := fn.MapE(ex, parseHailstone)
	if err != nil {
		t.Fatal(err)
	}
	want := p3{24, 13, 10}
	if got := findCollider(stones); got != want {
		t.Errorf("findCollider(ex) = %v, want %v", got, want)
	}
}
