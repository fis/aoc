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

package day11

import "testing"

var (
	ex1 = []string{
		"you: bbb ccc",
		"bbb: ddd eee",
		"ccc: ddd eee fff",
		"ddd: ggg",
		"eee: out",
		"fff: out",
		"ggg: out",
		"hhh: ccc fff iii",
		"iii: out",
	}
	ex2 = []string{
		"svr: aaa bbb",
		"aaa: fft",
		"fft: ccc",
		"bbb: tty",
		"tty: ccc",
		"ccc: ddd eee",
		"ddd: hub",
		"hub: fff",
		"eee: dac",
		"dac: fff",
		"fff: ggg hhh",
		"ggg: out",
		"hhh: out",
	}
)

func TestCountPaths(t *testing.T) {
	g := buildGraph(ex1)
	want := 5
	if got := newPathCounter(g, "you").countTo("out"); got != want {
		t.Errorf("pc(ex1, you).countTo(out) = %d, want %d", got, want)
	}
}

func TestCountValidPaths(t *testing.T) {
	g := buildGraph(ex2)
	want := 2
	if got := newValidPathCounter(g, "svr").countTo("out"); got != want {
		t.Errorf("vpc(ex2, svr).countTo(out) = %d, want %d", got, want)
	}
}
