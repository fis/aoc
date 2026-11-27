package day13

import (
	"testing"
)

var ex = []string{
	"[1,1,3,1,1]",
	"[1,1,5,1,1]",
	"",
	"[[1],[2,3,4]]",
	"[[1],4]",
	"",
	"[9]",
	"[[8,7,6]]",
	"",
	"[[4,4],4,4]",
	"[[4,4],4,4,4]",
	"",
	"[7,7,7,7]",
	"[7,7,7]",
	"",
	"[]",
	"[3]",
	"",
	"[[[]]]",
	"[[]]",
	"",
	"[1,[2,[3,[4,[5,6,7]]]],8,9]",
	"[1,[2,[3,[4,[5,6,0]]]],8,9]",
}

func TestPart1(t *testing.T) {
	packets, err := parsePackets(ex)
	if err != nil {
		t.Fatal(err)
	}
	want := 13
	if got := part1(packets); got != want {
		t.Errorf("part1(ex) = %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	packets, err := parsePackets(ex)
	if err != nil {
		t.Fatal(err)
	}
	want := 140
	if got := part2(packets); got != want {
		t.Errorf("part2(ex) = %d, want %d", got, want)
	}
}
