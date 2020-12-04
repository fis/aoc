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

package day04

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const (
	part1 = `
ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in
`
	part2valid = `
pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
hcl:#623a2f

eyr:2029 ecl:blu cid:129 byr:1989
iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

hcl:#888785
hgt:164cm byr:2001 iyr:2015 cid:88
pid:545766238 ecl:hzl
eyr:2022

iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719
`
	part2invalid = `
eyr:1972 cid:100
hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926

iyr:2019
hcl:#602927 eyr:1967 hgt:170cm
ecl:grn pid:012533040 byr:1946

hcl:dab227 iyr:2012
ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277

hgt:59cm ecl:zzz
eyr:2038 hcl:74454a iyr:2023
pid:3556412378 byr:2007
`
)

func TestReadPassports(t *testing.T) {
	want := []passport{
		{"ecl": "gry", "pid": "860033327", "eyr": "2020", "hcl": "#fffffd", "byr": "1937", "iyr": "2017", "cid": "147", "hgt": "183cm"},
		{"iyr": "2013", "ecl": "amb", "cid": "350", "eyr": "2023", "pid": "028048884", "hcl": "#cfa07d", "byr": "1929"},
		{"hcl": "#ae17e1", "iyr": "2013", "eyr": "2024", "ecl": "brn", "pid": "760753108", "byr": "1931", "hgt": "179cm"},
		{"hcl": "#cfa07d", "eyr": "2025", "pid": "166559648", "iyr": "2011", "ecl": "brn", "hgt": "59in"},
	}
	if got, err := readPassports(strings.NewReader(part1)); err != nil {
		t.Errorf("readPassports: %v", err)
	} else if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("readPassports mismatch (-want +got):\n%s", diff)
	}
}

func TestCountValid(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		validator func(passport) bool
		want      int
	}{
		{name: "part 1", input: part1, validator: passport.valid, want: 2},
		{name: "part 2, valid", input: part2valid, validator: passport.strictlyValid, want: 4},
		{name: "part 2, invalid", input: part2invalid, validator: passport.strictlyValid, want: 0},
	}
	for _, test := range tests {
		data, err := readPassports(strings.NewReader(test.input))
		if err != nil {
			t.Errorf("%s: readPassports: %v", test.name, err)
		} else if got := countValid(data, test.validator); got != test.want {
			t.Errorf("%s: countValid = %d, want %d", test.name, got, test.want)
		}
	}
}

func TestStrictRules(t *testing.T) {
	type testCase struct {
		value string
		want  bool
	}
	tests := map[string][]testCase{
		"byr": []testCase{{"2002", true}, {"2003", false}},
		"hgt": []testCase{{"60in", true}, {"190cm", true}, {"190in", false}, {"190", false}},
		"hcl": []testCase{{"#123abc", true}, {"#123abz", false}, {"123abc", false}},
		"ecl": []testCase{{"brn", true}, {"wat", false}},
		"pid": []testCase{{"000000001", true}, {"0123456789", false}},
	}
	for _, rule := range strictRules {
		for _, test := range tests[rule.key] {
			got := rule.predicate(test.value)
			if got != test.want {
				t.Errorf("valid(%s:%s) = %v, want %v", rule.key, test.value, got, test.want)
			}
		}
		delete(tests, rule.key)
	}
	for key := range tests {
		t.Errorf("unexecuted tests for %s", key)
	}
}
