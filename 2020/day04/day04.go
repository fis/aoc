// Copyright 2020 Google LLC
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

// Package day04 solves AoC 2020 day 4.
package day04

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/fis/aoc-go/util"
)

func init() {
	util.RegisterSolver(4, util.ChunkSolver(solve))
}

func solve(data []string) ([]int, error) {
	passes, err := parsePassports(data)
	if err != nil {
		return nil, err
	}
	valid := countValid(passes, passport.valid)
	strict := countValid(passes, passport.strictlyValid)
	return []int{valid, strict}, nil
}

func countValid(data []passport, validator func(passport) bool) (valid int) {
	for _, p := range data {
		if validator(p) {
			valid++
		}
	}
	return
}

type passport map[string]string

func (p passport) valid() bool {
	for _, f := range []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} {
		if _, ok := p[f]; !ok {
			return false
		}
	}
	return true
}

func (p passport) strictlyValid() bool {
	for _, rule := range strictRules {
		if !rule.predicate(p[rule.key]) {
			return false
		}
	}
	return true
}

var (
	patHGT = regexp.MustCompile(`^(\d+)(cm|in)$`)
	patHCL = regexp.MustCompile(`^#[0-9a-f]{6}$`)
	patECL = regexp.MustCompile(`^(?:amb|blu|brn|gry|grn|hzl|oth)$`)
	patPID = regexp.MustCompile(`^\d{9}$`)

	strictRules = []struct {
		key       string
		predicate func(string) bool
	}{
		{"byr", validateNum(1920, 2002)},
		{"iyr", validateNum(2010, 2020)},
		{"eyr", validateNum(2020, 2030)},
		{"hgt", validateHGT},
		{"hcl", patHCL.MatchString},
		{"ecl", patECL.MatchString},
		{"pid", patPID.MatchString},
	}
)

func validateHGT(s string) bool {
	parts := patHGT.FindStringSubmatch(s)
	if len(parts) != 3 {
		return false
	}
	limits, ok := map[string]struct{ min, max int }{
		"cm": {150, 193},
		"in": {59, 76},
	}[parts[2]]
	return ok && numBetween(parts[1], limits.min, limits.max)
}

func validateNum(min, max int) func(string) bool {
	return func(s string) bool {
		return numBetween(s, min, max)
	}
}

func numBetween(s string, min, max int) bool {
	n, err := strconv.Atoi(s)
	return err == nil && n >= min && n <= max
}

func parsePassports(data []string) (out []passport, err error) {
	for _, text := range data {
		p, err := parsePassport(text)
		if err != nil {
			return nil, err
		}
		out = append(out, p)
	}
	return out, nil
}

func parsePassport(data string) (passport, error) {
	p := make(passport)
	for _, word := range util.Words(data) {
		kv := strings.SplitN(word, ":", 2)
		if len(kv) < 2 {
			return nil, fmt.Errorf("invalid datum: %q", word)
		}
		p[kv[0]] = kv[1]
	}
	return p, nil
}
