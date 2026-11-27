package day19

import (
	"strings"
	"testing"

	"github.com/fis/aoc/2018/cpu"
	"github.com/fis/aoc/util"
)

var example = `
#ip 0
seti 5 0 1
seti 6 0 2
addi 0 1 0
addr 1 2 3
setr 1 0 0
seti 8 0 4
seti 9 0 5
`

func TestRun(t *testing.T) {
	prog, err := cpu.ParseProg(util.Lines(strings.TrimPrefix(example, "\n")))
	if err != nil {
		t.Fatalf("ParseProg: %v", err)
	}
	want := [6]int{6, 5, 6, 0, 0, 9}
	s := cpu.State{}
	s.Run(prog)
	if s.R != want {
		t.Errorf("Run -> %v, want %v", s.R, want)
	}
}

func TestSumDiv(t *testing.T) {
	got := sumDiv(939)
	if got != 1256 {
		t.Errorf("sumDiv(939) = %d, want 1256", got)
	}
}
