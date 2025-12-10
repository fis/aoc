// Copyright 2025 Google LLC
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
	"slices"
	"testing"

	"github.com/fis/aoc/util"
)

var ex = []util.P{
	{7, 1},
	{11, 1},
	{11, 7},
	{9, 7},
	{9, 5},
	{2, 5},
	{2, 3},
	{7, 3},
}

func TestLargestRect(t *testing.T) {
	want := 50
	if got := largestRect(ex); got != want {
		t.Errorf("largestRect(ex) = %d, want %d", got, want)
	}
}

func TestLargestValid(t *testing.T) {
	want := 24
	if got := largestValid(ex); got != want {
		t.Errorf("largestValid(ex) = %d, want %d", got, want)
	}
}

func TestLargestValidReversed(t *testing.T) {
	rev := slices.Clone(ex)
	slices.Reverse(rev)
	want := 24
	if got := largestValid(rev); got != want {
		t.Errorf("largestValid(rev) = %d, want %d", got, want)
	}
}
