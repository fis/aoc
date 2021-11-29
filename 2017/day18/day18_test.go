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

package day18

import (
	"testing"
)

func TestPart1(t *testing.T) {
	code := [][]string{
		{"set", "a", "1"},
		{"add", "a", "2"},
		{"mul", "a", "a"},
		{"mod", "a", "5"},
		{"snd", "a", ""},
		{"set", "a", "0"},
		{"rcv", "a", ""},
		{"jgz", "a", "-1"},
		{"set", "a", "1"},
		{"jgz", "a", "-2"},
	}
	prog := parseCode(code)
	want := 4
	if got := part1(prog); got != want {
		t.Errorf("part1 = %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	code := [][]string{
		{"snd", "1", ""},
		{"snd", "2", ""},
		{"snd", "p", ""},
		{"rcv", "a", ""},
		{"rcv", "b", ""},
		{"rcv", "c", ""},
		{"rcv", "d", ""},
	}
	prog := parseCode(code)
	want := 3
	if got := part2(prog); got != want {
		t.Errorf("part2 = %d, want %d", got, want)
	}
}
