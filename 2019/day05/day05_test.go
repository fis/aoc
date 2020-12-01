package day05

import (
	"testing"

	"github.com/fis/aoc-go/intcode"
)

func TestEx1(t *testing.T) {
	_, mem := intcode.Run([]int64{1002, 4, 3, 4, 33}, nil)
	if mem[4] != 99 {
		t.Errorf("Last instruction not set to halt: %d", mem[4])
	}
}

func TestEx2(t *testing.T) {
	tests := []struct {
		prog  []int64
		specs []testSpec
	}{
		{
			prog:  []int64{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
			specs: []testSpec{{7, 0}, {8, 1}, {9, 0}},
		},
		{
			prog:  []int64{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
			specs: []testSpec{{6, 1}, {7, 1}, {8, 0}, {9, 0}},
		},
		{
			prog:  []int64{3, 3, 1108, -1, 8, 3, 4, 3, 99},
			specs: []testSpec{{7, 0}, {8, 1}, {9, 0}},
		},
		{
			prog:  []int64{3, 3, 1107, -1, 8, 3, 4, 3, 99},
			specs: []testSpec{{6, 1}, {7, 1}, {8, 0}, {9, 0}},
		},
	}
	for _, test := range tests {
		for _, spec := range test.specs {
			got, _ := intcode.Run(test.prog, []int64{spec.input})
			if len(got) != 1 {
				t.Errorf("%v / %d: invalid output: %v", test.prog, spec.input, got)
			} else if n := got[0]; n != spec.want {
				t.Errorf("%v / %d = %d, want %d", test.prog, spec.input, got, spec.want)
			}
		}
	}
}

func TestEx3(t *testing.T) {
	progs := [][]int64{
		{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
		{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
	}
	specs := []testSpec{{-1, 1}, {0, 0}, {1, 1}, {2, 1}}
	for _, prog := range progs {
		for _, spec := range specs {
			got, _ := intcode.Run(prog, []int64{spec.input})
			if len(got) != 1 {
				t.Errorf("%v / %d: invalid output: %v", prog, spec.input, got)
			} else if n := got[0]; n != spec.want {
				t.Errorf("%v / %d = %d, want %d", prog, spec.input, got, spec.want)
			}
		}
	}
}

func TestEx4(t *testing.T) {
	progs := [][]int64{
		{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
		{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
	}
	specs := []testSpec{{-1, 1}, {0, 0}, {1, 1}, {2, 1}}
	for _, prog := range progs {
		for _, spec := range specs {
			got, _ := intcode.Run(prog, []int64{spec.input})
			if len(got) != 1 {
				t.Errorf("%v / %d: invalid output: %v", prog, spec.input, got)
			} else if n := got[0]; n != spec.want {
				t.Errorf("%v / %d = %d, want %d", prog, spec.input, got, spec.want)
			}
		}
	}
}

func TestEx5(t *testing.T) {
	prog := []int64{
		3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
		999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99,
	}
	specs := []testSpec{{6, 999}, {7, 999}, {8, 1000}, {9, 1001}, {10, 1001}}
	for _, spec := range specs {
		got, _ := intcode.Run(prog, []int64{spec.input})
		if len(got) != 1 {
			t.Errorf("%v / %d: invalid output: %v", prog, spec.input, got)
		} else if n := got[0]; n != spec.want {
			t.Errorf("%v / %d = %d, want %d", prog, spec.input, got, spec.want)
		}
	}
}

type testSpec struct{ input, want int64 }
