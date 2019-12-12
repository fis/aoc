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
origin.  Just for fun, instead of painting every position of the
wires, the solution instead tests every horizontal (resp. vertical)
segment of wire 1 with every vertical (resp. horizontal) segment of
wire 2, and keeps track of the intersections.

### Part 2

Part 2 introduced a notion of a signal propagation delay to the wires,
and asked for the intersection where the combined delay was lowest.

While it would be certainly possible to generalize the part 1 solution
for this, the variant here forgot all about orientation of the
segments.  So instead part 2 is solved the boring way, by just looping
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
them.  After removing the visible set, the cycle repeats until no
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
