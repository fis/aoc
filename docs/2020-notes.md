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

Part 1 is solved simply by executing the instructions (with a bitmap
to detect cycles). It would be perfectly reasonable to solve part 2 in
`O(n^2)` time, by just attempting to modify each `jmp`/`nop`
instruction (at most `n`) in turn, and then executing the program
(`O(n)` steps) to see if it now halts or loops. In fact, the initial
Go version (see history of `2020/day08/day08.go` if curious) did
this. However, we can do better.

Let's trot GraphViz out again. The Go solution has a utility to show
the full potential control flow (via the `dot8` option to the runner
binary). The common example is rendered into `2020-day08-ex.png` in
this directory. Black edges are those of the unmodified program, while
red edges denote the alternative control flow if that instruction was
to be flipped. GraphViz is again useless on the full puzzle input.

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
a naïve implementation of Dijkstra's algorithm with a binary heap can
solve the problem in `O(n log n)` time.

Further, considering what Dijkstra's algorithm would do on a graph
with this specific structure in more detail will show the same
behavior can be implemented without the need for the conventional
priority queue, leading to an `O(n)` algorithm.

Since Dijkstra's algorithm explores the nodes of the graph in the
order of the lowest tentative distance, and since all the relevant
distances in this case are either 0 or 1, an execution of the
algorithm will have the following stages:

- Starting from the initial node, the algorithm will follow all the
  weight 0 (black) edges, i.e., the unmodified opcodes, marking all
  the instructions the initial execution would touch with a tentative
  distance of 0.
- For every `jmp`/`nop` instruction, it will also include their
  alternative interpretations as unvisited nodes with a tentative
  distance of 1. Since they all share the same distance, the order
  does not matter. Of course if the initial execution reaches one of
  these nodes, it will get a lower tentative distance of 0, meaning
  that path will be fruitless to examine.
- After the initial marking, it will arbitrarily pick one of the
  possible unvisited nodes with a distance of 1. It will again follow
  the black edges (execute the instructions), but this time we can
  ignore the alternative edges, since they would have a distance of 2,
  and by definition a shorter path to the halt node exists.
- If at any point the subsequent execution would encounter a node that
  has already been visited (with a tentative distance of either 0 or
  1), it will give up this path, since it would be no better than what
  has already been considered. It would then pick the next unvisited
  node with a distance of 1, and repeat the process.
- The algorithm would terminate when it reaches the halt node during
  one of these executions, and the path taken would then indicate the
  solution.

Since the algorithm visits at most `n` nodes, and for every node it
performs only `O(1)` operations, this yields an `O(n)` solution. The
Go code implements this algorithm directly using the program code as
the data structure, without explicitly constructing the graph. It
simply uses a bitmap of seen instructions, and a list of potential
branching points (along with the accumulator state), exploring from
each branch only the new unvisited instructions.

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

Part 2, on the other hand, asks to find the contiguous interval that
sums up to a given target value. Since the numbers are all positive
integers, this can be done simply by maintaining the bounds and sum of
a candidate interval. If its sum is too low, numbers are added to it
by moving the right edge forward. If too high, numbers are instead
removed from it by moving the left edge forward.

To show that this algorithm finds the target region, we can rely on
the following invariant, maintained by every step of the algorithm:
both the left and right edges of the candidate interval are never
advanced past the true edges of the target interval. Since each step
of the algorithm increments one of the two edges, it's clear that
after `O(n)` steps, the algorithm must have converged to the solution.

First, let's look at the left edge. If it has not yet reached the true
position, it doesn't matter whether we increment it or not, so the
only situation we need to consider is when the edge is at the true
location. Since we know that the right edge is currently either before
or at its true location, the sum of the candidate interval is not
larger than the target value: this is because the target interval is a
superset of the candidate. So we will never increment the left edge in
this case.

The situation for the right edge is essentially symmetric. If it
hasn't reached the final location yet, we may increment it or not, as
we please.  If it has, we know that the left edge is again before or
at its true location, and in this case the sum can not be less than
the target (because the target interval in this case is a subset). So
we will never increment the right edge past its true location either.

We can also show that the edges never cross. If it should happen that
they're ever co-located, the sum of the candidate interval is zero.
This is always below the target, so we will increment the right edge,
not the left.

(The Go implementation will crash and burn violently if the sum of all
values in the array is less than the target value. So don't do that.)
