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

package day14

import (
	"testing"

	"github.com/fis/aoc/util/fn"
)

var ex = []string{
	"498,4 -> 498,6 -> 496,6",
	"503,4 -> 502,4 -> 502,9 -> 494,9",
}

func TestAddSand(t *testing.T) {
	paths, err := fn.MapE(ex, parsePath)
	if err != nil {
		t.Fatal(err)
	}
	bmp, sourceX := buildBitmap(paths)
	want := 24
	if got := addSand(bmp, sourceX); got != want {
		t.Errorf("addSand(ex, %d) = %d, want %d", sourceX, got, want)
	}
}

func TestAddShadow(t *testing.T) {
	paths, err := fn.MapE(ex, parsePath)
	if err != nil {
		t.Fatal(err)
	}
	bmp, sourceX := buildBitmap(paths)
	want := 93
	if got := addShadow(bmp, sourceX); got != want {
		t.Errorf("addShadow(ex, %d) = %d, want %d", sourceX, got, want)
	}
}
