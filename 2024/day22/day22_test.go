package day22

import (
	"testing"
)

func TestAnalyze(t *testing.T) {
	ex := []int{1, 2, 3, 2024}
	want1, want2 := 37990510, 23
	if got1, got2 := analyze(ex); got1 != want1 || got2 != want2 {
		t.Errorf("analyze(%v) = (%d,%d), want (%d,%d)", ex, got1, got2, want1, want2)
	}
}

func TestStep(t *testing.T) {
	tests := []struct {
		n, t int
		want int
	}{
		{123, 1, 15887950},
		{15887950, 1, 16495136},
		{16495136, 1, 527345},
		{527345, 1, 704524},
		{704524, 1, 1553684},
		{1553684, 1, 12683156},
		{12683156, 1, 11100544},
		{11100544, 1, 12249484},
		{12249484, 1, 7753432},
		{7753432, 1, 5908254},
		{1, 2000, 8685429},
		{10, 2000, 4700978},
		{100, 2000, 15273692},
		{2024, 2000, 8667524},
	}
	for _, test := range tests {
		if got := stepN(test.n, test.t); got != test.want {
			t.Errorf("stepN(%d, %d) = %d, want %d", test.n, test.t, got, test.want)
		}
	}
}
