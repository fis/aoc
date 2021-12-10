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

package day10

import (
	"testing"
)

func TestCheck(t *testing.T) {
	tests := []struct {
		line          string
		wantCorrupted bool
		wantScore     int
	}{
		{line: "[({(<(())[]>[[{[]{<()<>>", wantCorrupted: false, wantScore: 288957},
		{line: "[(()[<>])]({[<{<<[]>>(", wantCorrupted: false, wantScore: 5566},
		{line: "{([(<{}[<>[]}>{[]{[(<()>", wantCorrupted: true, wantScore: 1197},
		{line: "(((({<>}<{<{<>}{[]{[]{}", wantCorrupted: false, wantScore: 1480781},
		{line: "[[<[([]))<([[{}[[()]]]", wantCorrupted: true, wantScore: 3},
		{line: "[{[{({}]{}}([{[{{{}}([]", wantCorrupted: true, wantScore: 57},
		{line: "{<[[]]>}<{[{[{[]{()[[[]", wantCorrupted: false, wantScore: 995444},
		{line: "[<(<(<(<{}))><([]([]()", wantCorrupted: true, wantScore: 3},
		{line: "<{([([[(<>()){}]>(<<{{", wantCorrupted: true, wantScore: 25137},
		{line: "<{([{{}}[<[[[<>{}]]]>[]]", wantCorrupted: false, wantScore: 294},
	}
	for _, test := range tests {
		corrupted, score := check(test.line)
		if corrupted != test.wantCorrupted || score != test.wantScore {
			t.Errorf("check(%q) = (%t, %d), want (%t, %d)", test.line, corrupted, score, test.wantCorrupted, test.wantScore)
		}
	}
}
