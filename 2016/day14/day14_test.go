package day14

import (
	"crypto/md5"
	"testing"
)

func TestFindKey(t *testing.T) {
	makeSimple := func() keystreamer { return newSimpleStreamer("abc") }
	makeStretched := func() keystreamer { return newStretchedStreamer("abc", 2016) }
	tests := []struct {
		kind     string
		streamer func() keystreamer
		keyIndex int
		want     int
	}{
		{"simple", makeSimple, 1, 39},
		{"simple", makeSimple, 2, 92},
		{"simple", makeSimple, 64, 22728},
		{"stretched", makeStretched, 1, 10},
		// Passes, but slowly:
		// {"stretched", makeStretched, 64, 22551},
	}
	for _, test := range tests {
		ks := test.streamer()
		if got := findKey(ks, test.keyIndex); got != test.want {
			t.Errorf("findKey(%s, %d) = %d, want %d", test.kind, test.keyIndex, got, test.want)
		}
	}
}

func TestFindTriple(t *testing.T) {
	tests := []struct {
		input string
		want  byte
	}{
		{"abc0", noTriple},
		{"abc18", 8},
		{"abc39", 0xe},
		{"abc92", 9},
	}
	for _, test := range tests {
		h := hash(md5.Sum([]byte(test.input)))
		if got := h.findTriple(); got != test.want {
			t.Errorf("hash(%s).findTriple() = %x, want %x", test.input, got, test.want)
		}
	}
}

func TestMatchPentad(t *testing.T) {
	tests := []struct {
		input string
		key   byte
		want  bool
	}{
		{"abc18", 8, false},
		{"abc816", 0xe, true},
		{"abc200", 9, true},
	}
	for _, test := range tests {
		h := hash(md5.Sum([]byte(test.input)))
		if got := h.matchPentad(test.key); got != test.want {
			t.Errorf("hash(%s).matchPentad(%x) = %t, want %t", test.input, test.key, got, test.want)
		}
	}
}
