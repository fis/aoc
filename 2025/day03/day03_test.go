package day03

import "testing"

var ex = []string{
	"987654321111111",
	"811111111111119",
	"234234234234278",
	"818181911112111",
}

func TestMaxJoltage2(t *testing.T) {
	want := []int{98, 89, 78, 92}
	for i, bank := range ex {
		if got := maxJoltage2(bank); got != want[i] {
			t.Errorf("maxJoltage(%s) = %d, want %d", bank, got, want[i])
		}
	}
}

func TestMaxJoltageN(t *testing.T) {
	want := []int{987654321111, 811111111119, 434234234278, 888911112111}
	for i, bank := range ex {
		if got := maxJoltageN(bank, 12); got != want[i] {
			t.Errorf("maxJoltage(%s, 12) = %d, want %d", bank, got, want[i])
		}
	}
}
