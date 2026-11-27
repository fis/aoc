package day07

import (
	"testing"

	"github.com/fis/aoc/util"
	"github.com/google/go-cmp/cmp"
)

func TestToplexSort(t *testing.T) {
	g := parseRules(util.Lines(example))
	want := []string{"C", "A", "B", "D", "F", "E"}
	got := toplexSort(g)
	if !cmp.Equal(got, want) {
		t.Errorf("topLexSort = %v, want %v", got, want)
	}
}

func TestTimedSort(t *testing.T) {
	g := parseRules(util.Lines(example))
	wantOrder, wantTime := []string{"C", "A", "B", "F", "D", "E"}, 15
	gotOrder, gotTime := timedSort(g, 2, 0)
	if !cmp.Equal(gotOrder, wantOrder) || gotTime != wantTime {
		t.Errorf("timedSort = (%v, %d), want (%v, %d)", gotOrder, gotTime, wantOrder, wantTime)
	}
}
