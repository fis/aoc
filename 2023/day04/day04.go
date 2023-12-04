// Copyright 2023 Google LLC
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

// Package day04 solves AoC 2023 day 4.
package day04

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util/fn"
)

func init() {
	glue.RegisterSolver(2023, 4, glue.LineSolver(glue.WithParser(parseCardFast, solve)))
}

func solve(cards []card) ([]string, error) {
	p1 := countPoints(cards)
	p2 := countCards(cards)
	return glue.Ints(p1, p2), nil
}

func countPoints(cards []card) (sum int) {
	for _, c := range cards {
		matches := c.countMatches()
		if matches > 0 {
			sum += 1 << (matches - 1)
		}
	}
	return sum
}

func countCards(cards []card) (sum int) {
	counts := make([]int, len(cards))
	for i, c := range cards {
		counts[i]++
		matches := c.countMatches()
		for j := 1; j <= matches; j++ {
			counts[i+j] += counts[i]
		}
	}
	return fn.Sum(counts)
}

const maxNumber = 99

type card struct {
	winners []byte
	numbers []byte
}

func (c card) countMatches() (matches int) {
	var bitmap [maxNumber + 1]bool
	for _, w := range c.winners {
		bitmap[w] = true
	}
	for _, n := range c.numbers {
		if bitmap[n] {
			matches++
		}
	}
	return matches
}

func parseCardSimple(line string) (card, error) {
	_, data, ok := strings.Cut(line, ":")
	if !ok {
		return card{}, fmt.Errorf("missing : in %q", line)
	}
	winData, numData, ok := strings.Cut(data, "|")
	if !ok {
		return card{}, fmt.Errorf("missing | in %q", line)
	}
	winners, err := parseNumbers(winData)
	if err != nil {
		return card{}, fmt.Errorf("bad winners list: %w", err)
	}
	numbers, err := parseNumbers(numData)
	if err != nil {
		return card{}, fmt.Errorf("bad numbers list: %w", err)
	}
	return card{winners: winners, numbers: numbers}, nil
}

func parseNumbers(text string) (numbers []byte, err error) {
	for _, s := range strings.Split(text, " ") {
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		if n < 0 || n > maxNumber {
			return nil, fmt.Errorf("number out of range: %d", n)
		}
		numbers = append(numbers, byte(n))
	}
	return numbers, nil
}

func parseCardFast(line string) (card, error) {
	s := line
	for len(s) > 0 && s[0] != ':' {
		s = s[1:]
	}
	if len(s) == 0 {
		return card{}, fmt.Errorf("missing : in %q", line)
	}
	s = s[1:]
	winners := make([]byte, 0, len(s)/3)
	for {
		if len(s) > 0 && s[0] == ' ' {
			s = s[1:]
			continue
		}
		if len(s) == 0 {
			return card{}, fmt.Errorf("missing | in %q", line)
		}
		if s[0] == '|' {
			s = s[1:]
			break
		}
		if s[0] < '0' || s[0] > '9' {
			return card{}, fmt.Errorf("unexpected byte %02x in %q", s[0], line)
		}
		if len(s) > 1 && s[1] >= '0' && s[1] <= '9' {
			winners = append(winners, 10*(s[0]-'0')+(s[1]-'0'))
			s = s[2:]
		} else {
			winners = append(winners, s[0]-'0')
			s = s[1:]
		}
	}
	numbers := winners[len(winners):]
	for len(s) > 0 {
		if s[0] == ' ' {
			s = s[1:]
			continue
		}
		if s[0] < '0' || s[0] > '9' {
			return card{}, fmt.Errorf("unexpected byte %02x in %q", s[0], line)
		}
		if len(s) > 1 && s[1] >= '0' && s[1] <= '9' {
			numbers = append(numbers, 10*(s[0]-'0')+(s[1]-'0'))
			s = s[2:]
		} else {
			numbers = append(numbers, s[0]-'0')
			s = s[1:]
		}
	}
	return card{winners: winners, numbers: numbers}, nil
}
