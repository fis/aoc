package day06

import (
	"testing"
)

var ex = []string{
	"eedadn",
	"drvtee",
	"eandsr",
	"raavrd",
	"atevrs",
	"tsrnev",
	"sdttsa",
	"rasrtv",
	"nssdts",
	"ntnada",
	"svetve",
	"tesnvt",
	"vntsnd",
	"vrdear",
	"dvrsen",
	"enarar",
}

func TestDecode(t *testing.T) {
	want1, want2 := "easter", "advent"
	if got1, got2 := decode(ex); got1 != want1 || got2 != want2 {
		t.Errorf("decode(ex) = (%s, %s), want (%s, %s)", got1, got2, want1, want2)
	}
}
