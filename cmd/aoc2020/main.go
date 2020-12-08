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

// Binary aoc2020 runs the AoC 2020 puzzle solutions.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/fis/aoc-go/2020/day07"
	"github.com/fis/aoc-go/2020/days"
)

var specials = map[string]func(){
	"dot7": dot7,
}

func main() {
	flag.Parse()

	if flag.NArg() >= 1 {
		if special := specials[flag.Arg(0)]; special != nil {
			special()
			return
		}
	}

	if flag.NArg() != 2 {
		fmt.Fprintln(os.Stderr, "Usage: aoc2020 [flags] N input.txt")
		os.Exit(1)
	}
	day, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Not a number: %q: %v\n", flag.Arg(0), err)
		os.Exit(1)
	}
	out, err := days.Solve(day, flag.Arg(1))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Solution failed: %v\n", err)
		os.Exit(1)
	}
	for _, s := range out {
		fmt.Println(s)
	}
}

func dot7() {
	examples := map[string][]string{
		"ex1": []string{
			"light red bags contain 1 bright white bag, 2 muted yellow bags.",
			"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
			"bright white bags contain 1 shiny gold bag.",
			"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
			"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
			"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
			"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
			"faded blue bags contain no other bags.",
			"dotted black bags contain no other bags.",
		},
		"ex2": []string{
			"shiny gold bags contain 2 dark red bags.",
			"dark red bags contain 2 dark orange bags.",
			"dark orange bags contain 2 dark yellow bags.",
			"dark yellow bags contain 2 dark green bags.",
			"dark green bags contain 2 dark blue bags.",
			"dark blue bags contain 2 dark violet bags.",
			"dark violet bags contain no other bags.",
		},
	}

	var rules []string
	if flag.NArg() >= 2 {
		var ok bool
		rules, ok = examples[flag.Arg(1)]
		if !ok {
			fmt.Fprintf(os.Stderr, "Unknown example: %q\n", flag.Arg(1))
			os.Exit(1)
		}
	} else {
		data, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		rules = strings.Split(strings.TrimSpace(string(data)), "\n")
	}

	day07.PrintRules(os.Stdout, rules, "shiny gold")
}
