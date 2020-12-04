// Copyright 2019 Google LLC
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

package day16

import (
	"testing"
)

func TestFFT(t *testing.T) {
	tests := []struct {
		sig    string
		phases int
		want   string
	}{
		{
			sig:    "12345678",
			phases: 4,
			want:   "01029498",
		},
		{
			sig:    "80871224585914546619083218645595",
			phases: 100,
			want:   "24176176",
		},
		{
			sig:    "19617804207202209144916044189917",
			phases: 100,
			want:   "73745418",
		},
		{
			sig:    "69317163492948606335995924319873",
			phases: 100,
			want:   "52432133",
		},
	}
	for _, test := range tests {
		sig := digits(test.sig)
		fft(sig, test.phases)
		got := undigits(sig[:8])
		if got != test.want {
			t.Errorf("fft(%s, %d) = %s, want %s", test.sig, test.phases, got, test.want)
		}
	}
}

func TestRFFT(t *testing.T) {
	tests := []struct {
		sig  string
		want string
	}{
		{
			sig:  "03036732577212944063491565474664",
			want: "84462026",
		},
		{
			sig:  "02935109699940807407585447034323",
			want: "78725270",
		},
		{
			sig:  "03081770884921959731165446850517",
			want: "53553731",
		},
	}
	for _, test := range tests {
		sig := digits(test.sig)
		out := rfft(sig, 100, 10000)
		got := undigits(out)
		if got != test.want {
			t.Errorf("rfft(%s, 100, 10000) = %s, want %s", test.sig, got, test.want)
		}
	}
}
