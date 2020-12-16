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

package day16

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var (
	ex1Rules = []string{
		"class: 1-3 or 5-7",
		"row: 6-11 or 33-44",
		"seat: 13-40 or 45-50",
	}
	ex1Nearby = []string{
		"7,3,47",
		"40,4,50",
		"55,2,20",
		"38,6,12",
	}
)

func TestFilterValid(t *testing.T) {
	wantTickets, wantRate := [][]int{{7, 3, 47}}, 71
	if rules, err := parseRules(ex1Rules); err != nil {
		t.Errorf("parseRules: %v", err)
	} else if tickets, err := parseTickets(ex1Nearby); err != nil {
		t.Errorf("parseTickets: %v", err)
	} else if gotTickets, gotRate := filterValid(tickets, rules); !cmp.Equal(gotTickets, wantTickets) || gotRate != wantRate {
		t.Errorf("filterValid = (%v, %d), want (%v, %d)", gotTickets, gotRate, wantTickets, wantRate)
	}
}

var (
	ex2Rules = []string{
		"class: 0-1 or 4-19",
		"row: 0-5 or 8-19",
		"seat: 0-13 or 16-19",
	}
	ex2Nearby = []string{
		"3,9,18",
		"15,1,5",
		"5,14,9",
	}
)

func TestPossibleFields(t *testing.T) {
	want := []uint{
		0b010, // col 1: can't be class due to ticket 1 or seat due to ticket 2
		0b011, // col 2: can't be seat due to ticket 3
		0b111, // col 3: could be anything
	}
	if rules, err := parseRules(ex2Rules); err != nil {
		t.Errorf("parseRules: %v", err)
	} else if tickets, err := parseTickets(ex2Nearby); err != nil {
		t.Errorf("parseTickets: %v", err)
	} else if got := possibleFields(tickets, rules); !cmp.Equal(got, want) {
		t.Errorf("possibleFields = %03b, want %03b", got, want)
	}
}

func TestFieldNames(t *testing.T) {
	want := []string{"row", "class", "seat"}
	if rules, err := parseRules(ex2Rules); err != nil {
		t.Errorf("parseRules: %v", err)
	} else if tickets, err := parseTickets(ex2Nearby); err != nil {
		t.Errorf("parseTickets: %v", err)
	} else if got, err := fieldNames(tickets, rules); err != nil {
		t.Errorf("fieldNames: %v", err)
	} else if !cmp.Equal(got, want) {
		t.Errorf("fieldNames = %v, want %v", got, want)
	}
}
