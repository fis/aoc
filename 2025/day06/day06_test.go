package day06

import "testing"

var ex = []string{
	"123 328  51 64 ",
	" 45 64  387 23 ",
	"  6 98  215 314",
	"*   +   *   +  ",
}

func TestWrong(t *testing.T) {
	want := 4277556
	if got := parseWorksheetWrong(ex).grandTotal(); got != want {
		t.Errorf("wrong(ex).grandtotal() = %d, want %d", got, want)
	}
}

func TestRight(t *testing.T) {
	want := 3263827
	if got := parseWorksheetRight(ex).grandTotal(); got != want {
		t.Errorf("right(ex).grandtotal() = %d, want %d", got, want)
	}
}
