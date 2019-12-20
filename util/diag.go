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
