# Notes on the solutions

## Day 1

Not much to say here. Part 1 involved evaluating a simple expression
on a list of values and printing the sum. Part 2 had a recurrence
relation of sorts.

## Day 2

Second day introduced the first dialect of the Intcode programming
language. Part 1 asked to run the example program with a specific
initial state, and report which value was left when it halted. Part 2
required finding a specific pair of inputs to yield a given output,
out of just 10000 possibilities.

The solution here uses a standalone Intcode interpreter, only
supporting the opcodes 1, 2 and 99.

## Day 3

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

## Day 4

For some reason, day 4 was surprisingly simple. Both parts simply
asked how many numbers (from a range of around half a million)
satisfied certain simple properties about their digits.

## Day 5

Day 5 was mostly about Intcode extensions. This day's solution is the
first to use the shared Intcode interpreter.

The first part introduced numeric input and output instructions, then
just asked for the output of the program given the input.

Second part extended that with conditional jumps and comparison and
equality operations, then asked for the output given another input.

## Day 6

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

## Day 7

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

## Day 8

Another simple day, decoding images sent as layers of three-state
pixels: white, black, or transparent.

The two parts asked for some statistics of the layers, and the word
obtained by compositing the images, respectively. The solutions are
straightforward. The image renderer writes back-to-front ignoring
transparent pixels, instead of searching front-to-back for the first
non-transparent one.

## Day 9

The "Intcode every other day" pattern continues. Day 9 introduced
allegedly the final missing Intcode features: a relative addressing
mode (to complement the existing immediate and indirect ones), and an
opcode to set the base register for the relative mode.

The tasks involved no other work than finishing the interpreter. The
answers were simply program outputs for inputs 1 and 2.

## Day 10

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

## Day 11

More Intcode again, this time controlling a simple robot capable of
reading the color of its current cell in a 2D grid (through the input
opcode), then painting the cell either black or white and turning 90
degrees, before taking another step.

Part 1 asked how many cells the input program would visit, when
started on an entirely black canvas. Part 2 asked for the word printed
when started on a single white pixel on a black background.

The solution uses a callback-based I/O mechanism for the Intcode
interpreter to drive the robot, but is otherwise very straightforward.

## Day 12

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

## Day 13

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

## Day 14

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

## Day 15

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

## Day 16

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

## Day 17

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
