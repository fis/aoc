package day05

import (
	"testing"
)

var ex = []string{
	"    [D]    ",
	"[N] [C]    ",
	"[Z] [M] [P]",
	" 1   2   3 ",
	"",
	"move 1 from 2 to 1",
	"move 3 from 1 to 3",
	"move 2 from 2 to 1",
	"move 1 from 1 to 2",
}

func TestApplyMoves9000(t *testing.T) {
	stacks, moves := parseInput(ex)
	want := "CMZ"
	stacks = applyMoves(stacks, moves, move.apply9000)
	if got := tops(stacks); got != want {
		t.Errorf("applyMoves(ex) -> %s, want %s", got, want)
	}
}

func TestApplyMoves9001(t *testing.T) {
	stacks, moves := parseInput(ex)
	want := "MCD"
	stacks = applyMoves(stacks, moves, move.apply9001)
	if got := tops(stacks); got != want {
		t.Errorf("applyMoves(ex) -> %s, want %s", got, want)
	}
}
