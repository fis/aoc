package day01

import (
	"testing"
)

func TestCalibrationValue(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"1abc2", 12},
		{"pqr3stu8vwx", 38},
		{"a1b2c3d4e5f", 15},
		{"treb7uchet", 77},
	}
	for _, test := range tests {
		if got := calibrationValue(test.input); got != test.want {
			t.Errorf("calibrationValue(%s) = %d, want %d", test.input, got, test.want)
		}
	}
}

func TestCalibrationValueEx(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"two1nine", 29},
		{"eightwothree", 83},
		{"abcone2threexyz", 13},
		{"xtwone3four", 24},
		{"4nineeightseven2", 42},
		{"zoneight234", 14},
		{"7pqrstsixteen", 76},
	}
	for _, test := range tests {
		if got := calibrationValueEx(test.input); got != test.want {
			t.Errorf("calibrationValueEx(%s) = %d, want %d", test.input, got, test.want)
		}
	}
}
