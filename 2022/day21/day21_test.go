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

package day21

import (
	"testing"
)

var ex = []string{
	"root: pppw + sjmn",
	"dbpl: 5",
	"cczh: sllz + lgvd",
	"zczc: 2",
	"ptdq: humn - dvpt",
	"dvpt: 3",
	"lfqf: 4",
	"humn: 5",
	"ljgn: 2",
	"sjmn: drzm * dbpl",
	"sllz: 4",
	"pppw: cczh / lfqf",
	"lgvd: ljgn * ptdq",
	"drzm: hmdt - zczc",
	"hmdt: 32",
}

func TestPart1(t *testing.T) {
	root, _ := parseJobs(ex)
	want := 152
	if got := root.num(); got != want {
		t.Errorf("root.num() = %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	root, human := parseJobs(ex)
	root.forceEqual()
	want := 301
	if got := human.num(); got != want {
		t.Errorf("human.num() = %d, want %d", got, want)
	}
}
