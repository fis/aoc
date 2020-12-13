# Advent of Code

This repository contains my solutions for
[Advent of Code](https://adventofcode.com/), mostly in Go.

Here's a map of what's in this repository.

- Documentation
  - Overly verbose notes on the puzzle solutions:
    - [2020 notes](docs/2020-notes.md)
- Go packages
  - `2018/*`, `2019/*`, `2020/*`: Individual AoC puzzle solutions. If there's
    any code  shared between multiple puzzles that's still specific to one year,
    it's also in here. The `days` subpackage for each year serves two functions:
    it imports all the days, and also contains a unit test to verify each puzzle
    using the puzzle inputs in `YYYY/days/testdata/`.
  - `cmd/aoc`: multipurpose binary to execute any of the puzzles.
  - `glue`: Framework code so that the individual puzzle solutions can register
    solvers (and possibly other related utilities) for the binary via init
    functions.
  - `util`: Utility code useful for solutions across years. Of special note are
    the types `util.Level` (for 2D roguelike style data) and `util.Graph` (for
    labeled digraphs).
