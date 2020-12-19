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

package day13

import (
	"strings"
	"testing"

	"github.com/fis/aoc-go/util"
)

var ex1 = `
/->-\        
|   |  /----\
| /-+--+-\  |
| | |  | v  |
\-+-/  \-+--/
  \------/
`

var ex2 = `
/>-<\  
|   |  
| /<+-\
| | | v
\>+</ |
  |   ^
  \<->/
`

func TestSimulate(t *testing.T) {
	tests := []struct {
		name         string
		level        string
		keepGoing    bool
		wantX, wantY int
	}{
		{name: "ex1", level: ex1, keepGoing: false, wantX: 7, wantY: 3},
		{name: "ex2", level: ex2, keepGoing: true, wantX: 6, wantY: 4},
	}
	for _, test := range tests {
		level := util.ParseLevelString(strings.TrimPrefix(test.level, "\n"), ' ')
		gotX, gotY := simulate(level, test.keepGoing)
		if gotX != test.wantX || gotY != test.wantY {
			t.Errorf("simulate(%s, %v) = (%d,%d), want (%d,%d)", test.name, test.keepGoing, gotX, gotY, test.wantX, test.wantY)
		}
	}
}
