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

// Package day07 solves AoC 2023 day 7.
package day07

import (
	"bytes"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2023, 7, glue.LineSolver(glue.WithParser(parseBid, solve)))
}

func solve(bids []bidInfo) ([]string, error) {
	p1 := totalWinnings(bids)
	jokerify(bids)
	p2 := totalWinnings(bids)
	return glue.Ints(p1, p2), nil
}

func totalWinnings(bids []bidInfo) (total int) {
	slices.SortFunc(bids, func(a, b bidInfo) int { return compareHands(a.hand, b.hand) })
	for i, bid := range bids {
		total += (i + 1) * bid.value
	}
	return total
}

func jokerify(bids []bidInfo) {
	for i := range bids {
		bids[i].hand.jokerify()
	}
}

type bidInfo struct {
	hand  handInfo
	value int
}

func (bid bidInfo) String() string {
	return fmt.Sprintf("%v:%d", bid.hand, bid.value)
}

type handInfo struct {
	cards [handSize]byte
	typ   handType
}

func (h handInfo) String() string {
	var s [handSize]byte
	for i, c := range h.cards {
		s[i] = rankLabels[c]
	}
	return fmt.Sprintf("%s/%s", s, handTypeNames[h.typ])
}

func compareHands(a, b handInfo) int {
	if a.typ != b.typ {
		return a.typ.strength() - b.typ.strength()
	}
	return bytes.Compare(a.cards[:], b.cards[:])
}

const (
	handSize = 5
	numRanks = 13
)

type handType int

const (
	highCardHand handType = iota
	onePairHand
	twoPairHand
	threeKindHand
	fullHouseHand
	fourKindHand
	fiveKindHand
)

func (ht handType) strength() int { return int(ht) }

var handTypeNames = [...]string{
	highCardHand:  "high-card",
	onePairHand:   "one-pair",
	twoPairHand:   "two-pair",
	threeKindHand: "three-kind",
	fullHouseHand: "full-house",
	fourKindHand:  "four-kind",
	fiveKindHand:  "five-kind",
}

var rankLabels = [numRanks]byte{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}

const jokerRank = 9

func parseBid(line string) (parsed bidInfo, err error) {
	cards, bid, ok := strings.Cut(line, " ")
	if !ok {
		return bidInfo{}, fmt.Errorf("bad bid format: %q", line)
	}
	parsed.hand, err = parseHand(cards)
	if err != nil {
		return bidInfo{}, fmt.Errorf("bad hand: %q: %w", line, err)
	}
	parsed.value, err = strconv.Atoi(bid)
	if err != nil {
		return bidInfo{}, fmt.Errorf("non-numeric bid value: %q: %w", line, err)
	}
	return parsed, nil
}

func parseHand(cards string) (parsed handInfo, err error) {
	if len(cards) != handSize {
		return handInfo{}, fmt.Errorf("expected %d cards, got %d: %q", handSize, len(cards), cards)
	}
	for i, c := range []byte(cards) {
		rank := bytes.IndexByte(rankLabels[:], c)
		if rank < 0 {
			return handInfo{}, fmt.Errorf("not a card: %q", c)
		}
		parsed.cards[i] = byte(rank)
	}
	parsed.typ = classifyHand(parsed.cards)
	return parsed, nil
}

func classifyHand(cards [handSize]byte) handType {
	var (
		cardCounts  [numRanks]byte
		countCounts [handSize]byte
	)
	for _, c := range cards {
		cardCounts[c]++
	}
	for _, cc := range cardCounts {
		if cc > 0 {
			countCounts[cc-1]++
		}
	}
	switch {
	case countCounts[4] == 1:
		return fiveKindHand
	case countCounts[3] == 1:
		return fourKindHand
	case countCounts[2] == 1 && countCounts[1] == 1:
		return fullHouseHand
	case countCounts[2] == 1:
		return threeKindHand
	case countCounts[1] == 2:
		return twoPairHand
	case countCounts[1] == 1:
		return onePairHand
	default:
		return highCardHand
	}
}

func (hi *handInfo) jokerify() {
	hi.typ = classifyWithJokers(hi.cards)
	for i, c := range hi.cards {
		if c == jokerRank {
			hi.cards[i] = 0
		} else if c < jokerRank {
			hi.cards[i] = c + 1
		}
	}
}

func classifyWithJokers(cards [handSize]byte) handType {
	var (
		cardCounts  [numRanks]byte
		countCounts [handSize]byte
		jokers      byte
	)
	for _, c := range cards {
		if c == jokerRank {
			jokers++
		} else {
			cardCounts[c]++
		}
	}
	for _, cc := range cardCounts {
		if cc > 0 {
			countCounts[cc-1]++
		}
	}
	switch {
	case countCounts[4] == 1 || countCounts[3] == 1 && jokers == 1 || countCounts[2] == 1 && jokers == 2 || countCounts[1] == 1 && jokers == 3 || jokers >= 4:
		return fiveKindHand
	case countCounts[3] == 1 || countCounts[2] == 1 && jokers == 1 || countCounts[1] == 1 && jokers == 2 || jokers >= 3:
		return fourKindHand
	case countCounts[2] == 1 && countCounts[1] == 1 || countCounts[1] == 2 && jokers == 1:
		return fullHouseHand
	case countCounts[2] == 1 || countCounts[1] == 1 && jokers == 1 || jokers >= 2:
		return threeKindHand
	case countCounts[1] == 2:
		return twoPairHand
	case countCounts[1] == 1 || jokers >= 1:
		return onePairHand
	default:
		return highCardHand
	}
}
