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

// Package day25 solves AoC 2020 day 25.
package day25

import (
	"fmt"
	"math"

	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2020, 25, glue.IntSolver(solve))
}

func solve(input []int) ([]string, error) {
	if len(input) != 2 {
		return nil, fmt.Errorf("expected two numbers, got %d", len(input))
	}
	key := findKey(input[0], input[1], pohligHellman)
	return glue.Ints(key), nil
}

func findKey(pub1, pub2 int, log func(b, a, m int) int) (key int) {
	loop1 := log(7, pub1, 20201227)
	return pow(pub2, loop1, 20201227)
}

// Algorithms for solving x in b^x = a (mod m).

func trialMultiplication(b, a, m int) (x int) {
	x = 0
	for c := 1; c != a; c = (c * b) % m {
		x++
	}
	return x
}

func babyStep(b, a, m int) (x int) {
	return babyStepWithOrder(b, a, m, m-1)
}

func babyStepWithOrder(b, a, m, n int) (x int) {
	logTable := map[int]int{}

	mm := int(math.Ceil(math.Sqrt(float64(n))))
	for j, a := 0, 1; j < mm; j++ {
		logTable[a] = j
		a = (a * b) % m
	}
	amm := pow(b, m-mm-1, m)
	for i := 0; i < m; i++ {
		if j, ok := logTable[a]; ok {
			return ((i*mm)%m + j) % m
		}
		a = (a * amm) % m
	}
	panic("impossible: did not find solution")
}

func pohligHellman(b, a, m int) (x int) {
	if m != 20201227 {
		panic("only implemented for the group modulus 20201227")
	}
	n := 20201226
	pi := []int{116099, 29, 3, 2} // factors of 20201226
	xi := make([]int, len(pi))
	for i, p := range pi {
		bi := pow(b, n/p, m)
		ai := pow(a, n/p, m)
		xi[i] = babyStepWithOrder(bi, ai, m, p)
	}
	for len(pi) > 1 {
		xa, xb := xi[0], xi[1]
		ma, mb := pi[0], pi[1]
		for ; ; xa += ma {
			if xa%mb == xb {
				xi[1] = xa
				pi[1] = ma * mb
				xi, pi = xi[1:], pi[1:]
				break
			}
		}
	}
	return xi[0]
}

func pow(b, e, m int) (p int) {
	p = 1
	for e > 0 {
		if e&1 == 1 {
			p = (p * b) % m
		}
		e >>= 1
		b = (b * b) % m
	}
	return p
}
