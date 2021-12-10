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

package day09

import (
	"strings"
	"testing"

	"github.com/fis/aoc/util"
)

const ex = `
2199943210
3987894921
9856789892
8767896789
9899965678
`

func TestRiskLevels(t *testing.T) {
	level := util.ParseLevelString(strings.TrimPrefix(ex, "\n"), ' ')
	want := 15
	if got := riskLevels(level); got != want {
		t.Errorf("riskLevel = %d, want %d", got, want)
	}
}

func TestBasinSizes(t *testing.T) {
	level := util.ParseLevelString(strings.TrimPrefix(ex, "\n"), ' ')
	want := 1134
	if got := basinSizes(level); got != want {
		t.Errorf("basinSizes = %d, want %d", got, want)
	}
}
