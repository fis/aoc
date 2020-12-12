// Copyright 2019 Google LLC
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

package days

import (
	"fmt"
	"testing"

	"github.com/fis/aoc-go/glue"
	"github.com/google/go-cmp/cmp"
)

func TestAllDays(t *testing.T) {
	tests := []struct {
		day  int
		want []string
	}{
		{
			day:  1,
			want: []string{"3399947", "5097039"},
		},
		{
			day:  2,
			want: []string{"4138687", "6635"},
		},
		{
			day:  3,
			want: []string{"1431", "48012"},
		},
		{
			day:  4,
			want: []string{"979", "635"},
		},
		{
			day:  5,
			want: []string{"16434972", "16694270"},
		},
		{
			day:  6,
			want: []string{"292387", "433"},
		},
		{
			day:  7,
			want: []string{"277328", "11304734"},
		},
		{
			day: 8,
			want: []string{
				"1215",
				"#    #  #  ##  ###  #  # ",
				"#    #  # #  # #  # #  # ",
				"#    #### #    #  # #### ",
				"#    #  # #    ###  #  # ",
				"#    #  # #  # #    #  # ",
				"#### #  #  ##  #    #  # ",
			},
		},
		{
			day:  9,
			want: []string{"3638931938", "86025"},
		},
		{
			day:  10,
			want: []string{"230", "1205"},
		},
		{
			day: 11,
			want: []string{
				"2184",
				"..##..#..#..##..#..#.####.####.###..#..#.. ",
				" #..#.#..#.#..#.#..#....#.#....#..#.#.#....",
				" #..#.####.#....####...#..###..#..#.##.....",
				".####.#..#.#....#..#..#...#....###..#.#... ",
				".#..#.#..#.#..#.#..#.#....#....#....#.#..  ",
				" #..#.#..#..##..#..#.####.####.#....#..#.  ",
			},
		},
		{
			day:  12,
			want: []string{"10198", "271442326847376"},
		},
		{
			day:  13,
			want: []string{"420", "21651"},
		},
		{
			day:  14,
			want: []string{"365768", "3756877"},
		},
		{
			day:  15,
			want: []string{"330", "352"},
		},
		{
			day:  16,
			want: []string{"42205986", "13270205"},
		},
		{
			day:  17,
			want: []string{"3920", "673996"},
		},
		{
			day:  18,
			want: []string{"7430", "1864"},
		},
		{
			day:  19,
			want: []string{"226", "7900946"},
		},
		{
			day:  20,
			want: []string{"692", "8314"},
		},
		{
			day:  21,
			want: []string{"19352864", "1142488337"},
		},
		{
			day:  22,
			want: []string{"5169", "74258074061935"},
		},
		{
			day:  23,
			want: []string{"22650", "17298"},
		},
		{
			day:  24,
			want: []string{"27777901", "2047"},
		},
		{
			day: 25,
			want: []string{
				`Santa notices your small droid, looks puzzled for a moment, realizes what has happened, and radios your ship directly.`,
				`"Oh, hello! You should be able to get in by typing 134227456 on the keypad at the main airlock."`,
			},
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("day=%02d", test.day), func(t *testing.T) {
			if got, err := glue.SolveFile(2019, test.day, fmt.Sprintf("testdata/day%02d.txt", test.day)); err != nil {
				t.Errorf("Solve: %v", err)
			} else if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("Solve mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
