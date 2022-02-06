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

// Package day07 solves AoC 2016 day 7.
package day07

import "github.com/fis/aoc/glue"

func init() {
	glue.RegisterSolver(2016, 7, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	p1, p2 := 0, 0
	for _, line := range lines {
		if hasTLS(line) {
			p1++
		}
		if hasSSL(line) {
			p2++
		}
	}
	return glue.Ints(p1, p2), nil
}

func hasTLS(ip string) (hasABBA bool) {
	hyper := false
	for i := 0; i+3 < len(ip); i++ {
		if ip[i] == '[' || ip[i] == ']' {
			hyper = ip[i] == '['
			continue
		}
		if ip[i] < 'a' || ip[i] > 'z' || ip[i+1] < 'a' || ip[i+1] > 'z' || ip[i+2] < 'a' || ip[i+2] > 'z' || ip[i+3] < 'a' || ip[i+3] > 'z' {
			continue
		}
		if hasABBA && !hyper {
			continue // no need to check here
		}
		if ip[i+1] != ip[i] && ip[i+2] == ip[i+1] && ip[i+3] == ip[i] {
			if hyper {
				return false
			}
			hasABBA = true
		}
	}
	return hasABBA
}

func hasSSL(ip string) bool {
	const (
		aba byte = 1
		bab byte = 2
		ssl byte = 3
	)
	blocks := [26][26]byte{}
	hyper := false
	for i := 0; i+2 < len(ip); i++ {
		if ip[i] == '[' || ip[i] == ']' {
			hyper = ip[i] == '['
			continue
		}
		if ip[i] < 'a' || ip[i] > 'z' || ip[i+1] < 'a' || ip[i+1] > 'z' || ip[i+2] < 'a' || ip[i+2] > 'z' {
			continue
		}
		if ip[i+1] != ip[i] && ip[i+2] == ip[i] {
			if !hyper {
				a, b := ip[i]-'a', ip[i+1]-'a'
				blocks[a][b] |= aba
				if blocks[a][b] == ssl {
					return true
				}
			} else {
				a, b := ip[i+1]-'a', ip[i]-'a'
				blocks[a][b] |= bab
				if blocks[a][b] == ssl {
					return true
				}
			}
		}
	}
	return false
}
