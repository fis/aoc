package day02

import (
	"testing"

	"github.com/fis/aoc-go/intcode"
	"github.com/google/go-cmp/cmp"
)

func TestExamples(t *testing.T) {
	tests := []struct {
		prog []int64
		want []int64
	}{
		{
			prog: []int64{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
			want: []int64{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
		},
		{
			prog: []int64{1, 0, 0, 0, 99},
			want: []int64{2, 0, 0, 0, 99},
		},
		{
			prog: []int64{2, 3, 0, 3, 99},
			want: []int64{2, 3, 0, 6, 99},
		},
		{
			prog: []int64{2, 4, 4, 5, 99, 0},
			want: []int64{2, 4, 4, 5, 99, 9801},
		},
		{
			prog: []int64{1, 1, 1, 4, 99, 5, 6, 0, 99},
			want: []int64{30, 1, 1, 4, 2, 5, 6, 0, 99},
		},
	}
	for _, test := range tests {
		_, got := intcode.Run(test.prog, nil)
		if !cmp.Equal(got, test.want) {
			t.Errorf("%v -> %v, want %v", test.prog, got, test.want)
		}
	}
}
