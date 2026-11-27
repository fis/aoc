package day02

import (
	"testing"
)

func TestTotalScoreShapes(t *testing.T) {
	guide := [][2]shape{
		{rock, paper},
		{paper, rock},
		{scissors, scissors},
	}
	want := 15
	if got := totalScoreShapes(guide); got != want {
		t.Errorf("totalScoreShapes(%v) = %d, want %d", guide, got, want)
	}
}

func TestTotalScoreRounds(t *testing.T) {
	guide := []guideLine{
		{rock, draw},
		{paper, loss},
		{scissors, win},
	}
	want := 12
	if got := totalScoreRounds(guide); got != want {
		t.Errorf("totalScoreRounds(%v) = %d, want %d", guide, got, want)
	}
}
