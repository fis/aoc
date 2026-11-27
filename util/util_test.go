package util

import (
	"slices"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestScanInts(t *testing.T) {
	tests := []struct {
		input string
		want  []int
	}{
		{
			input: "123456",
			want:  []int{123456},
		},
		{
			input: "123456\n",
			want:  []int{123456},
		},
		{
			input: "foo 123456 bar\n",
			want:  []int{123456},
		},
		{
			input: "-123456",
			want:  []int{-123456},
		},
		{
			input: "foo -123456 bar\n",
			want:  []int{-123456},
		},
		{
			input: "foo-bar",
			want:  []int(nil),
		},
		{
			input: "foo - bar",
			want:  []int(nil),
		},
		{
			input: "123 - 456",
			want:  []int{123, 456},
		},
		{
			input: "123 -456 789",
			want:  []int{123, -456, 789},
		},
		{
			input: "123,-456,789",
			want:  []int{123, -456, 789},
		},
		{
			input: "123, -456, 789",
			want:  []int{123, -456, 789},
		},
		{
			input: "123\n-456\n789\n",
			want:  []int{123, -456, 789},
		},
		{
			input: "123-456",
			want:  []int{123, -456},
		}}
	for _, test := range tests {
		if got, err := ScanAllInts(strings.NewReader(test.input)); err != nil {
			t.Errorf("ScanAllInts(%q): %v", test.input, err)
		} else if !cmp.Equal(got, test.want) {
			t.Errorf("ScanAllInts(%q) = %#v, want %#v", test.input, got, test.want)
		}
		if got := Ints(test.input); !cmp.Equal(got, test.want) {
			t.Errorf("Ints(%q) = %#v, want %#v", test.input, got, test.want)
		}
	}
}

func TestWithPrefix(t *testing.T) {
	data := []string{
		// N.B. must be sorted
		"bar",
		"foo",
		"foobar",
		"foobaz",
		"fooquux",
		"quux",
		"zuul",
	}
	tests := []struct {
		prefix string
		want   []string
	}{
		{"b", []string{"bar"}},
		{"bar", []string{"bar"}},
		{"f", []string{"foo", "foobar", "foobaz", "fooquux"}},
		{"foo", []string{"foo", "foobar", "foobaz", "fooquux"}},
		{"foob", []string{"foobar", "foobaz"}},
		{"foobar", []string{"foobar"}},
		{"q", []string{"quux"}},
		{"r", nil},
	}
	for _, test := range tests {
		if got := WithPrefix(data, test.prefix); !cmp.Equal(got, test.want) {
			t.Errorf("WithPrefix(..., %s) = %v, want %v", test.prefix, got, test.want)
		}
	}
}

func TestCircleM(t *testing.T) {
	tests := []struct {
		p    P
		dist int
		want []P
	}{
		{p: P{3, 7}, dist: 1, want: []P{{4, 7}, {3, 6}, {2, 7}, {3, 8}}},
		{p: P{5, 11}, dist: 3, want: []P{{8, 11}, {7, 10}, {6, 9}, {5, 8}, {4, 9}, {3, 10}, {2, 11}, {3, 12}, {4, 13}, {5, 14}, {6, 13}, {7, 12}}},
	}
	for _, test := range tests {
		if got := slices.Collect(test.p.CircleM(test.dist)); !cmp.Equal(got, test.want) {
			t.Errorf("%v.CircleM(%d) = %v, want %v", test.p, test.dist, got, test.want)
		}
	}
}
