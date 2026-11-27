package day10

import (
	"strings"
	"testing"

	"github.com/fis/aoc/util"
)

func TestFindBot(t *testing.T) {
	input, _ := util.ScanAllRegexp(strings.NewReader(strings.TrimPrefix(ex, "\n")), inputRegexp)
	g := buildGraph(input)
	want := 2
	if got := g.findBot(5, 2); got != want {
		t.Errorf("g.findBot(5, 2) = %d, want %d", got, want)
	}
}
