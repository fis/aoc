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

// Package day16 solves AoC 2020 day 16.
package day16

import (
	"fmt"
	"math/bits"
	"strconv"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2020, 16, glue.ChunkSolver(solve))
}

func solve(chunks []string) ([]string, error) {
	if len(chunks) != 3 {
		return nil, fmt.Errorf("expected 3 chunks, got %d", len(chunks))
	}
	yourTicketText := util.Lines(chunks[1])
	if len(yourTicketText) != 2 {
		return nil, fmt.Errorf("expected 2 lines for your ticket, got %d", len(yourTicketText))
	}
	if yourTicketText[0] != "your ticket:" {
		return nil, fmt.Errorf("invalid header for your ticket: %s", yourTicketText[0])
	}
	nearbyTicketText := util.Lines(chunks[2])
	if len(nearbyTicketText) < 2 {
		return nil, fmt.Errorf("expected 2+ lines for nearby tickets, got %d", len(nearbyTicketText))
	}
	if nearbyTicketText[0] != "nearby tickets:" {
		return nil, fmt.Errorf("invalid header for nearby tickets: %s", nearbyTicketText[0])
	}

	rules, err := parseRules(util.Lines(chunks[0]))
	if err != nil {
		return nil, err
	}
	yourTicket, err := parseTicket(yourTicketText[1])
	if err != nil {
		return nil, err
	}
	nearbyTickets, err := parseTickets(nearbyTicketText[1:])
	if err != nil {
		return nil, err
	}

	valid, errorRate := filterValid(nearbyTickets, rules)

	names, err := fieldNames(valid, rules)
	if err != nil {
		return nil, err
	}
	deps := 1
	for i, f := range names {
		if strings.HasPrefix(f, "departure") {
			deps *= yourTicket[i]
		}
	}

	return glue.Ints(errorRate, deps), nil
}

type interval struct {
	min, max int
}
type validationRule struct {
	field     string
	low, high interval
}

func parseRules(lines []string) (rules []validationRule, err error) {
	for _, line := range lines {
		parts := strings.SplitN(line, ": ", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid line: %s", line)
		}
		r := validationRule{field: parts[0]}
		if _, err := fmt.Sscanf(parts[1], "%d-%d or %d-%d", &r.low.min, &r.low.max, &r.high.min, &r.high.max); err != nil {
			return nil, fmt.Errorf("invalid interval spec: %s: %w", parts[1], err)
		}
		rules = append(rules, r)
	}
	return rules, nil
}

func parseTicket(line string) (nums []int, err error) {
	parts := strings.Split(line, ",")
	nums = make([]int, len(parts))
	for i, s := range parts {
		nums[i], err = strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
	}
	return nums, nil
}

func parseTickets(lines []string) (tickets [][]int, err error) {
	for _, line := range lines {
		t, err := parseTicket(line)
		if err != nil {
			return nil, err
		}
		tickets = append(tickets, t)
	}
	return tickets, nil
}

func filterValid(tickets [][]int, rules []validationRule) (validTickets [][]int, errorRate int) {
	for _, ticket := range tickets {
		valid := true
	next:
		for _, v := range ticket {
			for _, r := range rules {
				if (v >= r.low.min && v <= r.low.max) || (v >= r.high.min && v <= r.high.max) {
					continue next
				}
			}
			errorRate += v
			valid = false
		}
		if valid {
			validTickets = append(validTickets, ticket)
		}
	}
	return validTickets, errorRate
}

func fieldNames(tickets [][]int, rules []validationRule) ([]string, error) {
	possible := possibleFields(tickets, rules)
	fields := make([]uint, len(possible))
	free := (uint(1) << len(possible)) - 1
	if !search(fields, len(possible), free, possible) {
		return nil, fmt.Errorf("unsatisfiable")
	}
	names := make([]string, len(fields))
	for i, f := range fields {
		names[i] = rules[bits.TrailingZeros(f)].field
	}
	return names, nil
}

func possibleFields(tickets [][]int, rules []validationRule) (possible []uint) {
	possible = make([]uint, len(rules))
	for col := range possible {
		possible[col] = (uint(1) << len(rules)) - 1
	}
	for _, ticket := range tickets {
		for col, v := range ticket {
			for fi, r := range rules {
				if (v < r.low.min || v > r.low.max) && (v < r.high.min || v > r.high.max) {
					possible[col] &^= uint(1) << fi
				}
			}
		}
	}
	return possible
}

func search(fields []uint, unset int, free uint, possible []uint) (ok bool) {
	if unset == 0 {
		return true
	}
	type choice struct {
		column   int
		possible uint
	}
	choices := make([]choice, 0, len(possible))
	for i, p := range possible {
		if fields[i] == 0 {
			c := choice{column: i, possible: p & free}
			if c.possible == 0 {
				return false
			}
			choices = append(choices, c)
		}
	}
	next, nextSize := 0, bits.OnesCount(choices[0].possible)
	for i, c := range choices {
		s := bits.OnesCount(c.possible)
		if s < nextSize {
			next, nextSize = i, s
		}
	}
	for p := choices[next].possible; p != 0; p = clearLow(p) {
		c := low(p)
		fields[choices[next].column] = c
		if search(fields, unset-1, free^c, possible) {
			return true
		}
	}
	fields[choices[next].column] = 0
	return false
}

func low(u uint) uint {
	t := u ^ (u - 1)
	return t ^ (t >> 1)
}

func clearLow(u uint) uint {
	return u & (u - 1)
}
