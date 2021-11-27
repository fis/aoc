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

package day25

import (
	"strings"
	"testing"

	"github.com/fis/aoc/util"
)

var ex1 = `
0,0,0,0
3,0,0,0
0,3,0,0
0,0,3,0
0,0,0,3
0,0,0,6
9,0,0,0
12,0,0,0
`

var ex2 = `
-1,2,2,0
0,0,2,-2
0,0,0,-2
-1,2,0,0
-2,-2,-2,2
3,0,2,-1
-1,3,2,2
-1,0,-1,0
0,2,1,-2
3,0,0,0
`

var ex3 = `
1,-1,0,1
2,0,-1,0
3,2,-1,0
0,0,3,1
0,0,-1,-1
2,3,-2,0
-2,2,0,0
2,-2,0,-1
1,-1,0,-1
3,2,0,2
`

var ex4 = `
1,-1,-1,-2
-2,-2,0,1
0,2,1,3
-2,3,-2,1
0,2,3,-2
-1,-1,1,-2
0,-2,-1,0
-2,2,3,-1
1,2,2,0
-1,-2,0,-2
`

func TestConstallationSet(t *testing.T) {
	tests := []struct {
		name   string
		points string
		want   int
	}{
		{name: "ex1", points: ex1, want: 2},
		{name: "ex2", points: ex2, want: 4},
		{name: "ex3", points: ex3, want: 3},
		{name: "ex4", points: ex4, want: 8},
	}
	for _, test := range tests {
		cs := constellationSet{}
		cs.addAll(util.Lines(strings.TrimPrefix(test.points, "\n")))
		got := cs.size()
		if got != test.want {
			t.Errorf("%s: got %d constellations, want %d", test.name, got, test.want)
		}
	}
}
