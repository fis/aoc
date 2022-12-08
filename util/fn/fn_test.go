package fn

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSum(t *testing.T) {
	tests := []struct {
		data []int
		want int
	}{
		{data: nil, want: 0},
		{data: []int{123}, want: 123},
		{data: []int{123, 456}, want: 579},
		{data: []int{1, 2, 3, 4, 5, 6, 7}, want: 28},
	}
	for _, test := range tests {
		got := Sum(test.data)
		if got != test.want {
			t.Errorf("Sum(%v) = %d, want %d", test.data, got, test.want)
		}
	}
}

func TestCountIf(t *testing.T) {
	tests := []struct {
		data []int
		f    func(int) bool
		want int
	}{
		{
			data: []int{1, 2, 3, 4, 5, 6, 7},
			f:    func(x int) bool { return x%2 == 1 },
			want: 4,
		},
	}
	for _, test := range tests {
		got := CountIf(test.data, test.f)
		if got != test.want {
			t.Errorf("CountIf(%v, ...) = %d, want %d", test.data, got, test.want)
		}
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		data []int
		want int
	}{
		{data: []int{123}, want: 123},
		{data: []int{123, 456, 789}, want: 123},
		{data: []int{123, 789, 456}, want: 123},
		{data: []int{789, 456, 123}, want: 123},
	}
	for _, test := range tests {
		got := Min(test.data)
		if got != test.want {
			t.Errorf("Min(%v) = %d, want %d", test.data, got, test.want)
		}
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		data []int
		want int
	}{
		{data: []int{123}, want: 123},
		{data: []int{123, 456, 789}, want: 789},
		{data: []int{123, 789, 456}, want: 789},
		{data: []int{789, 456, 123}, want: 789},
	}
	for _, test := range tests {
		got := Max(test.data)
		if got != test.want {
			t.Errorf("Max(%v) = %d, want %d", test.data, got, test.want)
		}
	}
}

func TestHead(t *testing.T) {
	tests := []struct {
		data []int
		def  int
		want int
	}{
		{nil, 999, 999},
		{[]int{123}, 999, 123},
		{[]int{123, 456, 789}, 999, 123},
	}
	for _, test := range tests {
		if got := Head(test.data, test.def); got != test.want {
			t.Errorf("Head(%v, %d) = %d, want %d", test.data, test.def, got, test.want)
		}
	}
}

func TestMap(t *testing.T) {
	tests := []struct {
		data []int
		f    func(int) int
		want []int
	}{
		{
			data: []int{1, 2, 3},
			f:    func(x int) int { return x + 1 },
			want: []int{2, 3, 4},
		},
	}
	for _, test := range tests {
		got := Map(test.data, test.f)
		if !cmp.Equal(got, test.want) {
			t.Errorf("Map(%v, ...) = %v, want %v", test.data, got, test.want)
		}
	}
}

func TestMapE(t *testing.T) {
	oddError := errors.New("that's odd")
	f := func(i int) (int, error) {
		if i%2 == 1 {
			return 0, oddError
		}
		return i + 1, nil
	}
	tests := []struct {
		data    []int
		want    []int
		wantErr error
	}{
		{
			data: []int{2, 4, 6, 8},
			want: []int{3, 5, 7, 9},
		},
		{
			data:    []int{2, 4, 5, 6, 8},
			wantErr: oddError,
		},
	}
	for _, test := range tests {
		got, err := MapE(test.data, f)
		if err != nil && (test.wantErr == nil || !errors.Is(err, oddError)) {
			t.Errorf("MapE(%v, f): %v, wanted error %v", test.data, err, oddError)
		} else if err == nil && !cmp.Equal(got, test.want) {
			t.Errorf("MapE(%v, f) = %v, want %v", test.data, got, test.want)
		}
	}
}

func TestFilter(t *testing.T) {
	tests := []struct {
		data []int
		op   string
		p    func(int) bool
		want []int
	}{
		{
			data: []int{1, 2, 3, 4, 5},
			op:   "even?",
			p:    func(i int) bool { return i%2 == 0 },
			want: []int{2, 4},
		},
		{
			data: []int{1, 2, 3, 4, 5},
			op:   "odd?",
			p:    func(i int) bool { return i%2 == 1 },
			want: []int{1, 3, 5},
		},
		{
			data: []int{1, 2, 3, 4, 5},
			op:   ">10?",
			p:    func(i int) bool { return i > 10 },
			want: nil,
		},
		{
			data: []int{1, 2, 3, 4, 5},
			op:   "<10?",
			p:    func(i int) bool { return i < 10 },
			want: []int{1, 2, 3, 4, 5},
		},
		{
			data: []int{},
			op:   "any",
			p:    func(int) bool { return true },
			want: nil,
		},
	}
	for _, test := range tests {
		if got := Filter(test.data, test.p); !cmp.Equal(got, test.want) {
			t.Errorf("Filter(%v, %s) = %v, want %v", test.data, test.op, got, test.want)
		}
	}
}

func TestForEach(t *testing.T) {
	data := []int{1, 2, 3, 4}
	want := []int{2, 3, 4, 5}
	var got []int
	ForEach(data, func(i int) { got = append(got, i+1) })
	if !cmp.Equal(got, want) {
		t.Errorf("ForEach(%v, f) -> %v, want %v", data, got, want)
	}
}
