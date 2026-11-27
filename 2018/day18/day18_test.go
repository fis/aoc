package day18

import (
	"strings"
	"testing"

	"github.com/fis/aoc/util"
)

var example = `
.#.#...|#.
.....#|##|
.|..|...#.
..|#.....#
#.#|||#|#|
...#.||...
.|....|...
||...#|.#|
|.||||..|.
...#.|..|.
`

func TestEvolve(t *testing.T) {
	level := util.ParseLevelString(strings.TrimPrefix(example, "\n"), ' ')
	data, w, h := convertLevel(level)
	want := 1147
	got := value(evolve(data, w, h, 10))
	if got != want {
		t.Errorf("evolve(10) -> value %d, want %d", got, want)
	}
}
