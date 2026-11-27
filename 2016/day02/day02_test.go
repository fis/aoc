package day02

import (
	"testing"
)

func TestDecode(t *testing.T) {
	sheet := []string{
		"ULL",
		"RRDDD",
		"LURDL",
		"UUUUD",
	}
	tests := []struct {
		name    string
		decoder func([]string) string
		want    string
	}{
		{name: "decode", decoder: decode, want: "1985"},
		{name: "decodeCross", decoder: decodeCross, want: "5DB3"},
	}
	for _, test := range tests {
		got := test.decoder(sheet)
		if got != test.want {
			t.Errorf("%s(%v) = %s, want %s", test.name, sheet, got, test.want)
		}
	}
}
