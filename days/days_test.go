package days

import (
	"fmt"
	"testing"

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
			day:  25,
			want: []string{"134227456"},
		},
	}

	for _, test := range tests {
		got, err := Solve(test.day, fmt.Sprintf("testdata/day%02d.txt", test.day))
		if err != nil {
			t.Errorf("Day %d failed: %v", test.day, err)
		} else if diff := cmp.Diff(test.want, got); diff != "" {
			t.Errorf("Day %d mismatch (-want +got):\n%s", test.day, diff)
		}
	}
}
