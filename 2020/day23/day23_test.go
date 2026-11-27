package day23

import (
	"testing"
)

func TestRingMove(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"389125467", "289154673"},
		{"289154673", "546789132"},
		{"546789132", "891346725"},
		{"891346725", "467913258"},
		{"467913258", "136792584"},
		{"136792584", "936725841"},
		{"936725841", "258367419"},
		{"258367419", "674158392"},
		{"674158392", "574183926"},
		{"574183926", "837419265"},
	}
	for _, test := range tests {
		if r1, err := newRing(test.in); err != nil {
			t.Errorf("newRing(%s): %v", test.in, err)
		} else if got := r1.move().String(); got != test.want {
			t.Errorf("%s.move() = %s, want %s", test.in, got, test.want)
		}
	}
}

func TestRingKey(t *testing.T) {
	r, _ := newRing("837419265")
	want := 92658374
	got := r.key()
	if got != want {
		t.Errorf("(837419265).key() = %d, want %d", got, want)
	}
}

func TestBigRing(t *testing.T) {
	r, _ := newBigRing("389125467")
	for range 10000000 {
		r.move()
	}
	want := 149245887792
	got := r.key()
	if got != want {
		t.Errorf("big ring x10M key = %d, want %d", got, want)
	}
}
