package day05

import (
	"testing"
)

func TestDecodePass(t *testing.T) {
	tests := []struct {
		pass     string
		row, col int
	}{
		{"FBFBBFFRLR", 44, 5},
		{"BFFFBBFRRR", 70, 7},
		{"FFFBBBFRRR", 14, 7},
		{"BBFFBBFRLL", 102, 4},
	}
	for _, test := range tests {
		if row, col, err := decodePass(test.pass); err != nil {
			t.Errorf("decodePass(%s): %v", test.pass, err)
		} else if row != test.row || col != test.col {
			t.Errorf("decodePass(%s) = (%d, %d), want (%d, %d)", test.pass, row, col, test.row, test.col)
		}
	}
}

func TestSeatID(t *testing.T) {
	tests := []struct {
		row, col int
		want     int
	}{
		{44, 5, 357},
		{70, 7, 567},
		{14, 7, 119},
		{102, 4, 820},
	}
	for _, test := range tests {
		if got := seatID(test.row, test.col); got != test.want {
			t.Errorf("seatID(%d, %d) = %d, want %d", test.row, test.col, got, test.want)
		}
	}
}
