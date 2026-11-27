package day03

import (
	"strings"
	"testing"

	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/fn"
)

var ex = strings.TrimPrefix(`
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
`, "\n")

func TestSumNumbers(t *testing.T) {
	tests := []struct {
		symFilter byte
		f         func([]int) int
		ft        string
		want      int
	}{
		{
			symFilter: 0,
			f:         fn.Sum[[]int], ft: "sum",
			want: 4361,
		},
		{
			symFilter: '*',
			f:         gearRatio, ft: "gearRatio",
			want: 467835,
		},
	}
	lv := util.ParseLevelString(ex, '.')
	for _, test := range tests {
		if got := sumNumbers(lv, test.symFilter, test.f); got != test.want {
			t.Errorf("sumNumbers(ex, %d, %s) = %d, want %d", test.symFilter, test.ft, got, test.want)
		}
	}
}
