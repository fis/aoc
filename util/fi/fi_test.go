package fi

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
		got := Sum(S(test.data))
		if got != test.want {
			t.Errorf("Sum(%v) = %d, want %d", test.data, got, test.want)
		}
	}
}

func TestSumF(t *testing.T) {
	f := func(x int) int { return x + 1 }
	tests := []struct {
		data []int
		want int
	}{
		{data: nil, want: 0},
		{data: []int{123}, want: 124},
		{data: []int{123, 456}, want: 581},
		{data: []int{1, 2, 3, 4, 5, 6, 7}, want: 35},
	}
	for _, test := range tests {
		got := SumF(S(test.data), f)
		if got != test.want {
			t.Errorf("SumF(%v, f) = %d, want %d", test.data, got, test.want)
		}
	}
}

func TestProd(t *testing.T) {
	tests := []struct {
		data []int
		want int
	}{
		{data: nil, want: 1},
		{data: []int{123}, want: 123},
		{data: []int{123, 456}, want: 56088},
		{data: []int{1, 2, 3, 4, 5, 6, 7}, want: 5040},
		{data: []int{1, 2, 3, 0, 5, 6, 7}, want: 0},
	}
	for _, test := range tests {
		got := Prod(S(test.data))
		if got != test.want {
			t.Errorf("Prod(%v) = %d, want %d", test.data, got, test.want)
		}
	}
}

func TestProdF(t *testing.T) {
	f := func(x int) int { return x + 1 }
	tests := []struct {
		data []int
		want int
	}{
		{data: nil, want: 1},
		{data: []int{123}, want: 124},
		{data: []int{123, 456}, want: 56668},
		{data: []int{1, 2, 3, 4, 5, 6, 7}, want: 40320},
		{data: []int{1, 2, 3, -1, 5, 6, 7}, want: 0},
	}
	for _, test := range tests {
		got := ProdF(S(test.data), f)
		if got != test.want {
			t.Errorf("ProdF(%v, f) = %d, want %d", test.data, got, test.want)
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
		got := CountIf(S(test.data), test.f)
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
		got := Min(S(test.data))
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
		got := Max(S(test.data))
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
		if got := Head(S(test.data), test.def); got != test.want {
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
		got := CS(Map(S(test.data), test.f))
		if !cmp.Equal(got, test.want) {
			t.Errorf("Map(%v, ...) = %v, want %v", test.data, got, test.want)
		}
	}
}

func TestMapRange(t *testing.T) {
	f := func(i int) int { return 2 * i }
	tests := []struct {
		start, end int
		want       []int
	}{
		{start: 5, end: 10, want: []int{10, 12, 14, 16, 18}},
		{start: 5, end: 6, want: []int{10}},
		{start: 5, end: 5, want: nil},
	}
	for _, test := range tests {
		if got := CS(MapRange(test.start, test.end, f)); !cmp.Equal(got, test.want) {
			t.Errorf("MapRange(%d, %d, f) = %v, want %v", test.start, test.end, got, test.want)
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
		if got := CS(Filter(S(test.data), test.p)); !cmp.Equal(got, test.want) {
			t.Errorf("Filter(%v, %s) = %v, want %v", test.data, test.op, got, test.want)
		}
	}
}

func TestAll(t *testing.T) {
	f := func(i int) bool { return i%2 == 0 }
	tests := []struct {
		data []int
		want bool
	}{
		{data: []int{}, want: true},
		{data: []int{2}, want: true},
		{data: []int{2, 8, 12}, want: true},
		{data: []int{2, 8, 9, 12}, want: false},
		{data: []int{3, 9, 13}, want: false},
		{data: []int{3}, want: false},
	}
	for _, test := range tests {
		if got := All(S(test.data), f); got != test.want {
			t.Errorf("All(%v, f) = %t, want %t", test.data, got, test.want)
		}
	}
}

func TestAny(t *testing.T) {
	f := func(i int) bool { return i%2 == 0 }
	tests := []struct {
		data []int
		want bool
	}{
		{data: []int{}, want: false},
		{data: []int{2}, want: true},
		{data: []int{2, 8, 12}, want: true},
		{data: []int{2, 8, 9, 12}, want: true},
		{data: []int{3, 9, 13}, want: false},
		{data: []int{3}, want: false},
	}
	for _, test := range tests {
		if got := Any(S(test.data), f); got != test.want {
			t.Errorf("Any(%v, f) = %t, want %t", test.data, got, test.want)
		}
	}
}

func TestChain(t *testing.T) {
	s1 := Seq(1, 2, 3)
	s2 := Seq[int]()
	s3 := Seq(4, 5, 6)
	want := []int{1, 2, 3, 4, 5, 6}
	got := CS(Chain(s1, s2, s3))
	if !cmp.Equal(got, want) {
		t.Errorf("Chain({1, 2, 3}, {}, {4, 5, 6}) = %v, want %v", got, want)
	}
}
