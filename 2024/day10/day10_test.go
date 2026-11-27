package day10

import (
	"strings"
	"testing"

	"github.com/fis/aoc/util"
)

var ex = strings.TrimPrefix(`
89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732
`, "\n")

func TestTotalScore(t *testing.T) {
	level := util.ParseFixedLevel([]byte(ex))
	want1, want2 := 36, 81
	if got1, got2 := totalScore(level); got1 != want1 || got2 != want2 {
		t.Errorf("totalScore(ex) = (%d,%d), want (%d,%d)", got1, got2, want1, want2)
	}
}
