# AoC 2020 solution notes

## Background

In December 2019, I spent a while playing around with the Advent of Code
puzzles, including writing a set of least-effort solutions in Python
([github.com/fis/aoc-py](https://github.com/fis/aoc-py/)), along the slightly
more robust (and verbose) Go solutions in this repository. As part of that, I
wrote occasionally extensive notes about each of the puzzles, for no discernible
reason. The 2019 notes can be found in that other repository.

Now it's December 2020, and this year I'm only doing the Go solutions. Since
writing the notes was so much fun, here are the notes for this year.

## [Day 1](https://adventofcode.com/2020/day/1)

As usual, the contest starts very simply. Given the framework has built-in
support for processing integer-formatted inputs and outputs, the solution is
just a matter of testing each pair (part 1) or 3-tuple (part 2) for the desired
property (summing up to 2020).

A more concise language than Go would make this *very* short.

## [Day 2](https://adventofcode.com/2020/day/2)

Day 2 involves validating some passwords, another simple implementation task.
The solution gets a little fancy by defining a `policy` type, with
`validateSled` / `validateToboggan` methods for the different rules.

## [Day 3](https://adventofcode.com/2020/day/3)

Third day brings us a first instance of an AoC staple, the two-dimensional
roguelike-style level. The `util.Level` utilities from last year make short work
of this. The task just involves counting trees on a given rational-number slope.

## [Day 4](https://adventofcode.com/2020/day/4)

Another data validation task. It's been brought to my attention that some people
actually like the mundanity of it. In this case, the only difference between
parts 1 and 2 is the set of validation rules, with those of part 2 being much
more extensive. As a result, this ends up being the first day to score highly on
the twistiness scale, close to 4.

(See [/vis/readme.md](../vis/readme.md) for details of how this is measured.)

The Go solution does make a bit of an attempt to represent the rules in a
relatively compact manner. Some are simply validated using a regular expression
while others reuse a numeric upper/lower bound checker. The height test, which
supports two different units, is the most complex of these.

Day 4 also introduces convenience functions for reading paragraphs separated by
blank lines in the `util` package.

## [Day 5](https://adventofcode.com/2020/day/5)

On day 5, the description really does its best to obfuscate the simplicity of
the actual problem. Calling it *binary space partitioning* is a little
misleading, though: really, the input is just regular binary numbers, with
letters in place of digits.

The Go solution decodes the row and column separately. In the end, part 2 does
not require that either, so that separation is completely useless, and the
function could just return the seat ID directly. But since it already had
independent unit tests...

## [Day 6](https://adventofcode.com/2020/day/6)

It's hard to say anything special about day 6, given its simplicity. The Go
solution manages to be quite verbose nevertheless.

## [Day 7](https://adventofcode.com/2020/day/7)

After a long wait, graphs are back!

Day 7 involves a weighted DAG, where nodes are bag colors (see the puzzle
flavour text for what that's all about), and edges indicate how many bags of the
head type must be included in the tail. The two tasks revolve around the special
*shiny gold* bag: part 1 asks how many different bag types can contain it, and
part 2 asks how many bags of any kind are inside it. Accordingly, part 1 is
essentially counting the number of ancestor nodes, while part 2 involves
evaluating a recursively defined quantity.

The Go solutions for both parts are superficially similar to a depth-first
search on the graph, implemented as a recursive function. More specifically, the
first part is literally just that on the predecessor relation (i.e., following
edges backwards), counting the number of visited nodes. You might argue the
second part is not exactly a depth-first search, but instead a memoized
recursive function, but the structure is of course very similar nevertheless.

As is the rule with anything graph-shaped, this repository includes visual
representations of the example graphs rendered by GraphViz. See the
`2020-day07-ex1.png` and `2020-day07-ex2.png` files in this directory; the
subcommand `aoc plot 2020 7` is what produced them. The *shiny gold* bag is
(incongruously) red, the ancestor nodes relevant for part 1 are green, and the
descendant nodes relevant for part 2 are blue. Any irrelevant nodes (neither
ancestors, nor descendants) would be white, except that the examples have none.

Unfortunately, the results of this representation of the actual puzzle input are
not really usefully rendered by GraphViz tools, at least without fiddling. A
scaled-down (20%) version is included as `2020-day07-small.png` anyway, just to
prove the point. On the other hand, it does have some white nodes as well.

## [Day 8](https://adventofcode.com/2020/day/8)

Hearkening back to the [Intcode](https://esolangs.org/wiki/Intcode) days of
2019, day 8 introduces a (very) simple instruction set, with just an
accumulator, (unconditional) jumps and nops. The two-part task is to first
execute the given program (and terminate the first time an instruction is
encountered a second time), then figure out which single `jmp` can be changed to
`nop` (or vice versa) to make the program halt properly (by reaching the end of
the program).

Part 1 is solved simply by executing the instructions (with a bitmap to detect
cycles). It would be perfectly reasonable to solve part 2 in `O(n^2)` time, by
just attempting to modify each `jmp`/`nop` instruction (at most `n`) in turn,
and then executing the program (`O(n)` steps) to see if it now halts or loops.
In fact, the initial Go version (see history of `2020/day08/day08.go` if
curious) did this. However, we can do better.

Let's trot GraphViz out again. The Go solution has a utility to show the full
potential control flow (via the `dot8` option to the runner binary). The common
example is rendered into `2020-day08-ex.png` in this directory. Black edges are
those of the unmodified program, while red edges denote the alternative control
flow if that instruction was to be flipped. GraphViz is again useless on the
full puzzle input.

The graph form makes it clear that asymptotically faster solutions exist. In
particular, by treating the graph as a weighted digraph and giving each black
edge a weight of 0 and red edge a weight of 1, the graph will have the following
properties, following from the puzzle specification:

- There is no path of length 0 from the first instruction to the halt node.
  Otherwise the initial program would not loop indefinitely.
- There is exactly one path of length 1 (ignoring zero-length cycles). This
  follows from the fact that the puzzle specifies the program can be made to
  halt by changing a single unique instruction (i.e., following all black edges
  except for a single red one).

This means finding the shortest path in the resulting graph (and then summing up
all `acc` instructions on that path) solves the problem. Since the number of
vertices is `n`, and the number of edges can be at most `2n` (for a program that
contains only `jmp`/`nop` instructions), a naïve implementation of Dijkstra's
algorithm with a binary heap can solve the problem in `O(n log n)` time.

Further, considering what Dijkstra's algorithm would do on a graph with this
specific structure in more detail will show the same behavior can be implemented
without the need for the conventional priority queue, leading to an `O(n)`
algorithm.

Since Dijkstra's algorithm explores the nodes of the graph in the order of the
lowest tentative distance, and since all the relevant distances in this case are
either 0 or 1, an execution of the algorithm will have the following stages:

- Starting from the initial node, the algorithm will follow all the weight 0
  (black) edges, i.e., the unmodified opcodes, marking all the instructions the
  initial execution would touch with a tentative distance of 0.
- For every `jmp`/`nop` instruction, it will also include their alternative
  interpretations as unvisited nodes with a tentative distance of 1. Since they
  all share the same distance, the order does not matter. Of course if the
  initial execution reaches one of these nodes, it will get a lower tentative
  distance of 0, meaning that path will be fruitless to examine.
- After the initial marking, it will arbitrarily pick one of the possible
  unvisited nodes with a distance of 1. It will again follow the black edges
  (execute the instructions), but this time we can ignore the alternative edges,
  since they would have a distance of 2, and by definition a shorter path to the
  halt node exists.
- If at any point the subsequent execution would encounter a node that has
  already been visited (with a tentative distance of either 0 or 1), it will
  give up this path, since it would be no better than what has already been
  considered. It would then pick the next unvisited node with a distance of 1,
  and repeat the process.
- The algorithm will terminate when it reaches the halt node during one of these
  executions, and the path taken will then indicate the solution.

Since the algorithm visits at most `n` nodes, and for every node it performs
only `O(1)` operations, this yields an `O(n)` solution. The Go code implements
this algorithm directly using the program code as the data structure, without
explicitly constructing the graph. It simply uses a bitmap of seen instructions,
and a list of potential branching points (along with the accumulator state),
exploring from each branch only the new unvisited instructions.

## [Day 9](https://adventofcode.com/2020/day/9)

In the [advent calendar](https://adventofcode.com/2020), day 9 is separated from
the preceding 8 days by a blank line. Could this mean change is afoot, and the
next set of puzzles is going to be all different? In a word, no. (And day 10 is
separated by a much larger gap, but of course that was not visible until later.)

Both parts of day 9 can be solved using a sliding window.

In part 1, the task is to find the first number that's *not* the sum of one of
the preceding 25 numbers. This is simply a matter of keeping a map of the 25*24
= 600 valid sums, and testing each number against it in turn. If the number is
valid, the map of sums can be updated by removing all sums of the outgoing
number with the 24 remaining, and adding the new ones with the incoming number.

Part 2, on the other hand, asks to find the contiguous interval that sums up to
a given target value. Since the numbers are all positive integers, this can be
done simply by maintaining the bounds and sum of a candidate interval. If its
sum is too low, numbers are added to it by moving the right edge forward. If too
high, numbers are instead removed from it by moving the left edge forward.

To show that this algorithm finds the target region, we can rely on the
following invariant, maintained by every step of the algorithm: both the left
and right edges of the candidate interval are never advanced past the true edges
of the target interval. Since each step of the algorithm increments one of the
two edges, it's clear that after `O(n)` steps, the algorithm must have converged
to the solution.

First, let's look at the left edge. If it has not yet reached the true position,
it doesn't matter whether we increment it or not, so the only situation we need
to consider is when the edge is at the true location. Since we know that the
right edge is currently either before or at its true location, the sum of the
candidate interval is not larger than the target value: this is because the
target interval is a superset of the candidate. So we will never increment the
left edge in this case.

The situation for the right edge is essentially symmetric. If it hasn't reached
the final location yet, we may increment it or not, as we please.  If it has, we
know that the left edge is again before or at its true location, and in this
case the sum can not be less than the target (because the target interval in
this case is a subset). So we will never increment the right edge past its true
location either.

We can also show that the edges never cross. If it should happen that they're
ever co-located, the sum of the candidate interval is zero. This is always below
the target, so we will increment the right edge, not the left.

(The Go implementation will crash and burn violently if the sum of all values in
the array is less than the target value. So don't do that.)

## [Day 10](https://adventofcode.com/2020/day/10)

The biggest challenge on day 10 is to wade through the problem description. The
actual task is in fact quite simple, since the only ways to connect all these
adapters together are in ascending order.

The Go solution for part 1 boils down to just sorting the list of numbers, and
counting how many differences of 1 and 3 there are between two consecutive
numbers.

For part 2, there's the obvious dynamic programming approach to the recursive
definition. The number of ways one adapter can be plugged in to the device is
the sum of ways one, two or three higher-joltage adapters can be connected,
depending on how many of those three can be plugged into it. So we can just
iterate backwards and sum them up. Or the other way around, since the problem is
completely symmetric.

## [Day 11](https://adventofcode.com/2020/day/11)

Ooh, it's the first obvious cellular automaton of the year.

In [B/S terms](https://conwaylife.com/wiki/Rulestring), part 1 is called
`B0/S0123`, while part 2 is `B0/S01234`. But of course this description omits
the most interesting part of the puzzle, which is how the neighbourhood is
defined.  While part 1 is in a sense the standard
[Moore neighbourhood](https://conwaylife.com/wiki/Moore_neighbourhood), it also
has the special floor tiles that are always empty, so a seat may in practice
have less than 8 neighbours. Part 2 is even more unconventional, as it uses a
line-of-sight algorithm, so the neighbours can be very non-local.

The Go solution doesn't do anything fancy. It parses the level into a map of
every seat's neighbours, and then updates the occupancy state vector (which has
one element per seat) based on the rules, and counts what's left when it no
longer changes.

## [Day 12](https://adventofcode.com/2020/day/12)

Day 12 feels like
[Logo](https://en.wikipedia.org/wiki/Logo_(programming_language)) to me.
Unfortunately, there's nothing particularly puzzling about the task. We just go
through the motions.

## [Day 13](https://adventofcode.com/2020/day/13)

Without checking yet, I expect day 13 to score reasonably well on the twistiness
scale. (Again, see [/vis/readme.md](../vis/readme.md) for details of how this is
measured.)

Part 1 is trivial: given a timestamp, we just need to round it up to the next
multiple of each interval, and see which one yields the shortest wait.

Part 2 requires a little more thought. The task is to find the first time `T`
such that the `N`th bus in the schedule (starting from index 0), if it's in
service at all, departs at time `T+N`.

A little more formally, if we number all buses in service (let's call them
constraints) with index `i`, we have a system of congruences for the desired
time `T`:

```
T + N_i ≡ 0 (mod p_i),  or equivalently
T ≡ k_i (mod p_i),      where k_i = -N_i mod p_i,
```

and `p_i` are the bus IDs, which all coincidentally happen to be prime numbers.
This is a
[Chinese remainder problem](https://en.wikipedia.org/wiki/Chinese_remainder_theorem),
for which there are several algorithms. The one used here is on the simple end
of the scale, and basically repeatedly merges two congruences (modulo `a` and
`b`) into a single one (modulo `a*b`), until only a single congruence
remains (and gives the solution). The merging is done by simply iterating over
the possible values given by the first (larger) congruence, and selecting the
first one that satisfies the other.

Merging the constraints could be done more efficiently, with the extended
Euclidean algorithm. But the simple solution already only takes `0.002s` for my
puzzle input (`go test -count 1 -run TestAllDays/day=13 ./2020/days`), so it
hardly seems necessary.

## [Day 14](https://adventofcode.com/2020/day/14)

For day 14, there's an instruction set, but not really any more complicated than
before.

Part 1 is straightforward: just run through the instructions, and sum what was
set. A mask of the form `0011XX` is turned into two bit-masks, `001100` (having
a bit set for every `1`) and `001111` (the complement of having a bit set for
every `0`), which are then OR'd and AND'd, respectively, with the value.

For part 2, the mask semantics change: they now modify the address, and `X` now
means to propagate the value to all possible combinations. It feels like there
should be a clever data structure, perhaps some sort of a tree, that can be used
to represent the contents of the memory without explicitly broadcasting the
values. Unfortunately, the explicit solution is sufficiently fast (`0.018s`,
says `go test`) for me to not bother figuring that out.
