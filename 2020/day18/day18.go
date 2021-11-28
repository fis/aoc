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

// Package day18 solves AoC 2020 day 18.
package day18

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2020, 18, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	p1, p2 := 0, 0

	for _, line := range lines {
		e, err := parseExpr(&tokenizer{input: line})
		if err != nil {
			return nil, fmt.Errorf("invalid simple expression: %s: %w", line, err)
		}
		p1 += e.value()

		e, err = parseAdvanced(&tokenizer{input: line})
		if err != nil {
			return nil, fmt.Errorf("invalid advanced expression: %s: %w", line, err)
		}
		p2 += e.value()
	}

	return glue.Ints(p1, p2), nil
}

type tokenType int

const (
	tokLit tokenType = iota
	tokAdd
	tokMul
	tokOpen
	tokClose
	tokEOF
)

type token struct {
	typ tokenType
	val int
}

type tokenizer struct {
	input string
	q     []token
}

func (t *tokenizer) peek() (token, error) {
	if len(t.q) == 0 {
		if err := t.advance(); err != nil {
			return token{}, err
		}
	}
	return t.q[0], nil
}

func (t *tokenizer) pop() (token, error) {
	tok, err := t.peek()
	if err != nil {
		return token{}, err
	}
	t.q = t.q[1:]
	return tok, nil
}

func (t *tokenizer) advance() error {
	for len(t.input) > 0 && t.input[0] == ' ' {
		t.input = t.input[1:]
	}
	if len(t.input) == 0 {
		t.q = append(t.q, token{typ: tokEOF})
		return nil
	}
	switch t.input[0] {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return t.advanceLit()
	case '+', '*', '(', ')':
		types := map[byte]tokenType{'+': tokAdd, '*': tokMul, '(': tokOpen, ')': tokClose}
		t.q, t.input = append(t.q, token{typ: types[t.input[0]]}), t.input[1:]
		return nil
	default:
		return fmt.Errorf("unexpected byte: %q", t.input[0])
	}
}

func (t *tokenizer) advanceLit() error {
	tail := strings.TrimLeft(t.input, "0123456789")
	head := t.input[:len(t.input)-len(tail)]
	n, err := strconv.Atoi(head)
	if err != nil {
		return fmt.Errorf("invalid literal: %q", head)
	}
	t.q, t.input = append(t.q, token{typ: tokLit, val: n}), tail
	return nil
}

type atom interface {
	value() int
}

type expr struct {
	args []atom
	ops  []tokenType
}

func (e expr) value() int {
	v := e.args[0].value()
	for i, op := range e.ops {
		switch op {
		case tokAdd:
			v += e.args[i+1].value()
		case tokMul:
			v *= e.args[i+1].value()
		}
	}
	return v
}

type literal int

func (n literal) value() int {
	return int(n)
}

func parseExpr(input *tokenizer) (expr, error) {
	return parseList(
		input,
		[]tokenType{tokAdd, tokMul},
		[]tokenType{tokClose, tokEOF},
		func(input *tokenizer) (atom, error) { return parseAtom(input, parseExpr) },
	)
}

func parseAdvanced(input *tokenizer) (expr, error) {
	return parseList(
		input,
		[]tokenType{tokMul},
		[]tokenType{tokClose, tokEOF},
		parseAdvancedAdd,
	)
}

func parseAdvancedAdd(input *tokenizer) (atom, error) {
	return parseList(
		input,
		[]tokenType{tokAdd},
		[]tokenType{tokMul, tokClose, tokEOF},
		func(input *tokenizer) (atom, error) { return parseAtom(input, parseAdvanced) },
	)
}

func parseList(input *tokenizer, sep, end []tokenType, parseArg func(*tokenizer) (atom, error)) (e expr, err error) {
	a, err := parseArg(input)
	if err != nil {
		return expr{}, err
	}
	e.args = append(e.args, a)
next:
	for {
		tok, err := input.peek()
		if err != nil {
			return expr{}, err
		}
		for _, t := range sep {
			if tok.typ == t {
				e.ops = append(e.ops, tok.typ)
				input.pop()
				a, err := parseArg(input)
				if err != nil {
					return expr{}, err
				}
				e.args = append(e.args, a)
				continue next
			}
		}
		for _, t := range end {
			if tok.typ == t {
				return e, nil
			}
		}
		return expr{}, fmt.Errorf("expected one of %v, %v, got %v", sep, end, tok)
	}
}

func parseAtom(input *tokenizer, subExpr func(*tokenizer) (expr, error)) (atom, error) {
	tok, err := input.pop()
	if err != nil {
		return nil, err
	}
	switch tok.typ {
	case tokLit:
		return literal(tok.val), nil
	case tokOpen:
		e, err := subExpr(input)
		if err != nil {
			return nil, err
		}
		tok, err = input.pop()
		if err != nil {
			return nil, err
		}
		if tok.typ != tokClose {
			return nil, fmt.Errorf("expected ), got %v", tok)
		}
		return e, nil
	default:
		return nil, fmt.Errorf("expected literal or (, got %v", tok)
	}
}
