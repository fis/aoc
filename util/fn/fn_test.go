package fn

import (
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
