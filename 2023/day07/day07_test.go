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

package day07

import (
	"testing"

	"github.com/fis/aoc/util/fn"
)

var ex = []string{
	"32T3K 765",
	"T55J5 684",
	"KK677 28",
	"KTJJT 220",
	"QQQJA 483",
}

func TestTotalWinnings(t *testing.T) {
	bids, err := fn.MapE(ex, parseBid)
	if err != nil {
		t.Fatal(err)
	}
	want := 6440
	if got := totalWinnings(bids); got != want {
		t.Errorf("totalWinnings(ex) = %d, want %d", got, want)
	}
}

func TestTotalWinningsWithJokers(t *testing.T) {
	bids, err := fn.MapE(ex, parseBid)
	if err != nil {
		t.Fatal(err)
	}
	jokerify(bids)
	want := 5905
	if got := totalWinnings(bids); got != want {
		t.Errorf("totalWinnings(j(ex)) = %d, want %d", got, want)
	}
}

func TestClassifyWithJokers(t *testing.T) {
	tests := []struct {
		hand string
		want handType
	}{
		// example hands
		{"32T3K", onePairHand},
		{"T55J5", fourKindHand},
		{"KK677", twoPairHand},
		{"KTJJT", fourKindHand},
		{"QQQJA", fourKindHand},
		// more systematic test cases
		{"23456", highCardHand},
		{"22345", onePairHand},
		{"22334", twoPairHand},
		{"22234", threeKindHand},
		{"22233", fullHouseHand},
		{"22223", fourKindHand},
		{"22222", fiveKindHand},
		{"2345J", onePairHand},
		{"2344J", threeKindHand},
		{"2233J", fullHouseHand},
		{"2223J", fourKindHand},
		{"2222J", fiveKindHand},
		{"234JJ", threeKindHand},
		{"233JJ", fourKindHand},
		{"222JJ", fiveKindHand},
		{"23JJJ", fourKindHand},
		{"22JJJ", fiveKindHand},
		{"2JJJJ", fiveKindHand},
		{"JJJJJ", fiveKindHand},
	}
	for _, test := range tests {
		hi, err := parseHand(test.hand)
		if err != nil {
			t.Errorf("parseHand(%s): %v", test.hand, err)
		} else if got := classifyWithJokers(hi.cards); got != test.want {
			t.Errorf("classifyWithJokers(%s) = %v, want %v", test.hand, got, test.want)
		}
	}
}
