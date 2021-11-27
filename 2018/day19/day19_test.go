// Copyright 2020 Google LLC
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

package day19

import (
	"strings"
	"testing"

	"github.com/fis/aoc/2018/cpu"
	"github.com/fis/aoc/util"
)

var example = `
#ip 0
seti 5 0 1
seti 6 0 2
addi 0 1 0
addr 1 2 3
setr 1 0 0
seti 8 0 4
seti 9 0 5
`

func TestRun(t *testing.T) {
	prog, err := cpu.ParseProg(util.Lines(strings.TrimPrefix(example, "\n")))
	if err != nil {
		t.Fatalf("ParseProg: %v", err)
	}
	want := [6]int{6, 5, 6, 0, 0, 9}
	s := cpu.State{}
	s.Run(prog)
	if s.R != want {
		t.Errorf("Run -> %v, want %v", s.R, want)
	}
}

func TestSumDiv(t *testing.T) {
	got := sumDiv(939)
	if got != 1256 {
		t.Errorf("sumDiv(939) = %d, want 1256", got)
	}
}
