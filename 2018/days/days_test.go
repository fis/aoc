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

var tests = []glue.TestCase{
	{
		Day:  1,
		Want: []string{"582", "488"},
	},
	{
		Day:  2,
		Want: []string{"6448", "evsialkqyiurohzpwucngttmf"},
	},
	{
		Day:  3,
		Want: []string{"110891", "297"},
	},
	{
		Day:  4,
		Want: []string{"21956", "134511"},
	},
	{
		Day:  5,
		Want: []string{"10878", "6874"},
	},
	{
		Day:  6,
		Want: []string{"4186", "45509"},
	},
	{
		Day:  7,
		Want: []string{"BGKDMJCNEQRSTUZWHYLPAFIVXO", "941"},
	},
	{
		Day:  8,
		Want: []string{"48496", "32850"},
	},
	{
		Day:  9,
		Want: []string{"374287", "3083412635"},
	},
	{
		Day: 10,
		Want: []string{
			"#####   #        ####   #    #  #    #  #####      ###   #### ",
			"#    #  #       #    #  ##   #  #    #  #    #      #   #    #",
			"#    #  #       #       ##   #  #    #  #    #      #   #     ",
			"#    #  #       #       # #  #  #    #  #    #      #   #     ",
			"#####   #       #       # #  #  ######  #####       #   #     ",
			"#    #  #       #  ###  #  # #  #    #  #           #   #     ",
			"#    #  #       #    #  #  # #  #    #  #           #   #     ",
			"#    #  #       #    #  #   ##  #    #  #       #   #   #     ",
			"#    #  #       #   ##  #   ##  #    #  #       #   #   #    #",
			"#####   ######   ### #  #    #  #    #  #        ###     #### ",
			"10476"},
	},
	{
		Day:  11,
		Want: []string{"235,18", "236,227,12"},
	},
	{
		Day:  12,
		Want: []string{"1696", "1799999999458"},
	},
	{
		Day:  13,
		Want: []string{"40,90", "65,81"},
	},
	{
		Day:  14,
		Want: []string{"2157138126", "20365081"},
	},
	{
		Day:  15,
		Want: []string{"228730", "33621"},
	},
	{
		Day:  16,
		Want: []string{"529", "573"},
	},
	{
		Day:  17,
		Want: []string{"27736", "22474"},
	},
	{
		Day:  18,
		Want: []string{"621205", "228490"},
	},
	{
		Day:  19,
		Want: []string{"1256", "16137576"},
	},
	{
		Day:  20,
		Want: []string{"3545", "7838"},
	},
	{
		Day:  21,
		Want: []string{"11474091", "4520776"},
	},
	{
		Day:  22,
		Want: []string{"5400", "1048"},
	},
	{
		Day:  23,
		Want: []string{"619", "71631000"},
	},
	{
		Day:  24,
		Want: []string{"26937", "4893"},
	},
	{
		Day:  25,
		Want: []string{"420"},
	},
}

func TestAllDays(t *testing.T) {
	glue.RunTests(t, tests, 2018)
}

func BenchmarkAllDays(b *testing.B) {
	glue.RunBenchmarks(b, tests, 2018)
}
