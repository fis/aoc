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

package day13

import (
	"testing"
)

var ex = []string{
	"[1,1,3,1,1]",
	"[1,1,5,1,1]",
	"",
	"[[1],[2,3,4]]",
	"[[1],4]",
	"",
	"[9]",
	"[[8,7,6]]",
	"",
	"[[4,4],4,4]",
	"[[4,4],4,4,4]",
	"",
	"[7,7,7,7]",
	"[7,7,7]",
	"",
	"[]",
	"[3]",
	"",
	"[[[]]]",
	"[[]]",
	"",
	"[1,[2,[3,[4,[5,6,7]]]],8,9]",
	"[1,[2,[3,[4,[5,6,0]]]],8,9]",
}

func TestPart1(t *testing.T) {
	packets, err := parsePackets(ex)
	if err != nil {
		t.Fatal(err)
	}
	want := 13
	if got := part1(packets); got != want {
		t.Errorf("part1(ex) = %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	packets, err := parsePackets(ex)
	if err != nil {
		t.Fatal(err)
	}
	want := 140
	if got := part2(packets); got != want {
		t.Errorf("part2(ex) = %d, want %d", got, want)
	}
}
