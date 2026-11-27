package day08

import "testing"

var example = []int{2, 3, 0, 3, 10, 11, 12, 1, 1, 0, 1, 99, 2, 1, 1, 2}

func TestChecksum(t *testing.T) {
	n, _ := parseLicense(example)
	want := 138
	got := n.checksum()
	if got != want {
		t.Errorf("checksum = %d, want %d", got, want)
	}
}

func TestValue(t *testing.T) {
	n, _ := parseLicense(example)
	want := 66
	got := n.value()
	if got != want {
		t.Errorf("value = %d, want %d", got, want)
	}
}
