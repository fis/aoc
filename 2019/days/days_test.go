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
	"testing"

	"github.com/fis/aoc/glue"
)

func TestAllDays(t *testing.T) {
	tests := []glue.TestCase{
		{
			Day:  1,
			Want: []string{"3399947", "5097039"},
		},
		{
			Day:  2,
			Want: []string{"4138687", "6635"},
		},
		{
			Day:  3,
			Want: []string{"1431", "48012"},
		},
		{
			Day:  4,
			Want: []string{"979", "635"},
		},
		{
			Day:  5,
			Want: []string{"16434972", "16694270"},
		},
		{
			Day:  6,
			Want: []string{"292387", "433"},
		},
		{
			Day:  7,
			Want: []string{"277328", "11304734"},
		},
		{
			Day: 8,
			Want: []string{
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
			Day:  9,
			Want: []string{"3638931938", "86025"},
		},
		{
			Day:  10,
			Want: []string{"230", "1205"},
		},
		{
			Day: 11,
			Want: []string{
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
			Day:  12,
			Want: []string{"10198", "271442326847376"},
		},
		{
			Day:  13,
			Want: []string{"420", "21651"},
		},
		{
			Day:  14,
			Want: []string{"365768", "3756877"},
		},
		{
			Day:  15,
			Want: []string{"330", "352"},
		},
		{
			Day:  16,
			Want: []string{"42205986", "13270205"},
		},
		{
			Day:  17,
			Want: []string{"3920", "673996"},
		},
		{
			Day:  18,
			Want: []string{"7430", "1864"},
		},
		{
			Day:  19,
			Want: []string{"226", "7900946"},
		},
		{
			Day:  20,
			Want: []string{"692", "8314"},
		},
		{
			Day:  21,
			Want: []string{"19352864", "1142488337"},
		},
		{
			Day:  22,
			Want: []string{"5169", "74258074061935"},
		},
		{
			Day:  23,
			Want: []string{"22650", "17298"},
		},
		{
			Day:  24,
			Want: []string{"27777901", "2047"},
		},
		{
			Day: 25,
			Want: []string{
				`Santa notices your small droid, looks puzzled for a moment, realizes what has happened, and radios your ship directly.`,
				`"Oh, hello! You should be able to get in by typing 134227456 on the keypad at the main airlock."`,
			},
		},
	}
	glue.RunTests(t, tests, 2019)
}
