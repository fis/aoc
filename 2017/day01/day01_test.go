package day01

import "testing"

func TestPart1(t *testing.T) {
	tests := []struct {
		digits string
		want   int
	}{
		{digits: "1122", want: 3},
		{digits: "1111", want: 4},
		{digits: "1234", want: 0},
		{digits: "91212129", want: 9},
	}
	for _, test := range tests {
		got := checksum([]byte(test.digits), 1)
		if got != test.want {
			t.Errorf("checksum(%s, 1) = %d, want %d", test.digits, got, test.want)
		}
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		digits string
		want   int
	}{
		{digits: "1212", want: 6},
		{digits: "1221", want: 0},
		{digits: "123425", want: 4},
		{digits: "123123", want: 12},
		{digits: "12131415", want: 4},
	}
	for _, test := range tests {
		got := checksum([]byte(test.digits), len(test.digits)/2)
		if got != test.want {
			t.Errorf("checksum(%s, %d) = %d, want %d", test.digits, len(test.digits)/2, got, test.want)
		}
	}
}
