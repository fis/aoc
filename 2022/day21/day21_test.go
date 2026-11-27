package day21

import (
	"testing"
)

var ex = []string{
	"root: pppw + sjmn",
	"dbpl: 5",
	"cczh: sllz + lgvd",
	"zczc: 2",
	"ptdq: humn - dvpt",
	"dvpt: 3",
	"lfqf: 4",
	"humn: 5",
	"ljgn: 2",
	"sjmn: drzm * dbpl",
	"sllz: 4",
	"pppw: cczh / lfqf",
	"lgvd: ljgn * ptdq",
	"drzm: hmdt - zczc",
	"hmdt: 32",
}

func TestPart1(t *testing.T) {
	root, _ := parseJobs(ex)
	want := 152
	if got := root.num(); got != want {
		t.Errorf("root.num() = %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	root, human := parseJobs(ex)
	root.forceEqual()
	want := 301
	if got := human.num(); got != want {
		t.Errorf("human.num() = %d, want %d", got, want)
	}
}
