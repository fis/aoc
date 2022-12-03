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

package day03

import (
	"testing"
)

var ex = []string{
	"vJrwpWtwJgWrhcsFMMfFFhFp",
	"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
	"PmmdzqPrVvPwwTWBwg",
	"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
	"ttgJtRGJQctTZtZT",
	"CrZsJsPPZsGzwwsLwLmpwMDw",
}

func TestPart1(t *testing.T) {
	sacks := pack(ex)
	want := 157
	if got := part1(sacks); got != want {
		t.Errorf("part1(%v) = %d, want %d", sacks, got, want)
	}
}

func TestPart2(t *testing.T) {
	sacks := pack(ex)
	want := 70
	if got := part2(sacks); got != want {
		t.Errorf("part2(%v) = %d, want %d", sacks, got, want)
	}
}
