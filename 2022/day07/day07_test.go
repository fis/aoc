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

package day07

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var ex = []string{
	"$ cd /",
	"$ ls",
	"dir a",
	"14848514 b.txt",
	"8504156 c.dat",
	"dir d",
	"$ cd a",
	"$ ls",
	"dir e",
	"29116 f",
	"2557 g",
	"62596 h.lst",
	"$ cd e",
	"$ ls",
	"584 i",
	"$ cd ..",
	"$ cd ..",
	"$ cd d",
	"$ ls",
	"4060174 j",
	"8033020 d.log",
	"5626152 d.ext",
	"7214296 k",
}

func TestParseListing(t *testing.T) {
	want := []int{584, 94853, 24933642, 48381165}
	got := parseListing(ex[1:])
	if !cmp.Equal(got, want) {
		t.Errorf("parseListing(ex) = %v, want %v", got, want)
	}
}

func TestPart1(t *testing.T) {
	sizes := parseListing(ex[1:])
	want := 95437
	if got := part1(sizes); got != want {
		t.Errorf("part1(%v) = %d, want %d", sizes, got, want)
	}
}

func TestPart2(t *testing.T) {
	sizes := parseListing(ex[1:])
	want := 24933642
	if got := part2(sizes); got != want {
		t.Errorf("part2(%v) = %d, want %d", sizes, got, want)
	}
}
