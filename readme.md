# AoC 2019 (Python)

This repository contains my solutions for the 2019 iteration of
[Advent of Code](https://adventofcode.com/). See below for extended
notes.

The sibling repository [aoc2019-go](https://github.com/fis/aoc2019-go)
contains versions in Go, if that's more up your alley.

## [Day 1](https://adventofcode.com/2019/day/1)

Not much to say here. Part 1 involved evaluating a simple expression
on a list of values and printing the sum. Part 2 had a recurrence
relation of sorts.

## [Day 2](https://adventofcode.com/2019/day/2)

Second day introduced the first dialect of the
[Intcode](https://esolangs.org/wiki/Intcode) programming
language. Part 1 asked to run the example program with a specific
initial state, and report which value was left when it halted. Part 2
required finding a specific pair of inputs to yield a given output,
out of just 10000 possibilities.

The solution here uses a standalone Intcode interpreter, only
supporting the opcodes 1, 2 and 99.

## [Day 3](https://adventofcode.com/2019/day/3)

Day 3 considered the intersections of two wires, described as a
sequence of movement instructions (right 75, down 30, ...).

### Part 1

In part 1, the task was simply to find the intersection closest to the
origin. Just for fun, instead of painting every position of the
wires, the solution instead tests every horizontal (resp. vertical)
segment of wire 1 with every vertical (resp. horizontal) segment of
wire 2, and keeps track of the intersections.

### Part 2

Part 2 introduced a notion of a signal propagation delay to the wires,
and asked for the intersection where the combined delay was lowest.

While it would be certainly possible to generalize the part 1 solution
for this, the variant here forgot all about orientation of the
segments. So instead part 2 is solved the boring way, by just looping
over each wire step and keeping track of the delay, then locating the
position with the lowest combined delay.

## [Day 4](https://adventofcode.com/2019/day/4)

For some reason, day 4 was surprisingly simple. Both parts simply
asked how many numbers (from a range of around half a million)
satisfied certain simple properties about their digits.

## [Day 5](https://adventofcode.com/2019/day/5)

Day 5 was mostly about Intcode extensions. This day's solution is the
first to use the shared Intcode interpreter.

The first part introduced numeric input and output instructions, then
just asked for the output of the program given the input.

Second part extended that with conditional jumps and comparison and
equality operations, then asked for the output given another input.

## [Day 6](https://adventofcode.com/2019/day/6)

Sixth day was another surprisingly simple one, involving what were in
the narrative described as orbits, but really it was just an arbitrary
directed tree.

### Part 1

The first subproblem asked for the total number of child
relationships, in the sense that a node was counted both for its
direct parent and its indirect grandparent. The solution simply walks
the tree depth-first, and keeps a running count, incrementing it for
each node by its depth (so counting it as the child of each its
ancestors).

### Part 2

Part 2 wanted to know the shortest path between two nodes. The
solution just finds the unique paths from the root to both nodes by
walking the tree, then trims them by removing their common prefix and
pastes them together.

## [Day 7](https://adventofcode.com/2019/day/7)

Back to the land of Intcode, though with no new instructions.

### Part 1

The setup for the first part called for five independent Intcode
interpreters, each fed a configuration setting and then the previous
interpreter's output. The task was to find the configuration resulting
in the biggest final output.

The solution simply tests each valid configuration, and runs the
programs sequentially.

### Part 2

The twist: a feedback loop. Instead of halting after the output, the
programs continued running, with the output from the last instance fed
back to the first one's input.

The solution here is kind of funny: it actually runs all the programs
as parallel threads, and uses Python's thread-safe queues as the
pipes. It's highly likely there's hardly any actual parallelism (and
4/5 interpreters will just be blocking for input at any one time), but
still.

## [Day 8](https://adventofcode.com/2019/day/8)

Another simple day, decoding images sent as layers of three-state
pixels: white, black, or transparent.

The two parts asked for some statistics of the layers, and the word
obtained by compositing the images, respectively. The solutions are
straightforward. The image renderer writes back-to-front ignoring
transparent pixels, instead of searching front-to-back for the first
non-transparent one.

## [Day 9](https://adventofcode.com/2019/day/9)

The "Intcode every other day" pattern continues. Day 9 introduced
allegedly the final missing Intcode features: a relative addressing
mode (to complement the existing immediate and indirect ones), and an
opcode to set the base register for the relative mode.

The tasks involved no other work than finishing the interpreter. The
answers were simply program outputs for inputs 1 and 2.

## [Day 10](https://adventofcode.com/2019/day/10)

Day 10 resists classification, though both parts were related to a
kind of a line-of-sight computation for ideal, dimensionless asteroids
located on an integer grid.

### Part 1

The first task was to find the asteroid from which the most other
asteroids were visible. The solution involves a visibility checker,
which starts from the delta vector between the two asteroids, reduces
it to the most simplified form based on its GCD, and then walks over
the map with that step to figure if there is another asteroid blocking
the view.

### Part 2

Part 2 was slightly odd: the task was to remove asteroids from the map
using a clockwise sweep, with non-visible asteroids surviving until
the next sweep. The answer was formed from the coordinates of the
200th removed asteroid.

The solution implements the sweep by first locating all visible
asteroids, then sorting them based on the angle of the line connecting
them. After removing the visible set, the cycle repeats until no
asteroids remain.

## [Day 11](https://adventofcode.com/2019/day/11)

More Intcode again, this time controlling a simple robot capable of
reading the color of its current cell in a 2D grid (through the input
opcode), then painting the cell either black or white and turning 90
degrees, before taking another step.

Part 1 asked how many cells the input program would visit, when
started on an entirely black canvas. Part 2 asked for the word printed
when started on a single white pixel on a black background.

The solution uses a callback-based I/O mechanism for the Intcode
interpreter to drive the robot, but is otherwise very straightforward.

## [Day 12](https://adventofcode.com/2019/day/12)

Today's tasks were on the topic of simulating a discrete, rather
non-physical universe, with interesting laws of gravity and motion.

### Part 1

Part 1 asked for a value derived from the state of the system after
1000 steps of simulation. The solution is the obvious one: apply the
gravity and velocity updates just as stated in the problem
description. The gravity update is done over all distinct pairs,
updating both endpoints, though could as well be done by looping over
the full cartesian product and updating only one side.

### Part 2

For part 2, the task was to determine how many steps will elapse until
the system returns to exactly its initial state. The initial state was
chosen to make that (probably) infeasibly long to just run step by
step.

Since the operations on each of the three spatial dimensions are
entirely independent (gravity has no distance component that would tie
them together), this was easily solvable by just finding out
independently for each of the dimensions how many steps their cycle
is. The full solution is then just the least common multiple of all
three.

## [Day 13](https://adventofcode.com/2019/day/13)

It's an odd day, and we all know what that means: more Intcode!

Today's puzzle input was a simple Breakout clone, where you first had
to count blocks in the initial state, then play the game to the end
(destroy all the blocks) and report the final score.

There are three parts to the solution this time:

- Part 1 simply renders the initial view using ASCII characters (just
  for illustration), then counts blocks.
- Part 2 (interactive mode) is a curses-based implementation that can
  be played manually. Press the left or right arrows to move the
  paddle, any other key to keep it in place and just advance one time
  step. This part is currently commented out.
- Part 2 (demo mode) just runs the game in headless mode until it
  halts, using a simple controller that makes the paddle attempt to
  follow the ball.

## [Day 14](https://adventofcode.com/2019/day/14)

The solution here approaches day 14 as a graph problem.

The puzzle input specified which combination and quantities of
chemicals produced N units of a different chemical, with the
conditions that there was just one reaction per output chemical, and
that no reaction could run partially.

The solution turns the input list into a dependency graph where the
directed edges lead from (e.g.) `A` to `B` and `C` if the (one)
reaction that produces `A` requires `B` and `C` as inputs. Edge
weights denote how many units are needed. Each node also holds the
amount of output units produced by the reaction.

The graph is always a DAG, with the property that the special chemical
`FUEL` is the only node with no in-edges (not needed by any other
reaction), and the special raw material `ORE` the only node with no
out-edges (not produced by any reaction).

Part 1 asked for the minimum amount of `ORE` needed to produce at
least one unit of `FUEL`. The solution computes this by performing a
topological sort of the DAG (with
[Kahn's algorithm](https://doi.org/10.1145%2F368996.369025)), then
iterating over the ordered nodes and keeping track of how much each
chemical is needed. The topological sort guarantees that each node is
visited only after all its ancestors, meaning the full quantity of
that chemical required is known.

Part 2 inverted the problem, and wanted to know the maximum amount of
`FUEL` that can be produced from a trillion `ORE`. To reduce this to a
solved problem, the solution just does a binary search to find the
point where the ore consumption to generate a given amount of fuel
exceeds a trillion.

Check out the `day14-*.png` files to see the dependency graphs. They
were created by the `day14-dot.py` script, which writes the graph in
Graphviz .dot file format.

## [Day 15](https://adventofcode.com/2019/day/15)

Odd is still the new Intcode.

The setting in day 15 involves using an Intcode-driven robot to map a
space, then measure the distance to a target square and time how long
it takes to fill the empty space.

The initial solution here did the mapping using a breadth-first search
starting from the origin, because it naturally gave the answer (length
of shortest path) to part 1. For simplicity, it also moved the robot
back to origin between expanding each of the new fringe nodes. Turns
out the map is actually a labyrinth, so this approach was abysmally
slow, taking a full minute to run. Still, it did the job, producing
the solutions I submitted. This version is still in the revision
history.

The very first solution for part 1 further used to stop short after it
found the target square. For part 2, the full map is needed, so that
version no longer exists at all. Stopping early doesn't really affect
the runtime: the target square is in the far corner of the labyrinth.

The current solution does a depth-first search to build the full map,
which is far more suited to the motions of the robot. With this
solution, it takes significantly less than a second to discover the
full map.

To answer the questions, once the full map is in memory, there's a
simple breadth-first search routine to compute distances to all
(non-wall, reachable) tiles from a given starting point. Part 1 answer
is just the origin-to-target distance, while part 2 is the maximum
distance to any tile from the target square.

## [Day 16](https://adventofcode.com/2019/day/16)

The topic of the day was the Flawed Frequency Transmission
algorithm. Denoting the `i`'th (1-based) digit of the original signal
by `s(i)` and the output signal by `t(i)`, one phase (step) of the FFT
algorithm is represented by the equations:

    t(j) = |sum_(i=1..N) s(i) * W(⌊i / j⌋ mod 4)| mod 10
    W(0..3) = {0, 1, 0, -1}

### Part 1

Part 1 asked for the first 8 digits after 100 phases of FFT on the
puzzle input. There are no immediately obvious speed-ups for this: the
first digits are affected in a somewhat complex manner by the entire
sequence. The solution implements the FFT as a simple translation of
the above equations.

### Part 2

Part 2 required extracting an 8-digit substring after 100 phases
iterated on the puzzle input repeated 10000 times. This is probably
not computationally reasonable by the simple method.

The solution here uses a trick, based on the sequence of values of `W`
used for the digits of the latter half of the signal in the FFT phase.
Note how the expression `W(⌊i / j⌋ mod 4)` behaves when `j` is large
compared to `i`:

    0 <= i < j:   ⌊i / j⌋ = 0  |  W(⌊i / j⌋ mod 4) = 0
    j <= i < 2j:  ⌊i / j⌋ = 1  |  W(⌊i / j⌋ mod 4) = 1

What this means is, for the latter half of the signal (where `i < 2j`
for all digits), we can simplify the algorithm to:

    t(j) = |sum_(i=1..N) s(i) * W(⌊i / j⌋ mod 4)| mod 10
         = |sum_(i=1..j-1) s(i) * 0 + sum_(i=j..N) s(i) * 1| mod 10
         = |sum_(i=j..N) s(i)| mod 10
         = (sum_(i=j..N) s(i)) mod 10

In other words, the `j`'th output digit in the latter half is simply
the sum of all the input digits from its own position to the end of
the signal, and is not affected by the earlier digits at all. This
means that, as long as the message we want to extract is in the latter
half, we can ignore all digits before the message, and further do the
update of the remaining signal in linear time simply by maintaining a
running sum of the digits from right to left.

By a strange coincidence, all the examples and the puzzle input asked
for values in the latter half of the signal. So the solution
constructs only the relevant part of the 10000-fold repeated signal
(in reverse order, for convenience), and then updates it in-place for
100 phases to obtain the answer.

## [Day 17](https://adventofcode.com/2019/day/17)

We'll all be so weirded out when day 25 rolls out and *isn't* related
to Intcode. However, today is not that day.

That being said, the interactions with the Intcode program today were
quite limited. The puzzle input, when executed as-is, simply produced
a map (as an ASCII image) of a (self-intersecting) path. For part 1,
it was enough to look up coordinates of all the intersections.

The program could also be modified to start running a simulation for
moving a robot along that path, accepting as input a very simple
program for the robot: one main function consisting of calls to three
functions A, B and C, together with definitions for those functions as
turtle graphics instructions (90-degree turns, N units forward).

The (not so challenging) challenge was to fit the required movement
rules within the 20-character length limits of the main program and
each of the functions.

Initially, I was planning to identify repeating patterns visually, by
copying the map into a vector graphics program, and producing linked
(optionally rotated) duplicates of the functions that could then be
manipulated to fit the path.

"Unfortunately," it turned out there was a very obvious single path
the robot could take (basically, crossing each intersection straight
through), and when that path was written as a linear program, it was
easy to spot reoccurring substrings and add line breaks to break it
down to functions. See comments of `day17.py` for that breakdown.

Just for the pretty pictures, the finished drawing is also available.
See `day17-plan.png`, or `day17-plan.svgz` for the Inkscape source
image. Mind the
[Inkscape bug](https://bugs.launchpad.net/inkscape/+bug/1462051) I
kept hitting when editing it.

## [Day 18](https://adventofcode.com/2019/day/18)

Ah, the famous traveling burglar problem.

More seriously, today's problem (of finding the shortest path to
collect a set of keys, with constraints that some keys are needed to
reach others) is very close to the *sequential ordering problem*
(SOP), also known as the *(asymmetric) traveling salesman problem with
precedence constraints* (TSP-PC, PCATSP, ...). In the general case,
it's an [NP-hard](http://mathworld.wolfram.com/NP-HardProblem.html)
problem. While the instances in the puzzle are small enough for exact
solvers, it could still be relatively expensive.

Fortunately the scenarios in the puzzle are quite constrained. The
solution here first does a breadth-first search from each key (and the
entry points), figuring out the shortest (mostly, the only) path, as
well as which keys and doors are on that path. This is used to build a
graph of connections between keys (and from the start nodes to each
key).

To find the shortest path to collect all the keys, the solution
performs what's essentially Dijkstra's algorithm on a graph that has a
vertex for every possible *N*-tuple of vertices and subset of keys,
and an edge where it's possible to follow an edge of the original
graph, given the keys and doors involved.

Formally:

```
Let G = (V, E) be the graph of keys and entry points, where for each
edge e = (u, v) ∈ E there are associated values:
- W(e), a positive integer weight
- K(e), the set of keys that will be collected by following that edge
- D(e), the set of keys that are needed to unlock all doors on the way

Now define G' = (V', E') and a modified W'(e) where:
- V' = V^N × 2^K.
- e' = ((u_1 .. u_n, K_u), (v_1 .. v_n, K_v)) ∈ E' if and only if:
  - there is an i ∈ [0, N) s.t. e = (u_i, v_i) ∈ E, u_j = v_j ∀j≠i
  - K_v = K_u ∪ K(e)
  - K_v ⊆ D(e)
- W'(e') = W(e)

The shortest path to collect all keys in G is the same as the shortest
path in G' from (e_1 .. e_n, ∅) to any node (v_1 .. v_n, K), where
(e_1 .. e_n) is the set of entry points of G and K is the set of keys.
```

## [Day 19](https://adventofcode.com/2019/day/19)

Intcode, the gift that keeps on giving. Although the day 19 puzzle was
again remarkably simple.

The provided input program could be used to query whether grid squares
were part of a beam or not. Part 1 asked for the number of points that
were, out of the top-left 50x50 square. This was trivial to calculate
simply by probing all the squares, in a reasonable amount of time.

The answer to part 2 was based on when a 100x100 square first fit
within the beam. A full scan would probably have been too slow, but it
was quite easy to follow the beam just by tracing the left/right edges
down, keep a history of the beam edges for 100 most recent rows, and
then compare whether a box of that size with its lower-left corner at
the beam's left edge would fit.

## [Day 20](https://adventofcode.com/2019/day/20)

A graph puzzle with a twist.

### Part 1

Part 1 was straightforward. The solution here parses the input ASCII
map, locates all portal labels, and for each label computes which
other labels are reachable along the normal paths. Since this is done
with a breadth-first search, it also gives the distance of the
shortest path to each reachable label.

Next, a weighted undirected graph is built, with labels as vertices
and edges between all reachable labels with the distances as weights.
Adding an edge of weight 1 between the matching inner and outer labels
makes the graph a complete description of the first part labyrinth.

The shortest path from `AA` to `ZZ` is now just the shortest path in
the grahp. The solution here uses
[Dijkstra's algorithm](https://doi.org/10.1007%2FBF01386390).

### Part 2

The twist in part 2 is that the inner portals, instead of leading
directly to the corresponding outer portal, leads to a recursive copy
of the maze. Only the outermost level contains the `AA` and `ZZ`
portals.

While this makes the graph technically infinite, the existing
description can still be used. We just need to add a level adjustment
value to each edge: `0` for following a regular path, `+1` for
stepping from an inner portal to an outer one, and `-1` for the
inverse operation.

Since we're stopping early, Dijkstra's algorithm can run on the new
graph almost unmodified. We just keep track of the recursion level in
addition to the node, and ignore edges that would lead to a negative
level.

The solution also uses a maximum distance limit to terminate in case
(as in example 2) there's no path to the exit node at all, but a cycle
that can be used to recurse forever.

## [Day 21](https://adventofcode.com/2019/day/21)

These puzzles are getting to be Intcode-related in name only. It's
probably all just a crafty way of offloading the work from validating
the solution from the puzzle site to the user's computer.

### Part 1

In part 1, the Intcode computer was controlling a *springdroid*, which
could read whether there's ground in front of it in the next four
squares (sensors `A` through `D`), do a sequence of boolean operations
(up to 15), and based on the results either decide to jump (a distance
of 4 steps) or continue walking.

The solution implements the logic of "jump as soon as you detect a
hole in the next three squares, but only if there's ground ahead":

    J = !(A & B & C) & D

This works well for simple holes: the eager jump is delayed until it's
safe.

### Part 2

Unfortunately, the simple logic doesn't work in the presence of traps:

       ---v
    --^   -v
    #####.#.##.#####
      1  234  5

We eagerly jump on square 1, because there's a hole (sensor `C`
reports square 2 is missing), and it seems safe (sensor `D` reports
square 3 is present). Unfortunately, it's a trap: there was a hole
immediately after (square 4), but also four steps in (square 5).  The
program from part 1 considers the jump unsafe and walks into the hole.

Fortunately, the springdroid also has an extended sensor mode, which
allows it to see up to nine tiles ahead (sensors `A` through `I`).

The logic for the part 2 solution starts identical to part 1, but adds
a further check: it also inhibits jumping if the square it landed on
looks like a trap. It detects a trap by checking whether it would be
forced to jump (sensor `E` reporting a hole), but the jump would be
unsafe (sensor `H` reporting a hole as well):

    J = !(A & B & C) & D & !(!E & !H)
      = !(A & B & C) & D & (E | H)

## [Day 22](https://adventofcode.com/2019/day/22)

Similarly to day 16, today's puzzle had a part 1 with an obvious
answer (just apply the specified operations), and part 2 where the
simple way was obviously computationally infeasible.

The puzzle input was a sequence of operations to shuffle a deck of *N*
cards, consisting of the following operations:

* `deal`: reverse the order of the cards
* `cut K`: move the first *K* (or all but the last |*K*| if *K* is
  negative) cards from the top to the bottom of the deck in same order
* `interleave K`: place each original card *K* steps apart, wrapping
  around from the end of the deck

The task of part 1 was to find the location of card 2019, after
shuffling a deck of 10007 cards. The solution here simply applies the
shuffling operations in order, then finds the card in the result.

For part 2, the deck was upgraded to 119315717514047 cards, and the
task was to apply the shuffle 101741582076661 times, then tell which
card ended up in location 2020.

The key insights for the solution here are:

* Since only the card that ends up in location 2020 is important, we
  can just trace backwards through the shuffling operations where that
  card came from. This makes the first big number (size of deck)
  irrelevant.
* The effect of any of the (inverted) shuffling operations on a card's
  location `x` can be represented as `f(x) = (A*x + B) % N`,
  where the parameters `A` and `B` are in `[0, N)`.
* The composition of two such operations can be represented in the
  same form. This means we can turn the entire shuffle into one
  operation, and then decompose the repetitions into successive
  "squarings" (figuring out the effect of repeating the original
  shuffle 1, 2, 4, 8, ... times) and "multiplications" (applying the
  current inverse operation to the current position), making also the
  second big number (number of repetitions) irrelevant.

The reverse operations of the shuffle steps on a single location `x`
are:

```
deal(x)          = N-1 - x
cut(x, K)        = (x + K) % N
interleave(x, K) = (x * K^-1) % N
```

For the last operation, `K^-1` denotes the
*[modular inverse](http://mathworld.wolfram.com/ModularInverse.html)*
of `K` (modulo `N`).

The required compositions of transformations are therefore:

```
Let f(x) = (A*x + B) % N and g(x) = (C*x + D) % N be arbitrary transformations.

deal(f(x))
= N-1 - (A*x + B) % N
= (((-A) % N)*x + N-1 - B) % N

cut(f(x), K)
= ((A*x + B) % N + K) % N
= (A*x + (B + K) % N) % N

interleave(f(x), K)
= ((A*x + B) % N * K^-1) % N
= ((A * K^-1 % N)*x + (B * K^-1) % N) % N

g(f(x))
= (C*((A*x + B) % N) + D) % N
= ((C * A % N)*x + (C * B + D) % N) % N

f(f(x))
= ((A * A % N)*x + (A * B + B) % N) % N
= ((A^2 % N)*x + ((A + 1) * B) % N) % N
```

Or in terms of the updates to the parameters:

```
               A                B
deal           (-A) % N         B
cut K          A                (B + K) % N
interleave K   (A * K^-1) % N   (B * K^-1) % N
f . g          (A * C) % N      (B * C + D) % N
f^2            A^2 % N          ((A + 1) * B) % N
```

## [Day 23](https://adventofcode.com/2019/day/23)

Nothing to say, really. The solution here is pretty crummy: it just
starts each Intcode interpreter as a separate thread, and in the main
thread runs a switch process. The threads communicate by message
passing.

There is a certain amount of non-determinism involved in the exact
sequence of messages, and especially in the idle network detection,
which takes a conservative approach of requiring five unsuccessful
receive attempts without an intervening send before considering a
machine idle.

A coöperative multitasking approach is probably the better option
here. You can check out the corresponding
[Go solution](https://github.com/fis/aoc2019-go/blob/master/day23/day23.go)
for an approach where each of the Intcode machines is executed in
round-robin order until they next request input. In that scheme, the
NAT can determine the network is idle as soon as a single full loop
passes with no send or (successful) receive operations.

## [Day 24](https://adventofcode.com/2019/day/24)

The puzzle for day 24 involved a two-dimensional
[totalistic cellular automaton](http://mathworld.wolfram.com/TotalisticCellularAutomaton.html),
quite similar to the well-known (Conway's)
[Game of Life](http://mathworld.wolfram.com/GameofLife.html),
though with slightly different rules.

Let `c` be the current state of the cell (0/1), and `|n|` be the
number of live cells in the surrounding 4-neighborhood. After one
step:

```
  c  |n|  new state
  0  0    0 - no change
  0  1-2  1 - becomes live
  0  3-4  0 - no change
  1  0    0 - dies
  1  1    1 - no change
  1  2-4  0 - dies
```

Or, represented graphically:

```
 .    .    #    .    .    #    .    .    #    #    .    #    #    .    #    #
...  ..#  ...  #..  ...  ..#  #.#  ..#  #..  ...  #..  #.#  ..#  #.#  #..  #.#
 .    .    .    .    #    .    .    #    .    #    #    .    #    #    #    #

 .    #    #    #    #    #    #    #    #    #    #    .    .    .    .    .


 .    .    #    .    .    #    .    .    #    #    .    #    #    .    #    #
.#.  .##  .#.  ##.  .#.  .##  ###  .##  ##.  .#.  ##.  ###  .##  ###  ##.  ###
 .    .    .    .    #    .    .    #    .    #    #    .    #    #    #    #

 .    #    #    #    #    .    .    .    .    .    .    .    .    .    .    .
```

Or in [B/S notation](https://conwaylife.com/wiki/Rulestring),
`B12/S1V`.

Part 1 simply asked for the first repeating state when the initial
state (puzzle input) was simulated on a regular 5x5 grid, with cells
outside the grid considered permanently empty.

Part 2 was more interesting: the 5x5 grid was swapped with an infinite
construction, where the middle cell was replaced by another 5x5 grid,
and so forth. For a single grid, its border cells were considered
adjacent to (one or two of) the four cells surrounding the (missing)
middle one in the containing level. Correspondingly, each of those
four cells in this grid were adjacent to all five of the relevant
border cells. The task was to simulate a modest number of steps (200)
and report the total number of live cells.

The solution here does not go for anything fancy. States are
represented naturally by 25-bit integers, and the simulation uses a
pre-generated list of neighbors for each cell. For part 2, the list
also contains a value (-1, 0, 1) to indicate which level they're
referring to, and new levels are added to the state list when they
become nonzero.

## [Day 25](https://adventofcode.com/2019/day/25)

Surprisingly, day 25 wasn't surprising: it did feature Intcode.

Only in a very minor way, though. It wasn't really a programming
puzzle as such, but a text adventure implemented in Intcode. The
"solution" here contains a way to run the program interactively, as
well as a quick and dirty loop to get through the door with a sensor
by trying all possible combinations of items.

You might also have a look at `day25-map.png`, which is the map of the
adventure game level.
