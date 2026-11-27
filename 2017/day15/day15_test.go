package day15

import (
	"testing"
)

const (
	xA = 65
	xB = 8921
)

func TestJudge(t *testing.T) {
	tests := []struct {
		N    int
		want int
	}{
		{N: 5, want: 1},
		{N: 40000000, want: 588},
	}
	for _, test := range tests {
		if got := judge(xA, xB, test.N); got != test.want {
			t.Errorf("judge(..., %d) = %d, want %d", test.N, got, test.want)
		}
	}
}

func TestJudge2(t *testing.T) {
	judges := []struct {
		name string
		f    func(xA, xB, N int) int
	}{
		{name: "judge2", f: judge2},
		{name: "judge2p", f: judge2p},
	}
	tests := []struct {
		N    int
		want int
	}{
		{N: 5, want: 0},
		{N: 1055, want: 0},
		{N: 1056, want: 1},
		{N: 5000000, want: 309},
	}
	for _, test := range tests {
		for _, j := range judges {
			if got := j.f(xA, xB, test.N); got != test.want {
				t.Errorf("%s(..., %d) = %d, want %d", j.name, test.N, got, test.want)
			}
		}
	}
}

func BenchmarkRelative(b *testing.B) {
	b.Run("judge", func(b *testing.B) {
		for range b.N {
			judge(xA, xB, 40000000)
		}
	})
	b.Run("judge2", func(b *testing.B) {
		for range b.N {
			judge2(xA, xB, 5000000)
		}
	})
	b.Run("judge2p", func(b *testing.B) {
		for range b.N {
			judge2p(xA, xB, 5000000)
		}
	})
}

func BenchmarkParallelism(b *testing.B) {
	b.Run("judge2", func(b *testing.B) {
		judge2(xA, xB, b.N)
	})
	b.Run("judge2p", func(b *testing.B) {
		judge2p(xA, xB, b.N)
	})
}
