package day13

import (
	"testing"

	"github.com/fis/aoc/util"
	"github.com/google/go-cmp/cmp"
)

var ex = struct {
	points []util.P
	folds  []foldSpec
}{
	points: []util.P{
		{6, 10}, {0, 14}, {9, 10}, {0, 3}, {10, 4}, {4, 11}, {6, 0}, {6, 12}, {4, 1},
		{0, 13}, {10, 12}, {3, 4}, {3, 0}, {8, 4}, {1, 10}, {2, 14}, {8, 10}, {9, 0},
	},
	folds: []foldSpec{{vertical: false, at: 7}, {vertical: true, at: 5}},
}

func TestFold(t *testing.T) {
	points := append([]util.P(nil), ex.points...)

	want1 := 17
	ex.folds[0].apply(points)
	if got1 := countVisible(points); got1 != want1 {
		t.Errorf("visible after first fold = %d, want %d", got1, want1)
	}

	want2 := []string{`#####`, `#   #`, `#   #`, `#   #`, `#####`}
	ex.folds[1].apply(points)
	if got2 := printPoints(points); !cmp.Equal(got2, want2) {
		t.Errorf("points after folding = %v, want %v", got2, want2)
	}
}
