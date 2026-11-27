package knot

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{input: ``, want: "a2582a3a0e66e6e86e3812dcb672a272"},
		{input: `AoC 2017`, want: "33efeb34ea91902bb2f59c9920caa6cd"},
		{input: `1,2,3`, want: "3efbe78a8d82f29979031a4aa0b16a9d"},
		{input: `1,2,4`, want: "63960835bcdc130f0b66d7ff4f6a5a8e"},
	}
	for _, test := range tests {
		if got := fmt.Sprintf("%x", Hash(N, Rounds, test.input)); got != test.want {
			t.Errorf("Hash(%d, %d, %q) = %q, want %q", N, Rounds, test.input, got, test.want)
		}
	}
}
