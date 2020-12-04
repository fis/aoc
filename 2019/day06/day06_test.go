// Copyright 2019 Google LLC
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

package day06

import (
	"testing"
)

func TestExample1(t *testing.T) {
	data := []string{
		"COM)B",
		"B)C",
		"C)D",
		"D)E",
		"E)F",
		"B)G",
		"G)H",
		"D)I",
		"E)J",
		"J)K",
		"K)L",
	}
	om := parseOrbits(data)
	count := om.countOrbits()
	if count != 42 {
		t.Errorf("countOrbits(%v) = %d, want 42", data, count)
	}
}

func TestExample2(t *testing.T) {
	data := []string{
		"COM)B",
		"B)C",
		"C)D",
		"D)E",
		"E)F",
		"B)G",
		"G)H",
		"D)I",
		"E)J",
		"J)K",
		"K)L",
		"K)YOU",
		"I)SAN",
	}
	om := parseOrbits(data)
	dist := om.transfers("YOU", "SAN")
	if dist != 4 {
		t.Errorf("transfers(%v, YOU, SAN) = %d, want 4", data, dist)
	}
}
