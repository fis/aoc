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

// Binary aoc provides all the supported AoC actions (solving, plotting, ...).
package main

import (
	"github.com/fis/aoc/glue"

	_ "github.com/fis/aoc/2016" // solvers
	_ "github.com/fis/aoc/2017" // solvers
	_ "github.com/fis/aoc/2018" // solvers
	_ "github.com/fis/aoc/2019" // solvers
	_ "github.com/fis/aoc/2020" // solvers
	_ "github.com/fis/aoc/2021" // solvers
	_ "github.com/fis/aoc/2022" // solvers
	_ "github.com/fis/aoc/2023" // solvers
)

func main() {
	glue.Main()
}
