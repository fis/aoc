package day08

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRender(t *testing.T) {
	layers := decode([]byte("0222112222120000"), 2, 2)
	img := render(layers)
	want := []byte{0, 1, 1, 0}
	if !cmp.Equal(img, want) {
		t.Errorf("render(%v) = %v, want %v", layers, img, want)
	}
}
