package day12

import (
	"testing"
)

var ex = []string{
	"0 <-> 2",
	"1 <-> 1",
	"2 <-> 0, 3, 4",
	"3 <-> 2, 4",
	"4 <-> 2, 3, 6",
	"5 <-> 6",
	"6 <-> 4, 5",
}

func TestPartition(t *testing.T) {
	want1, want2 := 6, 2
	g, err := buildGraph(ex)
	if err != nil {
		t.Fatal(err)
	}
	vertGroup, groupVerts := partition(g)
	zero, _ := g.V("0")
	got1, got2 := len(groupVerts[vertGroup[zero]]), len(groupVerts)
	if got1 != want1 || got2 != want2 {
		t.Errorf("part1 = %d, want %d; part2 = %d, want %d", got1, want1, got2, want2)
	}
}
