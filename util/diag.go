// Copyright 2019 Google LLC
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

package util

import (
	"flag"
	"fmt"
)

var verbose = flag.Bool("verbose", false, "turn on diagnostic outputs")

// Diag prints its arguments to the diagnostic output stream.
func Diag(a ...interface{}) {
	if !*verbose {
		return
	}
	fmt.Print(a...)
}

// Diagln prints its arguments, followed by a newline, to the diagnostic output stream.
func Diagln(a ...interface{}) {
	if !*verbose {
		return
	}
	fmt.Println(a...)
}

// Diagf formats data to the diagnostic output stream.
func Diagf(format string, a ...interface{}) {
	if !*verbose {
		return
	}
	fmt.Printf(format, a...)
}

// IsDiag tests if verbose output has been requested.
func IsDiag() bool {
	return *verbose
}
