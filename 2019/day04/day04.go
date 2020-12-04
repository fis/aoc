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

// Package day04 solves AoC 2019 day 4.
package day04

import (
	"fmt"
	"strconv"

	"github.com/fis/aoc-go/util"
)

func Solve(path string) ([]string, error) {
	inputRange, err := util.ReadIntRows(path)
	if err != nil {
		return nil, err
	}
	if len(inputRange) != 2 {
		return nil, fmt.Errorf("expected [low, high], got %v", inputRange)
	}

	a1, a2 := 0, 0
	for n := inputRange[0]; n <= inputRange[1]; n++ {
		digits := strconv.Itoa(n)
		if validate1([]byte(digits)) {
			a1++
		}
		if validate2([]byte(digits)) {
			a2++
		}
	}

	return []string{strconv.Itoa(a1), strconv.Itoa(a2)}, nil
}

func validate1(digits []byte) bool {
	double := false
	for i := 0; i+1 < len(digits); i++ {
		a, b := digits[i], digits[i+1]
		if b < a {
			return false
		}
		if b == a {
			double = true
		}
	}
	return double
}

func validate2(digits []byte) bool {
	double := false
	for i := 0; i+1 < len(digits); i++ {
		a, b := digits[i], digits[i+1]
		if b < a {
			return false
		}
		if b == a && (i == 0 || digits[i-1] != a) && (i+2 == len(digits) || digits[i+2] != b) {
			double = true
		}
	}
	return double
}
