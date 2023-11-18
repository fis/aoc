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

// Package ix contains small integer arithmetic functions in the style of the standard `math` package.
package ix

type signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type integer interface {
	signed | unsigned
}

// Abs returns the absolute value of x.
func Abs[T signed](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

// Sign returns -1, 0 or 1 if x is less than, equal, or greater than zero, respectively.
func Sign[T signed](x T) T {
	if x < 0 {
		return -1
	} else if x > 0 {
		return 1
	}
	return 0
}

// CeilDiv returns the result of the integer division a / b except rounded up (for nonnegative integers).
func CeilDiv[T integer](a, b T) T {
	return (a + b - 1) / b
}

// GCD returns the greatest common divisor of the two arguments.
func GCD[T integer](a, b T) T {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// LCM returns the least common multiple of the two arguments.
func LCM[T integer](a, b T) T {
	return a / GCD(a, b) * b
}

// Sqrt returns the integer square root of y. For a nonnegative argument,
// it is the floor of the mathematical square root. Panics for negative numbers.
func Sqrt(y int) int {
	if y < 0 {
		panic("sqrt(neg)")
	} else if y <= 1 {
		return y
	}
	x0 := y / 2
	x1 := (x0 + y/x0) / 2
	for x1 < x0 {
		x0 = x1
		x1 = (x0 + y/x0) / 2
	}
	return x0
}
