package day09

import (
	"strings"
	"testing"

	"github.com/fis/aoc/util"
)

const ex = `
2199943210
3987894921
9856789892
8767896789
9899965678
`

func TestRiskLevels(t *testing.T) {
	level := util.ParseLevelString(strings.TrimPrefix(ex, "\n"), ' ')
	want := 15
	if got := riskLevels(level); got != want {
		t.Errorf("riskLevel = %d, want %d", got, want)
	}
}

func TestBasinSizes(t *testing.T) {
	level := util.ParseLevelString(strings.TrimPrefix(ex, "\n"), ' ')
	want := 1134
	if got := basinSizes(level); got != want {
		t.Errorf("basinSizes = %d, want %d", got, want)
	}
}
