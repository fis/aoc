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
			want: []string{"582", "488"},
		},
		{
			day:  2,
			want: []string{"6448", "evsialkqyiurohzpwucngttmf"},
		},
		{
			day:  3,
			want: []string{"110891", "297"},
		},
		{
			day:  4,
			want: []string{"21956", "134511"},
		},
		{
			day:  5,
			want: []string{"10878", "6874"},
		},
		{
			day:  6,
			want: []string{"4186", "45509"},
		},
		{
			day:  7,
			want: []string{"BGKDMJCNEQRSTUZWHYLPAFIVXO", "941"},
		},
		{
			day:  8,
			want: []string{"48496", "32850"},
		},
		{
			day:  9,
			want: []string{"374287", "3083412635"},
		},
		{
			day: 10,
			want: []string{
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
			day:  17,
			want: []string{"27736", "22474"},
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("day=%02d", test.day), func(t *testing.T) {
			if got, err := glue.SolveFile(2018, test.day, fmt.Sprintf("testdata/day%02d.txt", test.day)); err != nil {
				t.Errorf("Solve: %v", err)
			} else if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("Solve mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
