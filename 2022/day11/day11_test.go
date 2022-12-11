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

package day11

import (
	"strings"
	"testing"

	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/fn"
)

var ex = strings.TrimPrefix(`
Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1
`, "\n")

func TestMonkeyBusiness(t *testing.T) {
	monkeys, err := fn.MapE(util.Chunks(ex), parseMonkey)
	if err != nil {
		t.Fatal(err)
	}
	want := 10605
	if got := monkeyBusiness(monkeys, 20); got != want {
		t.Errorf("monkeyBusiness(ex, 20) = %d, want %d", got, want)
	}
}

func TestWorryingMonkeyBusiness(t *testing.T) {
	monkeys, err := fn.MapE(util.Chunks(ex), parseMonkey)
	if err != nil {
		t.Fatal(err)
	}
	want := 2713310158
	if got := worryingMonkeyBusiness(monkeys, 10000); got != want {
		t.Errorf("worryingMonkeyBusiness(ex, 10000) = %d, want %d", got, want)
	}
}
