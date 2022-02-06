// Copyright 2022 Google LLC
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

package day07

import (
	"testing"
)

func TestHasTLS(t *testing.T) {
	tests := []struct {
		ip   string
		want bool
	}{
		{ip: "abba[mnop]qrst", want: true},
		{ip: "abcd[bddb]xyyx", want: false},
		{ip: "aaaa[qwer]tyui", want: false},
		{ip: "ioxxoj[asdfgh]zxcvbn", want: true},
	}
	for _, test := range tests {
		if got := hasTLS(test.ip); got != test.want {
			t.Errorf("hasTLS(%s) = %t, want %t", test.ip, got, test.want)
		}
	}
}

func TestHasSSL(t *testing.T) {
	tests := []struct {
		ip   string
		want bool
	}{
		{ip: "aba[bab]xyz", want: true},
		{ip: "xyx[xyx]xyx", want: false},
		{ip: "aaa[kek]eke", want: true},
		{ip: "zazbz[bzb]cdb", want: true},
	}
	for _, test := range tests {
		if got := hasSSL(test.ip); got != test.want {
			t.Errorf("hasSSL(%s) = %t, want %t", test.ip, got, test.want)
		}
	}
}
