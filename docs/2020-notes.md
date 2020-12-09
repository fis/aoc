# AoC 2020 solution notes

## Background

In December 2019, I spent a while playing around with the Advent of
Code puzzles, including writing a set of least-effort solutions in
Python ([github.com/fis/aoc-py](https://github.com/fis/aoc-py/)),
along the slightly more robust (and verbose) Go solutions in this
repository. As part of that, I wrote occasionally extensive notes
about each of the puzzles, for no discernible reason. The 2019 notes
can be found in that other repository.

Now it's December 2020, and this year I'm only doing the Go solutions.
Since writing the notes was so much fun, here are the notes for this
year.

## [Day 1](https://adventofcode.com/2020/day/1)

As usual, the contest starts very simply. The solution relies on the
`util.ReadIntRows` helper from last year, after which it's just a
matter of testing each pair (part 1) or 3-tuple (part 2) for the
desired property (summing up to 2020).

A more concise language than Go would make this *very* short.

## [Day 2](https://adventofcode.com/2020/day/2)

Day 2 involves validating some passwords, another simple
implementation task. The solution gets a little fancy by defining a
`policy` type, with `validateSled` / `validateToboggan` methods for
the different rules.

## [Day 3](https://adventofcode.com/2020/day/3)

Third day brings us a first instance of an AoC staple, the
two-dimensional roguelike-style level. The `util.Level` utilities from
last year make short work of this. The task just involves counting
trees on a given rational-number slope.

## [Day 4](https://adventofcode.com/2020/day/4)

Another data validation task. It's been brought to my attention that
some people actually like the mundanity of it. In this case, the only
difference between parts 1 and 2 is the set of validation rules, with
those of part 2 being much more extensive. As a result, this ends up
being the first day to score highly on the twistiness scale, close to
4.

TODO: Provide a link for twistiness.

The Go solution does make a bit of an attempt to represent the rules
in a relatively compact manner. Some are simply validated using a
regular expression while others reuse a numeric upper/lower bound
checker. The height test, which supports two different units, is the
most complex of these.

Day 4 also introduces convenience functions for reading paragraphs
separated by blank lines in the `util` package.

## [Day 5](https://adventofcode.com/2020/day/5)

On day 5, the description really does its best to obfuscate the
simplicity of the actual problem. Calling it *binary space
partitioning* is a little misleading, though: really, the input is
just regular binary numbers, with letters in place of digits.

The Go solution decodes the row and column separately. In the end,
part 2 does not require that either, so that separation is completely
useless, and the function could just return the seat ID directly. But
since it already had independent unit tests...

## [Day 6](https://adventofcode.com/2020/day/6)

It's hard to say anything special about day 6, given its simplicity.
The Go solution manages to be quite verbose nevertheless.

## [Day 7](https://adventofcode.com/2020/day/7)

After a long wait, graphs are back!

Day 7 involves a weighted DAG, where nodes are bag colors (see the
puzzle flavour text for what that's all about), and edges indicate how
many bags of the head type must be included in the tail. The two tasks
revolve around the special *shiny gold* bag: part 1 asks how many
different bag types can contain it, and part 2 asks how many bags of
any kind are inside it. Accordingly, part 1 is essentially counting
the number of ancestor nodes, while part 2 involves evaluating a
recursively defined quantity.

The Go solutions for both parts are superficially similar to a
depth-first search on the graph, implemented as a recursive function.
More specifically, the first part is literally just that on the
predecessor relation (i.e., following edges backwards), counting the
number of visited nodes. You might argue the second part is not
exactly a depth-first search, but instead a memoized recursive
function, but the structure is of course very similar nevertheless.

As is the rule with anything graph-shaped, this repository includes
visual representations of the example graphs rendered by GraphViz. See
the `2020-day07-ex1.png` and `2020-day07-ex2.png` files in this
directory, and the special `dot7` option to the `cmd/aoc2020` binary
that generated them. The *shiny gold* bag is (incongruously) red, the
ancestor nodes relevant for part 1 are green, and the descendant nodes
relevant for part 2 are blue. Any irrelevant nodes (neither
ancestors, nor descendants) would be white, except that the examples
have none.

Unfortunately, the results of this representation of the actual puzzle
input are not really usefully rendered by GraphViz tools, at least
without fiddling. A scaled-down (20%) version is included as
`2020-day07-small.png` anyway, just to prove the point. On the other
hand, it does have some white nodes as well.

## [Day 8](https://adventofcode.com/2020/day/8)

Hearkening back to the [Intcode](https://esolangs.org/wiki/Intcode)
days of 2019, day 8 introduces a (very) simple instruction set, with
just an accumulator, (unconditional) jumps and nops. The two-part task
is to first execute the given program (and terminate the first time an
instruction is encountered a second time), then figure out which
single `jmp` can be changed to `nop` (or vice versa) to make the
program halt properly (by reaching the end of the program).

The Go solution takes the path of least resistance. Part 1 is solved
simply by executing the instructions (with a bitmap to detect cycles),
while part 2 just attempts to modify each relevant instruction in
turn, re-executing the program a second time. This is clearly an
`O(n^2)` operation, since there is at most `n` instructions to change,
and each execution will simulate at most `n` instructions as well.

More as an excuse to get GraphViz out again, the Go solution also
contains a utility to show the full potential control flow. The common
example is rendered into `2020-day08-ex.png` in this directory. Black
edges are those of the unmodified program, while red edges denote the
alternative control flow if that instruction was to be flipped.
GraphViz is again useless on the full puzzle input.

The graph form makes it clear that asymptotically faster solutions
exist. In particular, by treating the graph as a weighted digraph and
giving each black edge a weight of 0 and red edge a weight of 1, the
graph will have the following properties, following from the puzzle
specification:

- There is no path of length 0 from the first instruction to the halt
  node. Otherwise the initial program would not loop indefinitely.
- There is exactly one path of length 1 (ignoring zero-length cycles).
  This follows from the fact that the puzzle specifies the program can
  be made to halt by changing a single unique instruction (i.e.,
  following all black edges except for a single red one).

This means finding the shortest path in the resulting graph (and then
summing up all `acc` instructions on that path) solves the problem.
Since the number of vertices is `n`, and the number of edges can be at
most `2n` (for a program that contains only `jmp`/`nop` instructions),
Dijkstra's algorithm can solve the problem in `O(n log n)` time. Doing
this is left as an exercise for the reader.

Given the special structure of the graph, even better solutions may
exist. Uncovering those is also left as an exercise.


## [Day 9](https://adventofcode.com/2020/day/9)

In the [advent calendar](https://adventofcode.com/2020), day 9 is
separated from the preceding 8 days by a blank line. Could this mean
change is afoot, and the next set of puzzles is going to be all
different? In a word, no. (And day 10 is separated by a much larger
gap, but of course that was not visible until later.)

Both parts of day 9 can be solved using a sliding window.

In part 1, the task is to find the first number that's *not* the sum
of one of the preceding 25 numbers. This is simply a matter of keeping
a map of the 25*24 = 600 valid sums, and testing each number against
it in turn. If the number is valid, the map of sums can be updated by
removing all sums of the outgoing number with the 24 remaining, and
adding the new ones with the incoming number.

Part 2, on the other hand, asks to find the contiguous range that sums
up to a given target value. Since the numbers are nonnegative, this
can be done simply by maintaining the bounds and sum of a candidate
window. If its sum is too low, new numbers are included by moving the
right edge forward. If too high, old numbers are dropped by moving the
left edge forward instead. The proof that this is a valid algorithm is
left as an exercise for the reader.
