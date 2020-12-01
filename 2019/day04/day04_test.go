package day04

import (
	"testing"
)

func TestValidate1(t *testing.T) {
	tests := []struct {
		digits string
		want   bool
	}{
		{"111111", true},
		{"223450", false},
		{"123789", false},
	}
	for _, test := range tests {
		got := validate1([]byte(test.digits))
		if got != test.want {
			t.Errorf("validate1(%s) = %v, want %v", test.digits, got, test.want)
		}
	}
}

func TestValidate2(t *testing.T) {
	tests := []struct {
		digits string
		want   bool
	}{
		{"112233", true},
		{"123444", false},
		{"111122", true},
	}
	for _, test := range tests {
		got := validate2([]byte(test.digits))
		if got != test.want {
			t.Errorf("validate2(%s) = %v, want %v", test.digits, got, test.want)
		}
	}
}
