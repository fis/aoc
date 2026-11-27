package day04

import (
	"testing"

	"github.com/fis/aoc/util"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		passphrase string
		want       bool
	}{
		{passphrase: "aa bb cc dd ee", want: true},
		{passphrase: "aa bb cc dd aa", want: false},
		{passphrase: "aa bb cc dd aaa", want: true},
	}
	for _, test := range tests {
		if got := part1(util.Words(test.passphrase)); got != test.want {
			t.Errorf("part1(%s) = %t, want %t", test.passphrase, got, test.want)
		}
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		passphrase string
		want       bool
	}{
		{passphrase: "abcde fhgij", want: true},
		{passphrase: "abcde xyz ecdab", want: false},
		{passphrase: "a ab abc abd abf abj", want: true},
		{passphrase: "iiii oiii ooii oooi oooo", want: true},
		{passphrase: "oiii ioii iioi iiio", want: false},
	}
	for _, test := range tests {
		if got := part2(util.Words(test.passphrase)); got != test.want {
			t.Errorf("part2(%s) = %t, want %t", test.passphrase, got, test.want)
		}
	}
}
