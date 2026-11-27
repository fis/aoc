package day02

import (
	"testing"
)

func TestChecksum(t *testing.T) {
	ex := []string{
		"abcdef",
		"bababc",
		"abbcde",
		"abcccd",
		"aabcdd",
		"abcdee",
		"ababab",
	}
	want := 12
	got := checksum(ex)
	if got != want {
		t.Errorf("checksum = %d, want %d", got, want)
	}
}

func TestDiff1(t *testing.T) {
	tests := []struct {
		left, right string
		wantOK      bool
		wantS       string
	}{
		{left: "abcde", right: "axcye", wantOK: false, wantS: ""},
		{left: "fghij", right: "fguij", wantOK: true, wantS: "fgij"},
	}
	for _, test := range tests {
		if ok, s := diff1(test.left, test.right); ok != test.wantOK || s != test.wantS {
			t.Errorf("diff1(%q, %q) = (%v, %q), want (%v, %q)", test.left, test.right, ok, s, test.wantOK, test.wantS)
		}
	}
}

func TestFindBox(t *testing.T) {
	ex := []string{
		"abcde",
		"fghij",
		"klmno",
		"pqrst",
		"fguij",
		"axcye",
		"wvxyz",
	}
	want := "fgij"
	got := findBox(ex)
	if got != want {
		t.Errorf("findBox = %q, want %q", got, want)
	}
}
