// Copyright 2025 Google LLC
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

package ix

import "cmp"

type Rat struct {
	P, Q int
}

func RatOf(p, q int) Rat {
	if q == 0 {
		panic("invalid rat")
	}
	if q < 0 {
		p, q = -p, -q
	}
	return Rat{p, q}.simplify()
}

func RatU(p, q int) Rat {
	if q < 0 {
		p, q = -p, -q
	}
	return Rat{p, q}
}

func RatI(n int) Rat            { return Rat{n, 1} }
func (r Rat) Add(s Rat) Rat     { return Rat{r.P*s.Q + s.P*r.Q, r.Q * s.Q}.simplify() }
func (r Rat) AddI(n int) Rat    { return Rat{r.P + r.Q, r.Q} }
func (r Rat) Sub(s Rat) Rat     { return Rat{r.P*s.Q - s.P*r.Q, r.Q * s.Q}.simplify() }
func (r Rat) SubI(n int) Rat    { return Rat{r.P - r.Q, r.Q} }
func (r Rat) Mul(s Rat) Rat     { return Rat{r.P * s.P, r.Q * s.Q}.simplify() }
func (r Rat) MulI(n int) Rat    { return Rat{r.P * n, r.Q}.simplify() }
func (r Rat) MulIU(n int) Rat   { return Rat{r.P * n, r.Q} }
func (r Rat) DivI(n int) Rat    { return Rat{r.P, r.Q * n}.simplify() }
func (r Rat) Compare(s Rat) int { return cmp.Compare(r.P*s.Q, s.P*r.Q) }
func (r Rat) IsValid() bool     { return r.Q != 0 }
func (r Rat) IsInt() bool       { return r.P%r.Q == 0 }

func (r Rat) Floor() int {
	if r.P >= 0 {
		return r.P / r.Q
	} else {
		return (r.P - r.Q + 1) / r.Q
	}
}

func (r Rat) Ceil() int {
	if r.P >= 0 {
		return (r.P + r.Q - 1) / r.Q
	} else {
		return r.P / r.Q
	}
}

func (r Rat) simplify() Rat {
	if d := GCD(Abs(r.P), r.Q); d > 1 {
		return Rat{r.P / d, r.Q / d}
	}
	return r
}

func RatSign(p, q int) int {
	switch {
	case p == 0, q == 0:
		return 0
	case p > 0 && q > 0, p < 0 && q < 0:
		return 1
	default:
		return -1
	}
}
