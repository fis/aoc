package day12

import (
	"testing"

	"github.com/fis/aoc/util/fn"
)

func TestRunProgram(t *testing.T) {
	ex := []string{
		"cpy 41 a",
		"inc a",
		"inc a",
		"dec a",
		"jnz a 2",
		"dec a",
	}
	prog, err := fn.MapE(ex, parseInst)
	if err != nil {
		t.Fatal(err)
	}
	want := 42
	if got := runProgram(prog, 0); got != want {
		t.Errorf("runProgram(ex) = %d, want %d", got, want)
	}
}
