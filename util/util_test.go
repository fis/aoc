// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
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
	tests := []struct{
		prefix string
		want []string
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
