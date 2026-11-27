package day01

import (
	"testing"
)

func TestMaxCalories(t *testing.T) {
	data := [][]int{
		{1000, 2000, 3000},
		{4000},
		{5000, 6000},
		{7000, 8000, 9000},
		{10000},
	}
	want1, want2 := 24000, 45000
	if got1, got2 := maxCalories(data); got1 != want1 || got2 != want2 {
		t.Errorf("maxCalories(%v) = (%d, %d), want (%d, %d)", data, got1, got2, want1, want2)
	}
}
