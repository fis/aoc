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

package day12

import (
	"testing"
)

var ex = []string{
	"Sabqponm",
	"abcryxxl",
	"accszExk",
	"acctuvwj",
	"abdefghi",
}

func TestShortestPath(t *testing.T) {
	m, start, end := readMap(ex)
	want := 31
	if got := shortestPath(m, start, end); got != want {
		t.Errorf("shortestPath(ex) = %d, want %d", got, want)
	}
}

func TestScenicPath(t *testing.T) {
	m, _, end := readMap(ex)
	want := 29
	if got := scenicPath(m, end); got != want {
		t.Errorf("scenicPath(ex) = %d, want %d", got, want)
	}
}
