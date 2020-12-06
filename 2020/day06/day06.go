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

// Package day06 solves AoC 2020 day 6.
package day06

import (
	"strconv"
	"strings"

	"github.com/fis/aoc-go/util"
)

func Solve(path string) ([]string, error) {
	data, err := util.ReadChunks(path)
	if err != nil {
		return nil, err
	}
	answers := make([][]answerSet, len(data))
	for i, chunk := range data {
		lines := strings.Split(strings.TrimSpace(chunk), "\n")
		answers[i] = parseGroup(lines)
	}

	var any, all int
	for _, a := range answers {
		any += countMerged(a, mergeAny)
		all += countMerged(a, mergeAll)
	}

	return []string{strconv.Itoa(any), strconv.Itoa(all)}, nil
}

type answerSet [26]bool

func parseGroup(lines []string) (answers []answerSet) {
	for _, s := range lines {
		answers = append(answers, parseAnswerSet(s))
	}
	return answers
}

func parseAnswerSet(s string) (answers answerSet) {
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			answers[c-'a'] = true
		}
	}
	return answers
}

func countMerged(answers []answerSet, merger func([]answerSet) answerSet) (count int) {
	merged := merger(answers)
	for _, a := range merged {
		if a {
			count++
		}
	}
	return count
}

func mergeAny(in []answerSet) (out answerSet) {
	for _, answers := range in {
		for i, a := range answers {
			out[i] = out[i] || a
		}
	}
	return out
}

func mergeAll(in []answerSet) (out answerSet) {
	for i := range out {
		out[i] = true
	}
	for _, answers := range in {
		for i, a := range answers {
			out[i] = out[i] && a
		}
	}
	return out
}
