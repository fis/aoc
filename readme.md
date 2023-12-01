# Advent of Code

This repository contains my solutions for
[Advent of Code](https://adventofcode.com/), mostly in Go.

Here's a map of what's in this repository.

- Documentation
  - Puzzle diaries, i.e., notes written during the contests:
    - [2019 notes](docs/2019-notes.md)
    - [2020 notes](docs/2020-notes.md)
    - [2021 notes](docs/2021-notes.md)
    - [2022 notes](docs/2022-notes.md)
  - Some notes on the earlier puzzles, solved outside the actual contest time:
    - [2017 notes](docs/2017-notes.md)
- Go packages
  - `20??/*`: Individual AoC puzzle solutions. If there's any code shared
    between multiple puzzles that's still specific to one year, it's also in
    here. The top-level subpackage for each year serves two functions: it
    imports all the days (so other packages can get them all as a group), and
    also contains a unit test to verify each puzzle using the puzzle inputs and
    outputs in the `testdata/YYYY/dayDD.{txt,out}` files.
  - `cmd/aoc`: Multipurpose binary to execute any of the puzzles.
  - `glue`: Framework code so that the individual puzzle solutions can register
    solvers (and possibly other related utilities) for the binary via init
    functions.
  - `util`: Utility code useful for solutions across years. Of special note are
    the types `util.Level` (for 2D roguelike style data) and `util.Graph` (for
    labeled digraphs). There are also two packages below this one:
    - `util/fn`: Very non-idiomatic-Go higher order functions, for conciseness.
    - `util/ix`: Integer functions that show up a lot in AoC puzzles.
- Python code
  - `2019-py`: The initial 2019 solutions I wrote in Python, before starting
    this whole Go adventure. May contain assorted odds and ends as well.
  - `vis`: Data visualization code. See [vis/readme.md](vis/readme.md).
- [Esoteric languages](https://esolangs.org/)
  - `eso/*/*.blsq`: Solutions in [Burlesque](https://mroman.ch/burlesque/).
  - `eso/*/*.b98`: Solutions in
    [(Be)funge 98](https://esolangs.org/wiki/Funge-98).
