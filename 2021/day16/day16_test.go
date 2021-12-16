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

package day16

import "testing"

func TestParsePacket(t *testing.T) {
	tests := []struct {
		data string
		want int
	}{
		{data: "D2FE28", want: 6},
		{data: "38006F45291200", want: 1 + 6 + 2},
		{data: "EE00D40C823060", want: 7 + 2 + 4 + 1},
		{data: "8A004A801A8002F478", want: 16},
		{data: "620080001611562C8802118E34", want: 12},
		{data: "C0015000016115A2E0802F182340", want: 23},
		{data: "A0016C880162017C3686B18A3D4780", want: 31},
	}
	for _, test := range tests {
		p, _ := parsePacket(bitReaderHex(test.data))
		if got := p.versionSum(); got != test.want {
			t.Errorf("%s -> %d, want %d", test.data, got, test.want)
		}
	}
}

func TestEval(t *testing.T) {
	tests := []struct {
		data string
		want int
	}{
		{data: "C200B40A82", want: 3},
		{data: "04005AC33890", want: 54},
		{data: "880086C3E88112", want: 7},
		{data: "CE00C43D881120", want: 9},
		{data: "D8005AC2A8F0", want: 1},
		{data: "F600BC2D8F", want: 0},
		{data: "9C005AC2F8F0", want: 0},
		{data: "9C0141080250320F1802104A08", want: 1},
	}
	for _, test := range tests {
		p, _ := parsePacket(bitReaderHex(test.data))
		if got := p.eval(); got != test.want {
			t.Errorf("%s -> %d, want %d", test.data, got, test.want)
		}
	}
}
