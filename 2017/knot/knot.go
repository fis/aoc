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

// Package knot implements the knot hash function (AoC 2017 days 10, 14).
package knot

// N is the standard list length for knot hashes.
const N = 256

// Rounds is the standard number of rounds for a knot hash.
const Rounds = 64

// Size is the number of bytes in a knot hash with the standard parameters.
const Size = 16

// List generates the initial knot hash state list, of the given size.
func List(N int) []byte {
	list := make([]byte, N)
	for i := 0; i < N; i++ {
		list[i] = byte(i)
	}
	return list
}

// Hash applies the knot hash function to the input string.
func Hash(N, rounds int, input string) []byte {
	list := List(N)
	lengths := append([]byte(input), []byte{17, 31, 73, 47, 23}...)
	pos, skip := 0, 0
	for i := 0; i < rounds; i++ {
		pos, skip = Round(pos, skip, list, lengths)
	}
	return compact(list, 16)
}

// Round applies a single knot hash Round to the given list.
func Round(pos, skip int, list []byte, lengths []byte) (newPos, newSkip int) {
	N := len(list)
	for _, length := range lengths {
		reverse(list, pos, int(length))
		pos = (pos + int(length) + skip) % N
		skip++
	}
	return pos, skip
}

func compact(list []byte, factor int) (hash []byte) {
	N := len(list) / factor
	hash = make([]byte, N)
	for i := 0; i < N; i++ {
		v := byte(0)
		for j := 0; j < factor; j++ {
			v ^= list[i*factor+j]
		}
		hash[i] = v
	}
	return hash
}

func reverse(list []byte, pos, length int) {
	N := len(list)
	for i := 0; i < length/2; i++ {
		pa, pb := (pos+i)%N, (pos+length-1-i)%N
		a, b := list[pa], list[pb]
		list[pa], list[pb] = b, a
	}
}
