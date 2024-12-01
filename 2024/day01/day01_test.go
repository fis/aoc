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

package day01

import (
	"slices"
	"testing"
)

var (
	exA = slices.Sorted(slices.Values([]int{3, 4, 2, 1, 3, 3}))
	exB = slices.Sorted(slices.Values([]int{4, 3, 5, 3, 9, 3}))
)

func TestTotalDistance(t *testing.T) {
	want := 11
	got := totalDistance(exA, exB)
	if got != want {
		t.Errorf("totalDistance(exA, exB) = %d, want %d", got, want)
	}
}

func TestSimilarity(t *testing.T) {
	want := 31
	got := similarity(exA, exB)
	if got != want {
		t.Errorf("similarity(exA, exB) = %d, want %d", got, want)
	}
}
