package day09

import (
	"testing"

	"github.com/fis/aoc2019-go/intcode"
	"github.com/google/go-cmp/cmp"
)

func TestExamples(t *testing.T) {
	tests := []struct {
		prog []int64
		want []int64
	}{
		{
			prog: []int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
			want: []int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
		},
		{
			prog: []int64{1102, 34915192, 34915192, 7, 4, 7, 99, 0},
			want: []int64{1219070632396864},
		},
		{
			prog: []int64{104, 1125899906842624, 99},
			want: []int64{1125899906842624},
		},
	}
	for _, test := range tests {
		vm := intcode.VM{}
		vm.Load(test.prog)
		got := vm.Run([]int64{})
		if !cmp.Equal(got, test.want) {
			t.Errorf("%v -> %v, want %v", test.prog, got, test.want)
		}
	}
}
