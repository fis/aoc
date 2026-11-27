package day18

import (
	"strings"
	"testing"

	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/fn"
)

var ex = strings.TrimPrefix(`
5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0
`, "\n")

func TestFindPath(t *testing.T) {
	incoming, _ := fn.MapE(util.Lines(ex), util.ParseP)
	level := simulate(7, incoming[:12])
	want := 22 + 1 // number of tiles vs. number of steps
	if got := len(findPath(level, nil)); got != want {
		t.Errorf("len(findPath(simulate(7, ex[:12]))) = %d, want %d", got, want)
	}
}

func TestFindBreak(t *testing.T) {
	incoming, _ := fn.MapE(util.Lines(ex), util.ParseP)
	level := simulate(7, incoming[:12])
	path := findPath(level, nil)
	want := util.P{6, 1}
	if got := findBreak(level, path, incoming[12:]); got != want {
		t.Errorf("findBreak(...) = %v, want %v", got, want)
	}
}
