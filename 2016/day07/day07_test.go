package day07

import (
	"testing"
)

func TestHasTLS(t *testing.T) {
	tests := []struct {
		ip   string
		want bool
	}{
		{ip: "abba[mnop]qrst", want: true},
		{ip: "abcd[bddb]xyyx", want: false},
		{ip: "aaaa[qwer]tyui", want: false},
		{ip: "ioxxoj[asdfgh]zxcvbn", want: true},
	}
	for _, test := range tests {
		if got := hasTLS(test.ip); got != test.want {
			t.Errorf("hasTLS(%s) = %t, want %t", test.ip, got, test.want)
		}
	}
}

func TestHasSSL(t *testing.T) {
	tests := []struct {
		ip   string
		want bool
	}{
		{ip: "aba[bab]xyz", want: true},
		{ip: "xyx[xyx]xyx", want: false},
		{ip: "aaa[kek]eke", want: true},
		{ip: "zazbz[bzb]cdb", want: true},
	}
	for _, test := range tests {
		if got := hasSSL(test.ip); got != test.want {
			t.Errorf("hasSSL(%s) = %t, want %t", test.ip, got, test.want)
		}
	}
}
